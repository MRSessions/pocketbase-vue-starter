/**
 * main.ts
 *
 * Bootstraps Vuetify and other plugins then mounts the App`
 */

// Components
import App from './App.vue'

// Composables
import { createApp } from 'vue'

// Plugins
import { registerPlugins } from '@/plugins'
import { pocketBaseSymbol } from '@/pocketbase/injectionSymbols'
import pbServer from '@/pocketbase'

const app = createApp(App)

registerPlugins(app)
app.provide(pocketBaseSymbol, pbServer)

app.mount('#app')
