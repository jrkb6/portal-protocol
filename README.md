# Portal protocol, a decentralized identity system

This repository is part of my Master's thesis research at the Technical University of Munich (TUM). It brings together all related modules, including identity registry, resolver contracts, and zero-knowledge circuits for private claims, as well as various simulations and demonstrations related to decentralized identity verification and finance.

## Submodules:

### 1. go-sdk
The `go-sdk` module includes Identity registry, resolver contracts, zero-knowledge circuits for private claims, and simulations including deployments, registration, claims, attestations, revocations, zk verifications. Part of the SDK also serves as a verifier plugin for third-party services.
- **Detailed Documentation**: [go-sdk README](./go-sdk/README.md)

### 2. portal-service
The `portal-service` is a backend service designed to verify Portal attestations & private-claims for third-party services. It interfaces with blockchain technologies to provide a secure and reliable verification process.
- **Detailed Documentation**: [portal-service README](./portal-service/README.md)

### 3. mock-defi
The `mock-defi` application serves as a demonstration of decentralized finance (De-Fi) application, focusing on showcasing the login feature using the Portal system.
- **Detailed Documentation**: [mock-defi README](./mock-defi/README.md)

For more detailed information on each module, please refer to the provided links to the respective README files.

