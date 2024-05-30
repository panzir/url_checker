# Concurrent URL Checker with Rate Limiting
## 1. Run the program (with additional flags or without)<br>
###  #go run main.go 2> log.txt 1> output.csv
    - log.txt or another file name for errors logging
    - output.csv or another file name for results
##  Additional flags:
change url-source file / list of urls:
###    --file=urls.txt<br>
change number of workers / max number of simultaneous URL checks:
###    --workers=200<br>
change max requests per second / max rate limit:
###    --rps=200<br>
change max possible answer time / milliseconds wait for response from URL:
###   --wait=1000<br>

## 2. Output format
url; response_status; response_time<br><br>
Examples: <br>
- url answered without errors after 123 milliseconds<br>
  http: //notfake.notfake; 200 OK; 123<br>
- url not answered <br>
  https: //fake.fake; ; 0<br>

### Summary at the end of output file
file:  urls.txt<br>
workers:  100<br>
milliseconds wait for response from URL:  1000<br>
urls:  500 <br>
duration: 3.935202075s <br>
rps: 127  / limit:  200 <br>

