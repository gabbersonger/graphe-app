<script lang="ts">
    import { modalData } from "@/components/Modals/data";
    import { app_modal } from "@/lib/managers/appManager";
    import { onMount } from "svelte";

    onMount(() => {
        function handleKeyup(event: KeyboardEvent) {
            if ($app_modal && event.key == "Escape") {
                $app_modal = "";
            }
        }
        document.addEventListener("keyup", handleKeyup);
        return () => {
            document.removeEventListener("keyup", handleKeyup);
        };
    });
</script>

<div class="screen-wrapper" class:open={$app_modal != ""}>
    <!-- svelte-ignore a11y-click-events-have-key-events -->
    <div class="overlay" on:click={() => ($app_modal = "")}></div>

    <div class="modal">
        <div class="container">
            {#each modalData as m}
                {#if $app_modal == m.name}
                    <svelte:component this={m.modal} />
                {/if}
            {/each}
        </div>
    </div>
</div>

<style>
    .screen-wrapper {
        --size-modal-max-width: 500px;
        --size-modal-margin: 15vh;

        position: absolute;
        inset: 0;
        z-index: 2;
        isolation: isolate;
        display: flex;
        flex-direction: column;
        justify-content: flex-start;
        align-items: center;
        padding: var(--size-modal-margin) 0;
    }

    .screen-wrapper:not(.open) {
        display: none;
    }

    .overlay {
        position: absolute;
        inset: 0;
        background-color: rgba(0, 0, 0, 0.4);
        /* NOTE: blur applied to background in App.svelte */
    }

    .modal {
        position: relative;
        width: 90%;
        max-width: var(--size-modal-max-width);
        max-height: calc(100vh - 2 * var(--size-modal-margin));
        background: var(--clr-background);
        color: var(--clr-text);
        border-radius: 0.5rem;
        overflow: hidden;
    }

    .modal .container {
        position: relative;
        width: 100%;
        max-height: 100%;
        overflow: scroll;
    }
</style>
