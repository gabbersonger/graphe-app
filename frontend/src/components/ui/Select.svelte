<script lang="ts">
    import { Select as SelectPrimitive } from "bits-ui";
    import { Check, ChevronDown } from "lucide-svelte";

    type T = $$Generic;

    export let selected: SelectPrimitive.Props<T>["selected"];
    export let items: SelectPrimitive.Props<T>["items"];
    export let onSelectedChange: SelectPrimitive.Props<T>["onSelectedChange"] =
        undefined;

    export let placeholder: string;
    export let label: string;
    export let disabled: T[] = [];
</script>

<SelectPrimitive.Root portal={null} bind:selected {onSelectedChange}>
    <SelectPrimitive.Trigger class="select-trigger" {...$$restProps}>
        <SelectPrimitive.Value {placeholder} />
        <ChevronDown />
    </SelectPrimitive.Trigger>
    <SelectPrimitive.Content class="select-content">
        <SelectPrimitive.Group class="select-group">
            {#if label != undefined}
                <SelectPrimitive.Label class="select-label">
                    {label}
                </SelectPrimitive.Label>
            {/if}
            {#if items}
                {#each items as item}
                    <SelectPrimitive.Item
                        class="select-item"
                        value={item.value}
                        label={item.label}
                        disabled={disabled.includes(item.value)}
                    >
                        <slot {item}>
                            {item.label}
                        </slot>
                        <SelectPrimitive.ItemIndicator class="select-indicator">
                            <Check />
                        </SelectPrimitive.ItemIndicator>
                    </SelectPrimitive.Item>
                {/each}
            {/if}
        </SelectPrimitive.Group>
    </SelectPrimitive.Content>
    <SelectPrimitive.Input />
</SelectPrimitive.Root>

<style>
    :global(.select-trigger),
    :global(.select-content) {
        border: 1px solid var(--clr-background-sub);
        border-radius: 0.2rem;
        background: var(--clr-background);
        font-size: 0.9rem;
        color: var(--clr-text);
    }

    :global(.select-trigger) {
        min-width: 8rem;
        min-width: 15ch;
        padding: 0.5rem;
        cursor: pointer;
        outline: none;

        display: flex;
        justify-content: space-between;
        align-items: center;
        gap: 1rem;
        font-family: var(--font-system);
    }

    :global(.select-trigger [data-placeholder]) {
        color: var(--clr-text-sub);
    }

    :global(.select-trigger:focus) {
        outline: 3px solid var(--clr-background-dark);
        outline-offset: 2px;
    }

    :global(.select-content) {
        display: flex;
        flex-direction: column;
        gap: 0.2rem;
        z-index: 1;
        box-shadow: 0 0 10px 0 rgba(0, 0, 0, 0.2);
        max-height: 13rem;
        overflow: scroll;
    }

    :global(.select-group) {
        display: flex;
        flex-direction: column;
        gap: 0.2rem;
        padding: 0.2rem;
    }

    :global(.select-label),
    :global(.select-item) {
        display: flex;
        flex-direction: row;
        justify-content: space-between;
        align-items: center;
        padding: 0.5rem;
        border-radius: 0.2rem;
        font-family: var(--font-system);
    }

    :global(.select-label) {
        font-weight: 500;
        color: var(--clr-text-highlight);
    }

    :global(.select-item:not([data-disabled])) {
        cursor: pointer;
    }

    :global(.select-item[data-highlighted]) {
        background: var(--clr-background-sub);
    }

    :global(.select-item[data-disabled]) {
        color: var(--clr-text-muted);
        pointer-events: none;
    }

    :global(.select-indicator) {
        height: 1em;
        width: 1em;
    }

    :global(.select-indicator > svg),
    :global(.select-trigger > svg) {
        height: 1em;
        color: inherit;
    }
</style>
