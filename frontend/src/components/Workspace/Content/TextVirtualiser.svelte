<script lang="ts">
    import type { ScriptureSection } from "!/graphe/internal/data";
    import type { BibleRef } from "@/lib/Scripture/types";

    export let data: Array<ScriptureSection>;
    export let current_verse: BibleRef | undefined = undefined;

    $: console.log(data);
    $: console.log(current_verse);
</script>

hi
