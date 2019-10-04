# Simple client api - comparation between multiple languages

1. Kotlin+Springboot http://localhost:5000/v1/clients
2. Elixir  http://localhost:4000/v1/clients
3. Golang  http://localhost:3000/v1/clients
4. Nodejs  http://localhost:3000/v1/clients
5. Clojure http://localhost:3000/v1/clients
6. Rust    http://localhost:9001/v1/clients
7. Python  http://localhost:9000/v1/clients

the endpoints will return something like this:   
```
[
  {"id":1,"name":"aaaa","bithDate":"yyyy-MM-dd","email":""},
  {"id":2,"name":"bbbb","bithDate":"yyyy-MM-dd","email":""}
]
```

# TODOS
- cache with redis
- service discovery
- api gateway 
- log tracing
- message broker producer/consumer with rabbitmq


### Memory usage
| seq | language | app | initial | overloaded |
| --- | --- | --- | --- | ---  |
| 1 | kotlin + springboot | clientsb |  519 MB | 672.4 MB |
| 2 | elixir + phoenix    | clientex |  51 MB  | 386.3 MB |
| 3 | golang              | clientgo |  5.6 MB | 30 MB |
| 4 | javascript + nodejs | clientjs | 36.8 MB | 36.8 MB |
| 5 | clojure             | clojure_rest_api | 221.1 MB | 221.5 MB |
|   | clojure subprocess  |   | 355.2 MB | 414.9 MB |
| 6 | rust                | clientrs | 9.4mb | 9.4mb |
| 7 | python              | clientpy | 23.1mb | 29.9mb |


### wrk -t12 -c400 -d30s http://localhost:PORT/v1/clients - 5 rows
|     | project     | Latency Avg | Stdev    | Max      | +/- Stdev | Req/Sec Avg | Stdev  | Max    | +/- Stdev | requests  |  Readed    | Requests/sec | Transfer/sec |
| --- | ----------- | ----------- | -------- | -------- | --------- | ----------- | ------ | ------ | --------- | --------- |  ----------| ------------ | ------------ |
| 1   | Kotlin      | 83.22ms     | 158.97ms | 2.00s    | 93.98%    | 337.30      | 101.43 | 595.00 | 69.84%    | 118344    |  53.72MB   | 3932.11      | 1.78MB       |
| 2   | Elixir      | 25.38ms     | 5.14ms   | 81.21ms  | 72.23%    | 1.31k       | 91.74  | 1.78k  | 78.83%    | 468782    |  119.45MB  | 15577.89     | 3.97MB       |
| 4   | Javascript  | 98.78ms     | 8.77ms   | 342.61ms | 93.23%    | 331.17      | 43.07  | 590.00 | 90.92%    | 117415    |  83.65MB   | 3901.77      | 2.78MB       |
| 5   | Clojure     | 134.84ms    | 72.00ms  | 1.17s    | 83.26%    | 251.85      | 50.83  | 0.88k  | 76.66%    | 88215     |  47.11MB   | 2930.85      | 1.57MB       |

### wrk -t12 -c400 -d30s http://localhost:PORT/v1/clients - 100 rows 
|     | project     | Latency Avg | Stdev    | Max      | +/- Stdev | Req/Sec Avg | Stdev  | Max    | +/- Stdev | requests  |  Readed    | Requests/sec | Transfer/sec |
| --- | ----------- | ----------- | -------- | -------- | --------- | ----------- | ------ | ------ | --------- | --------- |  ----------| ------------ | ------------ |
| 2   | Kotlin      | 126.71ms    | 224.67ms | 2.00s    | 93.26%    | 224.40      | 73.44  | 414.00 | 67.41%    | 78327     |  515.05MB  | 2603.93      | 17.12MB      |
| 2   | Elixir      | 73.21ms     | 20.13ms  | 259.69ms | 69.99%    | 453.50      | 45.83  | 595.00 | 72.57%    | 162002    |  1.20GB    | 5390.62      | 40.83MB      |
| 3   | Golang      | 93.79ms     | 163.38ms | 1.98s    | 88.47%    | 0.99k       | 179.00 | 1.74k  | 67.52%    | 355362    |  3.14GB    | 11811.47     | 106,75       |
| 6   | Rust        | 17.78ms     | 61.98ms  | 1.68s    | 98.42%    | 1.04k       | 660.48 | 2.85k  | 58.94%    | 366504    |  2.39GB    | 12179.52     | 81.49MB      |
| 7   | Python      | 1.69s       | 229.77ms | 1.99s    | 93.43%    | 12.87       | 9.81   | 70.00  | 73.47%    | 2495      |  21.92MB   | 82.91        | 745.78KB     |


# WRK on console

1. SpringBoot + Kotlin:   
5 rows:   
```   
Thread Stats   Avg      Stdev     Max   +/- Stdev   
    Latency    83.22ms  158.97ms   2.00s    93.98%   
    Req/Sec   337.30    101.43   595.00     69.84%  
  118344 requests in 30.10s, 53.72MB read   
  Socket errors: connect 0, read 0, write 0, timeout 111   
Requests/sec:   3932.11   
Transfer/sec:      1.78MB   
```
106 rows:   
```
   Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency   126.71ms  224.67ms   2.00s    93.26%
    Req/Sec   224.40     73.44   414.00     67.41%
  78327 requests in 30.08s, 515.05MB read
  Socket errors: connect 0, read 0, write 0, timeout 97
Requests/sec:   2603.93
Transfer/sec:     17.12MB
```
2. Elixir + Phoenix:   
5 rows:   
```
   Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency    25.38ms    5.14ms  81.21ms   72.23%
    Req/Sec     1.31k    91.74     1.78k    78.83%
  468782 requests in 30.09s, 119.45MB read
  Non-2xx or 3xx responses: 468782
Requests/sec:  15577.89
Transfer/sec:      3.97MB
```
106 rows:       
```
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency    73.21ms   20.13ms 259.69ms   69.99%
    Req/Sec   453.50     45.83   595.00     72.57%
  162002 requests in 30.05s, 1.20GB read
Requests/sec:   5390.62
Transfer/sec:     40.83MB
```
100 rows:    
```
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency    72.10ms   17.01ms 135.77ms   63.69%
    Req/Sec   459.08     41.55   660.00     74.06%
  164649 requests in 30.06s, 1.15GB read
Requests/sec:   5477.29
Transfer/sec:     39.10MB

```
3. Golang + Gorilla Mux:   
100 rows:   
```
 Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency    93.79ms  163.38ms   1.98s    88.47%
    Req/Sec     0.99k   179.00     1.74k    67.52%
  355362 requests in 30.09s, 3.14GB read
Requests/sec:  11811.47
Transfer/sec:    106.75MB
```
4. Javascript - Nodejs   
5 rows:    
```
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency    98.78ms    8.77ms 342.61ms   93.23%
    Req/Sec   331.17     43.07   590.00     90.92%
  117415 requests in 30.09s, 83.65MB read
Requests/sec:   3901.77
Transfer/sec:      2.78MB
```
5. Clojure:   
5 rows   
```
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency   134.84ms   72.00ms   1.17s    83.26%
    Req/Sec   251.85     50.83     0.88k    76.66%
  88215 requests in 30.10s, 47.11MB read
Requests/sec:   2930.85
Transfer/sec:      1.57MB
```
6. Rust   
100 rows   
```
 Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency    17.78ms   61.98ms   1.68s    98.42%
    Req/Sec     1.04k   660.48     2.85k    58.94%
  366504 requests in 30.09s, 2.39GB read
  Socket errors: connect 0, read 366500, write 0, timeout 13
Requests/sec:  12179.52
Transfer/sec:     81.49MB
```
7. Python   
100 rows   
```
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency     1.69s   229.77ms   1.99s    93.43%
    Req/Sec    12.87      9.81    70.00     73.47%
  2495 requests in 30.09s, 21.92MB read
  Socket errors: connect 0, read 58, write 0, timeout 120
Requests/sec:     82.91
Transfer/sec:    745.78KB
```