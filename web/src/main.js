import { createApp } from 'vue'
import ElementPlus from 'element-plus'
import 'element-plus/dist/index.css'
import router from './router'
import {createPinia} from 'pinia'
import piniaPluginPersistedstate from 'pinia-plugin-persistedstate'
import App from './App.vue'
import './assets/styles/global.css'
import * as ElementPlusIconsVue from '@element-plus/icons-vue'
import axios from 'axios'

const pinia = createPinia()

// 关键：先 use 插件，再 use pinia
pinia.use(piniaPluginPersistedstate)

const app = createApp(App)
app.use(pinia)
app.use(ElementPlus)
app.use(router)
app.config.globalProperties.$http = axios
//注册icon
for (const [key, component] of Object.entries(ElementPlusIconsVue)) {
    app.component(key, component)
  }
app.mount('#app')


