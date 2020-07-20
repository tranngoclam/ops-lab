sed -i -- "s/__IP_ADDRESS__/$PRIVATE_IP_ADDRESS/g" /opt/base.hcl
nomad agent -config /opt/base.hcl -config /opt/server.hcl
