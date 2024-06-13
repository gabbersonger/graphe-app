<script lang="ts">
    import type { ComponentType } from "svelte";
    import { type Icon, Undo2 } from "lucide-svelte";

    export let backButton: () => void = null;

    type $T = $$Generic;
    export let items: Array<$T>;

    export let rowData: { number: number; maxwidth: number };

    export let onItemClick: (index: number) => void;

    export let icon: ComponentType<Icon> = null;
    export let subheading: (index: number) => string = null;
    export let heading: (index: number) => string = null;

    export let isActive: (index: number) => boolean = (_) => false;
    export let isFocused: (index: number) => boolean = (_) => false;
</script>

<div
    class="container"
    style:--row-length={rowData.number.toString()}
    style:--item-maxwidth={rowData.maxwidth.toString()}
>
    {#if backButton}
        <button on:click={backButton} tabindex="0">
            <span class="subheading">
                <Undo2 />Return
            </span>
            <span>Go Back</span>
        </button>
    {/if}

    {#each items as _, index}
        <button
            on:click={() => onItemClick(index)}
            tabindex="0"
            class:active={isActive(index)}
            class:focus={isFocused(index)}
        >
            {#if subheading || icon}
                <span class="subheading">
                    {#if icon}<svelte:component this={icon} />{/if}
                    {#if subheading}{subheading(index)}{/if}
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
        width: calc(
            (100% - ((var(--row-length) - 1) * 1rem)) / var(--row-length)
        );
        min-width: calc(var(--item-maxwidth) * 1rem);
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
