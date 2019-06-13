FROM golang

ENV GO111MODULE on

WORKDIR /go/release
RUN go get -u github.com/bilibili/kratos/tool/kratos && cd $GOPATH/src && \
    kratos new service-test --proto && cd service-test && \
    sed -ie 's/addr = "0.0.0.0:9000"/addr = "0.0.0.0:9747"/g' configs/grpc.toml && \
    sed -ie 's/addr = "0.0.0.0:8000"/addr = "0.0.0.0:8747"/g' configs/http.toml

CMD [ "go", "run", "cmd/main.go", "-conf", "configs/" ]