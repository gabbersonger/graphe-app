import ModalSearch from "@/components/Workspace/Modals/content/ModalSearch.svelte";
import ModalVersion from "@/components/Workspace/Modals/content/ModalVersion.svelte";
import ModalText from "@/components/Workspace/Modals/content/ModalText.svelte";

export const modalData = [
  {
    name: "search",
    modal: ModalSearch,
  },
  {
    name: "version",
    modal: ModalVersion,
  },
  {
    name: "text",
    modal: ModalText,
  },
] as const;

export type ModalName = (typeof modalData)[number]["name"];
