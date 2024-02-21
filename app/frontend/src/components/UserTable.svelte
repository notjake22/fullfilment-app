<script lang="ts">
    import {GetUserTableInfo, SetNewUser, EditClient} from "../../wailsjs/go/main/App";
    import {database} from "../../wailsjs/go/models";
    import AddUser from "./layovers/AddUser.svelte";
    import EditUser from "./layovers/EditUser.svelte";
    import LoadingDots from "./layovers/LoadingDots.svelte";
    import Toasts from "./alerts/Toasts.svelte";
    import {addToast} from "./alerts/ts/store";

    let users: database.User[] = []
    let errMsg: string = ""
    let loadingTable: boolean = false
    let newUser: database.User

    let editUser: boolean = false
    let editedUser: database.User

    function openEditPopup(user: database.User): void {
        editedUser = user
        editUser = true
    }

    async function handleEditSubmit(event): Promise<void>{
        const editedClient = event.detail;
        loadingTable = true;
        console.log(editedClient);
        editedClient.balance = editedClient.balance * 100;
        try {
            await EditClient(editedClient).then(r => {
                if (!r) {
                    addToast({
                        message:"Error editing client",
                        type: "error",
                        dismissible: true,
                        timeout: 5000
                    })
                    console.error("there was an issue editing user")
                } else {
                    addToast({
                        message:"Client edited",
                        type: "success",
                        dismissible: true,
                        timeout: 3000
                    })
                    LoadTable()
                    return
                }
            })
        } catch (e) {
            addToast({
                message:"Error editing client: " + e.toString(),
                type: "error",
                dismissible: true,
                timeout: 5000
            })
            console.error(e)
        }

        loadingTable = false
    }

    async function handleSubmit(event): Promise<void> {
        newUser = event.detail
        console.log(newUser)
        newUser.balance = newUser.balance * 100;
        loadingTable = true
        try {
            await SetNewUser(newUser).then(r => {
                if (!r) {
                    addToast({
                        message:"Error adding client",
                        type: "error",
                        dismissible: true,
                        timeout: 5000
                    })
                } else {
                    addToast({
                        message:"Client added",
                        type: "success",
                        dismissible: true,
                        timeout: 3000
                    })
                    LoadTable()
                    return
                }
            })
        } catch (e) {
            addToast({
                message:"Error adding client: "+ e.toString(),
                type: "error",
                dismissible: true,
                timeout: 5000
            })
            console.log(e)
        }
        loadingTable = false
    }

    async function LoadTable(): Promise<void> {
        loadingTable = true
        try {
            await GetUserTableInfo().then(r => {
                console.log(r)
                let newUsers: database.User[] = []
                for (let i = 0; i < r.length; i++) {
                    newUsers.push({
                        id: r[i].id,
                        name: r[i].name,
                        balance: (r[i].balance/100),
                        phone: r[i].phone,
                        currentlyStoredItems: r[i].currentlyStoredItems,
                        totalOrderHistory: r[i].totalOrderHistory,
                    })
                }
                newUsers = newUsers.sort((a, b) => {
                    if (parseInt(a.id) < parseInt(b.id)) {
                        return -1
                    }
                })
                users = newUsers
                console.log(users)
            })
        } catch (e) {
            addToast({
                message:"Error loading clients: "+ e.toString(),
                type: "error",
                dismissible: true,
                timeout: 10000
            })
        }
        loadingTable = false
    }

    LoadTable().then((r => {}))
</script>

<div style="position: fixed; left: 50%">
    <Toasts />
</div>


<div style="justify-content: center; align-items: center;">
    {#if loadingTable}
        <div style="position: fixed; top: 95%; left: 55%">
            <LoadingDots />
        </div>
    {/if}
</div>

<div class="dashboard-container">
    <AddUser on:submit={handleSubmit}/>
    <div class="table-container">
        <table >
            <thead>
            <tr>
                <th>ID</th>
                <th>Name</th>
                <th>Phone</th>
                <th>Balance</th>
                <th>Currently Stored Items</th>
                <th>Total Order History</th>
                <th>Actions</th>
            </tr>
            </thead>
            <tbody>
            {#each users as user}
                <tr>
                    <td>{user.id}</td>
                    <td>{user.name}</td>
                    <td>{user.phone}</td>
                    <td>{user.balance}</td>
                    <td>{user.currentlyStoredItems}</td>
                    <td>{user.totalOrderHistory}</td>
                    <td>
                        <button class="edit-button" on:click={() => openEditPopup(user)}>Edit</button>
                    </td>
                </tr>
            {/each}
            </tbody>
        </table>
    </div>
    <div style="color: black">{errMsg}</div>
</div>

{#if editUser}
    <EditUser bind:isOpen={editUser} bind:user={editedUser} on:submit={handleEditSubmit}/>
{/if}


<style>
    .dashboard-container {
        display: flex;
        justify-content: center;
        align-items: center;
        height: calc(100vh - 0px); /* Subtracting the banner height */
    }

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

    .overlay-form button {
        padding: 10px;
        border: none;
        border-radius: 5px;
        background-color: #007BFF;
        color: #fff;
        font-size: 16px;
        cursor: pointer;
    }

    .overlay-form button:hover {
        background-color: #0056b3;
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