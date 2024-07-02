<script lang="ts">
    import ModalWrapper from "@/components/Workspace/Modals/ModalWrapper.svelte";

    let contentElement: HTMLDivElement;
</script>

<ModalWrapper bind:contentElement>
    <div slot="header" class="input-container">Functions</div>

    <div slot="content">content</div>
</ModalWrapper>
