## Identity registry and resolver contracts

### Prerequisites

1. ganache v7.8.0 (@ganache/cli: 0.9.0, @ganache/core: 0.9.0)
2. go version go1.20.5 darwin/arm64
3. solc Version: 0.8.20+commit.a1b79de6.Darwin.appleclang
4. abigen version 1.10.16-unstable
5. IPFS-kubo (for testing ipfs claims)
    - `docker pull ipfs/kubo`
    - Configure the ipfs node:
    - Set local volume mappings
        - `export ipfs_staging=</absolute/path/to/>`
        - `export ipfs_data=</absolute/path/to/>`
    - Run the docker container by mapping volumes and ports
        - 4001 -> P2P TCP/QUIC transports,
        - 5001 -> RPC API,
        - 8080->Gateway

    - `docker run -d --name ipfs_host -v $ipfs_staging:/export -v $ipfs_data:/data/ipfs -p 4001:4001 -p 4001:4001/udp -p 127.0.0.1:8080:8080 -p 127.0.0.1:5001:5001 ipfs/kubo:latest`
    - Check the status of the ipfs node by running the following command: ```docker logs -f ipfs_host```
   ```
   # Expected logs when the system is started.
   RPC API server listening on /ip4/0.0.0.0/tcp/5001
   WebUI: http://0.0.0.0:5001/webui
   Gateway server listening on /ip4/0.0.0.0/tcp/8080
   Daemon is ready
   ```
6. Connect to peers:```docker exec ipfs_host ipfs swarm peers```
7. Stop container after the test: ```docker stop ipfs_host```
8. Update `ipfsConn` constant in main.go to your ipfs node address.

### Go-sdk setup

1. Run ganache network
   first `ganache -v -m "much repair shock carbon improve miss forget sock include bullet interest solution"`
2. Run go bindings (if not exist under go-sdk/contracts repository): `go run main.go -bindings true`
3. Overwrite go bindings under `go-sdk/contracts` in
   packages `go-sdk/contracts/IdentityInterface/` `go-sdk/contracts/IdentityManager/` `go-sdk/contracts/IdentityRegistry/`
4. Deploy contracts with: `go run main.go -deploy true`

command will output something like:

```
Registry contract deployed at: 0xf3585FCD969502624c6A8ACf73721d1fce214E83
Manager contract deployed at: 0x2e144aF3Bde9B518C7C65FBE170c07c888f1fF1a
```

5. Update `managerAddr` and `registryAddr` in main.go constants with the deployment command output addresses.
6. Register identity: ```go run main.go -register true```
7. Set public claim for user: ```go run main.go -claim true```
8. Test individual ipfs claim: ```go run main.go -ipfs true```
9. Test on-chain ipfs claims: ```go run main.go -ipfs-on-chain true```
10. Test zero knowledge proove & verify from ipfs serialized circuits: ```go run main.go -circuits true```

### How to run simulations:

1. Run ganache network
   first `ganache -v -m "much repair shock carbon improve miss forget sock include bullet interest solution"`
2. `go run main.go -sim deployment` -> this will deploy registry, and identity contracts (attestor and claimOwner) and
   register those identities.
3. If you’re using the same ganache network, just copy `.env.example` to `.env` file. Otherwise, change the contract
   addresses in the .env file with the deployment simulation outputs.
4. `go run main.go -sim mimcDeployment` this will deploy mimccombined circuit and store metadata into registry contract.
5. `go run main.go -sim merkleDeployment` this will deploy Merklecombined circuit and store metadata into registry
   contract.
6. `go run main.go -sim mimcClaim` this will generate an age private claim for user with the previously deployed mimc
   circuit.
7. `go run main.go -sim claim privateMerkle` this will generate an age private claim for user with the previously
   deployed merkle circuit.
8. `go run main.go -sim claim publicOnChain` this will generate an age public onchain claim for user.
9. `go run main.go -sim claim publicIpfs` this will generate an age public ipfs claim for user.
10. `go run main.go -sim attestation privateMerkleLive` simulate merkleclaim live verification.
11. `go run main.go -sim attestation publicOnChain` attestor now attest previously generated age public claim of the
   user. User then stores it in identity contract attestation.
12. `go run main.go -sim attestation publicIpfs` attestor now attest previously generated age public claim of the user.
   User then stores it in identity contract attestation.
13. `go run main.go -sim attestation private` attestor now attest previously generated private age claim of the user.
   User then stores it in identity contract attestation.
14. `go run main.go -sim revocation` remove the last attestation. (add revocation map of the attestor identity contract)

### Running simulations.sh

Running all simulations **locally** together and saving the output to a file. Make sure .env file is updated and ipfs
node is running.

1. `chmod +x simulations.sh`
2. `./simulations.sh`
3. Make sure you kill the ganache process after the simulation is done. Otherwise, you will get an error when you run
   the simulation again.
   `pkill -f "node $(which ganache)" `