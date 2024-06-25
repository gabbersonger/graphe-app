<script lang="ts">
    import Select from "@/components/ui/Select.svelte";

    import type { Select as SelectPrimitive } from "bits-ui";
    import { EventsEmit } from "!wails/runtime/runtime";
    import { graphe_settings } from "@/lib/stores";
    import { fontData } from "@/static/fonts";
    import { GrapheLog } from "@/lib/utils";

    function makeSelectValue(
        value: string,
    ): SelectPrimitive.Props<string>["selected"] {
        const font = fontData.find((f) => f.name == value);
        if (!font) {
            GrapheLog(
                "error",
                `Invalid font (${value}) passed to \`makeSelectValue\``,
            );
        }
        return {
            label: font.name,
            value: font.css_line,
        };
    }

    const languages = [
        {
            name: "System",
            category: "english",
            option: makeSelectValue($graphe_settings.appearence.font.system),
            text: "This is an example of some system text.",
        },
        {
            name: "Greek",
            category: "greek",
            option: makeSelectValue($graphe_settings.appearence.font.greek),
            text: "ἐν ἀρχῇ ἐποίησεν ὁ θεὸς τὸν οὐρανὸν καὶ τὴν γῆν...",
        },
        {
            name: "Hebrew",
            category: "hebrew",
            option: makeSelectValue($graphe_settings.appearence.font.hebrew),
            text: "בְּרֵאשִׁ֖ית בָּרָ֣א אֱלֹהִ֑ים אֵ֥ת הַשָּׁמַ֖יִם וְאֵ֥ת הָאָֽרֶץ׃",
        },
        {
            name: "English",
            category: "english",
            option: makeSelectValue($graphe_settings.appearence.font.english),
            text: "In the beginning God created the heavens and the earth...",
        },
    ] as const;

    function onFontChange(
        lang: (typeof languages)[number]["name"],
        font: string,
    ) {
        EventsEmit(
            "graphe:setting",
            ["appearence", "font", lang.toLowerCase()],
            font,
        );
    }
</script>

<div class="font-selector">
    {#each languages as lang}
        <div class="font-selection">
            <div class="wrapper">
                <div class="font-selection-heading">{lang.name}</div>
                <div
                    class="font-selection-example"
                    data-lang={lang.name.toLowerCase()}
                >
                    {lang.text}
                </div>
            </div>

            <Select
                bind:selected={lang.option}
                onSelectedChange={(v) => onFontChange(lang.name, v.label)}
                items={fontData
                    .filter((f) => f.language == lang.category)
                    .map((f) => makeSelectValue(f.name))}
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
        font-family: var(--font-system);
        font-weight: 500;
        color: var(--clr-text);
        padding-bottom: 0.2rem;
    }

    .font-selection-example {
        color: var(--clr-text);
    }

    .font-selection-example[data-lang="hebrew"] {
        font-family: var(--font-hebrew);
    }

    .font-selection-example[data-lang="greek"] {
        font-family: var(--font-greek);
    }

    .font-selection-example[data-lang="english"] {
        font-family: var(--font-english);
    }

    .font-selection-example[data-lang="system"] {
        font-family: var(--font-system);
    }
</style>
