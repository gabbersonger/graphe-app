<script lang="ts">
    import SettingsSidebar from "@/components/Settings/SettingsSidebar.svelte";
    import SettingsContent from "@/components/Settings/SettingsContent.svelte";
    import { X } from "lucide-svelte";
    import { Events } from "@wailsio/runtime";
    import { windowSettingsManager } from "@/lib/managers/window_settings";

    function closeSettings() {
        Events.Emit({ name: "graphe:mode", data: "workspace" });
    }
</script>

<div class="settings" use:windowSettingsManager>
    <nav>
        <div class="wrapper">
            <SettingsSidebar />
        </div>
    </nav>
    <main>
        <div class="wrapper">
            <SettingsContent />
        </div>
    </main>
    <button on:click={closeSettings}>
        <X />
    </button>
</div>

<style>
    .settings {
        --size-settings-width: 270px;
        --size-settings-padding-outer: 60px;
        --size-settings-padding-inner: 24px;

        position: relative;
        width: 100%;
        height: 100%;
        background: var(--clr-background-dark);
        display: grid;
        grid-template-columns: var(--size-settings-width) 1fr;
        grid-template-rows: 1fr;
        grid-template-areas: "sidebar" "content";
        user-select: none;
        -webkit-user-select: none;
        cursor: default;
    }

    nav,
    nav .wrapper,
    main {
        position: relative;
        width: 100%;
        height: 100%;
        isolation: isolate;
    }

    nav {
        --wails-draggable: drag;
        user-select: none;
        -webkit-user-select: none;
        cursor: default;

        padding-top: var(--size-navbar-height);
        padding-left: var(--size-settings-padding-outer);
        padding-right: var(--size-settings-padding-inner);
    }

    main {
        --scroll-width: 10px;
        padding-left: var(--size-settings-padding-inner);
        padding-right: var(--size-settings-padding-outer);
        background: var(--clr-background);
        overflow-x: hidden;
        overflow-y: scroll;
    }

    main::-webkit-scrollbar {
        width: var(--scroll-width);
    }

    main::-webkit-scrollbar-corner {
        background: transparent;
    }

    main::-webkit-scrollbar-thumb {
        background: var(--clr-text-sub);
        -webkit-transition: 0.125s;
        transition: 0.125s;
        border-radius: 4px !important;
    }

    main::-webkit-scrollbar-track {
        background: 0 0;
    }

    main .wrapper {
        position: relative;
        width: 100%;
        padding-block: var(--size-navbar-height);
    }

    button {
        --size-button-size: 35px;
        position: fixed;
        width: var(--size-button-size);
        height: var(--size-button-size);
        top: var(--size-navbar-height);
        right: calc(
            (var(--size-settings-padding-outer) - var(--size-button-size)) / 2
        );
        background: none;
        color: var(--clr-text-muted);
        border: 1px solid var(--clr-text-muted);
        border-radius: 100%;
        display: flex;
        justify-content: center;
        align-items: center;
    }

    button:hover {
        color: var(--clr-text-sub);
        border-color: var(--clr-text-sub);
        cursor: pointer;
    }

    @media (max-width: 700px) {
        .settings {
            --size-settings-width: 200px;
            --size-settings-padding-outer: 15px;
            --size-settings-padding-inner: 15px;
        }

        button {
            display: none;
        }
    }
</style>
