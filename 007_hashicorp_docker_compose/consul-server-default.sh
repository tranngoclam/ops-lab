consul agent -server \
  -client 0.0.0.0 \
  -data-dir /data \
  -dns-port 53 \
  -node $NODE \
  -retry-join $CONSUL_SERVER_BOOTSTRAP \
  -ui
