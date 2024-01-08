import {ethers} from 'ethers';
import {registryStore, web3Store, loading, showLoginModal, portalVerified} from "./stores";
import contractAbi from './lib/contracts/IdentityRegistry.json';
import identityAbi from './lib/contracts/IdentityManager.json';

export async function connectWallet() {
    showLoginModal.set(false);
    console.log("in connectWallet");
    const provider = new ethers.BrowserProvider(window.ethereum, 'any');
    await provider.send('eth_requestAccounts', []);
    const signer = await provider.getSigner();
    const account = await signer.getAddress();
    const chainId = await provider.getNetwork().then(network => network.chainId);

    web3Store.set({
        provider,
        signer,
        account,
        chainId,
        identityAddress: null,
    });
    // initialize registry contract
    const contractAddress = import.meta.env.VITE_REGISTRY_CONTRACT_ADDRESS;
    const contract = new ethers.Contract(contractAddress, contractAbi, signer);
    console.log("Contract initialized", contract);
    loading.set(true);

    // get identity address from registry
    const identityAddress = await getIdentity(account, contract);
    const identityContract = new ethers.Contract(identityAddress, identityAbi, signer);

    registryStore.set({
        contract,
        contractAddress,
        contractAbi,
        identityContract,
    });
    loading.set(false);
    showLoginModal.set(true);
}

async function getIdentity(node: string, contract: ethers.Contract) {
    if (!contract) {
        console.error("Contract is not initialized.");
        return;
    }
    try {
        const result = await contract.manager(node);
        console.log(`getIdentity ${node}:`, result);
        web3Store.update($web3Store => {
            return {...$web3Store, identityAddress: result};
        });

        return result;
    } catch (error) {
        console.error("Error in getting manager:", error);
    }

}


// async function identityRegistered(node: string, contract: ethers.Contract) {
//     try {
//         // Assuming 'manager' is the function to check registration
//         // and it's available as a method on the ethers contract instance
//         let managerAddress = await contract.manager(node);
//         console.log(`Manager of ${node}:`, managerAddress);
//
//         if (managerAddress === ethers.ZeroAddress) {
//             console.log(`${node} is not registered`);
//             // deploy identity contract
//             managerAddress = await deployIdentityContract(identityAbi, '');
//
//         }
//
//         // Update the store or state with the manager address or registration status
//         // For example, you might want to do something like this:
//         web3Store.update(current => {
//             return {...current, identityAddress: managerAddress};
//         });
//
//         // Additional logic based on whether the manager address is what you expect
//         // for a registered or unregistered identity
//
//     } catch (error) {
//         console.error("Error in getting manager:", error);
//     }
// }

// async function deployIdentityContract(abi: any, byteCodePath: string): Promise<String> {
//     // get provider ans signer from web3store
//     const {provider, signer} = web3Store;
//     if (!bytecode) {
//         throw new Error("Bytecode is required for contract deployment.");
//     }
//     const registryAddress = import.meta.env.VITE_REGISTRY_CONTRACT_ADDRESS;
//     // Create a ContractFactory
//     const factory = new ethers.ContractFactory(abi, bytecode, signer);
//
//     try {
//         // Deploy the contract
//         const contract = await factory.deploy(registryAddress);  // Add constructor arguments if any
//
//         // Wait for the deployment transaction to be mined
//         await contract.deployed();
//
//         // Return the address of the deployed contract
//         return contract.address;
//     } catch (error) {
//         console.error("Contract deployment failed:", error);
//         throw error;  // Rethrow or handle error as needed
//     }
// }


// function readBytecode(filePath: string) {
//     try {
//         return readFileSync(filePath, 'utf8');
//     } catch (error) {
//         console.error("Error reading bytecode file:", error);
//         throw error;
//     }
// }

export function disconnectWallet() {
    web3Store.set({provider: null, signer: null, account: null, chainId: null, identityAddress: null});
    registryStore.set({contract: null, contractAddress: '', contractAbi: null, identityContract: null});
    loading.set(false);
    showLoginModal.set(false);
    portalVerified.set(false);
}
export function displayAddress(address:string) {
    return `${address.substring(0, 6)}...${address.substring(address.length - 4)}`;
}