// Utilities
import { defineStore } from "pinia";
import { ref, inject } from "vue";
import { useLocalStorage } from "@vueuse/core";
import { pocketBaseSymbol } from "@/pocketbase/injectionSymbols";



export const useAppStore = defineStore("app", () => {
  const adminUserSetup = ref({
    email: '',
    password: '',
    passwordConfirm: '',
  });

  const appSettings = ref(
    useLocalStorage("appSettings", {
      theme: "dark",
      name: "PocketBase Vue Starter",
      isSetup: false,
    })
  );

  // Actions
  // done for initial PocketBase setup
  const pb = inject(pocketBaseSymbol);
  const checkIsSetup = async () => {
    if (!appSettings.value.isSetup) {
      try {
        const initCheck = await pb?.send("/api/init-check", {});
        appSettings.value.isSetup = initCheck.isSetup;
      } catch (e) {
        appSettings.value.isSetup = false;
        throw new Error(
          "Something whent wrong with processing the request. " +
            "Please make sure that PocketBase is running and that the API is accessible."
        );
      }
    }
  };

  const registerAdmin = async () => {
    try {
      const register = await pb?.send("/api/admins", {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        body: adminUserSetup.value,
      });
      console.log(register.email);
      if (register.email === adminUserSetup.value.email) {
        appSettings.value.isSetup = true;
        return true;
      }
    } catch (e) {
      throw new Error(
        "Something whent wrong with processing the request. " +
          "Please make sure that PocketBase is running and that the API is accessible."
      );
    }
  };

  return { adminUserSetup, appSettings, checkIsSetup, registerAdmin };
});
