# telnet连接测试客户端

操作说明
```
$ ./telnet-client client 192.168.75.118 9092
###################Telnet Client##################
连接成功，按Ctrl + C 退出。
^C

$ ./telnet-client client 192.168.75.111 9093
###################Telnet Client##################
[2020-09-08 10:56:19]   error   dial tcp 192.168.75.111:9093: connect: connection refused
###################Error Exist##################

```