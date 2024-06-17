<script lang="ts">
    import { Select as SelectPrimitive } from "bits-ui";
    import { Check, ChevronDown } from "lucide-svelte";

    export let selected: SelectPrimitive.Props<string>["selected"];
    export let items: SelectPrimitive.Props<string>["items"];
    export let placeholder: string;
    export let label: string = undefined;
    export let disabled: string[] = [];
</script>

<SelectPrimitive.Root portal={null} bind:selected>
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
            {#each items as item}
                <SelectPrimitive.Item
                    class="select-item"
                    value={item.value}
                    label={item.label}
                    disabled={disabled.includes(item.value)}
                >
                    {item.label}
                    <SelectPrimitive.ItemIndicator class="select-indicator">
                        <Check />
                    </SelectPrimitive.ItemIndicator>
                </SelectPrimitive.Item>
            {/each}
        </SelectPrimitive.Group>
    </SelectPrimitive.Content>
    <SelectPrimitive.Input />
</SelectPrimitive.Root>

<style global>
    .select-trigger,
    .select-content {
        border: 1px solid var(--clr-background-sub);
        border-radius: 0.2rem;
        background: var(--clr-background);
        font-size: 0.9rem;
        color: var(--clr-text);
    }

    .select-trigger {
        min-width: 8rem;
        min-width: 15ch;
        padding: 0.5rem;
        cursor: pointer;
        outline: none;

        display: flex;
        justify-content: space-between;
        align-items: center;
        gap: 1rem;
    }

    .select-trigger [data-placeholder] {
        color: var(--clr-text-sub);
    }

    .select-trigger:focus {
        outline: 3px solid var(--clr-background-dark);
        outline-offset: 2px;
    }

    .select-content {
        display: flex;
        flex-direction: column;
        gap: 0.2rem;
        z-index: 1;
        box-shadow: 0 0 10px 0 rgba(0, 0, 0, 0.2);
        max-height: 13rem;
        overflow: scroll;
    }

    .select-group {
        display: flex;
        flex-direction: column;
        gap: 0.2rem;
        padding: 0.2rem;
    }

    .select-label,
    .select-item {
        display: flex;
        flex-direction: row;
        justify-content: space-between;
        align-items: center;
        padding: 0.5rem;
        border-radius: 0.2rem;
    }

    .select-label {
        font-weight: 500;
        color: var(--clr-text-highlight);
    }

    .select-item:not([data-disabled]) {
        cursor: pointer;
    }

    .select-item[data-highlighted] {
        background: var(--clr-background-sub);
    }

    .select-item[data-disabled] {
        color: var(--clr-text-muted);
        pointer-events: none;
    }

    .select-indicator {
        height: 1em;
        width: 1em;
    }

    .select-indicator > svg,
    .select-trigger > svg {
        height: 1em;
        color: inherit;
    }
</style>
