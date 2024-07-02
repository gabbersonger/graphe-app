<script lang="ts">
    import { Popover } from "bits-ui";
    import Command from "@/components/Settings/ui/Command.svelte";
    import { CirclePlus, Pencil, X, Check, Lock } from "lucide-svelte";
    import Button from "@/components/ui/Button.svelte";

    export let shortcut: string = null;
    export let locked = false;
    export let onChange: (value: string) => void;

    let popoverOpen = false;
    let hover = false;

    function confirmClick() {
        popoverOpen = false;
    }
</script>

<div
    role="alertdialog"
    class="shortcut-setter"
    class:hover={hover || popoverOpen}
    on:mouseenter={() => (hover = true)}
    on:pointerenter={() => (hover = true)}
    on:mouseover={() => (hover = true)}
    on:pointerover={() => (hover = true)}
    on:focus={() => (hover = true)}
    on:mouseout={() => (hover = false)}
    on:pointerout={() => (hover = false)}
    on:blur={() => (hover = false)}
>
    {#if shortcut}
        <div class="current"><Command text={shortcut} /></div>
    {:else}
        <div class="empty">–</div>
    {/if}

    {#if !locked}
        <Popover.Root bind:open={popoverOpen}>
            <Popover.Trigger asChild let:builder>
                <div
                    use:builder.action
                    {...builder}
                    class={shortcut ? "edit" : "prompt"}
                >
                    {#if shortcut}
                        <Pencil />
                    {:else}
                        <CirclePlus />Create
                    {/if}
                </div>
            </Popover.Trigger>
            <Popover.Content
                sideOffset={3}
                side="bottom"
                align="end"
                class="shortcut-popover"
            >
                <Popover.Arrow class="arrow" />
                <div class="shortcut-tray">Perform shortcut...</div>
                <div class="button-tray">
                    <Button
                        icon={X}
                        secondary={true}
                        on:click={() => (popoverOpen = false)}>Cancel</Button
                    >
                    <Button icon={Check} on:click={confirmClick}>
                        Confirm
                    </Button>
                </div>
            </Popover.Content>
        </Popover.Root>
    {:else}
        <div class="locked">
            <Lock />
        </div>
    {/if}
</div>

<style>
    .shortcut-setter {
        position: relative;
        width: 100%;
        height: 100%;
        display: flex;
        flex-direction: row;
        align-items: center;
        gap: 0.5rem;
    }

    .shortcut-setter:not(.hover) .edit,
    .shortcut-setter.hover .empty,
    .shortcut-setter:not(.hover) .prompt {
        display: none;
    }

    .edit {
        color: var(--clr-text-muted);
        cursor: pointer;
    }

    .locked {
        color: var(--clr-text-muted);
        cursor: no-drop;
    }

    .prompt {
        display: flex;
        justify-content: center;
        align-items: center;
        gap: 0.2rem;
        width: min-content;
        background: none;
        color: var(--clr-text-muted);
        border: 1px dashed var(--clr-text-muted);
        border-radius: 0.2rem;
        padding: 0.2rem;
        font-size: 0.7rem;
        cursor: pointer;
    }

    .prompt > :global(svg),
    .edit > :global(svg),
    .locked > :global(svg) {
        height: 1em;
        width: 1em;
    }

    :global(.arrow) {
        border-top: 1px solid var(--clr-background-dark);
        border-left: 1px solid var(--clr-background-dark);
    }

    :global(.shortcut-popover) {
        background: var(--clr-background-sub);
        border: 1px solid var(--clr-background-dark);
        border-radius: 0.4rem;
        width: 15rem;
        padding: 1rem;
        font-family: var(--font-system);
        font-size: 1rem;
        user-select: none;
        -webkit-user-select: none;
        outline: none;
    }

    :global(.shortcut-popover) > :global(.shortcut-tray) {
        min-height: 4rem;
        background: var(--clr-background);
        border-radius: 0.4rem;
        margin-bottom: 1rem;
        display: flex;
        flex-direction: row;
        align-items: center;
        justify-content: center;
        font-size: 0.7rem;
        color: var(--clr-text-muted);
    }

    :global(.shortcut-popover) > :global(.button-tray) {
        display: flex;
        flex-wrap: wrap;
        gap: 0.5rem;
        flex-direction: row;
        justify-content: center;
    }
</style>
