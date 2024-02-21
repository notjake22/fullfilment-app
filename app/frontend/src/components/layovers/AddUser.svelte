<script lang="ts">
    import { createEventDispatcher } from 'svelte';
    import {database} from "../../../wailsjs/go/models";
    import {blur} from "svelte/transition"

    let isOpen: boolean = false;
    let name: string = '';
    let phone: string = '';
    let balance: number = 0;
    let dispatchUser: database.User

    const dispatch = createEventDispatcher();

    function submit() {
        dispatchUser = { id: '', name: name, phone: phone, balance: balance, currentlyStoredItems: 0, totalOrderHistory: 0}
        dispatch('submit', dispatchUser);
        isOpen = false;
    }
</script>

{#if isOpen}
    <article transition:blur>
        <div class="overlay" on:click={() => (isOpen = false)}>
            <div class="popup" on:click={(event) => event.stopPropagation()}>
                <h2>New User Details</h2>
                <input bind:value={name} type="text" placeholder="Name" />
                <input bind:value={phone} type="email" placeholder="Phone" />
                <input bind:value={balance} type="number" placeholder="Balance" />
                <button on:click={submit}>Submit</button>
            </div>
        </div>
    </article>
{/if}

<button class="add-user-button" on:click={() => (isOpen = true)}>Add Client</button>

<style>
    h2 {
        color: #333333;
    }

    .overlay {
        display: flex;
        justify-content: center;
        align-items: center;
        position: fixed;
        top: 0;
        bottom: 0;
        left: 0;
        right: 0;
        background-color: rgba(0, 0, 0, 0.5);
    }
    .popup {
        background-color: white;
        padding: 20px;
        border-radius: 15px;
        box-shadow: 0 2px 10px rgba(0, 0, 0, 0.1);
        max-width: 350px;
        width: 100%;
    }
    .popup input {
        display: block;
        width: 100%;
        margin-bottom: 10px;
        padding: 10px;
        border-radius: 5px;
        border: 1px solid #ccc;
        box-sizing: border-box;
    }
    .popup button {
        background-color: #007AFF;
        color: white;
        border: none;
        border-radius: 5px;
        padding: 10px 20px;
        cursor: pointer;
    }

    .add-user-button {
        position: absolute;
        top: 20px;
        right: 40px;
        padding: 10px 20px;
        border: none;
        border-radius: 5px;
        background-color: #007BFF;
        color: #fff;
        font-size: 1em;
        font-weight: bold;
        cursor: pointer;
    }

    .add-user-button:hover {
        background-color: #0056b3;
    }
</style>