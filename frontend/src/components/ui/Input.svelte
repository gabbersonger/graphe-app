<script lang="ts">
    import { onMount } from "svelte";
    import type { ComponentType } from "svelte";
    import type { Icon } from "lucide-svelte";

    export let value: string;
    export let placeholder: string = "";
    export let icon: ComponentType<Icon> = null;

    export let focus: boolean = false;

    let inputEl: HTMLInputElement;
    onMount(() => {
        if (focus) inputEl.focus();
    });
</script>

<div class="container">
    {#if icon}
        <label for="input">
            <svelte:component this={icon} />
        </label>
    {/if}
    <input
        id="input"
        type="text"
        autocomplete="off"
        spellcheck="false"
        bind:value
        on:keypress
        on:focus={() => (focus = true)}
        on:blur={() => (focus = false)}
        {placeholder}
        bind:this={inputEl}
    />
    <div class="focus-ring"></div>
</div>

<style>
    .container {
        position: relative;
        width: 100%;
        background: var(--clr-background-sub);
        border-radius: 0.4em;
        display: flex;
        flex-direction: row;
        align-items: center;
        padding: 0 1em;
        gap: 1em;
        border: 1px solid var(--clr-background-dark);
    }

    .container label {
        height: 1.5rem;
        aspect-ratio: 1;
        color: var(--clr-text-sub);
    }

    .container label > :global(svg) {
        height: 100%;
        width: 100%;
    }

    .container input {
        width: 100%;
        background: transparent;
        border: none;
        padding: 1rem 0;
        font-family: inherit;
        font-size: 1rem;
        color: var(--clr-text);
        caret-color: var(--clr-main);
        outline: none;
    }

    .container input::placeholder {
        color: var(--clr-text-sub);
    }

    .container .focus-ring {
        display: none;
        position: absolute;
        inset: 0;
        background: transparent;
        pointer-events: none;
        border-radius: 0.3rem;
        outline: 0.2rem solid var(--clr-background-sub);
        outline-offset: 0.2em;
    }

    .container input:focus + .focus-ring {
        display: block;
    }
</style>
