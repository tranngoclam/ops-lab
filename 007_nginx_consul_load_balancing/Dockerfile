FROM centos:7

RUN yum -y update && yum clean all && yum install -y \
    bash \
    curl \
    dnsutils \
    iputils-ping \
    net-tools \
    jq \
    supervisor \
    tar \
    unzip \
    zip

ADD https://github.com/mattn/goreman/releases/download/v0.3.4/goreman_linux_amd64.zip /tmp/goreman.zip
RUN cd /bin && unzip /tmp/goreman.zip && chmod +x /bin/goreman && rm /tmp/goreman.zip

# Install Consul
ADD https://releases.hashicorp.com/consul/1.8.0/consul_1.8.0_linux_amd64.zip /tmp/consul.zip
RUN cd /bin && unzip /tmp/consul.zip && chmod +x /bin/consul && rm /tmp/consul.zip

# Install NGINX
ADD http://nginx.org/download/nginx-1.18.0.tar.gz /tmp/nginx.tar.gz
RUN mkdir /bin/nginx && cd $_ && tar -xf /tmp/nginx.tar.gz && rm /tmp/nginx.tar.gz

# Install Go
ADD https://golang.org/dl/go1.14.6.linux-amd64.tar.gz /tmp/go.tar.gz
RUN tar -C /usr/local -xzf /tmp/go.tar.gz && rm /tmp/go.tar.gz

# Build Server
ENV GOPATH=/go
ENV PATH=/go/bin:/usr/local/go/bin:/usr/local/sbin:/usr/local/bin:/usr/sbin:/usr/bin:/sbin:/bin
WORKDIR /go
COPY server.go $GOPATH
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o /bin/server server.go

COPY consul.sh server.sh nginx.sh /opt/
ADD Procfile /root/Procfile

EXPOSE 5000 8300 8301 8302 8400 8500

ENTRYPOINT [ "goreman" ]
CMD [ "-f", "/root/Procfile", "start" ]
