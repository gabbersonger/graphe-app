<script lang="ts">
    import Input from "@/components/ui/Input.svelte";
    import ModalWrapper from "@/components/Modals/ModalWrapper.svelte";
    import ModalButtons from "@/components/Modals/ModalButtons.svelte";
    import { BookOpenText } from "lucide-svelte";

    import { bibleData, versionData } from "@/lib/Scripture/data";
    import { app_version } from "@/lib/appManager";
    import { createBibleRef } from "@/lib/Scripture/ref";
    import { EventsEmit } from "!wails/runtime/runtime";

    let value: string = "";
    let isInputFocused = true;

    let selected_book: number = -1;

    $: available_books = versionData[$app_version].books.map((x) => ({
        book_number: x.book_number,
        ...bibleData[x.book_number - 1],
    }));

    function goto(book: number, chapter: number) {
        const ref = createBibleRef(bibleData[book - 1].name, chapter);
        EventsEmit("app:goto", ref);
        EventsEmit("ui:modal:close");
    }

    function clickBackButton() {
        selected_book = -1;
    }

    function clickBook(index: number) {
        selected_book = available_books[index].book_number;
        if (bibleData[selected_book - 1].num_chapters == 1) {
            goto(selected_book, 1);
        }
    }

    function clickChapter(index: number) {
        goto(selected_book, index + 1);
    }
</script>

<ModalWrapper>
    <div slot="header">
        <Input
            bind:value
            bind:focus={isInputFocused}
            icon={BookOpenText}
            placeholder="Choose a Passage"
        />
    </div>

    <div slot="content">
        {#if selected_book == -1}
            <ModalButtons
                items={available_books}
                rowData={{ number: 6, maxwidth: 7.2 }}
                onItemClick={clickBook}
                subheading={(index) => available_books[index].name}
                heading={(index) => available_books[index].abbreviation}
            />
        {:else}
            <ModalButtons
                backButton={clickBackButton}
                items={Array(bibleData[selected_book - 1].num_chapters)}
                rowData={{ number: 6, maxwidth: 7.2 }}
                onItemClick={clickChapter}
                subheading={(index) =>
                    `${bibleData[selected_book - 1].abbreviation} ${index + 1}`}
                heading={(index) => (index + 1).toString()}
            />
        {/if}
    </div>
</ModalWrapper>
