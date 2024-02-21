<script lang="ts">
    import {createEventDispatcher} from "svelte";
    import {database} from "../../../wailsjs/go/models";
    import { blur } from "svelte/transition"

    export let isOpen: boolean = false;
    export let itemInfo: database.ItemDataModel

    const dispatch = createEventDispatcher()
    function submit(event: boolean) {
        dispatch('submit', event)
        isOpen = false
    }
</script>

{#if isOpen}
    <article transition:blur>
        <div class="overlay" on:click={() => (isOpen = false)}>
            <div class="popup" on:click={(event) => event.stopPropagation()}>
                <h2>Are you sure you want to delete {itemInfo.itemName}</h2>
                <button class="edit-button" on:click={() => {submit(true)}}>Yes</button>
                <button class="edit-button" on:click={() => {submit(false)}}>No</button>
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
    .popup label{
        color: black;
        text-align: left;
    }
    .popup button {
        background-color: #007AFF;
        color: white;
        border: none;
        border-radius: 5px;
        padding: 10px 20px;
        cursor: pointer;
    }

    button.edit-button {
        background-color: #007BFF;
        color: white;
        padding: 0.5em 1em;
        border: none;
        border-radius: 4px;
        cursor: pointer;
        transition: background-color 0.3s ease;
    }

    button.edit-button:hover {
        background-color: #0056b3;
    }
</style>