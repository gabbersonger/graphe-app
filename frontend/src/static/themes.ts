import { GrapheLog } from "@/lib/utils";

export const themeData = [
  {
    name: "serika dark",
    colors: {
      background: "#323437",
      backgroundSub: "#2c2e31",
      backgroundDark: "#1a1b1c",
      main: "#e2b714",
      text: "#d1d0c5",
      textSub: "#646669",
      textHighlight: "#fff",
      textMuted: "#494B4D",
    },
  },
  {
    name: "hanok",
    colors: {
      background: "#d8d2c3",
      backgroundSub: "#cdc0af",
      backgroundDark: "#b8a48b",
      main: "#743D44",
      text: "#393b3b",
      textSub: "#8b6f5c",
      textHighlight: "#000",
      textMuted: "#9b8c81",
    },
  },
  {
    name: "catppuccin",
    colors: {
      background: "#1e1e2e",
      backgroundSub: "#181825",
      backgroundDark: "#08080C",
      main: "#cba6f7",
      text: "#cdd6f4",
      textSub: "#7f849c",
      textHighlight: "#f2cdcd",
      textMuted: "#545767",
    },
  },
  {
    name: "olive",
    colors: {
      background: "#e9e5cc",
      backgroundSub: "#d4cfbc",
      backgroundDark: "#b7b39e",
      main: "#92946f",
      text: "#373731",
      textSub: "#515148",
      textHighlight: "#171714",
      textMuted: "#979787",
    },
  },
  {
    name: "pulse",
    colors: {
      background: "#252525",
      backgroundSub: "#181818",
      backgroundDark: "#090909",
      main: "#17b8bd",
      text: "#e5f4f4",
      textSub: "#AEBFBF",
      textHighlight: "#F4FFFF",
      textMuted: "#53565a",
    },
  },
  {
    name: "zinc",
    colors: {
      background: "#FFFFFF",
      backgroundSub: "#F4F4F5",
      backgroundDark: "#D1D1D1",
      main: "#18181B",
      text: "#09090B",
      textSub: "#717179",
      textHighlight: "#000000",
      textMuted: "#B8B8BC",
    },
  },
] as const;

export type ThemeName = (typeof themeData)[number]["name"];

// NOTE: assumes correct #______ format
const hexToSelectionRGBA = (hex: string): string => {
  const r = parseInt(hex.slice(1, 3), 16);
  const g = parseInt(hex.slice(3, 5), 16);
  const b = parseInt(hex.slice(5, 7), 16);
  return `rgba(${r}, ${g}, ${b}, 0.99)`;
};

export const createThemeStyles = (themeName: ThemeName) => {
  let theme = themeData.find((t) => t.name == themeName);
  if (theme == undefined) {
    GrapheLog(
      "error",
      `[Static Themes] Invalid theme passed to \`createThemeStyles\` (theme: \`${themeName}\`)`,
    );
    return [];
  }
  return [
    {
      variable: "--clr-background",
      value: theme.colors.background,
    },
    {
      variable: "--clr-background-sub",
      value: theme.colors.backgroundSub,
    },
    {
      variable: "--clr-background-dark",
      value: theme.colors.backgroundDark,
    },
    {
      variable: "--clr-main",
      value: theme.colors.main,
    },
    {
      variable: "--clr-selection",
      value: hexToSelectionRGBA(theme.colors.main),
    },
    {
      variable: "--clr-text",
      value: theme.colors.text,
    },
    {
      variable: "--clr-text-sub",
      value: theme.colors.textSub,
    },
    {
      variable: "--clr-text-highlight",
      value: theme.colors.textHighlight,
    },
    {
      variable: "--clr-text-muted",
      value: theme.colors.textMuted,
    },
  ];
};

export const createThemeStylesString = (themeName: ThemeName) => {
  return createThemeStyles(themeName).reduce(
    (acc, cur) => acc + `${cur.variable}: ${cur.value};\n`,
    "",
  );
};
