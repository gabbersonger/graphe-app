<script lang="ts">
    import Select from "@/components/ui/Select.svelte";

    import type { Select as SelectPrimitive } from "bits-ui";
    import { EventsEmit } from "!wails/runtime/runtime";
    import { graphe_settings } from "@/lib/stores";
    import { fontData } from "@/static/fonts";

    function makeSelectValue(
        value: string,
    ): SelectPrimitive.Props<string>["selected"] {
        return {
            label: value,
            value: value,
        };
    }

    const languages = [
        {
            name: "System",
            category: "english",
            value: makeSelectValue($graphe_settings.appearence.font.system),
            text: "This is an example of some system text.",
        },
        {
            name: "Greek",
            category: "greek",
            value: makeSelectValue($graphe_settings.appearence.font.greek),
            text: "ἐν ἀρχῇ ἐποίησεν ὁ θεὸς τὸν οὐρανὸν καὶ τὴν γῆν...",
        },
        {
            name: "Hebrew",
            category: "hebrew",
            value: makeSelectValue($graphe_settings.appearence.font.hebrew),
            text: "בְּרֵאשִׁ֖ית בָּרָ֣א אֱלֹהִ֑ים אֵ֥ת הַשָּׁמַ֖יִם וְאֵ֥ת הָאָֽרֶץ׃",
        },
        {
            name: "English",
            category: "english",
            value: makeSelectValue($graphe_settings.appearence.font.english),
            text: "In the beginning God created the heavens and the earth...",
        },
    ] as const;

    function onFontChange(
        lang: (typeof languages)[number]["name"],
        value: string,
    ) {
        EventsEmit(
            "graphe:setting",
            ["appearence", "font", lang.toLowerCase()],
            value,
        );
    }
</script>

<div class="font-selector">
    {#each languages as lang, i}
        <div class="font-selection">
            <div class="wrapper">
                <div class="font-selection-heading">{lang.name}</div>
                <div class="font-selection-example" data-lang="hebrew">
                    {lang.text}
                </div>
            </div>

            <Select
                bind:selected={lang.value}
                onSelectedChange={(v) => onFontChange(lang.name, v.value)}
                items={fontData
                    .filter((f) => f.language == lang.category)
                    .map((f) => {
                        return { value: f.name, label: f.name };
                    })}
                placeholder="Choose a font"
                label="Font Family"
            />
        </div>
    {/each}
</div>

<style>
    .font-selector {
        display: flex;
        flex-direction: column;
        gap: 0.7rem;
    }

    .font-selection {
        display: flex;
        flex-direction: row;
        align-items: center;
        justify-content: space-between;
        gap: 0.5rem;
        border: 1px solid var(--clr-background-dark);
        border-radius: 0.2rem;
        padding: 0.5rem;
        font-size: 0.8rem;
    }

    .font-selection-heading {
        font-weight: 500;
        color: var(--clr-text);
        padding-bottom: 0.2rem;
    }

    .font-selection-example {
        color: var(--clr-text);
    }
</style>
