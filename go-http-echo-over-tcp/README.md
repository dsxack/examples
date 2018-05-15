# go_http_echo_over_tcp

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
Time taken for tests:   112.309 seconds
Complete requests:      1000000
Failed requests:        0
Total transferred:      2171002880 bytes
Total body sent:        2087000000
HTML transferred:       1990000000 bytes
Requests per second:    8904.03 [#/sec] (mean)
Time per request:       11.231 [ms] (mean)
Time per request:       0.112 [ms] (mean, across all concurrent requests)
Transfer rate:          18877.61 [Kbytes/sec] received
                        18147.18 kb/s sent
                        37024.80 kb/s total

Connection Times (ms)
              min  mean[+/-sd] median   max
Connect:        0    3   2.6      3      41
Processing:     0    8   6.4      6     123
Waiting:        0    6   5.9      5     103
Total:          0   11   6.5     10     131

Percentage of the requests served within a certain time (ms)
  50%     10
  66%     12
  75%     13
  80%     14
  90%     18
  95%     23
  98%     30
  99%     35
 100%    131 (longest request)
```
