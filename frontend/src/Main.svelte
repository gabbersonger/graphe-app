<script lang="ts">
    import { DataDB } from "!wails/internal/data";

    async function Init() {
        const a = await DataDB.GetScriptureSection({
            version: "esv",
            start: 1_001_001,
            end: 1_002_001,
        });
        return a;
    }
    console.log(Init());
</script>

hi
