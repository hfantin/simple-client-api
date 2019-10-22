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
|     | project     | Latency Avg | Stdev    | Max      | +/- Stdev | Req/Sec Avg | Stdev  | Max    | +/- Stdev | requests  |  Readed    | Requests/sec | Transfer/sec | mem           |
| --- | ----------- | ----------- | -------- | -------- | --------- | ----------- | ------ | ------ | --------- | --------- |  ----------| ------------ | ------------ |---------------|
| 1   | Kotlin      | 114.42ms    | 198.10ms | 2.00s    | 93.90%    | 232.97      | 65.53  | 410.00 | 68.09%    | 81301     |  564.23MB  | 2701.33      | 18.87MB      | 509 ~ 726 MB  |
| 2   | Elixir      | 68.93ms     | 17.62ms  | 223.57ms | 73.96%    | 485.49      | 49.42  | 2.12k  | 80.97%    | 172680    |  1.21GB    | 5770.42      | 41.20MB      | 50 ~ 317 MB   |
| 3   | Golang      | 87.20ms     | 144.51ms | 1.38s    | 87.64%    | 1.02k       | 208.93 | 2.46k  | 73.22%    | 361334    |  2.42GB    | 12006.20     | 82.18        | 5.3 ~ 19.4 MB |
| 4   | Javascript  | 235.96ms    | 19.18ms  | 493.55ms | 92.35%    | 139.36      | 79.06  | 333.00 | 59.07%    | 49300     |  421.92MB  | 1638.45      | 14.02MB      |  35.9 ~ 92 MB |
| 5   | Clojure     | 157.69ms    | 51.83ms  | 550,22ms | 86.39%    | 203.89      | 48.34  | 530.00 | 78.27%    | 72495     |  581.16MB  | 2410.85      | 19.33MB      |  539 ~ 662 MB |
| 6   | Rust        | 19.23ms     | 71.06ms  | 1.68s    | 98.14%    | 1.03k       | 580.22 | 2.52k  | 61.22%    | 367449    |  2.49GB    | 11222.90     | 84.88MB      | 9.4 ~ 9.4 MB  |
| 7   | Python      | 1.64s       | 223.95ms | 1.97s    | 92.12%    | 12.04       | 9.27   | 60.00  | 77.30%    | 2559      |  22.40MB   | 85.13        | 765.76KB     | 23.2 ~ 28,8 MB |


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
100 rows:   
```
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency   114.42ms  198.10ms   2.00s    93.90%
    Req/Sec   232.97     65.53   410.00     68.09%
  81301 requests in 30.10s, 564.23MB read
  Socket errors: connect 0, read 0, write 0, timeout 105
Requests/sec:   2701.33
Transfer/sec:     18.75MB
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
100 rows:       
```
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency    68.37ms   17.62ms 223.57ms   73.96%
    Req/Sec   485.49     49.42     2.12k    80.97%
  173680 requests in 30.10s, 1.21GB read
Requests/sec:   5770.42
Transfer/sec:     41.20MB

```
3. Golang + Gorilla Mux:   
100 rows:   
```
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency    87.20ms  144.51ms   1.38s    87.64%
    Req/Sec     1.02k   208.93     2.46k    73.32%
  361334 requests in 30.10s, 2.42GB read
Requests/sec:  12006.20
Transfer/sec:     82.18MB

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
100 rows: 
```
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency   235.96ms   19.18ms 493.55ms   92.35%
    Req/Sec   139.36     79.06   333.00     59.07%
  49300 requests in 30.09s, 421.92MB read
Requests/sec:   1638.45
Transfer/sec:     14.02MB

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
100 rows: 
```
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency   157.69ms   51.83ms 550.22ms   86.39%
    Req/Sec   203.89     48.34   530.00     78.27%
  72495 requests in 30.07s, 581.16MB read
Requests/sec:   2410.85
Transfer/sec:     19.33MB

```

6. Rust   
100 rows   
```
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency    19.23ms   71.06ms   1.68s    98.14%
    Req/Sec     1.03k   580.22     2.52k    61.22%
  367449 requests in 30.06s, 2.49GB read
  Socket errors: connect 0, read 367503, write 0, timeout 32
Requests/sec:  12222.90
Transfer/sec:     84.88MB
```
7. Python   
100 rows   
```
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency     1.64s   223.95ms   1.97s    92.12%
    Req/Sec    12.04      9.27    60.00     77.30%
  2559 requests in 30.06s, 22.48MB read
  Socket errors: connect 0, read 72, write 0, timeout 134
Requests/sec:     85.13
Transfer/sec:    765.76KB


```