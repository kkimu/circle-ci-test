#!bin/bash
docker run --rm --volumes-from myapp-gopath -v $PWD:/go/src/app golang:1.6.2 go get github.com/labstack/echo/...
docker run --rm --volumes-from myapp-gopath -v $PWD:/go/src/app golang:1.6.2 go get github.com/satori/go.uuid
docker run --rm --volumes-from myapp-gopath -v $PWD:/go/src/app golang:1.6.2 go get gopkg.in/go-playground/validator.v8
docker run --rm --volumes-from myapp-gopath -v $PWD:/go/src/app golang:1.6.2 go get github.com/go-sql-driver/mysql
docker run --rm --volumes-from myapp-gopath -v $PWD:/go/src/app golang:1.6.2 go get github.com/jinzhu/gorm

docker-compose up
