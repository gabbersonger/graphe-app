<script lang="ts">
    import Select from "@/components/ui/Select.svelte";
    import type { Select as SelectPrimitive } from "bits-ui";
    import { fontData } from "@/static/fonts";
    import { graphe_settings } from "@/lib/stores";
    import type { SettingsValues } from "!/graphe/internal";
    import FontSelect from "./FontSelect.svelte";

    const language_info = [
        {
            name: "system",
            category: "english",
            text: "This is an example of some system text.",
        },
        {
            name: "greek",
            category: "greek",
            text: "ἐν ἀρχῇ ἐποίησεν ὁ θεὸς τὸν οὐρανὸν καὶ τὴν γῆν...",
        },
        {
            name: "hebrew",
            category: "hebrew",
            text: "בְּרֵאשִׁ֖ית בָּרָ֣א אֱלֹהִ֑ים אֵ֥ת הַשָּׁמַ֖יִם וְאֵ֥ת הָאָֽרֶץ׃",
        },
        {
            name: "english",
            category: "english",
            text: "In the beginning God created the heavens and the earth...",
        },
    ] as const;

    type Languages = (typeof language_info)[number]["name"];
    const language_selects: Map<
        Languages,
        {
            selected: SelectPrimitive.Props<string>["selected"];
            items: SelectPrimitive.Props<string>["items"];
        }
    > = new Map();

    function makeLanguageSelects(settings: SettingsValues | undefined) {
        language_selects.clear();
        if (settings == undefined) return;

        for (const language of language_info) {
            const available_fonts = fontData.filter(
                (f) => f.language == language.category,
            );
            const selected_font = settings.appearence.font[language.name];
            const selected_font_data = available_fonts.find(
                (f) => f.name == selected_font,
            );
            if (selected_font_data == undefined) {
                GrapheLog(
                    "error",
                    `[FontSelector] Invalid font (type: \`${language.name}\`, font: \`${selected_font}\`)`,
                );
                return;
            }

            language_selects.set(language.name, {
                selected: {
                    label: selected_font_data.name,
                    value: selected_font_data.css_line,
                },
                items: available_fonts.map((font) => ({
                    label: font.name,
                    value: font.css_line,
                })),
            });
        }
    }
    $: makeLanguageSelects($graphe_settings);

    function GrapheLog(arg0: string, arg1: string) {
        throw new Error("Function not implemented.");
    }
</script>

<div class="font-selector">
    {#each language_info as language}
        <div class="font-selection">
            <div class="wrapper">
                <div class="font-selection-heading">{language.name}</div>
                <div
                    class="font-selection-example"
                    data-lang={language.name.toLowerCase()}
                >
                    {language.text}
                </div>
            </div>

            {#if language_selects.has(language.name)}
                <FontSelect
                    name={language.name}
                    values={language_selects.get(language.name)}
                />
            {/if}
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
        text-transform: capitalize;
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
