FROM centos:7 AS base

# Install Utilities
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

# Install Nomad
ADD https://releases.hashicorp.com/nomad/0.12.0/nomad_0.12.0_linux_amd64.zip /tmp/nomad.zip
RUN cd /bin && unzip /tmp/nomad.zip && chmod +x /bin/nomad && rm /tmp/nomad.zip

# Bootstrap
FROM base AS bootstrap

COPY consul-server-bootstrap.sh nomad-server-default.sh nomad-client-default.sh server.hcl client.hcl /opt/
ADD Procfile-bootstrap /root/Procfile

ENTRYPOINT [ "goreman" ]
CMD [ "-f", "/root/Procfile", "start" ]

# Default
FROM base AS default

COPY consul-server-default.sh nomad-server-default.sh nomad-client-default.sh server.hcl client.hcl /opt/
ADD Procfile /root/Procfile

ENTRYPOINT [ "goreman" ]
CMD [ "-f", "/root/Procfile", "start" ]
