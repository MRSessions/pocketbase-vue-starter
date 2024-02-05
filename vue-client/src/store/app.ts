// Utilities
import { defineStore } from "pinia";
import { ref, inject } from "vue";
import { useLocalStorage } from "@vueuse/core";



export const useAppStore = defineStore("app", () => {

  const appSettings = ref(
    useLocalStorage("appSettings", {
      theme: "dark",
      name: "PocketBase Vue Starter",
    })
  );

  return { appSettings };
});
