<script lang="ts">
    import Input from "@/components/ui/Input.svelte";
    import ModalWrapper from "./ModalWrapper.svelte";

    import { bibleData } from "@/lib/Scripture/data";

    let value = "";

    let selected_book: number = -1;
    let selected_chapter: number = -1;
</script>

<ModalWrapper>
    <div slot="header">
        <Input bind:value placeholder="Passage" autofocus={true} />
    </div>
    <div slot="content">
        {#if selected_book < 0}
            <div class="books-wrapper">
                {#each bibleData as book, i}
                    <div class="block" on:click={() => (selected_book = i)}>
                        {book.abbreviation}
                    </div>
                {/each}
            </div>
        {/if}

        {#if selected_book >= 0}
            <div class="chapter-wrapper">
                <div class="block" on:click={() => (selected_book = -1)}>
                    Back
                </div>
                {#each Array(bibleData[selected_book].num_chapters) as _, i}
                    <div class="block">
                        {i + 1}
                    </div>
                {/each}
            </div>
        {/if}
    </div>
</ModalWrapper>

<style>
    .books-wrapper,
    .chapter-wrapper {
        display: flex;
        flex-direction: row;
        flex-wrap: wrap;
        justify-content: center;
        align-items: center;
        gap: 1rem;
    }

    .block {
        height: 2rem;
        aspect-ratio: 2;
        font-size: 0.8rem;

        display: flex;
        align-items: center;
        justify-content: center;

        border: 1px solid var(--clr-background-sub);
        border-radius: 0.2em;
        color: var(--clr-text);

        cursor: pointer;
    }

    .block:hover {
        background: var(--clr-background-sub);
        color: var(--clr-text-highlight);
    }

    /* .block {
        width: max(7.3%, 45px);
        aspect-ratio: 3 / 2;
        border-radius: 0.1rem;
        display: flex;
        align-items: center;
        justify-content: center;
        background: var(--clr-background-sub);
        color: var(--clr-text-sub);
        cursor: pointer;
    }

    .block:hover {
        background: var(--clr-background-dark);
        color: var(--clr-text);
    } */
</style>
