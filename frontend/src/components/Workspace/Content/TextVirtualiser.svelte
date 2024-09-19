<script lang="ts">
    import type { ScriptureSection } from "!/graphe/internal/data";
    import type { ScriptureRef } from "!/graphe/internal/scripture";

    export let data: Array<ScriptureSection>;
    export let current_verse: ScriptureRef | undefined = undefined;

    $: console.log(data);
    $: console.log(current_verse);
</script>

hi
