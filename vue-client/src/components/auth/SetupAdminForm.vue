<template>
  <v-container>
    <v-row class="pt-12" justify="center">
      <v-col cols="12" sm="12" md="8">
        <v-card variant="tonal">
          <v-card-title>
            <span class="headline">Setup</span>
          </v-card-title>
          <v-card-text class="pb-0">
            <v-form ref="form" v-model="valid" lazy-validation>
              <v-text-field v-model="appStore.adminUserSetup.email" label="E-mail" required
                :rules="[rules.required, rules.email]"></v-text-field>
              <v-text-field v-model="appStore.adminUserSetup.password"
                :append-icon="showPassword ? 'mdi-eye' : 'mdi-eye-off'" :type="showPassword ? 'text' : 'password'"
                label="Password" hint="At least 10 characters" counter required
                @click:append="showPassword = !showPassword" :rules="[rules.min]"></v-text-field>
              <v-text-field v-model="appStore.adminUserSetup.passwordConfirm"
                :append-icon="showConfirmPassword ? 'mdi-eye' : 'mdi-eye-off'"
                :type="showConfirmPassword ? 'text' : 'password'" label="Confirm Password" hint="At least 10 characters"
                counter required @click:append="showConfirmPassword = !showConfirmPassword"
                :rules="[rules.min, rules.passwordMatch]"></v-text-field>
            </v-form>
          </v-card-text>
          <div class="pb-6 px-6">
            <v-btn size="x-large" variant="tonal" :disabled="!valid" color="success" @click="setup" block>
              Setup
            </v-btn>
          </div>
        </v-card>
      </v-col>
    </v-row>
  </v-container>
</template>

<script setup lang="ts">
import { ref } from 'vue';
import { useAppStore } from '@/store/app';
import router from '@/router';

const appStore = useAppStore()

const rules = {
  required: (v: string) => !!v || 'Required.',
  min: (v: string) => v.length >= 10 || 'At least 10 characters',
  passwordMatch: () => (appStore.adminUserSetup.passwordConfirm === appStore.adminUserSetup.password) || 'Passwords must match',
  email: (v: string) => /^\w+([.-]?\w+)*@\w+([.-]?\w+)*(\.\w{2,3})+$/.test(v) || 'E-mail must be valid',
};

const valid = ref(false);
const showPassword = ref(false);
const showConfirmPassword = ref(false);

async function setup() {
  try {
    const registerSuccess = await appStore.registerAdmin()
    if (registerSuccess) {
      router.push('/')
    }
  } catch (error) {
    if (error instanceof Error) {
      router.push('/n/error?msg=' + error.message)
    }
  }
}

</script>
