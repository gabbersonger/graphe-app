import ModalSearch from "@/components/Modals/ModalSearch.svelte";
import ModalFunctions from "@/components/Modals/ModalFunctions.svelte";
import ModalChooseText from "@/components/Modals/ModalChooseText.svelte";
import ModalChoosePassage from "@/components/Modals/ModalChoosePassage.svelte";

export const modalData = [
  {
    name: "search",
    modal: ModalSearch,
  },
  {
    name: "functions",
    modal: ModalFunctions,
  },
  {
    name: "chooseText",
    modal: ModalChooseText,
  },
  {
    name: "choosePassage",
    modal: ModalChoosePassage,
  },
] as const;

export type ModalName = (typeof modalData)[number]["name"];
