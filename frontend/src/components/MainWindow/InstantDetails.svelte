<script lang="ts">
    import { app_instantDetails } from "@/lib/appManager";
</script>

{#if $app_instantDetails}
    <div class="container">
        <div>
            <b>{$app_instantDetails.english}</b>
            <span class="word">{$app_instantDetails.text}</span>
            {$app_instantDetails.translit}
            <span class="count">[{$app_instantDetails.inflected_count}x]</span>
        </div>
        {#each $app_instantDetails.dictionary as dictionary_word}
            <div>
                — <span class="word">{dictionary_word.form}</span>
                {dictionary_word.gloss}
                <span class="count">[{$app_instantDetails.lexeme_count}x]</span>
            </div>
        {/each}
        {#each $app_instantDetails.strongs as strongs}
            <div>{strongs.num} {strongs.grammar}</div>
        {/each}
    </div>
{/if}

<style>
    .container {
        position: absolute;
        bottom: 1em;
        right: 1em;
        width: calc(100% - 2em);
        max-width: 50ch;
        height: auto;
        padding: 1em;
        background: var(--clr-text);
        /* outline: 0.2em solid var(--clr-background-sub);
        outline-offset: 0.2em; */
        color: var(--clr-background);
        display: flex;
        flex-direction: column;
        border-radius: 0.5em;
        font-size: 0.8em;
        line-height: 1.5;
    }

    .word {
        font-family: var(--font-content);
        font-size: 1.2em;
    }

    .count {
        color: var(--clr-text-muted);
    }
</style>
