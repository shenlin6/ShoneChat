#!/bin/bash
reso_addr='registry.cn-hangzhou.aliyuncs.com/shone-chat/im-api-dev'
tag='latest'

container_name="shone-chat-im-api-test"

docker stop ${container_name}

docker rm ${container_name}

docker rmi ${reso_addr}:${tag}

docker pull ${reso_addr}:${tag}


docker run -p 8882:8882  --name=${container_name} -d ${reso_addr}:${tag}