<script lang="ts">
    import { GitBranch } from "lucide-svelte";
    import { GetEnvironmentInfo } from "!wails/go/app/App";
</script>

<div class="version-info">
    <GitBranch />
    {#await GetEnvironmentInfo()}
        0.0.0
    {:then environment_info}
        {environment_info.version}
    {/await}
</div>

<style>
    .version-info {
        position: absolute;
        bottom: 1rem;
        right: 1rem;
        display: flex;
        align-items: center;
        gap: 0.1rem;
        font-family: monospace;
        font-size: 0.8rem;
        color: var(--clr-text-sub);
        pointer-events: none;
        background: var(--clr-background);
        border-radius: 0.1rem;
        padding: 0.2rem;
        user-select: none;
        -webkit-user-select: none;
        pointer-events: none;
    }

    .version-info > :global(svg) {
        height: 1em;
    }
</style>
