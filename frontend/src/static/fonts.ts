import { SettingsValues_Appearence_Font } from "!/graphe/internal/settings";
import { GrapheLog } from "@/lib/utils";

export const fontData = [
  {
    name: "System",
    language: "english",
    css_line: "system-ui, -apple-system, Roboto",
  },
  {
    name: "Neuton",
    language: "english",
    css_line: '"Neuton"',
  },
  {
    name: "Geist",
    language: "english",
    css_line: '"Geist", sans-serif',
  },
  {
    name: "SBL Greek",
    language: "greek",
    css_line: '"SBL Greek"',
  },
  {
    name: "SBL Hebrew",
    language: "hebrew",
    css_line: '"SBL Hebrew"',
  },
] as const;

export type FontName = (typeof fontData)[number]["name"];

export function createFontStyles(fonts: SettingsValues_Appearence_Font) {
  const system = fontData.find((f) => f.name == fonts.system);
  const greek = fontData.find((f) => f.name == fonts.greek);
  const hebrew = fontData.find((f) => f.name == fonts.hebrew);
  const english = fontData.find((f) => f.name == fonts.english);
  if (system == undefined) {
    GrapheLog(
      "error",
      `Invalid system font passed to \`createFontStyles\` (font: \`${fonts.system}\`)`,
    );
    return [];
  }
  if (greek == undefined) {
    GrapheLog(
      "error",
      `Invalid greek font passed to \`createFontStyles\`: "${fonts.greek}"`,
    );
    return [];
  }
  if (hebrew == undefined) {
    GrapheLog(
      "error",
      `Invalid hebrew font passed to \`createFontStyles\`: "${fonts.hebrew}"`,
    );
    return [];
  }
  if (english == undefined) {
    GrapheLog(
      "error",
      `Invalid english font passed to \`createFontStyles\`: "${fonts.english}"`,
    );
    return [];
  }
  return [
    {
      variable: "--font-system",
      value: system.css_line,
    },
    {
      variable: "--font-greek",
      value: greek.css_line,
    },
    {
      variable: "--font-hebrew",
      value: hebrew.css_line,
    },
    {
      variable: "--font-english",
      value: english.css_line,
    },
  ];
}
