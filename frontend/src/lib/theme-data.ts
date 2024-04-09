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
    },
  },
] as const;

export type ThemeName = (typeof themeData)[number]["name"];
export const defaultTheme: ThemeName = "hanok";
