<div class="container">
    <div class="wrapper">
        <div class="max-width">
            <slot />
        </div>
    </div>
</div>

<style>
    .container {
        --size-max-width: 80ch;
        position: relative;
        height: 100%;
        width: 100%;
        overflow: scroll;
    }

    .wrapper {
        position: absolute;
        inset: 0;
        overflow: scroll;
    }

    .max-width {
        position: relative;
        width: 90%;
        max-width: var(--size-max-width);
        margin: 0 auto;
        padding: 1rem 0;
    }
</style>
