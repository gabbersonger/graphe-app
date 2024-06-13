<script lang="ts">
    import ModalResultView from "../ModalResultView.svelte";

    import { BookOpenText, StickyNote } from "lucide-svelte";
    import { createRef, isValidRef } from "@/lib/Scripture/ref";
    import { app_version } from "@/lib/managers/appManager";
    import { EventsEmit } from "!wails/runtime/runtime";
    import { bibleData, versionData } from "@/lib/Scripture/data";
    import { getVersionBookIndex } from "@/lib/Scripture/version";
    import type { BibleVersion } from "@/lib/Scripture/types";

    type Mode = "book" | "chapter" | "verse";
    let mode: Mode = "book";

    let value = "";
    let search_results: { value: string | number; display: string }[] = [];
    $: search_results = filterResults(value);
    let selected_book = 0;
    let selected_chapter = 0;

    const possible_books: Map<string, number> = new Map();
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

    function updateMode(query: string): string {
        if (query.length == 0) {
            return "";
        }

        // Get the longest book prefix and where it ends
        let capture_book = 0;
        let capture_end_index = 0;
        for (const [book_string, book_number] of possible_books.entries()) {
            const index = query.indexOf(book_string);
            if (index == 0 && index + book_string.length > capture_end_index) {
                capture_book = book_number;
                capture_end_index = index + book_string.length;
            }
        }

        // If there was a book prefix...
        if (capture_end_index > 0) {
            const post_capture = query.slice(capture_end_index).trim();
            const groups = post_capture.split(/[:\.]/);

            if (post_capture.length == query.length - capture_end_index) {
                if (post_capture.length == 0) {
                    selected_book = capture_book;
                    mode = "chapter";
                    return "";
                }
            } else if (groups.length == 1) {
                mode = "chapter";
                return post_capture;
            } else if (groups.length == 2) {
                const potential_chapter = Number(groups[0]);
                const book_data =
                    versionData[$app_version].books[
                        getVersionBookIndex($app_version, selected_book)
                    ];
                if (
                    !isNaN(potential_chapter) &&
                    potential_chapter > 0 &&
                    potential_chapter <= book_data.num_chapters &&
                    book_data.num_verses[potential_chapter - 1] != 0
                ) {
                    selected_chapter = potential_chapter;
                    mode = "verse";
                    return groups[1];
                }
            }
        }

        mode = "book";
        selected_book = 0;
        return query;
    }

    function normaliseString(string: string): string {
        return string.trim().toLowerCase();
    }

    const SEARCH_BOOKS = versionData[$app_version].books.map(
        (b: (typeof versionData)[BibleVersion]["books"][number]) => ({
            book_number: b.book_number,
            string:
                normaliseString(bibleData[b.book_number - 1].name) +
                "|" +
                normaliseString(bibleData[b.book_number - 1].short) +
                "|" +
                bibleData[b.book_number - 1].abbreviations
                    .map((a: string) => normaliseString(a))
                    .join("|"),
        }),
    );

    type ResultData = { value: string | number; display: string };

    function createBookResults(query: string): ResultData[] {
        let temp_results: (ResultData & { query_location: number })[] = [];
        for (let i = 0; i < SEARCH_BOOKS.length; i++) {
            const book_num = SEARCH_BOOKS[i].book_number;
            const book_string = SEARCH_BOOKS[i].string;
            const found_at = book_string.indexOf(query);
            if (found_at < 0) continue;
            temp_results.push({
                value: book_num,
                display: bibleData[book_num - 1].name,
                query_location: found_at,
            });
        }
        temp_results.sort((a, b) => a.query_location - b.query_location);
        let results: ResultData[] = [];
        for (let i = 0; i < temp_results.length; i++)
            results.push({
                value: temp_results[i].value,
                display: temp_results[i].display,
            });
        return results;
    }

    function createChapterResults(query: string): ResultData[] {
        let results: ResultData[] = [];
        if (query.length == 0) results.push({ value: "Back", display: "Back" });
        const v_b_index = getVersionBookIndex($app_version, selected_book);
        const book_data = versionData[$app_version].books[v_b_index];
        for (let i = 0; i < book_data.num_chapters; i++) {
            if (book_data.num_verses[i] == 0) continue;
            if (`${i + 1}`.indexOf(query) < 0) continue;
            results.push({
                value: i + 1,
                display: `${bibleData[selected_book - 1].name} ${i + 1}`,
            });
        }
        return results;
    }

    function createVerseResults(query: string): ResultData[] {
        let results: ResultData[] = [];
        if (query.length == 0) results.push({ value: "Back", display: "Back" });

        const v_b_index = getVersionBookIndex($app_version, selected_book);
        const book_data = versionData[$app_version].books[v_b_index];
        const ref_start =
            String(selected_book) + String(selected_chapter).padStart(3, "0");
        for (let i = 0; i < book_data.num_verses[selected_chapter - 1]; i++) {
            const possible_ref = parseInt(
                ref_start + String(i + 1).padStart(3, "0"),
            );
            if (!isValidRef($app_version, possible_ref)) continue;
            if (`${i + 1}`.indexOf(query) < 0) continue;
            results.push({
                value: i + 1,
                display: `${bibleData[selected_book - 1].name} ${selected_chapter}:${i + 1}`,
            });
        }
        return results;
    }

    function filterResults(query: string) {
        query = normaliseString(query);
        query = updateMode(query);

        if (mode == "book") return createBookResults(query);
        else if (mode == "chapter") return createChapterResults(query);
        else if (mode == "verse") return createVerseResults(query);
    }

    function goto(
        book: number,
        chapter: number,
        verse: number | "start" = "start",
    ) {
        const ref = createRef($app_version, book, chapter, verse);
        EventsEmit("app:goto", ref);
        EventsEmit("app:modal:close");
    }

    function chooseResult(index: number) {
        // Deal with clicking back button
        if (search_results[index].value == "Back") {
            if (mode == "verse") {
                mode = "chapter";
                selected_chapter = 0;
            } else {
                mode = "book";
                selected_book = 0;
                selected_chapter = 0;
            }
            if (value == "") search_results = filterResults("");
            else value = "";
            return;
        }

        // Deal with all other buttons
        const result = search_results[index].value as number;
        if (mode == "book") {
            selected_book = result;
            mode = "chapter";
            value = "";
            search_results = filterResults(value);
        } else if (mode == "chapter") {
            goto(selected_book, result);
        } else if (mode == "verse") {
            goto(selected_book, selected_chapter, result);
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
