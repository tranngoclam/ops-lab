consul agent -server \
    -advertise $PRIVATE_IP_ADDRESS \
    -bind $PRIVATE_IP_ADDRESS \
    -bootstrap \
    -client 0.0.0.0 \
    -data-dir /data \
    -enable-local-script-checks \
    -node $NODE \
    -ui
