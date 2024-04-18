<script lang="ts">
    import type { ComponentType } from "svelte";
    import type { Icon } from "lucide-svelte";

    type $T = $$Generic;
    export let items: Array<$T>;

    export let onItemClick: (index: number) => void;

    export let icon: ComponentType<Icon> = null;
    export let subheading: (index: number) => string = null;
    export let heading: (index: number) => string = null;

    export let isActive: (index: number) => boolean = (_) => false;
    export let isFocused: (index: number) => boolean = (_) => false;
</script>

<div class="container">
    {#each items as _, index}
        <button
            on:click={() => onItemClick(index)}
            tabindex="0"
            class:active={isActive(index)}
            class:focus={isFocused(index)}
        >
            {#if subheading}
                <span class="subheading">
                    {#if icon}<svelte:component this={icon} />{/if}
                    {subheading(index)}
                </span>
            {/if}
            {#if heading}
                <span class="heading">{heading(index)}</span>
            {/if}
        </button>
    {/each}
</div>

<style>
    .container {
        display: flex;
        flex-direction: row;
        flex-wrap: wrap;
        justify-content: flex-start;
        gap: 1rem;
    }

    button {
        position: relative;
        width: calc((100% - 3rem) / 4);
        min-width: 10rem;
        aspect-ratio: 2 / 1;
        background: none;
        border: 2px solid var(--clr-background-sub);
        border-radius: 0.2rem;
        color: var(--clr-text);
        font-size: 1rem;
        display: flex;
        flex-direction: column;
        align-items: flex-start;
        padding: 1em;
        text-align: left;
    }

    button .subheading {
        font-size: 0.7em;
        color: var(--clr-text-sub);
        padding-bottom: 0.4em;
        display: flex;
        align-items: center;
        gap: 0.5em;
    }

    button .subheading > :global(svg) {
        height: 1em;
        width: 1em;
    }

    button .heading {
        font-size: 1em;
        text-align: left;
        line-height: 1.3;
    }

    button:not(.active):hover {
        background: var(--clr-background-sub);
        border-color: var(--clr-background-dark);
        cursor: pointer;
    }

    button.active {
        color: var(--clr-text-highlight);
        color: var(--clr-text);
        background: var(--clr-background-sub);
    }

    button.active .subheading {
        color: var(--clr-text);
        color: var(--clr-main);
    }

    button:focus,
    button.focus {
        border-color: var(--clr-main);
        outline: none;
    }
</style>
