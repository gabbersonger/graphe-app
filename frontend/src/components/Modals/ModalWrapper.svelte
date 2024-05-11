<script lang="ts">
    let headerHeight: number;
</script>

<div class="wrapper">
    <div class="header" bind:clientHeight={headerHeight}>
        <slot name="header" />
    </div>
    <div class="content" style:--size-modal-header={`${headerHeight}px`}>
        <slot name="content" />
    </div>
</div>

<style>
    .wrapper {
        --modal-wrapper-scroll-width: 10px;

        position: relative;
    }

    .content {
        max-height: calc(
            100vh - 2 * var(--size-modal-margin) - var(--size-modal-header)
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
</style>
