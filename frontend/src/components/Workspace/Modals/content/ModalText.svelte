<script lang="ts">
    import ModalResultView from "@/components/Workspace/Modals/ModalResultView.svelte";
    import { BookOpenText, StickyNote } from "lucide-svelte";

    import { workspace_version } from "@/lib/stores";
    import {
        BookData,
        ScriptureService,
        type VersionData,
    } from "!/graphe/internal/scripture";
    import { GrapheEvent, GrapheLog } from "@/lib/utils";

    let initialised = false;
    let current_version: string | undefined = undefined;
    $: init($workspace_version);

    let version_data: VersionData;
    let bible_data: BookData[];
    let possible_books: Map<string, number> = new Map();
    let search_books: Array<{
        book_number: number;
        search_string: string;
    }> = [];

    async function init(version: string | undefined) {
        if (current_version == version) return;
        current_version = version;
        initialised = false;
        if (version == undefined) return;

        let data: any = await ScriptureService.GetVersionData(version);
        version_data = data as VersionData;

        data = await ScriptureService.GetBibleData();
        bible_data = data as BookData[];

        possible_books.clear();
        search_books = [];

        for (let i = 0; i < version_data.books.length; i++) {
            const book_number = version_data.books[i].book_number;
            const book_data = bible_data[book_number - 1];
            if (!book_data)
                GrapheLog(
                    "error",
                    `[ModalText] Invalid book number found in version data (version: \`${current_version}\`, book_number: ${book_number}`,
                );

            possible_books.set(book_data.name.toLowerCase(), book_number);
            possible_books.set(book_data.short.toLowerCase(), book_number);
            for (let j = 0; j < book_data.abbreviations.length; j++) {
                possible_books.set(
                    book_data.abbreviations[j].toLowerCase(),
                    book_number,
                );
            }

            search_books.push({
                book_number: book_number,
                search_string:
                    normaliseString(book_data.name) +
                    "|" +
                    normaliseString(book_data.short) +
                    "|" +
                    book_data.abbreviations
                        .map((a: string) => normaliseString(a))
                        .join("|"),
            });
        }

        initialised = true;
        search_results = filterResults(value);
    }

    // DISPLAYING RESULTS
    type Mode = "book" | "chapter" | "verse";
    type ResultData = { value: string | number; display: string };

    let mode: Mode = "book";
    let value = "";
    let search_results: ResultData[] = [];
    $: search_results = filterResults(value);

    function filterResults(query: string): ResultData[] {
        query = normaliseString(query);
        query = updateMode(query);

        if (mode == "book") return createBookResults(query);
        else if (mode == "chapter") return createChapterResults(query);
        else if (mode == "verse") return createVerseResults(query);
        return [];
    }

    function normaliseString(string: string): string {
        return string.trim().toLowerCase();
    }

    function updateMode(query: string): string {
        if (query.length == 0 || version_data == undefined) {
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
                const book_data = version_data.books.find(
                    (b) => b.book_number == selected_book,
                );
                if (!book_data) return "";
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

    function createBookResults(query: string): ResultData[] {
        let temp_results: (ResultData & { query_location: number })[] = [];
        for (let i = 0; i < search_books.length; i++) {
            const found_at = search_books[i].search_string.indexOf(query);
            if (found_at < 0) continue;
            temp_results.push({
                value: search_books[i].book_number,
                display: bible_data[search_books[i].book_number - 1].name,
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

        const book_data = version_data.books.find(
            (vb) => vb.book_number == selected_book,
        );
        if (book_data == undefined) {
            GrapheLog(
                "error",
                `[ModalText] Invalid book number selected (version: \`${current_version}\` selected_book_number: ${selected_book})`,
            );
            return [];
        }

        for (let i = 0; i < book_data.num_chapters; i++) {
            if (book_data.num_verses[i] == 0) continue;
            if (`${i + 1}`.indexOf(query) < 0) continue;
            results.push({
                value: i + 1,
                display: `${bible_data[selected_book - 1].name} ${i + 1}`,
            });
        }
        return results;
    }

    function createVerseResults(query: string): ResultData[] {
        let results: ResultData[] = [];
        if (query.length == 0) results.push({ value: "Back", display: "Back" });

        const book_data = version_data.books.find(
            (vb) => vb.book_number == selected_book,
        );
        if (book_data == undefined) {
            GrapheLog(
                "error",
                `[ModalText] Invalid book number selected (version: \`${current_version}\` selected_book_number: ${selected_book})`,
            );
            return [];
        }

        const ref_start =
            String(selected_book) + String(selected_chapter).padStart(3, "0");
        for (let i = 0; i < book_data.num_verses[selected_chapter - 1]; i++) {
            const possible_ref = parseInt(
                ref_start + String(i + 1).padStart(3, "0"),
            );
            const is_valid = ScriptureService.IsRefValid(
                possible_ref,
                current_version ?? "",
            ); // FIX: this only works because IsRefValid is fast (no await)
            if (!is_valid) continue;
            if (`${i + 1}`.indexOf(query) < 0) continue;
            results.push({
                value: i + 1,
                display: `${bible_data[selected_book - 1].name} ${selected_chapter}:${i + 1}`,
            });
        }
        return results;
    }

    // CHOOSING BOOK
    let selected_book = 0;
    let selected_chapter = 0;
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
            goto(selected_book, result, 1);
            // TODO: fix the above line to go to the first verse of the chapter
        } else if (mode == "verse") {
            goto(selected_book, selected_chapter, result);
        }
    }

    async function goto(book: number, chapter: number, verse: number) {
        const ref = await ScriptureService.CreateRef(book, chapter, verse);
        if ($workspace_version == undefined) return;
        const valid = await ScriptureService.IsRefValid(
            ref,
            $workspace_version,
        );
        if (!valid)
            return GrapheLog(
                "error",
                `[ModalText] Invalid reference made (ref: ${ref}, version: \`${$workspace_version}\`)`,
            );
        GrapheEvent("window:workspace:goto", ref);
        GrapheEvent("window:workspace:modal:close");
    }
</script>

{#if initialised}
    <ModalResultView
        icon={mode == "book" ? BookOpenText : StickyNote}
        placeholder="Choose a Passage"
        bind:value
        results={search_results}
        {chooseResult}
        noResults="We couldn't find a passage that matches your search"
    />
{/if}
