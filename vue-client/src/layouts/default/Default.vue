<template>
  <v-app :theme="appStore.appSettings.theme">
    <default-app-bar />
    <default-view />
  </v-app>
</template>

<script lang="ts" setup>
import { useAppStore } from '@/store/app';
import DefaultAppBar from './AppBar.vue'
import DefaultView from './View.vue'
import router from '@/router';

const appStore = useAppStore()

try {
  await appStore.checkIsSetup()
  if (!appStore.appSettings.isSetup) {
    router.push('/admin/setup-admin')
  }
} catch (error) {
  if (error instanceof Error) {
    router.push('/n/error?msg=' + error.message)
  }
}
</script>
