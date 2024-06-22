<script lang="ts">
    import { settings_section } from "@/lib/stores";
    import { settingsData } from "@/components/Settings/data";
</script>

{#each settingsData as s}
    {#if $settings_section == s.name}
        <svelte:component this={s.content} />
    {/if}
{/each}
