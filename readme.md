## How to Run

1. Clone this repository
2. download library using command Go get github.com/julienschmidt/httprouter
3. run app using command Go run app.go

## Curl Testing:
``` 
curl --location 'http://127.0.0.1:8080/range-fizzbuzz?from=1&to=100' 

```

## Features:
1. New Endpoint Get range-fizzbuzz.
2. Log are safe in fizzbuzz.log.
3. Implement Validation Range, which `from` Argument must be 1 at least and the `to` Argument must no more than 100.
4. Implement Gracefull shutdown within 10s after get signal interrupt.