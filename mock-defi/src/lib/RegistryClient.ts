import type {Contract} from "ethers";
import type {CircuitMetadata} from "./types/portal";

export class RegistryClient {
    private readonly contract: Contract;

    constructor(contract: Contract) {
        this.contract = contract;
    }

    async getCircuit(circuitId: string): Promise<CircuitMetadata> {
        try {
            console.log("Searching for circuit:", circuitId);
            const result = await this.contract.getCircuit(circuitId);
            console.log("Circuit:", result);
            return {
                name: result.name,
                statement: result.statement,
                contractAddress: result.contractAddress,
                fields: result.fields
            };
        } catch (error) {
            console.error("Error getting circuit:", error);
        }
        return Promise.reject("Error getting circuit");
    }

    async manager(node: string): Promise<string> {
        try {
            const result = await this.contract.manager(node);
            console.log(`Manager of ${node}:`, result);
            return result;
        } catch (error) {
            console.error("Error in getting manager:", error);
        }
        return Promise.reject(`Error getting manager for ${node}`);
    }

}




