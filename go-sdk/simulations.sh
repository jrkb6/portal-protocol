#!/bin/bash

echo "Restart ganache if it is running..."
pkill -f "node $(which ganache)"
echo "Starting ganache..."
ganache -v -m "much repair shock carbon improve miss forget sock include bullet interest solution" -g 0x2CB417800 > ../sim_results/ganache_logs/ganache.log &
echo "Running simulations..."
echo "Simulation results will be saved in ../sim_results/ganache_logs"
echo "Running deployment simulation..."
go run main.go -sim deployment > ../sim_results/ganache_logs/deployment.log
go run main.go -sim merkleDeployment > ../sim_results/ganache_logs/merkleDeployment.log
echo "Running private claim simulation..."
go run main.go -sim claim privateMerkle > ../sim_results/ganache_logs/merkleClaim.log
echo "Running public on-chain claim simulation..."
go run main.go -sim claim publicOnChain> ../sim_results/ganache_logs/onchainClaim.log
echo "Running public ipfs claim simulation..."
go run main.go -sim claim publicIpfs > ../sim_results/ganache_logs/ipfsClaim.log
echo "Running live verification simulation..."
go run main.go -sim claim privateMerkleLive > ../sim_results/ganache_logs/merkleClaimLiveVerification.log
echo "Running public on-chain attestation simulation..."
go run main.go -sim attestation publicOnChain > ../sim_results/ganache_logs/onChainAttestation.log
echo "Running public ipfs attestation simulation..."
go run main.go -sim attestation publicIpfs > ../sim_results/ganache_logs/ipfsAttestation.log
echo "Running private merkle attestation simulation..."
go run main.go -sim attestation private > ../sim_results/ganache_logs/merkleAttestation.log
echo "Running private revocation simulation..."
go run main.go -sim revocation > ../sim_results/ganache_logs/merkleRevocation.log

echo "Killing ganache..."
#pkill -f "node $(which ganache)"
echo "Done!"