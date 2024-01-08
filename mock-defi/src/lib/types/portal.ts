export type PrivateClaim = {
    value: string; // commitment depending on public or private
    statement: string; // phi
    ipfsCircuitMetadata: string; // circuit ipfs cid
    eventHash: string; // verification event hash
    timestamp: bigint;
    id: number;
    version: number;
};

export type PublicClaim = {
    value: string; // value
    statement: string; // phi
    timestamp: bigint;
    id: number;
    version: number;
};

export type Attestation = {
    signature: string; // hexadecimal signature signs(t, claim={commiment/public_value}, statement),
    attestor: string; // issuer hex string representation of address
    expires: bigint; // unix timestamp
    timestamp: bigint;
    claimId: number;
    id: number;
};

export type Revocation = {
    attestedTo: string; // hex string address
    attestationId: number;
    status: string;
    timestamp: bigint;
};
export type CircuitMetadata = {
    name: string;
    statement: string;
    contractAddress: string;
    fields: string[];
};

export type CircuitOnIpfs = {
    metadata: CircuitMetadata;
    ccs: Uint8Array;
    srs: Uint8Array;
    pk: Uint8Array;
    vk: Uint8Array;
    proof: Uint8Array;
    wt: Uint8Array;
};
