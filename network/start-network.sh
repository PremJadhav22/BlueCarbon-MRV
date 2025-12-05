#!/bin/bash

# Clean up previous setup
echo "Cleaning up previous network..."
docker-compose down
rm -rf crypto-config channel-artifacts

# Generate crypto material
echo "Generating crypto material..."
cryptogen generate --config=./crypto-config.yaml

# Generate genesis block
echo "Generating genesis block..."
configtxgen -profile BlueCarbonOrdererGenesis -channelID system-channel -outputBlock ./channel-artifacts/genesis.block

# Create channel configuration
echo "Creating channel configuration..."
configtxgen -profile BlueCarbonChannel -outputCreateChannelTx ./channel-artifacts/channel.tx -channelID bluecarbon-channel

# Start network
echo "Starting network..."
docker-compose up -d

# Wait for network to start
sleep 10

# Create channel
echo "Creating channel..."
docker exec cli peer channel create -o orderer.example.com:7050 -c bluecarbon-channel -f /opt/gopath/src/channel-artifacts/channel.tx --tls true --cafile /opt/gopath/src/crypto/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem

# Join peer to channel
echo "Joining peer to channel..."
docker exec cli peer channel join -b bluecarbon-channel.block

echo "Network started successfully!"
echo "Orderer: localhost:7050"
echo "Peer: localhost:7051"