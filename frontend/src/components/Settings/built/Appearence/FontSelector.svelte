<script>
    import Input from "@/components/ui/Input.svelte";

    const options = [
        {
            name: "System",
            text: "This is an example of some system text.",
        },
        {
            name: "Greek",
            text: "ἐν ἀρχῇ ἐποίησεν ὁ θεὸς τὸν οὐρανὸν καὶ τὴν γῆν...",
        },
        {
            name: "Hebrew",
            text: "בְּרֵאשִׁ֖ית בָּרָ֣א אֱלֹהִ֑ים אֵ֥ת הַשָּׁמַ֖יִם וְאֵ֥ת הָאָֽרֶץ׃",
        },
        {
            name: "English",
            text: "In the beginning God created the heavens and the earth...",
        },
    ];
</script>

<div class="font-selector">
    {#each options as option}
        <div class="font-selection">
            <div class="wrapper">
                <div class="font-selection-heading">{option.name}</div>
                <div class="font-selection-example" data-lang="hebrew">
                    {option.text}
                </div>
            </div>

            <Input />
        </div>
    {/each}
</div>

<style>
    .font-selector {
        display: flex;
        flex-direction: column;
        gap: 0.7rem;
    }

    .font-selection {
        display: flex;
        flex-direction: row;
        align-items: center;
        justify-content: space-between;
        gap: 0.5rem;
        border: 1px solid var(--clr-background-dark);
        border-radius: 0.2rem;
        padding: 0.5rem;
        font-size: 0.8rem;
    }

    .font-selection-heading {
        font-weight: 500;
        color: var(--clr-text);
        padding-bottom: 0.2rem;
    }

    .font-selection-example {
        color: var(--clr-text-sub);
    }
</style>
