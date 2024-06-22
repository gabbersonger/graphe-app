<script lang="ts">
    let headerHeight: number = 0;
    let footerHeight: number = 0;

    export let contentElement: HTMLDivElement;
</script>

<div class="wrapper">
    <div class="header" bind:clientHeight={headerHeight}>
        <slot name="header" />
    </div>
    <div
        class="content"
        class:rounded={"footer" in $$slots}
        style:--size-modal-header={`${headerHeight}px`}
        style:--size-modal-footer={`${footerHeight}px`}
        bind:this={contentElement}
    >
        <slot name="content" {contentElement} />
    </div>
    {#if "footer" in $$slots}
        <div class="footer" bind:clientHeight={footerHeight}>
            <slot name="footer" />
        </div>
    {/if}
</div>

<style>
    .wrapper {
        --modal-wrapper-scroll-width: 10px;

        position: relative;
        background: var(--clr-background-sub);
    }

    .header,
    .content {
        background: var(--clr-background);
    }

    .content {
        max-height: calc(
            100vh - 2 * var(--size-modal-margin) - var(--size-modal-header) -
                var(--size-modal-footer)
        );
        overflow-x: hidden;
        overflow-y: scroll;
    }

    .content::-webkit-scrollbar {
        width: var(--modal-wrapper-scroll-width);
    }

    .content::-webkit-scrollbar-corner {
        background: transparent;
    }

    .content::-webkit-scrollbar-thumb {
        background: var(--clr-text-sub);
        -webkit-transition: 0.125s;
        transition: 0.125s;
        border-radius: 4px !important;
    }

    .content::-webkit-scrollbar-track {
        background: 0 0;
    }

    .content.rounded {
        border-bottom-left-radius: 0.5rem;
        border-bottom-right-radius: 0.5rem;
    }
</style>
