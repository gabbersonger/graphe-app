<script lang="ts">
    export let text: string;
</script>

<div class="command">
    {#each text as letter}
        <span>{letter}</span>
    {/each}
</div>

<style>
    .command {
        display: inline;
        padding: 0.1rem 0.2rem;
        margin-inline: 0.1rem;

        background: var(--clr-background-sub);
        border: 1px solid var(--clr-background-dark);
        border-radius: 0.1rem;

        font-family: var(--font-system);
        font-size: 0.7rem;
        color: var(--clr-text-sub);
    }

    .command span:not(:last-child) {
        margin-right: 0.2rem;
    }
</style>
