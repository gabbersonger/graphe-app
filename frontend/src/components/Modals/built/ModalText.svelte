<script lang="ts">
    import ModalResultView from "../ModalResultView.svelte";

    import { BookOpenText, StickyNote } from "lucide-svelte";
    import { createRef } from "@/lib/Scripture/ref";
    import { app_version } from "@/lib/appManager";
    import { EventsEmit } from "!wails/runtime/runtime";
    import { bibleData, versionData } from "@/lib/Scripture/data";
    import { onMount } from "svelte";
    import { getVersionBookIndex } from "@/lib/Scripture/version";

    type Mode = "book" | "chapter" | "verse";
    let mode: Mode = "book";

    let value = "";
    let search_results: { value: string | number; display: string }[] = [];
    $: search_results = filterResults(value);
    let selected_book = 0;
    let selected_chapter = 0;

    const possible_books: Map<string, number> = new Map();
    onMount(() => {
        for (let i = 0; i < versionData[$app_version].books.length; i++) {
            const book_number = versionData[$app_version].books[i].book_number;
            const book_data = bibleData[book_number - 1];

            possible_books.set(book_data.name.toLowerCase(), book_number);
            possible_books.set(book_data.short.toLowerCase(), book_number);
            for (let j = 0; j < book_data.abbreviations.length; j++) {
                possible_books.set(
                    book_data.abbreviations[j].toLowerCase(),
                    book_number,
                );
            }
        }
    });

    function updateMode(query: string): string {
        if (query.length == 0) return "";

        let capture_book = 0;
        let capture_end_index = 0;
        for (const [book_string, book_number] of possible_books.entries()) {
            const index = query.indexOf(book_string);
            if (index == 0 && index + book_string.length > capture_end_index) {
                capture_book = book_number;
                capture_end_index = index + book_string.length;
            }
        }

        if (capture_end_index > 0) {
            const post_capture = query.slice(capture_end_index).trim();
            if (post_capture.length == query.length - capture_end_index) {
                if (post_capture.length == 0) {
                    selected_book = capture_book;
                    mode = "chapter";
                    return "";
                } else {
                    mode = "book";
                    selected_book = 0;
                    return query;
                }
            }

            if (post_capture.includes(":") || post_capture.includes(".")) {
                const groups = post_capture.split(/[:\.]/);
                if (groups.length > 2) {
                    mode = "book";
                    return query;
                } else if (groups.length == 1) {
                    mode = "chapter";
                    return post_capture;
                }

                const potential_chapter = Number(groups[0]);
                if (isNaN(potential_chapter)) {
                    mode = "book";
                    return query;
                }

                const version_index = getVersionBookIndex(
                    $app_version,
                    selected_book,
                );
                if (
                    potential_chapter > 0 &&
                    potential_chapter <=
                        versionData[$app_version].books[version_index]
                            .num_chapters
                ) {
                    selected_chapter = potential_chapter;
                    mode = "verse";
                    return groups[1];
                }
            } else {
                mode = "chapter";
                return post_capture;
            }
        }
        mode = "book";
        return query;
    }

    function normaliseString(string: string): string {
        return string.toLowerCase();
    }

    const search_books = versionData[$app_version].books.map((b) => ({
        book_number: b.book_number,
        string:
            normaliseString(bibleData[b.book_number - 1].name) +
            "|" +
            normaliseString(bibleData[b.book_number - 1].short) +
            "|" +
            bibleData[b.book_number - 1].abbreviations
                .map((a: string) => normaliseString(a))
                .join("|"),
    }));

    function createNumberResults(num: number) {
        let results = ["Back"];
        for (let i = 0; i < num; i++) {
            results.push(`${i + 1}`);
        }
        return results;
    }

    function filterResults(query: string) {
        query = query.trim().toLowerCase();
        query = updateMode(query);

        if (mode == "book") {
            return search_books
                .map((x) => ({ ...x, index: x.string.indexOf(query) }))
                .filter((x) => x.index >= 0)
                .sort((a, b) => a.index - b.index)
                .map((x) => ({
                    value: x.book_number,
                    display: bibleData[x.book_number - 1].name,
                }));
        } else if (mode == "chapter") {
            const version_index = getVersionBookIndex(
                $app_version,
                selected_book,
            );
            const n_chapters =
                versionData[$app_version].books[version_index].num_chapters;
            return createNumberResults(n_chapters)
                .map((x) => ({ item: x, index: x.indexOf(query) }))
                .filter((x) => x.index >= 0)
                .sort((a, b) => a.index - b.index)
                .map((x) => ({
                    value: x.item,
                    display:
                        x.item == "Back"
                            ? x.item
                            : `${bibleData[selected_book - 1].name} ${x.item}`,
                }));
        } else if (mode == "verse") {
            const version_index = getVersionBookIndex(
                $app_version,
                selected_book,
            );
            const num_verses =
                versionData[$app_version].books[version_index].num_verses[
                    selected_chapter - 1
                ];

            return createNumberResults(num_verses)
                .map((x) => ({ item: x, index: x.indexOf(query) }))
                .filter((x) => x.index >= 0)
                .sort((a, b) => a.index - b.index)
                .map((x) => ({
                    value: x.item,
                    display:
                        x.item == "Back"
                            ? x.item
                            : `${bibleData[selected_book - 1].name} ${selected_chapter}:${x.item}`,
                }));
        }
    }

    function chooseResult(index: number) {
        if (search_results[index].value == "Back") {
            if (mode == "chapter") {
                value = "";
                mode = "book";
                selected_book = 0;
                selected_chapter = 0;
                search_results = filterResults(value);
            } else if (mode == "verse") {
                value = "";
                mode = "chapter";
                selected_chapter = 0;
                search_results = filterResults(value);
            }
            return;
        }

        if (mode == "book") {
            selected_book = search_results[index].value as number;
            mode = "chapter";
            value = "";
            search_results = filterResults(value);
        } else if (mode == "chapter") {
            const chapter = search_results[index].value as number;
            const ref = createRef($app_version, selected_book, chapter);
            EventsEmit("app:goto", ref);
            EventsEmit("ui:modal:close");
        } else if (mode == "verse") {
            const verse = search_results[index].value as number;
            const ref = createRef(
                $app_version,
                selected_book,
                selected_chapter,
                verse,
            );
            EventsEmit("app:goto", ref);
            EventsEmit("ui:modal:close");
        }
    }
</script>

<ModalResultView
    icon={mode == "book" ? BookOpenText : StickyNote}
    placeholder="Choose a Passage"
    bind:value
    results={search_results}
    {chooseResult}
    noResults="We couldn't find a passage that matches your search"
/>
