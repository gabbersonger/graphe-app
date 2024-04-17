<script lang="ts">
    import { EventsEmit } from "!wails/runtime/runtime";

    let val: number;
</script>

Choose Version

<button on:click={() => EventsEmit("app:version", "gnt")}>gnt</button>
<button on:click={() => EventsEmit("app:version", "hot")}>hot</button>
<button on:click={() => EventsEmit("app:version", "lxx")}>lxx</button>
<button on:click={() => EventsEmit("app:version", "esv")}>esv</button>

<br /><br />
<input type="number" bind:value={val} />
<button on:click={() => EventsEmit("app:goto", val)}>goto ref</button>
