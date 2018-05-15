# rust_http_echo_over_tcp

```
ab -c 100 -n 1000000 -p /etc/ssl/certs/WoSign.pem http://127.0.0.1:8000/
This is ApacheBench, Version 2.3 <$Revision: 1706008 $>
Copyright 1996 Adam Twiss, Zeus Technology Ltd, http://www.zeustech.net/
Licensed to The Apache Software Foundation, http://www.apache.org/

Benchmarking 127.0.0.1 (be patient)
Completed 100000 requests
Completed 200000 requests
Completed 300000 requests
Completed 400000 requests
Completed 500000 requests
Completed 600000 requests
Completed 700000 requests
Completed 800000 requests
Completed 900000 requests
Completed 1000000 requests
Finished 1000000 requests


Server Software:        
Server Hostname:        127.0.0.1
Server Port:            8000

Document Path:          /
Document Length:        1990 bytes

Concurrency Level:      100
Time taken for tests:   286.382 seconds
Complete requests:      1000000
Failed requests:        0
Total transferred:      2171000000 bytes
Total body sent:        2087000000
HTML transferred:       1990000000 bytes
Requests per second:    3491.84 [#/sec] (mean)
Time per request:       28.638 [ms] (mean)
Time per request:       0.286 [ms] (mean, across all concurrent requests)
Transfer rate:          7403.10 [Kbytes/sec] received
                        7116.66 kb/s sent
                        14519.77 kb/s total

Connection Times (ms)
              min  mean[+/-sd] median   max
Connect:        0    0   0.3      0      20
Processing:     0   28  11.4     26     350
Waiting:        0   28  11.1     25     349
Total:          0   29  11.4     26     350

Percentage of the requests served within a certain time (ms)
  50%     26
  66%     29
  75%     32
  80%     34
  90%     40
  95%     48
  98%     59
  99%     70
 100%    350 (longest request)
 ```
