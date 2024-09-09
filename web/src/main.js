
import './xxyy.css'
import './icon.css'
import './woff2/1.woff2'
import './woff2/2.woff2'
import './woff2/3.woff2'
import './woff2/4.woff2'
import './woff2/5.woff2'

import 'element-plus/dist/index.css'
import 'ant-design-vue/dist/reset.css';
import { createPinia } from 'pinia'
import { createApp } from 'vue'
import Antd from 'ant-design-vue';

import App from './App.vue'
import router from './router'

import 'mdui/mdui.css';
import 'mdui';
import 'mdui/components/icon.js';
import { setColorScheme } from 'mdui/functions/setColorScheme.js';
setColorScheme('#ffffff');









const app = createApp(App)




app.use(createPinia())
app.use(router)


app.use(Antd)

app.mount('#app')
