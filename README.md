# Concurrent URL Checker with Rate Limiting
## 1 Run the program (with additional flags or without)<br>
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

## 2 Output format
url; response_status; response_time<br><br>
Examples: <br>
- url answered without errors after 123 milliseconds<br>
  http: //notfake.notfake; 200 OK; 123<br>
- url not answered <br>
  https: //fake.fake; ; 0<br>

### Summary at the end of the output file
file:  urls.txt<br>
workers:  100<br>
milliseconds wait for response from URL:  1000<br>
urls:  500 <br>
duration: 3.935202075s <br>
rps: 127  / limit:  200 <br>

## 3 Logging errors
Error logs have 3 types of possible errors:
- ERROR_TYPE_1 / err1 - invalid url
- ERROR_TYPE_2 / err2 - scanner error
- ERROR_TYPE_3 / err3 - redirect

## 4 Unit tests
For testing run<br>
###  #go test main_test.go






