<script>
    import {web3Store,showLoginModal, portalVerified} from '../stores';

    import {onDestroy} from "svelte";
    import {disconnectWallet} from "../walletFunctions.ts";
    import WalletsModal from "./WalletsModal.svelte";
    import {displayAddress} from "../walletFunctions.ts";

    export let title = "";


   function showModal() {
        showLoginModal.set(true);
    }
    function closeWalletsModal() {
        showLoginModal.set(false);
    }

    let account = "";
    // Subscribe to the web3Store
    const unsubscribe = web3Store.subscribe(($web3Store) => {
        account = $web3Store?.account; // Update account whenever store changes
    });

    // Cleanup subscription when the component is unmounted
    onDestroy(() => {
        unsubscribe();
    });
</script>

<!--<div class="bg-white navbar bg-base-100 py-0.1">-->
<div class="header sticky top-0 navbar bg-white shadow-md">
    <div class="container mx-auto flex justify-between items-center">
        <div class="navbar-center">
            <a href="/" class="btn btn-ghost normal-case text-xl">{title}</a>
        </div>

    </div>
    <div class="navbar-end">

        {#if $portalVerified}
            <div class="flex items-center space-x-3">
                <div class="text-black font-medium">Wallet: {displayAddress(account)}</div>
                <button type="button" on:click={disconnectWallet}
                        class="text-gray-900 bg-white hover:bg-gray-100 border border-gray-200 focus:ring-4 focus:outline-none focus:ring-gray-100 font-medium rounded-lg text-sm px-5 py-2.5 text-center inline-flex items-center dark:focus:ring-gray-600 dark:bg-gray-800 dark:border-gray-700 dark:text-white dark:hover:bg-gray-700">
                    <svg aria-hidden="true" class="w-4 h-4 me-2" fill="none" stroke="currentColor" viewBox="0 0 24 24"
                         xmlns="http://www.w3.org/2000/svg">
                        <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2"
                              d="M13.828 10.172a4 4 0 00-5.656 0l-4 4a4 4 0 105.656 5.656l1.102-1.101m-.758-4.899a4 4 0 005.656 0l4-4a4 4 0 00-5.656-5.656l-1.1 1.1"></path>
                    </svg>
                    Disconnect
                </button>
            </div>
        {:else}
            <button type="button" on:click={showModal} data-modal-toggle="crypto-modal"
                    class="text-gray-900 bg-white hover:bg-gray-100 border border-gray-200 focus:ring-4 focus:outline-none focus:ring-gray-100 font-medium rounded-lg text-sm px-5 py-2.5 text-center inline-flex items-center dark:focus:ring-gray-600 dark:bg-gray-800 dark:border-gray-700 dark:text-white dark:hover:bg-gray-700">
                <svg aria-hidden="true" class="w-4 h-4 me-2" fill="none" stroke="currentColor" viewBox="0 0 24 24"
                     xmlns="http://www.w3.org/2000/svg">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2"
                          d="M13.828 10.172a4 4 0 00-5.656 0l-4 4a4 4 0 105.656 5.656l1.102-1.101m-.758-4.899a4 4 0 005.656 0l4-4a4 4 0 00-5.656-5.656l-1.1 1.1"></path>
                </svg>
                Connect wallet
            </button>
        {/if}

    </div>
</div>


<WalletsModal>

</WalletsModal>
