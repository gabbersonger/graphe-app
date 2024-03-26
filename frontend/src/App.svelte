<script>
    import { GetScriptureSections, GetScriptureWord } from "!wails/app/App";

    async function b() {
        let ranges = [];
        ranges.push({
            start: 40011001,
            end: 40011999,
        });
        let start = performance.now();
        let scripture = await GetScriptureSections("esv", ranges);
        let end = performance.now();
        console.log(scripture);
        console.log(`Took: ${end - start}ms`);
    }

    async function a() {
        let start = performance.now();
        let res = await GetScriptureWord(40004006, 26);
        let end = performance.now();
        console.log(res);
        console.log(`Took: ${end - start}ms`);
    }

    b();
</script>

<div class="app">
    <br /><br />
    <button on:click={() => a()}>test</button>
</div>

<style>
</style>
