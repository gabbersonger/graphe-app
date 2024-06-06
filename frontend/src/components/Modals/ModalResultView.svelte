<script lang="ts">
    import ModalWrapper from "@/components/Modals/ModalWrapper.svelte";
    import { onMount, type ComponentType } from "svelte";
    import {
        type Icon,
        Search,
        ArchiveX,
        ChevronLeft,
        CornerDownLeft,
        ArrowUp,
        ArrowDown,
    } from "lucide-svelte";

    export let icon: ComponentType<Icon> = null;
    export let placeholder: string;
    export let value: string;
    export let results: { value: any; display: string }[];
    export let chooseResult: (index: number) => void;
    export let noResults: string;

    let inputElem: HTMLInputElement;
    let contentElement: HTMLDivElement;
    let resultElems: HTMLButtonElement[] = [];

    let selected = 0;
    $: if (results.length > 0) {
        selected = results[0].display == "Back" ? 1 : 0;
    }

    const VIEW_GAP = 5;
    function ensureItemInView(item: HTMLButtonElement) {
        mouseEnterLocked = true;

        const top = contentElement.scrollTop;
        const bottom = top + contentElement.offsetHeight;
        const itemTop = item.offsetTop;
        const itemBottom = itemTop + item.offsetHeight;

        if (itemTop < top) {
            contentElement.scrollTop = itemTop - VIEW_GAP;
        } else if (itemBottom > bottom) {
            contentElement.scrollTop =
                itemBottom - contentElement.offsetHeight + VIEW_GAP;
        }
    }

    function handleKeyDown(e: KeyboardEvent) {
        switch (e.code) {
            case "ArrowDown":
                selected = (selected + 1) % results.length;
                ensureItemInView(resultElems[selected]);
                return;
            case "ArrowUp":
                selected = selected == 0 ? results.length - 1 : selected - 1;
                ensureItemInView(resultElems[selected]);
                return;
            case "Enter":
                e.preventDefault();
                if (results.length > 0) {
                    chooseResult(selected);
                }
                return;
        }
    }

    let mouseEnterLocked = false;
    function handleMouseEnter(index: number) {
        if (!mouseEnterLocked) {
            selected = index;
        } else {
            mouseEnterLocked = false;
        }
    }

    onMount(() => {
        inputElem.focus();

        addEventListener("keydown", handleKeyDown);

        return () => {
            removeEventListener("keydown", handleKeyDown);
        };
    });
</script>

<ModalWrapper bind:contentElement>
    <div slot="header" class="input-container">
        {#if icon}
            <label for="input">
                <Search />
            </label>
        {/if}
        <input
            bind:this={inputElem}
            id="input"
            type="text"
            {placeholder}
            autocomplete="off"
            spellcheck="false"
            bind:value
            on:keypress
        />
    </div>

    <div slot="content">
        {#if results.length > 0}
            <div class="results-container">
                {#each results as result, index}
                    <button
                        class="result"
                        class:selected={selected == index}
                        on:click={() => chooseResult(index)}
                        on:mouseenter={() => handleMouseEnter(index)}
                        bind:this={resultElems[index]}
                    >
                        {#if result.display != "Back"}
                            <div><svelte:component this={icon} /></div>
                        {:else}
                            <div><ChevronLeft /></div>
                        {/if}
                        <div>{result.display}</div>
                    </button>
                {/each}
            </div>
        {:else}
            <div class="no-results">
                <ArchiveX strokeWidth={1} />
                <div>{noResults}</div>
            </div>
        {/if}
    </div>

    <div slot="footer" class="footer">
        <div class="footer-block">
            <div class="footer-icon">
                <ArrowUp />
            </div>
            <div class="footer-icon">
                <ArrowDown />
            </div>
            <div class="footer-text">Navigate</div>
        </div>
        <div class="footer-block">
            <div class="footer-icon">
                <CornerDownLeft />
            </div>
            <div class="footer-text">Select</div>
        </div>
    </div>
</ModalWrapper>

<style>
    .input-container {
        position: relative;
        width: 100%;

        display: flex;
        flex-direction: row;
        align-items: center;
        padding: 0 1em;
        gap: 1em;

        background: transparent;
        border: none;
        border-bottom: 1px solid var(--clr-background-dark);
    }

    .input-container label {
        height: 1rem;
        aspect-ratio: 1;
        color: var(--clr-text-sub);
    }

    .input-container label > :global(svg) {
        height: 100%;
        width: 100%;
    }

    .input-container input {
        width: 100%;
        background: transparent;
        border: none;
        padding: 1rem 0;
        font-family: inherit;
        font-size: 0.9rem;
        color: var(--clr-text);
        caret-color: var(--clr-main);
        outline: none;
    }

    .input-container input::placeholder {
        color: var(--clr-text-sub);
    }

    .results-container {
        position: relative;
        width: 100%;
        display: flex;
        flex-direction: column;
        justify-content: flex-start;
        align-items: flex-start;
        padding-left: var(--modal-wrapper-scroll-width);
        padding-block: 0.4em;
    }

    .results-container .result {
        display: block;
        width: 100%;
        height: 2.5rem;
        border-radius: 0.2em;
        padding: 0 1em;

        background: none;
        border: none;
        font-size: 0.9em;
        color: var(--clr-text-sub);

        display: flex;
        justify-content: flex-start;
        align-items: center;
        gap: 0.75em;
    }

    .results-container .result :global(svg) {
        height: 1rem;
        width: 1rem;
    }

    .results-container .result.selected {
        background: var(--clr-background-sub);
        color: var(--clr-text);
    }

    .no-results {
        width: 100%;
        display: flex;
        flex-direction: column;
        align-items: center;
        justify-content: center;
        gap: 2em;
        padding-block: min(4em, 10vh);
        color: var(--clr-text-sub);
        user-select: none;
        -webkit-user-select: none;
        pointer-events: none;
    }

    .no-results :global(svg) {
        height: 4em;
        width: 4em;
    }

    .footer {
        padding: 1rem 1rem;
        display: flex;
        flex-direction: row;
        gap: 2em;
        font-size: 0.8em;
        color: var(--clr-text-sub);
    }

    .footer-block {
        display: flex;
        flex-direction: row;
        align-items: center;
        justify-content: center;
        gap: 0.5em;
    }

    .footer-icon {
        display: flex;
        align-items: center;
        justify-content: center;
        background: var(--clr-background);
        padding: 0.2em;
        border-radius: 0.2em;
        width: 1.5em;
        aspect-ratio: 1;
    }

    .footer-icon :global(svg) {
        height: 1em;
        width: 1em;
    }
</style>
