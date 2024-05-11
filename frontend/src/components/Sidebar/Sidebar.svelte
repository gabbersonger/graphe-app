<script lang="ts">
    import { ui_showSidebar, ui_sidebarSection } from "@/lib/uiManager";
    import Button from "@/components/ui/Button.svelte";
    import { sidebarData } from "@/components/Sidebar/data";
    import { PanelRightClose, X } from "lucide-svelte";
</script>

<div id="sidebar">
    <div class="shelf">
        <div class="invis">
            <Button icon={PanelRightClose} />
        </div>
        <div>
            <!-- TODO: make this into a dropdown on left -->
            {#each sidebarData as el}
                <Button
                    icon={el.icon}
                    active={$ui_sidebarSection == el.name}
                    on:click={() => ($ui_sidebarSection = el.name)}
                />
            {/each}
        </div>
        <Button
            icon={PanelRightClose}
            on:click={() => ($ui_showSidebar = false)}
        />
    </div>
    <div class="content">
        {#each sidebarData as el}
            <div class="wrapper" class:hidden={$ui_sidebarSection != el.name}>
                <svelte:component this={el.window} />
            </div>
        {/each}
    </div>
</div>

<style>
    #sidebar {
        position: relative;
        width: 100%;
        height: 100%;
        background: var(--clr-background-dark);
        color: var(--clr-text);
        display: grid;
        grid-template-rows: var(--size-navbar-height) 1fr;
        grid-template-columns: 1fr;
        grid-template-areas: "shelf" "content";
    }

    .shelf {
        grid-area: shelf;
        display: flex;
        flex-direction: row;
        align-items: center;
        justify-content: space-between;
        padding: 0 1em;
    }

    .shelf > div {
        display: flex;
        flex-direction: row;
        align-items: center;
        justify-content: center;
        gap: 1.2em;
    }

    .shelf .invis {
        opacity: 0;
        user-select: none;
        -webkit-user-select: none;
        pointer-events: none;
    }

    .content {
        grid-area: content;
        overflow: scroll;
        position: relative;
    }

    .content .wrapper {
        position: absolute;
        width: 100%;
        top: 0;
        padding: 0 1rem;
    }

    .content .wrapper.hidden {
        display: none;
    }
</style>
