<script lang="ts">
    import { createEventDispatcher } from 'svelte';
    import {blur} from "svelte/transition"

    type NewItem = {
        itemName: string
        imageUri: string
    }

    export let isOpen: boolean = false;
    let name: string = '';
    let imageUri: string = '';
    let newItem: NewItem
    const dispatch = createEventDispatcher();

    function submit(): void {
        newItem = {
            itemName: name,
            imageUri: imageUri
        }
        dispatch('submit', newItem)
        isOpen = false
    }
</script>

{#if isOpen}
    <article transition:blur>
        <div class="overlay" on:click={() => (isOpen = false)}>
            <div class="popup" on:click={(event) => event.stopPropagation()}>
                <h2>Enter Custom Item Details</h2>
                <input bind:value={name} type="text" placeholder="Name" />
                <input bind:value={imageUri} type="text" placeholder="Image URL (optional)" />
                <button on:click={submit}>Submit</button>
            </div>
        </div>
    </article>
{/if}

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
</style>