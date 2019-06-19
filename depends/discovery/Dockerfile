FROM alpine

WORKDIR /discovery
RUN apk --update add wget && \
    wget https://github.com/bilibili/discovery/releases/download/v1.0.1/discovery_1.0.1_Linux_x86_64.tar.gz && \
    tar -zxvf discovery_1.0.1_Linux_x86_64.tar.gz && \
    rm discovery_1.0.1_Linux_x86_64.tar.gz && \
    sed -ie 's/addr = "127.0.0.1:7171"/addr = "0.0.0.0:7171"/g' configs/discovery-example.toml

EXPOSE 7171

CMD ["./discovery", "-conf", "configs/discovery-example.toml", "-alsologtostderr"]