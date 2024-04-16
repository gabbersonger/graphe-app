import ModalSearch from "@/components/Modals/built/ModalSearch.svelte";
import ModalChooseText from "@/components/Modals/built/ModalChooseText.svelte";
import ModalChoosePassage from "@/components/Modals/built/ModalChoosePassage.svelte";
import ModalFunctions from "@/components/Modals/built/ModalFunctions.svelte";
import ModalAppearence from "@/components/Modals/built/ModalAppearence.svelte";

export const modalData = [
  {
    name: "search",
    modal: ModalSearch,
  },
  {
    name: "chooseText",
    modal: ModalChooseText,
  },
  {
    name: "choosePassage",
    modal: ModalChoosePassage,
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
