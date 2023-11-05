// uno.config.ts
import {
  defineConfig,
  presetIcons,
  presetUno,
} from "unocss";

export default defineConfig({
  rules: [["m-1", { margin: "1px" }]],
  presets: [
    presetUno(),
    presetIcons({
      scale: 1.2,
      cdn: "https://esm.sh/",
    }),
  ],
});
