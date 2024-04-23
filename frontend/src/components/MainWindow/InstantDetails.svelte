<script lang="ts">
    import { app_instantDetails } from "@/lib/appManager";
</script>

{#if $app_instantDetails}
    <div class="container">
        <p>
            <b>{$app_instantDetails.english}</b>
            <span class="word">{$app_instantDetails.text}</span>
            ({#each $app_instantDetails.dictionary as dictionary_word, index}
                {#if index != 0},&nbsp;{/if}
                <span class="word">{dictionary_word.form}</span>
            {/each})
            {$app_instantDetails.translit}
        </p>
        {#each $app_instantDetails.strongs as strongs}
            <p>{strongs.num} {strongs.grammar}</p>
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

    .container p {
        margin: 0;
    }

    .word {
        font-family: var(--font-content);
    }
</style>
