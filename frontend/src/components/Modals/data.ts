import ModalSearch from "@/components/Modals/built/ModalSearch.svelte";
import ModalVersion from "@/components/Modals/built/ModalVersion.svelte";
import ModalText from "@/components/Modals/built/ModalText.svelte";
import ModalFunctions from "@/components/Modals/built/ModalFunctions.svelte";
import ModalAppearence from "@/components/Modals/built/ModalAppearence.svelte";

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
  {
    name: "functions",
    modal: ModalFunctions,
  },
  {
    name: "appearence",
    modal: ModalAppearence,
  },
] as const;

export type ModalName = (typeof modalData)[number]["name"];
