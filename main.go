package main

import (
    "fmt"
    "flag"
    "bufio"
    "os"
    "time"
    "sync"
    "sync/atomic"
    "strconv"
    "errors"
    "net/http"
    "log"
)

var Wait int
var infoLog *log.Logger
var errorLog *log.Logger

type ping struct {
    url string
    
    err1 error //invalid url
    err2 error //scanner error
    err3 error //redirect
    status string 
        
    lines int
    duration time.Duration
}

func (p *ping) check() {
    var resp *http.Response
    var err1 error
    done := make(chan bool,1)
    
    timer1 := time.NewTimer(time.Duration(Wait) * time.Millisecond)
    start := time.Now()
    
    go func(r **http.Response) {
        *r, err1 = http.Get(p.url)
        p.err1 = err1
        done<-true
        close(done)
    }(&resp)
    
    select {
    case <-timer1.C:
        p.err1 = errors.New("Timer expired. No answer from URL")
    case <-done:
        timer1.Stop()
        p.duration = time.Since(start)
    }
    if p.err1 != nil {
        return
    }
    
    defer resp.Body.Close()
    p.status = resp.Status
    
    scanner := bufio.NewScanner(resp.Body)
    for i := 0; scanner.Scan() && i < 5; i++ {
        p.lines++
    }
    
    err2 := scanner.Err()
    if err2 != nil {
        p.err2 = err2
    }
    if p.lines < 5 {
        p.err3 = errors.New("Redirect")
    }
}

func (p *ping) output() {
    infoLog.Printf("%s;%s;%v", p.url, p.status, p.duration.Milliseconds()) 
    if p.err1 != nil {
        errorLog.Printf("URL: %s ERROR_TYPE_1: %v", p.url, p.err1) 
        return
    }
    if p.err2 != nil {
        errorLog.Printf("URL: %s ERROR_TYPE_2: %v", p.url, p.err2) 
        return
    }
    if p.err3 != nil {
        errorLog.Printf("URL: %s ERROR_TYPE_3: %v", p.url, p.err3) 
        return
    }
}

func main() {
    infoLog = log.New(os.Stdout, "", 0)
    errorLog = log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)
    infoLog.Printf("url;status;milliseconds")
    
    fileName := flag.String("file", "urls.txt", "list of urls")
    workersNumber := flag.String("workers", "100", "max number of simultaneous URL checks")
    rpsLimit := flag.String("rps", "200", "max rate limit")
    waitURL := flag.String("wait", "1000", "milliseconds wait for response from URL")
	flag.Parse()
    
    w, err := strconv.Atoi(*waitURL)
    if err != nil {
        errorLog.Fatal(err)
    }
    if w <= 0 {
        errorLog.Fatal(errors.New("can't wait <= 1 sec"))
    }
    Wait = w
    
    rps, err := strconv.Atoi(*rpsLimit)
    if err != nil {
        errorLog.Fatal(err)
    }
    
    workers, err := strconv.Atoi(*workersNumber)
    if err != nil {
        errorLog.Fatal(err)
    }
    if workers <= 0 {
        errorLog.Fatal(errors.New("max workers <= 0"))
    }
         
    file, err := os.Open(*fileName)
    if err != nil {
        errorLog.Fatal(err)
    }
    defer file.Close()
    r := bufio.NewReader(file)
    
    var wg sync.WaitGroup
    urls := make(chan string, workers)
    var ops uint64
    
    start := time.Now()
    
    wg.Add(1)
    go func() {
        var limit float64
        if rps <= 0 {
            limit = 1
        } else {
            limit = 1000./float64(rps)
        }
        limiter := time.Tick(time.Duration(limit) * time.Millisecond)
        for  {
            s, err := r.ReadString('\n')
            if err!= nil {
                close(urls) 
                wg.Done()
                return
            }
            <-limiter
            urls <- s[:len(s)-1]
        }
    }()    
    
    for w := 0; w < workers; w++ {
        wg.Add(1)
        go func() {
            for {
                url, ok := <-urls
                if ok != true && len(urls) == 0 {
                    wg.Done()
                    break
                }
                p := ping{url:url}
                p.check()
                atomic.AddUint64(&ops, 1)
                //fmt.Println("atomic:",ops)
                
                wg.Add(1)
                go p.output()
                wg.Done()
            }
        }()
    }
    
    wg.Wait()
    duration := time.Since(start)
    
    time.Sleep(1 * time.Second)
    fmt.Println("")
    fmt.Println("file: ", *fileName)
    fmt.Println("workers: ", workers)
    fmt.Println("milliseconds wait for response from URL: ", w)
    fmt.Println("")
    fmt.Println("urls: ", ops)
    fmt.Println("duration:", duration)
    fmt.Println("rps:",int(float64(ops)/duration.Seconds()), " / limit: ", rps)
}
