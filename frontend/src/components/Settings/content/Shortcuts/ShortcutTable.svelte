<script lang="ts">
    import Input from "@/components/ui/Input.svelte";
    import Button from "@/components/ui/Button.svelte";
    import ShortcutButton from "@/components/Settings/content/Shortcuts/ShortcutButton.svelte";
    import { TextSearch, RotateCcw, MoveDown, Lock } from "lucide-svelte";
    import { shortcutsData } from "@/components/Settings/content/Shortcuts/data";
    import { graphe_settings } from "@/lib/stores";
    import type { SettingsValues } from "!/graphe/internal";
    import { GrapheEvent } from "@/lib/utils";

    // Limit the height of the table
    let fullHeight: number, fullWidth: number;
    let tableBody: HTMLDivElement;
    $: if (tableBody && fullHeight > 0 && fullWidth > 0) {
        const offset = tableBody.getBoundingClientRect().top;
        tableBody.style.maxHeight = `${fullHeight - offset - 56}px`;
    }

    // Display a tag if there is more to scroll
    let fullyScrolled = false;
    function checkScroll() {
        if (tableBody != undefined)
            fullyScrolled =
                tableBody.scrollHeight <=
                tableBody.scrollTop + tableBody.clientHeight;
        else fullyScrolled = true;
    }
    $: if (fullHeight > 0 || fullWidth > 0 || filteredShortcuts.length >= 0)
        checkScroll();

    // Reset all the shortcuts
    function resetAllShortcuts() {
        GrapheEvent("graphe:setting:reset", ["shortcuts"]);
    }

    // Reset invidual shortcut
    function resetShortcut(shortcut: string) {
        GrapheEvent("graphe:setting:reset", ["shortcuts", shortcut]);
    }

    // Change the value for a shortcut
    function onShortcutChange(shortcut: string, value: string) {
        GrapheEvent("graphe:setting", {
            setting: ["shortcuts", shortcut],
            value: value,
        });
    }

    // Filter the displayed shortcuts
    let search: string = "";
    $: filteredShortcuts = filterShortcuts(search, $graphe_settings);
    function filterShortcuts(
        searchString: string,
        data: SettingsValues | undefined,
    ) {
        if (data == undefined) return [];
        return shortcutsData
            .filter((s) => {
                const a = s.description.toLowerCase().replaceAll(" ", "");
                const b = searchString.toLowerCase().replaceAll(" ", "");
                return a.includes(b);
            })
            .map((s) => {
                if ("value" in s) {
                    return { ...s, field: "", locked: true };
                } else {
                    return {
                        ...s,
                        value: data.shortcuts[s.field],
                        locked: false,
                    };
                }
            });
    }
</script>

<svelte:window bind:innerHeight={fullHeight} bind:innerWidth={fullWidth} />

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
    <div
        class="table-body"
        bind:this={tableBody}
        on:scroll={() => checkScroll()}
    >
        <table>
            {#each filteredShortcuts as shortcut}
                <tr>
                    <td>{shortcut.description}</td>
                    <td>
                        <ShortcutButton
                            shortcut={shortcut.value}
                            locked={shortcut.locked}
                            onChange={(v) =>
                                onShortcutChange(shortcut.field, v)}
                            onReset={() => resetShortcut(shortcut.field)}
                        />
                    </td>
                </tr>
            {:else}
                <div class="no-results">No results found</div>
            {/each}
        </table>
    </div>

    {#if !fullyScrolled}
        <div class="scroll-indicator"><MoveDown />Scroll For More</div>
    {/if}
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

    .table .scroll-indicator {
        position: absolute;
        bottom: 0;
        left: calc(50%);
        transform: translateX(-50%) translateY(50%);
        display: flex;
        flex-direction: row;
        align-items: center;
        gap: 0.5rem;
        background: var(--clr-background-sub);
        color: var(--clr-text);
        font-family: var(--font-system);
        font-size: 0.7rem;
        padding: 0.4rem;
        border: 1px solid var(--clr-background-dark);
        border-radius: 0.4rem;
    }

    .table .scroll-indicator > :global(svg) {
        height: 1em;
        width: 1em;
    }

    .no-results {
        width: 100%;
        display: flex;
        align-items: center;
        justify-content: center;
        padding: 2rem 1rem;
        background: var(--clr-background-sub);
        color: var(--clr-text-sub);
        font-size: 0.8rem;
    }
</style>
