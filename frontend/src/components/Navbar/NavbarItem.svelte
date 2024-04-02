<script lang="ts">
    import type { ComponentType } from "svelte";
    import type { Icon } from "lucide-svelte";

    export let icon: ComponentType<Icon> = null;
    export let text: string = "";
</script>

<!-- svelte-ignore a11y-click-events-have-key-events -->
<div class="item" on:click>
    {#if icon}
        <div class="icon">
            <svelte:component this={icon} />
        </div>
    {/if}
    {#if text}
        <div class="text">{text.toLowerCase()}</div>
    {/if}
</div>

<style>
    .item {
        display: flex;
        flex-direction: row;
        gap: 0.4rem;
        font-size: 0.8rem;
        line-height: 1rem;
        color: var(--clr-text-sub);
        cursor: pointer;
        --wails-draggable: no-drag;
    }

    .item:hover {
        transition: 0.2s color ease-in-out;
        color: var(--clr-text);
    }

    .item .icon {
        position: relative;
        height: 0.8rem;
        aspect-ratio: 1;
    }

    .item .icon > :global(svg) {
        height: 100%;
        width: 100%;
    }

    .item .text {
        display: flex;
        flex-direction: row;
        align-items: center;
        line-height: 0;
    }
</style>
