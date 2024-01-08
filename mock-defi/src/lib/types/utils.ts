import type {Attestation, PublicClaim, Revocation} from "./portal";
//import { createVerify } from "crypto";

export function attestationExpired(attestation: Attestation): boolean {
    const now = Math.floor(Date.now() / 1000);
    return  BigInt(now) > attestation.expires;

}

// export function verifySignature(signature:string, hash:string, algorithm:String="SHA256", signer:String, ): boolean {
//     const verifier = createVerify("SHA256");
//     verifier.update(Buffer.from(hash, 'hex'));
//     const signerBuffer = Buffer.from(signer, 'hex');
//     return verifier.verify(signerBuffer, Buffer.from(signature, 'hex'));
//
// }
