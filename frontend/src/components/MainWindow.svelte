<script lang="ts">
    import {
        biblePointToString,
        createBiblePoint,
    } from "@/lib/data/bibleReference";

    const p1 = createBiblePoint("Matthew", 2);
    const r1 = biblePointToString(p1, "long");

    const p2 = createBiblePoint("Matthew", 28, "end");
    const r2 = biblePointToString(p2, "long");
    console.log(r1, r2);
</script>

main
