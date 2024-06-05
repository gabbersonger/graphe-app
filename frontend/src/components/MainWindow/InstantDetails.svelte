<script lang="ts">
    import { app_instantDetails, app_version } from "@/lib/appManager";

    function count(s: string) {
        return parseInt(s).toLocaleString();
    }
</script>

{#if $app_instantDetails}
    <div class="container">
        {#if $app_version == "gnt"}
            <div>
                <div class="pill">{$app_instantDetails.fields["English"]}</div>
                <span class="word">{$app_instantDetails.text}</span>
                <span class="translit"
                    >{$app_instantDetails.fields["Translit"]}</span
                >
                <span class="count">
                    [{count($app_instantDetails.fields["InflectedCount"])}x]
                </span>
            </div>

            {#each $app_instantDetails.collections as c}
                <div class="indent">
                    — <span class="word">{c["Form"]}</span>
                    <b>{c["Strong"]} {c["Gloss"]}</b>
                    {c["Grammar"]}
                    <span class="count">
                        [{count(c["FormCount"])}x]
                    </span>
                </div>
            {/each}
        {:else if $app_version == "lxx"}
            <div>
                <div class="pill">{$app_instantDetails.fields["English"]}</div>
                <span class="word">{$app_instantDetails.text}</span>
                <span class="translit"
                    >{$app_instantDetails.fields["Translit"]}</span
                >
                <span class="count">
                    [{count($app_instantDetails.fields["InflectedCount"])}x]
                </span>
            </div>

            <div class="indent">
                — <span class="word">{$app_instantDetails.fields["Form"]}</span>
                <b
                    >{$app_instantDetails.fields["Strong"]}
                    {$app_instantDetails.fields["Gloss"]}</b
                >
                {$app_instantDetails.fields["Grammar"]}
                <span class="count">
                    [{count($app_instantDetails.fields["FormCount"])}x]
                </span>
            </div>
        {:else if $app_version == "esv"}
            <div>
                <span class="word">{$app_instantDetails.text}</span>
                <span class="count">
                    [{count($app_instantDetails.fields["EnglishCount"])}x]
                </span>
                <div class="pill">{$app_instantDetails.fields}</div>
            </div>
        {/if}
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
