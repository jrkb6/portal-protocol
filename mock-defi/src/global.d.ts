import {MetaMaskInpageProvider} from "@metamask/providers"

declare global {
    interface Window {
        ethereum?: MetaMaskInpageProvider
    }

    interface Web3Props {

        provider: ethers.providers.JsonRpcProvider | null;
        signer: ethers.providers.JsonRpcSigner | null;
        account: string | null;
        chainId: bigint | null;

        identityAddress: string | null;
    }


    interface RegistryProps {
        contract: ethers.Contract | null;
        contractAddress: string;
        contractAbi: any;
        identityContract: ethers.Contract | null;

    }

}