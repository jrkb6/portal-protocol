# portal-server

The Portal Server is a backend service designed to verify Portal attestations & private-claims for third-party services. It interfaces with blockchain technologies to provide a secure and reliable verification process.

## Getting Started

To get the server running, you need to set up your environment with the necessary configurations.

### Prerequisites

Ensure you have the following installed:
- A local or remote IPFS node
- A local or remote Ethereum RPC node
### Setup
1. ganache v7.8.0 (@ganache/cli: 0.9.0, @ganache/core: 0.9.0)
2. go version go1.20.5
3. IPFS-kubo (for testing ipfs claims)

### Installation

Clone the repository to your local machine:

```bash
git clone https://github.com/jrkb6/portal-server.git
cd portal-server

```
Create a .env file in the root of the project with the following contents:
```bash
OWNER_PRIVATE_KEY=<Your_Private_Key_Here>
REGISTRY_ADDRESS=0xf3585FCD969502624c6A8ACf73721d1fce214E83
IPFS_URL=http://localhost:5001
RPC_URL=http://localhost:8545
CHAIN_ID=1337
```
Install the dependencies:
```bash
go install
```
Run the server:
```bash
go run main.go
```
Server should be running on port 3000.
### Serving the static files for the frontend
Copy the frontend build files to the `build/` directory.


## Endpoints Overview

### 1. Private Merkle Claim Live Verification

- **Endpoint:** `api/claim/:sender?`
- **Method:** `GET`
- **Description:** This endpoint is used for live verification of private Merkle claims. It allows users to submit a claim name along with the sender's wallet address for verification.

#### Parameters:
- `sender`: Sender wallet address. Retrieved from the path parameter. 
- `claimName`: Name of the private merkle claim. Retrieved from the query parameter.

### 2. Attestation Verification

- **Endpoint:** `/verify/:sender?`
- **Method:** `GET`
- **Description:** This endpoint is used for the verification of attestations. Users can submit an attestation name along with the sender's address and claim type for verification.

#### Parameters:
- `sender`: Sender wallet address. Retrieved from the path parameter.
- `attestationName`: Name of the attestation. Retrieved from the query parameter.
- `claimType`: Type of the claim associated with the attestation. Retrieved from the query parameter.

### 3. User Registration Signature Request

- **Endpoint:** `/sign/:sender?`
- **Method:** `GET`
- **Description:** This endpoint is for initiating a signature request for user registration from the registry owner's address. It is used in the context of identity verification.

#### Parameters:
- `sender`: Sender wallet address. Retrieved from the path parameter. .
- `identity`: Identity address deployed by the user. Retrieved from the query parameter.



