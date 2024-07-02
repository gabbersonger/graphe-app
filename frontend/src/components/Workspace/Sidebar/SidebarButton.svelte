<script lang="ts">
    import type { ComponentType } from "svelte";
    import type { Icon } from "lucide-svelte";

    export let icon: ComponentType<Icon>;
    export let active: boolean = false;
</script>

<button class="item" class:active on:click>
    <svelte:component this={icon} />
</button>

<style>
    .item {
        color: var(--clr-text-sub);
        position: relative;
        height: 1.2rem;
        aspect-ratio: 1;
        cursor: pointer;
        background: none;
        border: none;
        padding: 0;
    }

    .item:hover {
        transition: 0.2s color ease-in-out;
        color: var(--clr-text);
    }

    .item > :global(svg) {
        height: 100%;
        width: 100%;
    }

    .item.active {
        color: var(--clr-main);
    }
</style>
