# Pull Docker image of official golang ethereum implementation
FROM ethereum/client-go:v1.13.15

# Copy sample account
COPY sample /root/.ethereum/keystore

COPY sample/password.txt /root

# Copy genesis.json file to /tmp
# Genesis file defined genesis block. Official documentation: https://geth.ethereum.org/docs/fundamentals/private-network
COPY genesis.json /tmp

# Initializing geth
RUN geth init /tmp/genesis.json \
    && rm -f ~/.ethereum/geth/nodekey

ENTRYPOINT ["geth"]
