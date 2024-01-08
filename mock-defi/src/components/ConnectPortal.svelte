<script>

    import {ManagerClient} from "../lib/ManagerClient.ts";
    import {registryStore, showLoginModal, web3Store, loading, portalVerified} from "../stores.ts";
    import {onMount} from "svelte";
    import {ServiceClient} from "$lib/ServiceClient.ts";
    import {displayAddress} from "../walletFunctions.ts";

    let managerClient;
    let credential = {};
    let verificationFailed = false;
    let failMessage = '';
    let options = [
        {name: "agePubOnChain", type: "public"},
        {name: "agePubIpfs", type: "ipfs"},
        {name: "agePriv", type: "private"},
        {name: "agePrivLive", type: "privateLive"},
        {name: "LocationMunich", type: "public"}
    ];
    onMount(async () => {

        const unsubscribe = registryStore.subscribe((registryProps) => {
            if (registryProps) {
                managerClient = new ManagerClient(registryProps.identityContract);
            }
        });

        return () => {
            unsubscribe();
        };
    });

    function closeModal() {
        showLoginModal.set(false);
    }

    async function loginWithPortalCredential(credential) {
        verificationFailed = false;
        if (!managerClient) {
            console.error('Contract not initialized');
            return;
        }
        // sleep 2 seconds to show fetching animation
        await new Promise(r => {
            loading.set(true);
            setTimeout(r, 3000)
        });
        loading.set(false);

        try {
            const attestation = await managerClient.loginWithPortalCredential(credential.name);
            await verifyAttestation(credential.name, credential.type);
        } catch (error) {
            console.error('Error querying attestation:', error);
            verificationFailed = true;
        }
    }

    async function verifyAttestation(credentialName, claimType) {
        let verified;
        const serviceClient = new ServiceClient();
        if (claimType === 'privateLive') {
            verified = await serviceClient.claimRequest($web3Store.account, credentialName);
            if (verified.verified) {
                portalVerified.set(true);
                closeModal();
            } else {
                console.log("Verification failed", verified.reason);
                failMessage = verified.reason;
                verificationFailed = true;
            }
        } else {
            verified = await serviceClient.verifyAttestation($web3Store.account, credentialName, claimType);
            if (verified.verified) {
                portalVerified.set(true);
                closeModal();
            } else {
                console.log("Verification failed", verified.reason);
                failMessage = verified.reason;
                verificationFailed = true;
            }

        }


    }


</script>

<div class="grid items-center">
    {#if verificationFailed}

        <div class="bg-red-100 border border-red-400 text-red-700 px-2 py-2 rounded relative" role="alert">
            <strong class="font-bold">Verification failed. {failMessage}</strong>
            <span class="block sm:inline"> Please try again.</span>
        </div>
    {/if}
    <div class="space-x-2 bg-green-100 p-2 rounded-lg">
        <p class="text-sm font-normal text-gray-700 dark:text-gray-300">Connected Wallet: <span
                class="font-semibold text-gray-900 dark:text-white">{displayAddress($web3Store.account)}</span></p>
    </div>

</div>

<div class="py-2">
    <div class="bg-gray-600 bg-opacity-50 overflow-y-auto h-full w-full" id="my-modal">
        <div class="top-20 mx-auto p-5 shadow-lg bg-white">
            <div class="mt-3 text-center">
                <div class="mx-auto flex items-center justify-center h-12 w-12 rounded-full bg-purple-100">
                    <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="1.5"
                         stroke="currentColor" class="w-6 h-6">
                        <path stroke-linecap="round" stroke-linejoin="round"
                              d="M2.25 18.75a60.07 60.07 0 0 1 15.797 2.101c.727.198 1.453-.342 1.453-1.096V18.75M3.75 4.5v.75A.75.75 0 0 1 3 6h-.75m0 0v-.375c0-.621.504-1.125 1.125-1.125H20.25M2.25 6v9m18-10.5v.75c0 .414.336.75.75.75h.75m-1.5-1.5h.375c.621 0 1.125.504 1.125 1.125v9.75c0 .621-.504 1.125-1.125 1.125h-.375m1.5-1.5H21a.75.75 0 0 0-.75.75v.75m0 0H3.75m0 0h-.375a1.125 1.125 0 0 1-1.125-1.125V15m1.5 1.5v-.75A.75.75 0 0 0 3 15h-.75M15 10.5a3 3 0 1 1-6 0 3 3 0 0 1 6 0Zm3 0h.008v.008H18V10.5Zm-12 0h.008v.008H6V10.5Z"/>
                    </svg>

                </div>
                <h3 class="text-lg leading-6 font-medium text-gray-900">TUM-DEFI demo requests the following
                    information</h3>
                {#if !$loading}
                    <div class="mt-2 px-7 py-3">
                        <p class="text-sm text-gray-500">Verify your age via Portal before login </p>
                        <form class="mt-4 max-w-sm mx-auto">
                            <label for="credentials"
                                   class="block mb-2 text-sm font-medium text-gray-900 dark:text-white">Select your
                                credential</label>

                            <select id="credentials" bind:value={credential}
                                    class="bg-gray-50 border border-gray-300 text-gray-900 text-sm rounded-lg focus:ring-blue-500 focus:border-blue-500 block w-full p-2.5 dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-blue-500 dark:focus:border-blue-500">

                                {#each options as option}
                                    <option value={option}>{option.name}</option>
                                {/each}
                            </select>
                        </form>


                        <div class="mt-4">

                            <button on:click={() => loginWithPortalCredential(credential)}
                                    class="px-4 py-2 bg-green-500 text-white text-base font-medium rounded-md w-full shadow-sm hover:bg-green-700 focus:outline-none focus:ring-2 focus:ring-green-600">
                                Approve
                            </button>
                            <button on:click={closeModal}
                                    class="mt-2 px-4 py-2 bg-red-500 text-white text-base font-medium rounded-md w-full shadow-sm hover:bg-red-700 focus:outline-none focus:ring-2 focus:ring-red-600">
                                Cancel
                            </button>

                        </div>
                    </div>
                {:else}
                    <div class="py-2 text-center">
                        <div aria-label="Verifying Credential" role="status">
                            <svg aria-hidden="true"
                                 class="inline w-8 h-8 text-gray-200 animate-spin dark:text-gray-600 fill-blue-600"
                                 viewBox="0 0 100 101" fill="none" xmlns="http://www.w3.org/2000/svg">
                                <path d="M100 50.5908C100 78.2051 77.6142 100.591 50 100.591C22.3858 100.591 0 78.2051 0 50.5908C0 22.9766 22.3858 0.59082 50 0.59082C77.6142 0.59082 100 22.9766 100 50.5908ZM9.08144 50.5908C9.08144 73.1895 27.4013 91.5094 50 91.5094C72.5987 91.5094 90.9186 73.1895 90.9186 50.5908C90.9186 27.9921 72.5987 9.67226 50 9.67226C27.4013 9.67226 9.08144 27.9921 9.08144 50.5908Z"
                                      fill="currentColor"/>
                                <path d="M93.9676 39.0409C96.393 38.4038 97.8624 35.9116 97.0079 33.5539C95.2932 28.8227 92.871 24.3692 89.8167 20.348C85.8452 15.1192 80.8826 10.7238 75.2124 7.41289C69.5422 4.10194 63.2754 1.94025 56.7698 1.05124C51.7666 0.367541 46.6976 0.446843 41.7345 1.27873C39.2613 1.69328 37.813 4.19778 38.4501 6.62326C39.0873 9.04874 41.5694 10.4717 44.0505 10.1071C47.8511 9.54855 51.7191 9.52689 55.5402 10.0491C60.8642 10.7766 65.9928 12.5457 70.6331 15.2552C75.2735 17.9648 79.3347 21.5619 82.5849 25.841C84.9175 28.9121 86.7997 32.2913 88.1811 35.8758C89.083 38.2158 91.5421 39.6781 93.9676 39.0409Z"
                                      fill="currentFill"/>
                            </svg>
                            <span class="text-xl font-medium text-gray-500">Verifying Credential On Portal...</span>
                        </div>
                    </div>
                {/if}
            </div>
        </div>
    </div>
</div>

