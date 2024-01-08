import {writable} from 'svelte/store';
import type {Writable} from 'svelte/store';
import contractAbi from './lib/contracts/IdentityRegistry.json';

export const web3Store: Writable<Web3Props> = writable({
    provider: null,
    signer: null,
    account: null,
    chainId: null,
    identityAddress: null,
});
export const registryStore: Writable<RegistryProps> = writable({
    contract: null,
    contractAddress: import.meta.env.VITE_REGISTRY_CONTRACT_ADDRESS,
    contractAbi: contractAbi,
    identityContract: null,
});
export const loading= writable(false);
export const portalVerified= writable(false);
export const showLoginModal= writable(false);