<script lang="ts">
    import { GrapheEvent } from "@/lib/utils";
    import { X } from "lucide-svelte";

    function closeSidebar() {
        GrapheEvent("window:workspace:sidebar", "toggle");
    }
</script>

<div id="sidebar">
    <div class="sidebar-top">
        <div class="header">
            <span>Analytics</span>
            <button on:click={closeSidebar}>
                <X />
            </button>
        </div>
        controls
    </div>
    <div class="sidebar-bottom">
        <div class="wrapper">list</div>
    </div>
</div>

<style>
    #sidebar {
        position: relative;
        width: 100%;
        height: 100%;
        background: var(--clr-background-sub);
        color: var(--clr-text);
        border-left: 1px solid var(--clr-background-dark);
        display: flex;
        flex-direction: column;
    }

    .sidebar-top {
        border-bottom: 1px solid var(--clr-background-dark);
    }

    #sidebar .sidebar-bottom {
        flex: 1;
        overflow: hidden;
    }

    #sidebar .sidebar-bottom .wrapper {
        overflow: scroll;
        height: 100%;
    }

    .sidebar-top,
    .sidebar-bottom .wrapper {
        position: relative;
        width: 100%;
        isolation: isolate;
        padding: 1rem;
    }

    .header {
        display: flex;
        justify-content: space-between;
        align-items: center;
        padding-bottom: 1rem;
    }

    .header span {
        font-family: var(--font-system);
        font-size: 0.7rem;
        font-weight: bold;
        color: var(--clr-text-sub);
        text-transform: uppercase;
    }

    button {
        --size-button-size: 35px;
        width: var(--size-button-size);
        height: var(--size-button-size);
        background: none;
        color: var(--clr-text-muted);
        border: 1px solid var(--clr-text-muted);
        border-radius: 100%;
        display: flex;
        justify-content: center;
        align-items: center;
    }

    button:hover {
        color: var(--clr-text-sub);
        border-color: var(--clr-text-sub);
        cursor: pointer;
    }
</style>
