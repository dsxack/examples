# java_http_echo_over_tcp

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
Time taken for tests:   118.174 seconds
Complete requests:      1000000
Failed requests:        0
Total transferred:      2171000135 bytes
Total body sent:        2087000000
HTML transferred:       1990000000 bytes
Requests per second:    8462.07 [#/sec] (mean)
Time per request:       11.817 [ms] (mean)
Time per request:       0.118 [ms] (mean, across all concurrent requests)
Transfer rate:          17940.58 [Kbytes/sec] received
                        17246.42 kb/s sent
                        35187.00 kb/s total

Connection Times (ms)
              min  mean[+/-sd] median   max
Connect:        0    5  57.1      1    3002
Processing:     1    7   8.1      6     317
Waiting:        0    6   7.9      5     316
Total:          3   12  58.1      7    3007

Percentage of the requests served within a certain time (ms)
  50%      7
  66%      9
  75%     10
  80%     11
  90%     13
  95%     14
  98%     17
  99%     22
 100%   3007 (longest request)
 ```
