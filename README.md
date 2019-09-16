# Simple client api - comparation between multiple languages

1. Springboot client:   
   wrk -t12 -c400 -d30s http://localhost:5000/v1/clients
2. Elixir client:   
   wrk -t12 -c400 -d30s http://localhost:4000/api/clients
3. Golang client:   
   wrk -t12 -c400 -d30s http://localhost:3000/v1/clients/all
4. Python 
   TODO

5. Nodejs
   wrk -t12 -c400 -d30s http://localhost:3000/v1/clients

6. Rust

7. wrk -t12 -c400 -d30s http://localhost:3000/v1/clients




the endpoints will return something like this:   
```
[
  {"id":1,"name":"aaaa","bithDate":"yyyy-MM-dd","email":""},
  {"id":2,"name":"bbbb","bithDate":"yyyy-MM-dd","email":""}
]
```

### Memory usage
| seq | language | app | initial | overloaded |
| --- | --- | --- | --- | ---  |
| 1 | kotlin + springboot | clientsb |  519 MB | 672.4 MB |
| 2 | elixir + phoenix    | clientex |  51 MB  | 386.3 MB |
| 3 | golang              | clientgo |  5.6 MB | 30 MB |
| 4 | python              | TODO |  |  |
| 5 | javascript + nodejs | clientjs | 36.8 MB | 36.8 MB |
| 6 | rust                | TODO |  |  |
| 7 | clojure             | clojure_rest_api | 221.1 MB | 221.5 MB |
|   | clojure subprocess  |   | 355.2 MB | 414.9 MB |

 ### WRK results - 12 threads and 400 http connections

1. SpringBoot + Kotlin:    

table with 5 rows:    
```   
Thread Stats   Avg      Stdev     Max   +/- Stdev   
    Latency    83.22ms  158.97ms   2.00s    93.98%   
    Req/Sec   337.30    101.43   595.00     69.84%  
  118344 requests in 30.10s, 53.72MB read   
  Socket errors: connect 0, read 0, write 0, timeout 111   
Requests/sec:   3932.11   
Transfer/sec:      1.78MB   
```
table with 106 rows:   
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
table with 5 rows:   
```
   Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency    25.38ms    5.14ms  81.21ms   72.23%
    Req/Sec     1.31k    91.74     1.78k    78.83%
  468782 requests in 30.09s, 119.45MB read
  Non-2xx or 3xx responses: 468782
Requests/sec:  15577.89
Transfer/sec:      3.97MB
```
table with 106 rows:    
```
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency    73.21ms   20.13ms 259.69ms   69.99%
    Req/Sec   453.50     45.83   595.00     72.57%
  162002 requests in 30.05s, 1.20GB read
Requests/sec:   5390.62
Transfer/sec:     40.83MB

```

3. Golang + Gorilla Mux:   
table with 5 rows:    
```
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency   240.66ms  409.91ms   2.00s    84.76%
    Req/Sec   199.88    403.17     3.98k    93.64%
  60682 requests in 30.08s, 29.43MB read
  Socket errors: connect 0, read 0, write 0, timeout 1789
  Non-2xx or 3xx responses: 8859
Requests/sec:   2017.05
Transfer/sec:      0.98MB

```
table with 106 rows:   
```
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency   192.28ms  335.69ms   2.00s    87.60%
    Req/Sec   303.46    330.86     1.82k    87.25%
  99431 requests in 30.06s, 675.91MB read
  Socket errors: connect 0, read 0, write 0, timeout 1034
  Non-2xx or 3xx responses: 29513
Requests/sec:   3307.67
Transfer/sec:     22.48MB

```
0bs.: Many errors: [mysql] 2019/09/15 11:20:54 packets.go:36: unexpected EOF

4. Python 
```
TODO
```


5. Javascript - Nodejs
table with 5 rows:    
```
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency    98.78ms    8.77ms 342.61ms   93.23%
    Req/Sec   331.17     43.07   590.00     90.92%
  117415 requests in 30.09s, 83.65MB read
Requests/sec:   3901.77
Transfer/sec:      2.78MB
```

6. Rust
```
```

7. Clojure:
table with 5 rows
```
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency   134.84ms   72.00ms   1.17s    83.26%
    Req/Sec   251.85     50.83     0.88k    76.66%
  88215 requests in 30.10s, 47.11MB read
Requests/sec:   2930.85
Transfer/sec:      1.57MB

```


|  seq  |  project  |  rows  |  Latency Avg  |  Stdev  |  Max  |  +/- Stdev  |  Req/Sec Avg  |  Stdev  |  Max  |  +/- Stdev  |  requests  |  Requests/sec  |  Transfer/sec  |
|---|---|---|---|---|----|---|---|---|---|---|---|---|---|
| 1 | Springboot | 5 | 83.22ms | 158.97ms | 2.00s | 93.98% | 337.30 | 101.43 | 595.00 | 69.84% | 118344 in 30.10s, 53.72MB read | 3932.11 | 1.78MB |
|   |        | 5 | 224.40 | 73.44 | 414.00 | 67.41% | 78327 requests in 30.08s, 515.05MB read | 2603.93 | 17.12MB |
| 2 | Elixir | 5 | 25.38ms | 5.14ms | 81.21ms | 72.23% | 1.31k | 91.74 | 1.78k | 78.83% | 468782 requests in 30.09s, 119.45MB read | 15577.89 | 3.97MB |
|   |        | 5 | 73.21ms | 20.13ms | 259.69ms | 69.99% | 453.50 | 45.83 | 595.00 | 72.57% | 162002 requests in 30.05s, 1.20GB read | 5390.62 | 40.83MB |
| 3 | Golang | 5 | 240.66ms | 409.91ms | 2.00s | 84.76% | 199.88 | 403.17 | 3.98k | 93.64% | 60682 requests in 30.08s, 29.43MB read | 2017.05 | 0.98MB |
|   |        | 106 | 192.28ms | 335.69ms | 2.00s | 87.60% | 303.46 | 330.86 | 1.82k | 87.25% | 9431 requests in 30.06s, 675.91MB read | 3307.67 | 22.48MB |
| 5 | Javascript | 5 | 98.78ms | 8.77ms | 342.61ms | 93.23% | 331.17 | 43.07 | 590.00 | 90.92% |   117415 requests in 30.09s, 83.65MB read | 3901.77 | 2.78MB |
| 7 | Clojure | 5 | 134.84ms | 72.00ms | 1.17s | 83.26% | 251.85 | 50.83 | 0.88k | 76.66% | 88215 in 30.10s, 47.11MB read | 2930.85 | 1.57MB |


        