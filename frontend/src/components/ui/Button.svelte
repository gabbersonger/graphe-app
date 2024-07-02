<script lang="ts">
    import type { ComponentType } from "svelte";
    import type { Icon } from "lucide-svelte";

    export let icon: ComponentType<Icon> = null;
    export let secondary = false;
</script>

<button on:click class:secondary>
    {#if icon}
        <div class="icon">
            <svelte:component this={icon} />
        </div>
    {/if}
    <div class="text">
        <slot />
    </div>
</button>

<style>
    button {
        display: flex;
        flex-direction: row;
        align-items: center;
        justify-content: center;
        gap: 0.3rem;
        background: none;
        color: var(--clr-text);
        outline: none;
        border: 1px solid var(--clr-background-dark);
        border-radius: 0.2rem;
        padding: 0.5rem;
        height: 2.2rem;
        font-family: var(--font-system);
        font-size: 0.7rem;
        cursor: pointer;
        line-height: 0;
    }

    button.secondary {
        color: var(--clr-text-sub);
    }

    button:hover {
        background: var(--clr-background-sub);
    }

    button .icon > :global(svg) {
        height: 1em;
        width: 1em;
    }
</style>
