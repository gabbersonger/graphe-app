<script lang="ts">
    import type { ComponentType } from "svelte";
    import type { Icon } from "lucide-svelte";

    export let icon: ComponentType<Icon> = null;
    export let placeholder = "";
    export let value = "";
</script>

<div class="input-wrapper" class:icon>
    {#if icon}
        <div class="icon">
            <svelte:component this={icon} />
        </div>
    {/if}

    <input type="text" {placeholder} bind:value />
</div>

<style>
    .input-wrapper {
        position: relative;
        width: 100%;
        max-width: 300px;
        height: 2.2rem;
        background: none;
        color: var(--clr-text);
        border: 1px solid var(--clr-background-dark);
        border-radius: 0.2rem;
        padding: 0.5rem;
        font-family: var(--font-system);
        font-size: 0.8rem;
    }

    .input-wrapper .icon > :global(svg) {
        height: 1rem;
        width: 1rem;
        color: var(--clr-text-sub);
    }

    .input-wrapper input {
        position: absolute;
        background: none;
        border: none;
        outline: none;
        color: inherit;
        inset: 0;
        padding: 0.5rem;
    }

    .input-wrapper input::placeholder {
        color: var(--clr-text-sub);
    }

    .input-wrapper.icon input {
        padding-left: 2rem;
    }
</style>
