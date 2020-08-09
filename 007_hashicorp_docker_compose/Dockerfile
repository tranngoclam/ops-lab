FROM centos:7 AS base

ENV CONSUL_VERSION=1.8.2
ENV NOMAD_VERSION=0.12.1
ENV DOCKER_VERSION=19.03.12

# Install Utilities
RUN yum update -y && yum install -y \
    bash \
    curl \
    dnsutils \
    iproute \
    iputils-ping \
    net-tools \
    jq \
    tar \
    unzip \
    yum-utils \
    zip

# Install Docker
RUN yum-config-manager --add-repo https://download.docker.com/linux/centos/docker-ce.repo
RUN yum install -y docker-ce docker-ce-cli containerd.io

# Install Goreman
ADD https://github.com/mattn/goreman/releases/download/v0.3.4/goreman_linux_amd64.zip /tmp/goreman.zip
RUN cd /bin && unzip /tmp/goreman.zip && chmod +x /bin/goreman && rm /tmp/goreman.zip

# Install Consul
ADD https://releases.hashicorp.com/consul/${CONSUL_VERSION}/consul_${CONSUL_VERSION}_linux_amd64.zip /tmp/consul.zip
RUN cd /bin && unzip /tmp/consul.zip && chmod +x /bin/consul && rm /tmp/consul.zip

# Install Nomad
ADD https://releases.hashicorp.com/nomad/${NOMAD_VERSION}/nomad_${NOMAD_VERSION}_linux_amd64.zip /tmp/nomad.zip
RUN cd /bin && unzip /tmp/nomad.zip && chmod +x /bin/nomad && rm /tmp/nomad.zip

COPY newenv.json nginx.nomad docker.sh nomad-server.sh nomad-client.sh base.hcl server.hcl client.hcl /opt/

# Bootstrap
FROM base AS bootstrap

COPY consul-server-bootstrap.sh /opt/
ADD Procfile-bootstrap /root/Procfile

ENTRYPOINT [ "goreman" ]
CMD [ "-f", "/root/Procfile", "start" ]

# Default
FROM base AS default

COPY consul-server-default.sh /opt/
ADD Procfile /root/Procfile

ENTRYPOINT [ "goreman" ]
CMD [ "-f", "/root/Procfile", "start" ]
