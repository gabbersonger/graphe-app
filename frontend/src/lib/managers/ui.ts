import { graphe_settings } from "@/lib/stores";
import { createThemeStyles, type ThemeName } from "@/static/themes";
import { createFontStyles } from "@/static/fonts";

export function uiManager(elem: HTMLElement) {
  const unsubscribe = graphe_settings.subscribe((value) => {
    if (value == null) return;
    if (elem == null || elem.parentElement == null) return;

    // Set theme
    const theme_name = value.appearence.theme as ThemeName;
    const theme_styles = createThemeStyles(theme_name);
    for (let i = 0; i < theme_styles.length; i++) {
      elem.parentElement.style.setProperty(
        theme_styles[i].variable,
        theme_styles[i].value,
      );
    }

    // Set font
    const font_styles = createFontStyles(value.appearence.font);
    for (let i = 0; i < font_styles.length; i++) {
      elem.parentElement.style.setProperty(
        font_styles[i].variable,
        font_styles[i].value,
      );
    }

    // Set zoom
    elem.parentElement.style.setProperty(
      "--zoom-factor",
      (value.appearence.zoom / 100).toFixed(1),
    );
  });

  return {
    destroy() {
      unsubscribe();
    },
  };
}
