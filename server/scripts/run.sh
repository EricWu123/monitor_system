#!/bin/bash

# 登录 admin 用户
curl -X POST -i http://127.0.0.1:8090/login --data '{
    "userName": "admin",
    "password": "BiTYKRZpk8VONUjnZ8XMeA=="
}' -c cookie.txt

# 查看系统信息
curl -X POST -i http://127.0.0.1:8090/query/system_info --data '{
    "begin": "1639305120",
    "end": "1639316241",
    "HostName": "wyq-System-Product-Name",
    "OS": "linux"
}' -b @cookie.txt


# 登录 guest 用户
curl -X POST -i http://127.0.0.1:8090/login --data '{
    "userName": "guest",
    "password": "PdGCLl/peFOFgjyo9ufF9w=="
}' -c cookie.txt

# 查看系统信息
curl -X POST -i http://127.0.0.1:8090/query/system_info --data '{
    "begin": "1639305120",
    "end": "1639316241",
    "HostName": "wyq-System-Product-Name",
    "OS": "linux"
}' -b @cookie.txt

# 无cookie
curl -X POST -i http://127.0.0.1:8090/query/system_info --data '{
    "begin": "1639305120",
    "end": "1639316241",
    "HostName": "wyq-System-Product-Name",
    "OS": "linux"
}'

