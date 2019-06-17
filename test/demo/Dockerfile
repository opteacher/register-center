FROM golang

ENV GO111MODULE on

ENV DISCOVERY_ADDR discovery:7171
ENV REGISTER_CENTER 127.0.0.1:8089

WORKDIR /go/cache
ADD go.mod .
ADD go.sum .
RUN go mod download

WORKDIR /go/release
ADD . .
RUN GOOS=linux CGO_ENABLED=0 go build -ldflags="-s -w" -installsuffix cgo -o app cmd/main.go && ls -l app

CMD ["./app", "-conf", "configs/"]