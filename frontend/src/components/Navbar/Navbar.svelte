<script>
    import NavbarItem from "@/components/Navbar/NavbarItem.svelte";
    import NavbarButton from "@/components/Navbar/NavbarButton.svelte";
    import {
        BookOpenText,
        PanelRight,
        PanelRightClose,
        Search,
        Quote,
    } from "lucide-svelte";
    import { ui_showSidebar } from "@/stores/app";
</script>

<div id="navbar">
    <div class="container">
        <div class="wrapper wrapper-left">
            <NavbarItem
                icon={BookOpenText}
                text="gnt"
                on:click={() => console.log("hi")}
            />

            <NavbarItem
                icon={Quote}
                text="matthew 20.12"
                on:click={() => console.log("hi")}
            />

            <div class="separator"></div>

            <NavbarItem icon={Search} on:click={() => console.log("hi")} />
        </div>

        <div class="wrapper">
            <NavbarButton
                icon={$ui_showSidebar ? PanelRightClose : PanelRight}
                on:click={() => {
                    $ui_showSidebar = !$ui_showSidebar;
                }}
            />
        </div>
    </div>
</div>

<style>
    #navbar {
        position: relative;
        width: 100%;
        height: 100%;
        display: flex;
        flex-direction: row;
        align-items: center;
        justify-content: flex-start;
        padding: 0 1em;
        user-select: none;
        -webkit-user-select: none;
    }

    .container {
        width: 100%;
        height: 100%;
        display: flex;
        flex-direction: row;
        align-items: center;
        justify-content: space-between;
    }

    .wrapper {
        display: flex;
        flex-direction: row;
        align-items: center;
        justify-content: flex-start;
        gap: 1.2em;
    }

    .wrapper-left {
        height: 28px;
        background: var(--clr-background-dark);
        border-radius: 0.4em;
        padding: 0 1.2em;
    }

    .separator {
        height: 1rem;
        width: 0.15rem;
        background: var(--clr-background-sub);
    }
</style>
