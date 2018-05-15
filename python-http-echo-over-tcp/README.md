# python_http_echo_over_tcp

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
Time taken for tests:   494.680 seconds
Complete requests:      1000000
Failed requests:        0
Total transferred:      2171000000 bytes
Total body sent:        2087000000
HTML transferred:       1990000000 bytes
Requests per second:    2021.51 [#/sec] (mean)
Time per request:       49.468 [ms] (mean)
Time per request:       0.495 [ms] (mean, across all concurrent requests)
Transfer rate:          4285.84 [Kbytes/sec] received
                        4120.01 kb/s sent
                        8405.85 kb/s total

Connection Times (ms)
              min  mean[+/-sd] median   max
Connect:        0    0   1.3      0      44
Processing:     1   49  11.8     49     190
Waiting:        1   47  11.5     47     154
Total:          5   49  11.3     49     190

Percentage of the requests served within a certain time (ms)
  50%     49
  66%     52
  75%     55
  80%     57
  90%     65
  95%     69
  98%     73
  99%     76
 100%    190 (longest request)
 ```
