#!/bin/bash

echo "-------- setup.sh start -----------"
docker stop `docker ps -q`
docker rm mysql nginx airmeet
docker run --name data-mysql -v /var/lib/mysql busybox
docker run --name myapp-gopath -itd -v /go busybox

#docker run --rm --volumes-from myapp-gopath -v $PWD:/go/src/app golang:1.6.2 go get github.com/labstack/echo/...
#docker run --rm --volumes-from myapp-gopath -v $PWD:/go/src/app golang:1.6.2 go get github.com/satori/go.uuid
#docker run --rm --volumes-from myapp-gopath -v $PWD:/go/src/app golang:1.6.2 go get gopkg.in/go-playground/validator.v8
#docker run --rm --volumes-from myapp-gopath -v $PWD:/go/src/app golang:1.6.2 go get github.com/go-sql-driver/mysql
#docker run --rm --volumes-from myapp-gopath -v $PWD:/go/src/app golang:1.6.2 go get github.com/jinzhu/gorm

docker run \
	--name nginx \
	-v /var/run/docker.sock:/tmp/docker.sock:ro \
	-p 80:80 \
	-d \
	jwilder/nginx-proxy


docker run \
	--name mysql \
	--volumes-from data-mysql \
	-v $PWD/mysql/conf:/etc/mysql/conf.d \
	-v $PWD/mysql/init:/docker-entrypoint-initdb.d \
	-e MYSQL_ROOT_PASSWORD=$DB_PASS \
	-e MYSQL_DATABASE=airmeet \
	-d \
	-t \
	-i \
	-p 3306:3306 \
	mysql mysqld

docker build -t airmeet:0.1 airmeet
docker run \
	--name airmeet \
	--link mysql \
	-e VIRTUAL_HOST=$LOCAL_VIRTUAL_HOST \
	-e DB_PASS=$DB_PASS \
	-d \
	-p 3000:3000 \
	airmeet:0.1



echo "-------- setup.sh end -----------"
