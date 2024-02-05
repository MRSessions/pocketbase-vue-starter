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

const appStore = useAppStore()
</script>
