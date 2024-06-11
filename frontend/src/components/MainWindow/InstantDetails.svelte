<script lang="ts">
    import type { database } from "!wails/go/models";
    import { app_instantDetails } from "@/lib/managers/appManager";

    type InstantDetailsData_Fields = Map<string, string>;

    type InstantDetailsData = {
        version: string;
        ref: number;
        word_number: number;
        text: string;
        fields?: InstantDetailsData_Fields;
        collections?: InstantDetailsData_Fields[];
    };

    let shown = false;
    let detail: InstantDetailsData = undefined;

    function addField(
        fields: InstantDetailsData_Fields,
        name: string,
        data: string,
    ) {
        if (name.includes("[int]")) {
            name = name.replace("[int]", "");
            data = parseInt(data).toLocaleString();
        }
        fields.set(name, data);
    }

    function handleData(data: database.ScriptureWordData) {
        if (data == null) {
            shown = false;
            return;
        }

        detail = {
            version: data.version,
            ref: data.ref,
            word_number: data.word_number,
            text: data.text,
        };

        if (data.fields && data.fields.length > 0) {
            detail.fields = new Map();
            for (let i = 0; i < data.fields.length; i++) {
                addField(
                    detail.fields,
                    data.fields[i].name,
                    data.fields[i].data,
                );
            }
        }
        if (data.collections && data.collections.length > 0) {
            detail.collections = [];
            for (let i = 0; i < data.collections.length; i++) {
                const collection = new Map();
                for (let j = 0; j < data.collections[i].length; j++) {
                    addField(
                        collection,
                        data.collections[i][j].name,
                        data.collections[i][j].data,
                    );
                }
                detail.collections.push(collection);
            }
        }

        shown = true;
    }

    $: handleData($app_instantDetails);
</script>

{#if shown}
    <div class="container">
        {#if detail.version == "gnt"}
            <div>
                <div class="pill">{detail.fields.get("English")}</div>
                <span class="word">{detail.text}</span>
                <span class="translit">{detail.fields.get("Translit")}</span>
                <span class="count">
                    [{detail.fields.get("InflectedCount")}x]
                </span>
            </div>

            {#each detail.collections as c}
                <div class="indent">
                    — <span class="word">{c.get("Form")}</span>
                    <b>{c.get("Strong")} {c.get("Gloss")}</b>
                    {c.get("Grammar")}
                    <span class="count">[{c.get("FormCount")}x]</span>
                </div>
            {/each}
        {:else if detail.version == "lxx"}
            <div>
                <div class="pill">{detail.fields.get("English")}</div>
                <span class="word">{detail.text}</span>
                <span class="translit">{detail.fields.get("Translit")}</span>
                <span class="count">
                    [{detail.fields.get("InflectedCount")}x]
                </span>
            </div>

            <div class="indent">
                — <span class="word">{detail.fields.get("Form")}</span>
                <b>
                    {detail.fields.get("Strong")}
                    {detail.fields.get("Gloss")}
                </b>
                {detail.fields.get("Grammar")}
                <span class="count">[{detail.fields.get("FormCount")}x]</span>
            </div>
        {:else if detail.version == "esv"}
            <div>
                <span class="word">{detail.text}</span>
                <span class="count">
                    [{detail.fields.get("EnglishCount")}x]
                </span>
            </div>

            {#each detail.collections as c}
                <div class="indent">
                    — <span class="word">Form</span>
                    <b>{c.get("Strong")} Gloss</b>
                    Grammar
                    <span class="count">[{c.get("StrongCount")}x]</span>
                </div>
            {/each}
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
