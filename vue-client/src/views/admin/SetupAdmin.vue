<template>
  <SetupAdminForm />
</template>

<script setup lang="ts">
import SetupAdminForm from '@/components/auth/SetupAdminForm.vue';
import { useAppStore } from '@/store/app';
import router from '@/router';

const appStore = useAppStore()

if(appStore.appSettings.isSetup) {
  router.push('/')
}
</script>
