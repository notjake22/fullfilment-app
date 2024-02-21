<script lang="ts">
    import {GetServiceSettings, SetNewServiceSettings} from "../../wailsjs/go/main/App";
    import {database} from "../../wailsjs/go/models";
    import LoadingDots from "./layovers/LoadingDots.svelte";
    import Toasts from "./alerts/Toasts.svelte";
    import {addToast} from "./alerts/ts/store";

    let price: number
    let loading: boolean = false

    async function HandleSettingsEdit(): Promise<void> {
        loading = true;
        const newSettings: database.ServiceSettings = {
            itemPrice: price * 100
        }
        try {
            await SetNewServiceSettings(newSettings)
            addToast({
                message:"Settings Edited",
                type: "success",
                dismissible: true,
                timeout: 3000
            })
        } catch (e) {
            addToast({
                message:"Issue editing settings",
                type: "error",
                dismissible: true,
                timeout: 5000
            })
            console.log(e)
        }
        loading = false;
    }

    async function LoadSettings(): Promise<void> {
        loading = true
        try {
            await GetServiceSettings().then(r => {
                price = r.itemPrice/100;
            })
        } catch (e) {
            console.log(e)
        }

        loading = false
    }

    LoadSettings().then(r => {})
</script>

<div style="position: fixed; left: 50%">
    <Toasts />
</div>

<div class="login-container">
    <div class="login-box">
        <h2 style="color: black">Service Settings</h2>
        <label for="itemPrice">Price Per Item</label>
        <input id="itemPrice" bind:value={price} type="number" placeholder="Per Item Price" />
        <button on:click={async ()=>{await HandleSettingsEdit().then(r => {})}}>Submit</button>
    </div>
    {#if loading}
        <div style="position: fixed; top: 95%; left: 55%">
            <LoadingDots />
        </div>
    {/if}
</div>

<style>
    .login-container {
        display: flex;
        justify-content: center;
        align-items: center;
        height: 100vh;
        font-family: 'Roboto', sans-serif;
        /*background-color: #f2f2f2;*/
    }

    .login-box {
        width: 500px;
        padding: 50px;
        border-radius: 10px;
        background-color: #fff;
        box-shadow: 0 0 10px rgba(0,0,0,0.1);
        box-sizing: border-box;
    }

    .login-box h2 {
        margin-bottom: 30px;
        text-align: center;
        color: black;
    }

    .login-box input {
        width: 100%;
        padding: 10px;
        margin-bottom: 10px;
        border: 1px solid #ddd;
        border-radius: 5px;
        font-size: 16px;
        box-sizing: border-box;
    }

    .login-box label{
        color: black;
        text-align: left;
    }

    .login-box button {
        width: 60%;
        padding: 10px;
        border: none;
        border-radius: 5px;
        background-color: #007BFF;
        color: #fff;
        font-size: 16px;
        cursor: pointer;
        transition: background-color 0.2s ease;
    }

    .login-box button:hover {
        background-color: #0056b3;
    }
</style>