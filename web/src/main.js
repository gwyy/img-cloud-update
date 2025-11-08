import { createApp } from 'vue'
import ElementPlus from 'element-plus'
import 'element-plus/dist/index.css'
import router from './router'
import App from './App.vue'
import './assets/styles/global.css'
import * as ElementPlusIconsVue from '@element-plus/icons-vue'
import axios from 'axios'

const app = createApp(App)
app.use(ElementPlus)
app.use(router)

//注册icon
for (const [key, component] of Object.entries(ElementPlusIconsVue)) {
    app.component(key, component)
  }
app.mount('#app')


