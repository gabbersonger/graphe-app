<script lang="ts">
    import type { data } from "!wails/go/models";
    import { workspace_instantDetailsData } from "@/lib/stores";
    import { GrapheLog } from "@/lib/utils";

    type InstantDetailsData_Fields = Map<string, string | number>;

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

    function handleData(data: data.ScriptureWordData) {
        if (data == null) {
            shown = false;
            return;
        }

        detail = {
            version: data.version,
            ref: data.ref,
            word_number: data.word_number,
            text: data.text,
            fields: new Map(),
            collections: [],
        };

        for (let i = 0; i < data.fields.length; i++) {
            const field = data.fields[i];
            switch (typeof field.data) {
                case "string":
                case "number":
                    detail.fields.set(field.name, field.data);
                    break;
                case "object":
                    if (
                        field.data != null &&
                        field.data.constructor.name == "Array"
                    ) {
                        for (let j = 0; j < field.data.length; j++) {
                            const collection = new Map();
                            for (let k = 0; k < field.data[j].length; k++) {
                                collection.set(
                                    field.data[j][k].name,
                                    field.data[j][k].data,
                                );
                            }
                            detail.collections.push(collection);
                        }
                        break;
                    }
                default:
                    GrapheLog(
                        "error",
                        `[InstantDetails] Invalid type in instant details data (field: \`${field.name}\`, type: \`${typeof field.data}\`)`,
                    );
            }
        }
        shown = true;
    }

    $: handleData($workspace_instantDetailsData);
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
        font-family: var(--font-system);
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
