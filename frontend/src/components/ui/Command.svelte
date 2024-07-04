<script lang="ts">
    import {
        Plus,
        CircleArrowOutUpLeft,
        ArrowUp,
        ArrowRightToLine,
        Space,
        ArrowDown,
        ArrowLeft,
        ArrowRight,
        Delete,
        ChevronUp,
        ArrowBigUp,
        Command,
    } from "lucide-svelte";

    export let styling = true;

    export let text: string;
    $: parts = text.split("+");

    const MAPPED_KEYS = {
        cmdorctrl: Command,
        shift: ArrowBigUp,
        control: ChevronUp,
        alt: Option,
        plus: Plus,
        backspace: Delete,
        left: ArrowLeft,
        right: ArrowRight,
        up: ArrowUp,
        down: ArrowDown,
        tab: ArrowRightToLine,
        space: Space,
        escape: CircleArrowOutUpLeft,
    };
</script>

{#if text}
    <div class:command={styling}>
        <div class="wrapper">
            {#each parts as part}
                {#if part in MAPPED_KEYS}
                    <svelte:component this={MAPPED_KEYS[part]} />
                {:else}
                    <span>{part}</span>
                {/if}
            {/each}
        </div>
    </div>
{/if}

<style>
    div {
        display: inline-block;
    }

    .command {
        display: inline-block;
        width: min-content;

        padding: 0.1rem 0.2rem;
        margin-inline: 0.1rem;

        background: var(--clr-background-sub);
        border: 1px solid var(--clr-background-dark);
        border-radius: 0.1rem;

        font-family: monospace;
        font-size: 0.7rem;
        color: var(--clr-text-sub);
    }

    .command span:not(:last-child) {
        margin-right: 0.2rem;
    }

    .wrapper {
        display: flex;
        flex-direction: row;
        align-items: center;
        justify-content: center;
        gap: 0.1em;
        line-height: 0;
    }

    div > :global(svg) {
        height: 1em;
        width: 1em;
        display: inline;
    }
</style>
