consul agent -server \
  -retry-join $PRIVATE_IP_ADDRESS \
  -client 0.0.0.0
