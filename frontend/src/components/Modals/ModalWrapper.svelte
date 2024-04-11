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
    .content {
        --size-modal-header-gap: 1rem;

        max-height: calc(
            100vh - 2 * var(--size-modal-margin) - 2 * var(--size-modal-padding) -
                var(--size-modal-header) - var(--size-modal-header-gap)
        );
        margin-top: var(--size-modal-header-gap);
        overflow: scroll;
    }
</style>
