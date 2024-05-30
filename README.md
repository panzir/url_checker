# Concurrent URL Checker with Rate Limiting
1. Run the program (with additional flags or without)
#go run main.go 2> log.txt 1> output.csv
- log.txt or another file name for errors logging
- output.csv or another file name for results
Additional flags:
- change url-source file / list of urls
--file=urls.txt
- change number of workers / max number of simultaneous URL checks
--workers=200
- change max requests per second / max rate limit
--rps=200
- change max possible answer time / milliseconds wait for response from URL
--wait=1000

2. 
