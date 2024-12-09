#!/bin/bash
reso_addr='registry.cn-hangzhou.aliyuncs.com/shonechat/social-api-dev'
tag='latest'

container_name="shonechat-social-api-test"

docker stop ${container_name}

docker rm ${container_name}

docker rmi ${reso_addr}:${tag}

docker pull ${reso_addr}:${tag}



docker run -p 8881:8881  --name=${container_name} -d ${reso_addr}:${tag}