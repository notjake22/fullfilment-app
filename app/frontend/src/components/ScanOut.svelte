<script lang="ts">
    import {
        AddNewCustomItem,
        GetUserTableInfo,
        ScanCancel,
        ScanInHandler,
        ScanOutCancel, ScanOutHandler
    } from "../../wailsjs/go/main/App";
    import {database} from "../../wailsjs/go/models";
    import LoadingDots from "./layovers/LoadingDots.svelte";
    import Toasts from "./alerts/Toasts.svelte";
    import {addToast} from "./alerts/ts/store";
    import AddCustomItem from "./layovers/AddCustomItem.svelte";

    type Users = {
        id: string
        name: string
    }

    type ItemData = {
        itemName: string
        itemImage: string
        itemCode: string
        totalItems: number
        id: string
    }

    let loading: boolean = false;
    let upc: string = "";
    let lastScannedUPC: string = "";
    let currentView: string = "Select Client";
    let users: Users[] = [
        {id: "0", name: "Select User"}
    ];
    let selectedUser: string = users[0].id;

    let scanInLoading: boolean = false;
    let showItemDetails: boolean = false;
    let currentItemDetails: ItemData;

    function startScanning(): void {
        if (selectedUser === "0") {
            return
        }
        currentView = "Scan Listener";
    }

    function onKeyDown(event: KeyboardEvent):void {
        if (currentView !== "Scan Listener") {
            return
        }
        if (event.key === "Enter") {
            HandleScannedInItem()
        } else {
            upc = upc + event.key
            console.log(upc)
        }

    }

    async function HandleScannedInItem(): Promise<void> {
        const itemPayload: database.ScanInRequest = {
            upc: upc,
            clientID: selectedUser
        }
        scanInLoading = true;
        showItemDetails = false;
        try {
            await ScanOutHandler(itemPayload).then(r => {
                console.log(r)
                currentItemDetails = {
                    itemName: r.itemName,
                    itemImage: r.imageUri,
                    itemCode: upc,
                    totalItems: 0,
                    id: r.ID
                }
                showItemDetails = true
                addToast({
                    message:"Item marked as scanned out",
                    type: "success",
                    dismissible: true,
                    timeout: 3000
                })
            })
        } catch (e) {
            if (e === "item not found") {
                addToast({
                    message:"Item not found",
                    type: "error",
                    dismissible: true,
                    timeout: 10000
                })
            } else {
                addToast({
                    message:"Error loading scanned item: " + e,
                    type: "error",
                    dismissible: true,
                    timeout: 10000
                })
                // todo add better error handling
                console.log(e)
            }
        }
        lastScannedUPC = upc;
        upc = ""
        scanInLoading = false
    }

    async function CancelScanIn(): Promise<void> {
        currentView = "Scan Listener";
        showItemDetails = false;
        loading = true
        try {
            await ScanOutCancel(currentItemDetails.id)
            addToast({
                message:"Scan in canceled",
                type: "success",
                dismissible: true,
                timeout: 2500
            })
        } catch (e){
            addToast({
                message:"Error canceling scan in: "+ e,
                type: "error",
                dismissible: true,
                timeout: 10000
            })
            console.log(e) //todo: do better handling
        }
        loading = false
    }


    async function LoadUsers(): Promise<void> {
        loading = true
        // reset array so prev loaded clients aren't there anymore
        let newUsers: Users[] = [
            {id: "0", name: "Select User"}
        ]
        selectedUser = users[0].id
        try {
            await GetUserTableInfo().then(r => {
                for (let i = 0; i < r.length; i++) {
                    newUsers.push({
                        id: r[i].id,
                        name: r[i].name
                    })
                    console.log(users)
                }
                users = newUsers
            })
        } catch (e) {
            addToast({
                message:"Error loading clients: "+ e,
                type: "error",
                dismissible: true,
                timeout: 10000
            })
            console.log(e)
        }
        loading = false
    }

    LoadUsers()
</script>

<Toasts />

{#if currentView === "Select Client"}
    <div class="title-header">
        <h1>Scan Out</h1>
    </div>
    <div class="select-container">
        <div class="select-box">
            <h2>Select Client to Start Scanning for</h2>
            <div class="user-selector">
                <select bind:value={selectedUser}>
                    {#each users as user (user.id)}
                        <option value={user.id}>{user.name}</option>
                    {/each}
                </select>
            </div>
            <button on:click={startScanning}>Start</button>
        </div>
    </div>
{/if}

{#if loading}
    <div style="position: fixed; top: 95%; left: 55%">
        <LoadingDots />
    </div>
{/if}

{#if currentView === "Scan Listener"}
    {#if showItemDetails}
        <div class="dashboard">
            <h2>{currentItemDetails.itemName}</h2>
            <img src="{currentItemDetails.itemImage}" alt="{currentItemDetails.itemName}" style="height: 200px; width: 200px"/>
            <p>Item Code: {currentItemDetails.itemCode}</p>
            <!--            <p>Total Items: {currentItemDetails.totalItems}</p>-->
            <button on:click={async ()=>{await CancelScanIn().then(r => {})}} class="cancel-button">Cancel</button>
        </div>
    {:else}
        <h1>You can now start scanning items</h1>
        <button class="back-button" on:click={() => (currentView = "Select Client")}>back</button>
        {#if scanInLoading}
            <div style="position: fixed; top: 50%; left: 50%">
                <LoadingDots />
            </div>
        {/if}
    {/if}
{/if}
<svelte:window on:keydown={onKeyDown} />
<style>
    .select-container {
        display: flex;
        justify-content: center;
        align-items: center;
        height: 100vh;
        font-family: 'Roboto', sans-serif;
        /*background-color: #f2f2f2;*/
    }

    .select-box {
        width: 500px;
        padding: 50px;
        border-radius: 10px;
        background-color: #fff;
        box-shadow: 0 0 10px rgba(0,0,0,0.1);
        box-sizing: border-box;
    }

    .select-box h2 {
        margin-bottom: 20px;
        text-align: center;
        color: black;
    }

    .user-selector {
        width: 100%;
        padding: 10px;
        margin-bottom: 10px;
        /*border: 1px solid #ddd;*/
        border-radius: 5px;
        font-size: 16px;
        box-sizing: border-box;
    }

    .select-box select {
        width: 100%;
        height: 100%;
        font-size: 1.2em;
        padding: 0.5em;
        /*padding: 10px;*/
        box-sizing: border-box;
        border: none;
        background: #ffffff;
        color: #000;
        box-shadow: 0 1px 3px rgba(0, 0, 0, 0.25);
        border-radius: 15px;
        appearance: none;
        outline: none;
    }

    .select-box button {
        width: 80%;
        padding: 10px;
        border: none;
        border-radius: 5px;
        background-color: #007BFF;
        color: #fff;
        font-size: 16px;
        cursor: pointer;
        transition: background-color 0.2s ease;
    }

    .select-box button:hover {
        background-color: #0056b3;
    }

    .title-header h1 {
        display: flex;
        justify-content: center;
        align-items: center;
        color: white;
        position: absolute;
        /*top: 50%;*/
        left: 42%;
    }

    h1 {
        display: flex;
        justify-content: center;
        align-items: center;
        color: white;
        position: absolute;
        /*top: 50%;*/
        left: 25%;
    }

    .back-button {
        position: absolute;
        top: 40px;
        right: 40px;
        padding: 8px 15px;
        border: none;
        border-radius: 5px;
        background-color: #007BFF;
        color: #fff;
        font-size: 1em;
        font-weight: bold;
        cursor: pointer;
    }

    .back-button:hover {
        background-color: #0056b3;
    }

    .dashboard {
        display: flex;
        flex-direction: column;
        align-items: center;
        justify-content: center;
        gap: 20px;
        width: 80%;
        max-width: 900px;
        margin: 50px auto;
        padding: 40px;
        box-shadow: 0px 2px 4px rgba(0, 0, 0, 0.1);
        border-radius: 10px;
        background-color: #f2f2f2;
    }

    .dashboard h2 {
        margin-bottom: 20px;
        text-align: center;
        color: black;
    }

    .dashboard p {
        color: black;
    }

    img {
        max-width: 100%;
        border-radius: 10px;
    }

    .cancel-button {
        padding: 10px 20px;
        background-color: #f44336;
        color: white;
        border: none;
        border-radius: 5px;
        cursor: pointer;
        font-size: 1em;
    }

    .cancel-button:hover {
        background-color: #e33e2b;
    }
</style>