<script lang="ts">
    import NavbarItem from "@/components/Navbar/NavbarItem.svelte";
    import Button from "@/components/ui/Button.svelte";
    import {
        BookOpenText,
        PanelRight,
        PanelRightClose,
        Search,
        Quote,
    } from "lucide-svelte";
    import { ui_showSidebar, ui_currentRef, ui_modal } from "@/stores/app";
    import { bibleRefToString } from "@/lib/Scripture/ref";

    let width: number;
</script>

<div id="navbar" bind:clientWidth={width}>
    <div class="container">
        {#if width > 450}
            <div class="wrapper"></div>
        {/if}

        <div class="wrapper wrapper-nav">
            <NavbarItem
                icon={BookOpenText}
                text="gnt"
                on:click={() => ($ui_modal = "chooseText")}
            />

            <NavbarItem
                icon={Quote}
                text={bibleRefToString($ui_currentRef, "chapter")}
                on:click={() => ($ui_modal = "choosePassage")}
            />

            <div class="separator"></div>

            <NavbarItem icon={Search} on:click={() => ($ui_modal = "search")} />
        </div>

        <div class="wrapper">
            <Button
                icon={$ui_showSidebar ? PanelRightClose : PanelRight}
                on:click={() => ($ui_showSidebar = !$ui_showSidebar)}
            />
        </div>
    </div>

    <div class="navbar-fade"></div>
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

    .wrapper-nav {
        height: 36px;
        background: var(--clr-background-dark);
        border-radius: 0.4em;
        padding: 0 1.8em;
    }

    .separator {
        height: 1rem;
        width: 0.15rem;
        background: var(--clr-background-sub);
    }

    .navbar-fade {
        --size-fade-height: 20px;
        position: absolute;
        width: 90vw;
        height: var(--size-fade-height);
        right: 5vw;
        bottom: calc(-1 * var(--size-fade-height));
        background: linear-gradient(
            180deg,
            var(--clr-background) 30%,
            rgba(0, 0, 0, 0) 100%
        );
        z-index: 1;
    }
</style>
