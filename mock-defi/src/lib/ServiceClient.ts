import type {Attestation} from "./types/portal";
import axios from 'axios';

export interface VerifyResponse {
    verified: boolean;
    attestation: Attestation;
    reason: string;
}

export interface ClaimResponse {
    verified: boolean;
    claim: string;
}

export interface ApiError {
    message: string;
    reason?: string;
}

export class ServiceClient {
    private apiUrl: string;

    constructor() {
        this.apiUrl = import.meta.env.VITE_SERVICE_URL || "http://localhost:3000";
    }

    private async request<T>(path: string, params: any): Promise<T | ApiError> {
        try {
            const response = await axios.get<T>(`${this.apiUrl}${path}`, {params});
            return response.data;
        } catch (error) {
            return this.handleError(error);
        }
    }

    private handleError(error: unknown): ApiError {
        if (axios.isAxiosError(error)) {
            console.error("Error in API request:", error.message);
            const status = error.response?.status;
            const reason = error.response?.data?.reason as string;
            return {
                message: `Error in API request: ${error.message}`,
                reason: status === 400 ? reason : undefined
            };
        }
        return {message: 'An unknown error occurred'};
    }

    public verifyAttestation(sender: string, attestationName: string, claimType: string): Promise<VerifyResponse | ApiError> {
        return this.request<VerifyResponse>('/verify/' + sender, {
            attestationName: attestationName,
            type: claimType,
        });
    }

    public claimRequest(sender: string, claimName: string): Promise<ClaimResponse | ApiError> {
        return this.request<ClaimResponse>('/claim/' + sender, {
            claimName: claimName,
        });
    }
}

