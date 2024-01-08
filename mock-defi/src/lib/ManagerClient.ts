import type {Attestation, PublicClaim, Revocation} from "./types/portal";
import type {Contract} from "ethers";

export class ManagerClient {
    private readonly contract: Contract;

    constructor(contract: Contract) {
        this.contract = contract;
    }

    async loginWithPortalCredential(name: string) {

        try {
            return this.getAttestation(name);

        } catch (error) {
            console.error('Error querying attestation:', error);
        }
    }

    async getClaim(key: string, type: string) {
        try {
            const result = await this.contract.getClaim(key, type);
            console.log(result);
            return result;
        } catch (error) {
            console.error('Error querying claim:', error);
        }
    }

    async getPublicClaim(key: string): Promise<PublicClaim> {
        try {
            const result = await this.contract.getPublicClaim(key);
            const mapped: PublicClaim = {
                value: result.value,
                statement: result.statement,
                timestamp: BigInt(result.timestamp),
                id: Number(result.id),
                version: Number(result.version)

            }
            console.log(result);
            return mapped;
        } catch (error) {
            console.error('Error querying public claim:', error);
        }
        return Promise.reject("Error querying public claim");
    }

    async getIPFSClaim(key: string): Promise<PublicClaim> {
        try {
            const result = await this.contract.getIPFSClaim(key);
            const mapped: PublicClaim = {
                value: result.value,
                statement: result.statement,
                timestamp: BigInt(result.timestamp),
                id: Number(result.id),
                version: Number(result.version)

            }
            console.log(result);
            return mapped;
        } catch (error) {
            console.error('Error querying public claim:', error);
        }
        return Promise.reject("Error querying public claim");
    }

    async getAttestation(key: string): Promise<Attestation> {
        try {
            const rawAttestation = await this.contract.attestations(key);

            const mapped: Attestation = {
                signature: rawAttestation.signature,
                attestor: rawAttestation.attestor,
                expires: BigInt(rawAttestation.expires),
                timestamp: BigInt(rawAttestation.timestamp),
                claimId: Number(rawAttestation.claimId),
                id: Number(rawAttestation.id)
            };
            console.log("Mapped:", mapped);
            return mapped;

        } catch (error) {
            console.error('Error querying attestation:', error);
        }
        return Promise.reject("Error querying attestation");
    }


    async getRevocation(key: string): Promise<Revocation> {
        try {
            const rawRevocation = await this.contract.revocations(key);

            const mapped: Revocation = {
                attestedTo: rawRevocation.attestedTo,
                attestationId: Number(rawRevocation.attestationId),
                status: rawRevocation.status,
                timestamp: BigInt(rawRevocation.timestamp)

            };
            console.log("Mapped:", mapped);
            return mapped;

        } catch (error) {
            console.error('Error querying revocation:', error);
        }
        return Promise.reject("Error querying revocation");
    }

}




