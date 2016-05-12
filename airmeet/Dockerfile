FROM golang:1.6.2

RUN mkdir -p /go/src/app
COPY . /go/src/app
WORKDIR /go/src/app

RUN go get github.com/labstack/echo/...
RUN go get github.com/k0kubun/pp
RUN go get github.com/satori/go.uuid
RUN go get gopkg.in/go-playground/validator.v8
RUN go get github.com/go-sql-driver/mysql
RUN go get github.com/jinzhu/gorm


EXPOSE 3000

CMD ["go","run","main.go","api.go","db.go","structs.go","errors.go"]
