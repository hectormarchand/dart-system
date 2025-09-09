import './assets/main.css'

import { createApp } from 'vue'
import App from './App.vue'
import router from './router'
import PrimeVue from 'primevue/config';
import Nora from '@primeuix/themes/nora';
import { definePreset } from '@primeuix/themes'

const app = createApp(App)

app.use(router)

const dartPreset = definePreset(Nora, {
  semantic: {
    primary: {
      50: '{orange.50}',
      100: '{orange.100}',
      200: '{orange.200}',
      300: '{orange.300}',
      400: '{orange.400}',
      500: '{orange.500}',
      600: '{orange.600}',
      700: '{orange.700}',
      800: '{orange.800}',
      900: '{orange.900}',
      950: '{orange.950}'
    }
  },
  components: {
    stepper: {
      steppanel: {
        background: 'transparent',
      },
    },
  },
})

app.use(PrimeVue, {
  theme: {
    preset: dartPreset
  },
});

app.mount('#app')
