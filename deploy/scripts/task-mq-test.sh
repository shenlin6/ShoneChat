#!/bin/bash
reso_addr='registry.cn-hangzhou.aliyuncs.com/shone-chat/task-mq-dev'
tag='latest'

container_name="shone-chat-task-mq-test"

docker stop ${container_name}

docker rm ${container_name}

docker rmi ${reso_addr}:${tag}

docker pull ${reso_addr}:${tag}

docker run  --name=${container_name} -d ${reso_addr}:${tag}