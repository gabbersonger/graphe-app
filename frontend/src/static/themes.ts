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
      main: "#513a2a",
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
] as const;

export type ThemeName = (typeof themeData)[number]["name"];
export const defaultTheme: ThemeName = "catppuccin";

export const createThemeStyles = (themeName: ThemeName) => {
  let theme = themeData.find((t) => t.name == themeName);
  return `
    --clr-background: ${theme.colors.background};
    --clr-background-sub: ${theme.colors.backgroundSub};
    --clr-background-dark: ${theme.colors.backgroundDark};
    --clr-main: ${theme.colors.main};
    --clr-text: ${theme.colors.text};
    --clr-text-sub: ${theme.colors.textSub};
    --clr-text-highlight: ${theme.colors.textHighlight};
    --clr-text-muted: ${theme.colors.textMuted};
  `;
};
