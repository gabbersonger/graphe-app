<script lang="ts">
    import { ui_sidebarSection } from "@/stores/app";
    import Button from "@/components/ui/Button.svelte";
    import { sidebarData } from "@/components/Sidebar/data";
</script>

<div id="sidebar">
    <div class="shelf">
        {#each sidebarData as el}
            <Button
                icon={el.icon}
                active={$ui_sidebarSection == el.name}
                on:click={() => ($ui_sidebarSection = el.name)}
            />
        {/each}
    </div>
    <div class="content">
        {#each sidebarData as el}
            {#if $ui_sidebarSection == el.name}
                <svelte:component this={el.window} />
            {/if}
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
        justify-content: center;
        gap: 1.2em;
    }

    .content {
        grid-area: content;
        padding: 0 1rem;
        overflow: scroll;
    }
</style>
