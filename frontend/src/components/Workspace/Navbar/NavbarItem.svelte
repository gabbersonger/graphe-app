<script lang="ts">
    import Command from "@/components/ui/Command.svelte";
    import type { ComponentType } from "svelte";
    import type { Icon } from "lucide-svelte";

    export let icon: ComponentType<Icon> | null = null;
    export let text: string = "";
    export let tooltip: string = "";
    export let command: string = "";
    export let selected: boolean = false;
    export let disabled: boolean = false;
</script>

<button class="item" on:click class:selected class:disabled>
    <div class="wrapper">
        {#if icon != null}
            <div class="icon">
                <svelte:component
                    this={icon}
                    strokeWidth={selected ? 3 : 1.5}
                />
            </div>
        {/if}
        {#if text}
            <div class="text">{text.toLowerCase()}</div>
        {/if}
    </div>

    {#if tooltip || command}
        <div class="tooltip-wrapper">
            <div class="tooltip">
                {tooltip}
                {#if command}
                    <Command text={command} styling={false} />
                {/if}
            </div>
        </div>
    {/if}
</button>

<style>
    .item {
        position: relative;
        font-size: 0.8rem;
        line-height: 1rem;
        font-family: var(--font-system);
        color: var(--clr-text-sub);
        cursor: pointer;
        background: none;
        border: none;
        --wails-draggable: no-drag;
    }

    .item .wrapper {
        position: relative;
        display: flex;
        flex-direction: row;
        align-items: center;
        gap: 0.4rem;
    }

    .item.selected:not(.disabled) {
        color: var(--clr-main);
    }

    .item.disabled {
        pointer-events: none;
        color: var(--clr-text-muted);
    }

    .item:not(.selected):hover {
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
        line-height: 0;
        white-space: nowrap;
    }

    .item .tooltip-wrapper {
        position: absolute;
        left: 50%;
        width: 100vw;
        pointer-events: none;
        padding-top: calc(1rem + 2px);
    }

    .item:not(:hover) .tooltip-wrapper {
        display: none;
    }

    .item .tooltip-wrapper .tooltip {
        --size-triangle-height: 0.6em;
        position: relative;
        height: 1.8rem;
        width: min-content;
        transform: translateX(-50%);
        background: var(--clr-text);
        color: var(--clr-background);
        padding: 0.5em 0.8em;
        border-radius: 0.2em;
        z-index: 2;
        margin-top: var(--size-triangle-height);
        white-space: nowrap;
        display: flex;
        flex-direction: row;
        align-items: center;
        gap: 0.4rem;
        line-height: 0;
    }

    .item .tooltip-wrapper .tooltip::after {
        content: "";
        display: block;
        height: 0;
        width: 0;
        position: absolute;
        top: calc(-1 * var(--size-triangle-height) + 1px);
        left: calc(50% - var(--size-triangle-height) / 1.2);
        border-left: calc(var(--size-triangle-height) / 1.2) solid transparent;
        border-right: calc(var(--size-triangle-height) / 1.2) solid transparent;
        border-bottom: var(--size-triangle-height) solid var(--clr-text);
    }
</style>
