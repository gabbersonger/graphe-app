<script lang="ts">
    import { GetScriptureSections } from "!wails/app/App";
    import {
        biblePointToRef,
        createBiblePointAtChapterStart,
        createBiblePointAtChapterEnd,
    } from "@/lib/data/bibleReference";

    async function onStartup() {
        let matt20 = createBiblePointAtChapterStart("Matthew", 20);
        let matt25 = createBiblePointAtChapterEnd("Matthew", 25);
        let rom3 = createBiblePointAtChapterStart("Romans", 3);
        let rom5 = createBiblePointAtChapterEnd("Romans", 5);

        let data = await GetScriptureSections("gnt", [
            { start: biblePointToRef(matt20), end: biblePointToRef(matt25) },
            { start: biblePointToRef(rom3), end: biblePointToRef(rom5) },
        ]);
        console.log(data);
    }
    onStartup();
</script>

<div class="app"></div>
