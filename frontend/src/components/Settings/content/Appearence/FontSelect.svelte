<script lang="ts">
    import Select from "@/components/ui/Select.svelte";
    import type { Select as SelectPrimitive } from "bits-ui";
    import { GrapheEvent } from "@/lib/utils";

    export let name: string;
    export let values:
        | {
              selected: SelectPrimitive.Props<string>["selected"];
              items: SelectPrimitive.Props<string>["items"];
          }
        | undefined;

    function onFontChange(font: string | undefined) {
        if (font != undefined) {
            GrapheEvent("graphe:setting", {
                setting: ["appearence", "font", name],
                value: font,
            });
        }
    }
</script>

{#if values != undefined}
    <Select
        bind:selected={values.selected}
        onSelectedChange={(v) => onFontChange(v?.label)}
        items={values.items}
        placeholder="Choose a font"
        label="Font Family"
    />
{/if}
