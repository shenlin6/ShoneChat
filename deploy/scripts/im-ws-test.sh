#!/bin/bash
reso_addr='registry.cn-hangzhou.aliyuncs.com/shone-chat/im-ws-dev'
tag='latest'

pod_ip="192.168.117.24"

container_name="shone-chat-im-ws-test"

docker stop ${container_name}

docker rm ${container_name}

docker rmi ${reso_addr}:${tag}

docker pull ${reso_addr}:${tag}

docker run -p 10090:10090 -e POD_IP=${pod_ip}  --name=${container_name} -d ${reso_addr}:${tag}