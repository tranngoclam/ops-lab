consul agent -server \
  -client 0.0.0.0 \
  -bootstrap-expect 3 \
  -data-dir /data \
  -ui \
  -dns-port 53 \
  -bind $PRIVATE_IP_ADDRESS \
  -advertise $PRIVATE_IP_ADDRESS \
  -node $NODE
