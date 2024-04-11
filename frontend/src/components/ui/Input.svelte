<script lang="ts">
    import { onMount } from "svelte";

    export let value: string;
    export let placeholder: string = "";
    export let autofocus: boolean = false;

    let inputEl: HTMLInputElement;
    onMount(() => {
        if (autofocus) inputEl.focus();
    });
</script>

<input type="text" bind:value {placeholder} bind:this={inputEl} />

<style>
    input {
        width: 100%;
        background: var(--clr-background);
        border: 1px solid var(--clr-background-sub);
        border-radius: 0.5rem;
        padding: 1rem;
        font-size: 1rem;
        color: var(--clr-text);
        caret-color: var(--clr-main);
    }

    input::placeholder {
        color: var(--clr-text-sub);
    }

    input:focus {
        outline: 0.2rem solid var(--clr-background-sub);
        outline-offset: 0.1rem;
    }
</style>
