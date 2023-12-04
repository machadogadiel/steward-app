// https://nuxt.com/docs/api/configuration/nuxt-config

export default defineNuxtConfig({
  devtools: { enabled: true },
  app: {
    head: {
      title: 'Gadiel.dev',
    }
  },
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
