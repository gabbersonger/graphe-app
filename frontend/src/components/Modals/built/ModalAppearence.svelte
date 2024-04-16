<script>
    import { EventsEmit } from "!wails/runtime/runtime";
    import { themeData } from "@/static/themes";
</script>

{#each themeData as theme}
    <button on:click={() => EventsEmit("ui:theme", theme.name)}>
        Theme: {theme.name}
    </button>
    <br />
{/each}
