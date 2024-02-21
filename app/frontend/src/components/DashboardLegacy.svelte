<script lang="ts">
    import UserTable from "./UserTable.svelte";
    import InventoryTable from "./InventoryTable.svelte";

    let currentPage: string

    function togglePage(newPage: string): void {
        currentPage = newPage
    }
</script>

<body>
<div class="dashboard-banner">Swift Fulfillment</div>
<div class="nav-bar">
    <div
            class="nav-item {currentPage === 'Users' ? 'active' : ''}"
            on:click={() => togglePage('Users')}>
        Users
    </div>
    <div
            class="nav-item {currentPage === 'Scan' ? 'active' : ''}"
            on:click={() => togglePage('Scan')}>
        Scan
    </div>
    <div
            class="nav-item {currentPage === 'Inventory' ? 'active' : ''}"
            on:click={() => togglePage('Inventory')}>
        Inventory
    </div>
    <div
            class="nav-item {currentPage === 'Admin' ? 'active' : ''}"
            on:click={() => togglePage('Admin')}>
        Admin Management
    </div>
</div>
<div class="component-controller-container">
    {#if currentPage === "Users"}
        <UserTable />
    {:else if currentPage === "Inventory"}
        <InventoryTable />
    {/if}
</div>
</body>

<style>
    body {
        font-family: 'Helvetica Neue', Arial, sans-serif;
        background-color: #f2f2f2;
        margin: 0;
        padding: 0;
    }

    .component-controller-container {
        /*display: flex;*/
        height: calc(100vh - 140px);
    }

    .dashboard-banner {
        background-color: #007BFF;
        color: #fff;
        width: 100%;
        padding: 20px 0;
        text-align: center;
        font-size: 1.5em;
        font-weight: bold;
    }

    .nav-bar {
        background-color: #fff;
        color: #333;
        display: flex;
        justify-content: space-around;
        align-items: center;
        padding: 20px 0;
        font-size: 1.2em;
        font-weight: 500;
        box-shadow: 0 2px 4px rgba(0, 0, 0, 0.1);
    }

    .nav-item {
        cursor: pointer;
        padding: 5px 10px;
        border-radius: 10px;
        transition: background-color 0.2s ease;
    }

    .nav-item:hover {
        background-color: #f2f2f2;
    }

    .nav-item.active {
        color: #007BFF;
    }
</style>