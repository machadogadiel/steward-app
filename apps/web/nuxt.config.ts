// https://nuxt.com/docs/api/configuration/nuxt-config

export default defineNuxtConfig({
  devtools: { enabled: true },
  modules: [
    [
      "@nuxt-modules/compression",
      {
        algorithm: "brotliCompress",
      },
    ],
  ],
  vite: {
    plugins: [],
    vue: {
      customElement: true,
    },
    vueJsx: {
      mergeProps: true,
    },
  },
});
