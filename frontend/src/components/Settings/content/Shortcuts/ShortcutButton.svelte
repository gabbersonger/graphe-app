<script lang="ts">
    import { Popover } from "bits-ui";
    import Command from "@/components/ui/Command.svelte";
    import {
        CirclePlus,
        Pencil,
        X,
        Check,
        Lock,
        Command as Meta,
        ArrowBigUp,
        ChevronUp,
        Option,
        Delete,
        CornerDownLeft,
        CircleArrowOutUpLeft,
        ArrowUp,
        ArrowRightToLine,
        Space,
        ArrowDown,
        ArrowLeft,
        ArrowRight,
        RotateCcw,
    } from "lucide-svelte";
    import Button from "@/components/ui/Button.svelte";

    export let shortcut: string = null;
    export let locked = false;
    export let onChange: (value: string) => void;
    export let onReset: () => void;

    let popoverOpen = false;
    let hover = false;

    function resetShortcutValue() {
        return {
            cmdorctrl: false,
            shift: false,
            control: false,
            alt: false,
            key: "",
        };
    }
    let shortcutValue = resetShortcutValue();

    const VALID_KEYS = {
        KeyA: "A",
        KeyB: "B",
        KeyC: "C",
        KeyD: "D",
        KeyE: "E",
        KeyF: "F",
        KeyG: "G",
        KeyH: "H",
        KeyI: "I",
        KeyJ: "J",
        KeyK: "K",
        KeyL: "L",
        KeyM: "M",
        KeyN: "N",
        KeyO: "O",
        KeyP: "P",
        KeyQ: "Q",
        KeyR: "R",
        KeyS: "S",
        KeyU: "U",
        KeyV: "V",
        KeyW: "W",
        KeyX: "X",
        KeyY: "Y",
        KeyZ: "Z",
        Digit0: "0",
        Digit1: "1",
        Digit2: "2",
        Digit3: "3",
        Digit4: "4",
        Digit5: "5",
        Digit6: "6",
        Digit7: "7",
        Digit8: "8",
        Digit9: "9",
        Backquote: "`",
        Minus: "-",
        Equal: "=",
        Backspace: "backspace",
        BracketLeft: "[",
        BracketRight: "]",
        Backslash: "\\",
        Semicolon: ";",
        Quote: "'",
        Enter: "enter",
        Comma: ",",
        Period: ".",
        Slash: "/",
        ArrowLeft: "left",
        ArrowRight: "right",
        ArrowUp: "up",
        ArrowDown: "down",
        Tab: "tab",
        " ": "space",
        Escape: "escape",
    };

    function keyListener(event: KeyboardEvent) {
        event.preventDefault();
        if (
            event.code in VALID_KEYS &&
            (event.metaKey || event.ctrlKey || event.altKey)
        ) {
            shortcutValue = {
                cmdorctrl: event.metaKey,
                shift: event.shiftKey,
                control: event.ctrlKey,
                alt: event.altKey,
                key: VALID_KEYS[event.code],
            };
            return;
        }
        shortcutValue = resetShortcutValue();
    }

    function onOpenChange(v: boolean) {
        if (v) {
            document.addEventListener("keydown", keyListener);
        } else {
            shortcutValue = resetShortcutValue();
            document.removeEventListener("keydown", keyListener);
        }
    }

    function resetClick() {
        popoverOpen = false;
        onReset();
    }

    function confirmClick() {
        popoverOpen = false;
        let value_array = [];
        for (const key in shortcutValue) {
            if (key == "key") {
                value_array.push(shortcutValue[key]);
            } else if (shortcutValue[key]) {
                value_array.push(key);
            }
        }
        onChange(value_array.join("+"));
        shortcutValue = resetShortcutValue();
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
        <Popover.Root
            bind:open={popoverOpen}
            {onOpenChange}
            closeOnEscape={false}
        >
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

                <div class="shortcut-tray">
                    {#if shortcutValue.key != ""}
                        <div class="value">
                            {#if shortcutValue.cmdorctrl}<Meta />{/if}
                            {#if shortcutValue.shift}<ArrowBigUp />{/if}
                            {#if shortcutValue.control}<ChevronUp />{/if}
                            {#if shortcutValue.alt}<Option />{/if}

                            {#if shortcutValue.key.length == 1}
                                <span>{shortcutValue.key}</span>
                            {:else if shortcutValue.key == "backspace"}
                                <Delete />
                            {:else if shortcutValue.key == "enter"}
                                <CornerDownLeft />
                            {:else if shortcutValue.key == "escape"}
                                <CircleArrowOutUpLeft />
                            {:else if shortcutValue.key == "space"}
                                <Space />
                            {:else if shortcutValue.key == "tab"}
                                <ArrowRightToLine />
                            {:else if shortcutValue.key == "up"}
                                <ArrowUp />
                            {:else if shortcutValue.key == "down"}
                                <ArrowDown />
                            {:else if shortcutValue.key == "left"}
                                <ArrowLeft />
                            {:else if shortcutValue.key == "right"}
                                <ArrowRight />
                            {/if}
                        </div>
                    {:else}
                        <div class="placeholder">Perform shortcut...</div>
                    {/if}
                </div>
                <div class="button-tray">
                    <Button icon={X} on:click={() => (popoverOpen = false)}>
                        Close
                    </Button>
                    <Button
                        icon={RotateCcw}
                        disabled={shortcut == null}
                        on:click={resetClick}>Reset</Button
                    >
                    <Button
                        icon={Check}
                        disabled={shortcutValue.key == ""}
                        on:click={confirmClick}
                    >
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

    /* POPOVER */
    :global(.arrow) {
        border-top: 1px solid var(--clr-background-dark);
        border-left: 1px solid var(--clr-background-dark);
    }

    :global(.shortcut-popover) {
        display: flex;
        flex-direction: row;
        padding: 1rem;
        gap: 1rem;
        background: var(--clr-background-sub);
        border: 1px solid var(--clr-background-dark);
        border-radius: 0.4rem;
        user-select: none;
        -webkit-user-select: none;
        outline: none;
        font-family: var(--font-system);
    }

    :global(.shortcut-popover .shortcut-tray) {
        width: 15rem;
        display: flex;
        flex-direction: row;
        align-items: center;
        justify-content: center;

        background: var(--clr-background);
        border-radius: 0.4rem;
        font-size: 0.7rem;
        color: var(--clr-text-muted);
        cursor: default;
    }

    :global(.shortcut-popover .shortcut-tray .value) {
        font-size: 1.5rem;
        line-height: 0;
        display: flex;
        flex-direction: row;
        align-items: center;
        justify-content: center;
        gap: 0.2rem;
    }

    :global(.shortcut-popover .shortcut-tray .value > svg) {
        height: 1.2rem;
        width: 1.2rem;
    }

    :global(.shortcut-popover .button-tray) {
        display: flex;
        flex-direction: column;
        gap: 0.5rem;
    }

    /*
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

    :global(.shortcut-popover) > :global(.close) {
        position: absolute;
        right: 0;
        top: 0;
        background: none;
        border: none;
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

    :global(.shortcut-popover) > :global(.shortcut-tray) > :global(.value) {
        font-family: var(--font-system);
        font-size: 1.5rem;
        line-height: 0;
        display: flex;
        flex-direction: row;
        align-items: center;
        justify-content: center;
        gap: 0.2rem;
    }

    :global(.shortcut-popover)
        > :global(.shortcut-tray)
        > :global(.value)
        > :global(svg) {
        height: 1.2rem;
        width: 1.2rem;
    }

    :global(.shortcut-popover) > :global(.button-tray) {
        display: flex;
        flex-wrap: wrap;
        gap: 0.5rem;
        flex-direction: row;
        justify-content: center;
    } */
</style>
