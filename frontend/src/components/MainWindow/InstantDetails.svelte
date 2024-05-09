<script lang="ts">
    import { app_instantDetails } from "@/lib/appManager";
</script>

{#if $app_instantDetails}
    <div class="container">
        <div>
            <div class="pill">{$app_instantDetails.english}</div>
            <span class="word">{$app_instantDetails.text}</span>
            <span class="translit">{$app_instantDetails.translit}</span>
            <span class="count">
                [{$app_instantDetails.inflected_count.toLocaleString()}x]
            </span>
        </div>

        {#each $app_instantDetails.dictionary as dictionary_word}
            <div class="indent">
                — <span class="word">{dictionary_word.form}</span>
                <b>{dictionary_word.strong} {dictionary_word.gloss}</b>
                {dictionary_word.grammar}
                <span class="count">
                    [{dictionary_word.count.toLocaleString()}x]
                </span>
            </div>
        {/each}
    </div>
{/if}

<style>
    .container {
        position: absolute;
        bottom: 1em;
        right: 1em;
        width: calc(100% - 2em);
        max-width: 55ch;
        height: auto;
        padding: 1em;
        background: var(--clr-text);
        color: var(--clr-background);
        display: flex;
        flex-direction: column;
        gap: 0.5em;
        border-radius: 0.5em;
        font-size: 0.8em;
    }

    .pill {
        display: inline-block;
        background: var(--clr-main);
        padding: 0.2em 0.7em;
        border-radius: 0.2em;
        color: var(--clr-background-dark);
        margin-right: 0.2em;
    }

    .word {
        font-family: var(--font-content);
        font-size: 1.2em;
    }

    .translit {
        font-style: italic;
    }

    .count {
        color: var(--clr-text-muted);
    }

    .indent {
        padding-left: 2em;
    }
</style>
