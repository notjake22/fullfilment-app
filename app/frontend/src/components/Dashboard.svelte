<script lang="ts">
    import UserTable from "./UserTable.svelte";
    import InventoryTable from "./InventoryTable.svelte";
    import AdminManagement from "./AdminManagement.svelte";
    import ScanIn from "./ScanIn.svelte";
    import ScanOut from "./ScanOut.svelte";
    import {fade} from "svelte/transition"
    let logoSrc = "../assets/images/swift-logo.png"

    type ComponentStructure = {
        name: string
        component: any
    }

    let components: ComponentStructure[] = [
        { name: 'Clients', component: UserTable },
        { name: 'Inventory', component: InventoryTable },
        { name: 'Scan In', component: ScanIn },
        { name: 'Scan Out', component: ScanOut },
        { name: 'Admin Management', component: AdminManagement },
    ];

    let selectedComponent = components[0];
</script>

<body>

<div class="logo">Swift Prep</div>
<div class="logo-image">
    <img alt="logo image" style="height: 100px;width: 100px;" src="https://cdn.discordapp.com/attachments/670383959942889472/1136732375632593077/swift-logo.png">
</div>

<div class="sidebar">
    {#each components as component (component.name)}
        <a transition:fade class="sidebar-item {selectedComponent === component ? 'selected' : ''}" on:click={() => (selectedComponent = component)}>{component.name}</a>
    {/each}
</div>

<div class="main-content">
    <svelte:component this={selectedComponent.component} name={selectedComponent.name} />
</div>

</body>



<style>
    body {
        font-family: 'Helvetica Neue', Arial, sans-serif;
        /*background-color: #f2f2f2;*/
        margin: 0;
        padding: 0;
    }

    .sidebar {
        z-index: 1;
        position: fixed;
        left: 0;
        top: 0;
        bottom: 0;
        width: 150px; /* smaller width */
        background-color: #f5f5f5;
        overflow: auto;
        display: flex; /* add flexbox */
        flex-direction: column; /* make flexbox vertical */
        justify-content: center; /* center items vertically */
    }

    .sidebar-item {
        display: block;
        padding: 1em;
        text-decoration: none;
        color: black;
        border-radius: 30px;
    }

    .sidebar-item.selected {
        background-color: #6699cc;
        color: white;
    }

    .main-content {
        margin-left: 200px; /* width of the sidebar */
        position: relative;
    }

    .logo {
        z-index: 2;
        position: absolute;
        top: 0;
        left: 0;
        padding: 0.85em;
        font-size: 1.5em;
        font-weight: bold;
        color: #1b2636;
    }

    .logo-image {
        z-index: 2;
        top: 50px;
        position: absolute;
        /*top: 0;*/
        left: 25px;
    }
</style>