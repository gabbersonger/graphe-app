<script lang="ts">
    import Input from "@/components/ui/Input.svelte";
    import Button from "@/components/ui/Button.svelte";
    import { TextSearch, RotateCcw } from "lucide-svelte";
    import ShortcutButton from "./ShortcutButton.svelte";

    let search: string;

    let fullHeight: number;
    let tableBody: HTMLDivElement;
    $: if (tableBody && fullHeight > 0) {
        const offset = tableBody.getBoundingClientRect().top;
        tableBody.style.maxHeight = `${fullHeight - offset - 56}px`;
    }

    function resetAllShortcuts() {
        console.log("TODO");
    }
</script>

<svelte:window bind:innerHeight={fullHeight} />

<div class="item-wrapper">
    <Input
        placeholder="Search for a setting"
        icon={TextSearch}
        bind:value={search}
    />

    <Button on:click={() => resetAllShortcuts()} icon={RotateCcw}>
        Reset All
    </Button>
</div>

<div class="table">
    <table>
        <tr>
            <th>Description</th>
            <th>Shortcut</th>
        </tr>
    </table>
    <div class="table-body" bind:this={tableBody}>
        <table>
            {#each Array(12) as a}
                <tr>
                    <td>Open the settings window</td>
                    <td><ShortcutButton shortcut="⌘4" /></td>
                </tr>
            {/each}
        </table>
    </div>
</div>

<style>
    .item-wrapper {
        display: flex;
        flex-wrap: wrap-reverse;
        flex-direction: row;
        justify-content: space-between;
        align-items: center;
        gap: 0.6rem;
        padding-bottom: 1rem;
    }

    .table {
        position: relative;
        width: 100%;
        border: 1px solid var(--clr-background-dark);
        border-radius: 0.2rem;
    }

    .table .table-body {
        overflow: scroll;
    }

    .table table {
        width: 100%;
        text-align: left;
        font-family: var(--font-system);
        font-weight: normal;
        border-collapse: collapse;
    }

    .table table th,
    .table table td {
        height: 2.5rem;
        padding: 0.5rem;
        border-bottom: 1px solid var(--clr-background-dark);
    }

    .table tr:last-child td {
        border-bottom: none;
    }

    .table th:nth-child(2),
    .table td:nth-child(2) {
        width: 100px;
    }

    .table th {
        font-weight: normal;
        font-size: 0.7rem;
        color: var(--clr-text-sub);
    }

    .table td {
        font-size: 0.9rem;
        color: var(--clr-text);
    }

    .table .table-body tr:hover {
        background: var(--clr-background-sub);
    }
</style>
