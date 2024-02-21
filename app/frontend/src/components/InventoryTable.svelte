<script lang="ts">
    import {DeleteInventoryItem, GetInventoryItems, GetUserTableInfo} from "../../wailsjs/go/main/App";
    import LoadingDots from "./layovers/LoadingDots.svelte";
    import Toasts from "./alerts/Toasts.svelte";
    import {addToast} from "./alerts/ts/store";
    import DeleteItem from "./layovers/DeleteItem.svelte";

    type InventoryItem = {
        id: number
        itemName: string
        upc: string
        associatedUser: string
        associatedUserID: string
    }

    type UserSelection = {
        id: string
        name: string
    }

    let loadingTable: boolean = false;
    let inventoryItems: InventoryItem[] = []
    let allInventoryItems: InventoryItem[] = []

    let selectedUser: string = "0";
    let users: UserSelection[] = []

    let popUpIsOpen: boolean = false;
    let selectedItem: InventoryItem

    const onUserChange = () => {
        let newTable: InventoryItem[] = []
        for (let i: number = 0; i < allInventoryItems.length; i++) {
            if (selectedUser === allInventoryItems[i].associatedUserID) {
                newTable.push(allInventoryItems[i])
            }
        }
        inventoryItems = newTable
    }

    function openEditPopup(item: InventoryItem) {
        selectedItem = item;
        popUpIsOpen = true;
    }

    async function handleDeleteSubmit(event): Promise<void> {
        if (!event.detail) {
            return
        }
        try {
            await DeleteInventoryItem(selectedItem.id.toString())
            addToast({
                message:"Item deleted",
                type: "success",
                dismissible: true,
                timeout: 3000
            })
            await LoadTable()
        } catch (e) {
            addToast({
                message:"Error deleting item: "+ e,
                type: "error",
                dismissible: true,
                timeout: 10000
            })
            console.log(e)
            return
        }
    }

    async function LoadTable(): Promise<void> {
        loadingTable = true
        let newItems: InventoryItem[] = []
        try {
            await GetInventoryItems().then(r => {
                for (let i = 0; i < r.length; i++) {
                    if (r[i].currentlyStored) {
                        const newItem: InventoryItem = {
                            id: parseInt(r[i].id),
                            itemName: r[i].itemName,
                            upc: r[i].upc,
                            associatedUser: r[i].ownerName,
                            associatedUserID: r[i].ownerID,
                        }
                        newItems.push(newItem)
                    }
                }
                // sort inventory by id, oldest to newest
                newItems = newItems.sort((a, b) => {
                    if (a.id < b.id) {
                        return -1
                    }
                })
                allInventoryItems = newItems
                inventoryItems = newItems
            })
            await loadUsers();
        } catch (e) {
            addToast({
                message:"Error loading inventory table: "+ e.toString(),
                type: "error",
                dismissible: true,
                timeout: 10000
            })
            console.log(e)
        }
        loadingTable = false
    }

    async function loadUsers(): Promise<void> {
        let newUsers: UserSelection[] = [
            {id: "0", name: "Filter by Client"}
        ];
        try {
            await GetUserTableInfo().then(r => {
                for (let i: number = 0; i < r.length; i ++) {
                    newUsers.push({
                        id: r[i].id,
                        name: r[i].name,
                    })
                }
            })
            users = newUsers;
            selectedUser = "0"
        } catch (e) {
            addToast({
                message:"Error loading users to filter: "+ e.toString(),
                type: "error",
                dismissible: true,
                timeout: 10000
            })
        }
    }

    LoadTable().then(r => {})
</script>

<Toasts />

<body>
    <div class="dashboard-container">
        <div class="user-selector">
            <select bind:value={selectedUser} on:change={onUserChange}>
                {#each users as user (user.id)}
                    <option value={user.id}>{user.name}</option>
                {/each}
            </select>
        </div>
        <div class="table-container">
            <table >
                <thead>
                <tr>
                    <th>ID</th>
                    <th>Item Name</th>
                    <th>Client Owner</th>
                    <th>UPC</th>
                    <th>Actions</th>
                </tr>
                </thead>
                <tbody>
                {#each inventoryItems as item}
                    <tr>
                        <td>{item.id}</td>
                        <td>{item.itemName}</td>
                        <td>{item.associatedUser}</td>
                        <!--                <td>{item.associatedUserID}</td>-->
                        <td>{item.upc}</td>
                        <td>
                            <button class="edit-button" on:click={() => openEditPopup(item)}>Delete</button>
                        </td>
                    </tr>
                {/each}
                </tbody>
            </table>
        </div>
    </div>
    {#if loadingTable}
        <div style="position: fixed; top: 95%; left: 55%">
            <LoadingDots />
        </div>
    {/if}

    {#if popUpIsOpen}
        <DeleteItem bind:isOpen={popUpIsOpen} bind:itemInfo={selectedItem} on:submit={handleDeleteSubmit}/>
    {/if}
</body>

<style>
    .dashboard-container {
        display: flex;
        justify-content: center;
        align-items: center;
        height: calc(100vh - 0px); /* Subtracting the banner height */
    }

    table {
        width: 100%;
        border-collapse: collapse;
        color: #333;
    }

    thead {
        border-radius: 15px;
        box-shadow: 0px 4px 8px rgba(0, 0, 0, 0.1);
    }

    thead th {
        position: sticky; /* make the table heads sticky */
        top: 0; /* table head will be placed from the top of the table and sticks to it */
    }

    /*.dashboard-table th,*/
    /*.dashboard-table td {*/
    /*    color: black;*/
    /*    padding: 20px;*/
    /*    text-align: left;*/
    /*}*/

    /*.dashboard-table th {*/
    /*    background-color: #007BFF;*/
    /*    color: black;*/
    /*}*/

    /*.dashboard-table tr {*/
    /*    border-bottom: 1px solid #ddd;*/
    /*}*/

    /*.dashboard-table tr:last-child {*/
    /*    border-bottom: none;*/
    /*}*/


    .table-container {
        width: 90%;
        height: 80vh;
        overflow-y: auto;
        margin: auto;
        /*padding-top: 30px;*/
        background: white;
        border-radius: 15px;
        box-shadow: 0px 4px 8px rgba(0, 0, 0, 0.1);
    }

    table {
        width: 100%;
        border-collapse: collapse;
        color: #333;
    }

    th {
        background-color: #007BFF;
        color: white;
    }

    th, td {
        padding: 15px;
        text-align: left;
        border-bottom: 1px solid #ddd;
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

    .user-selector {
        position: absolute;
        top: 15px;
        right: 40px;
        padding: 5px 10px;
        font-size: 1em;
        font-weight: bold;
        cursor: pointer;
        box-sizing: border-box;
        border: none;
        background: #ffffff;
        color: #000;
        box-shadow: 0 1px 3px rgba(0, 0, 0, 0.25);
        border-radius: 15px;
        appearance: none;
        outline: none;
    }
    .user-selector select {
        width: 100%;
        height: 100%;
        font-size: 1.0em;
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
</style>