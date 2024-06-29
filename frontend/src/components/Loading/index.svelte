<script lang="ts">
    import { EventsEmit } from "!wails/runtime/runtime";
    import { onMount } from "svelte";

    onMount(() => {
        // TODO: actually check stuff before progressing
        EventsEmit("graphe:mode", "settings");
    });
</script>

Loading...
