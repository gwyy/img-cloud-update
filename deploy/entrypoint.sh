#!/bin/bash

/usr/sbin/nginx &  
cd /opt/img-cloud-update/server/ && ./server &


echo "imgCloudUpdate ALL start!!!"
tail -f /dev/null