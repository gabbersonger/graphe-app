import { graphe_settings } from "@/lib/stores";
import { createThemeStyles, type ThemeName } from "@/static/themes";
import { createFontStyles } from "@/static/fonts";

export function uiManager(elem: HTMLElement) {
  const unsubscribe = graphe_settings.subscribe((value) => {
    // Set theme
    const theme_styles = createThemeStyles(value.appearence.theme as ThemeName);
    for (let i = 0; i < theme_styles.length; i++) {
      elem.style.setProperty(theme_styles[i].variable, theme_styles[i].value);
    }

    // Set font
    const font_styles = createFontStyles(value.appearence.font);
    for (let i = 0; i < font_styles.length; i++) {
      elem.style.setProperty(font_styles[i].variable, font_styles[i].value);
    }
  });

  return {
    destroy() {
      unsubscribe();
    },
  };
}
