<script lang="ts">
    import type { ComponentType } from "svelte";
    import type { Icon } from "lucide-svelte";

    export let icon: ComponentType<Icon>;
</script>

<!-- svelte-ignore a11y-click-events-have-key-events -->
<div class="item" on:click>
    <svelte:component this={icon} />
</div>

<style>
    .item {
        color: var(--clr-text-sub);
        position: relative;
        height: 1.2rem;
        aspect-ratio: 1;
        cursor: pointer;
    }

    .item:hover {
        transition: 0.2s color ease-in-out;
        color: var(--clr-text);
    }

    .item > :global(svg) {
        height: 100%;
        width: 100%;
    }
</style>
