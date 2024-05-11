import ModalSearch from "@/components/Modals/built/ModalSearch.svelte";
import ModalVersion from "@/components/Modals/built/ModalVersion.svelte";
import ModalText from "@/components/Modals/built/ModalText.svelte";

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
