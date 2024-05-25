import './assets/main.css'
import './xxyy.css'
import 'ant-design-vue/dist/reset.css';
import { createPinia } from 'pinia'
import { createApp } from 'vue'
import Antd from 'ant-design-vue';
import App from './App.vue'
import router from './router'
import store from './stores';

import 'mdui/mdui.css';
import 'mdui';








const app = createApp(App)


app.use(store)
app.use(createPinia())
app.use(router)


app.use(Antd).mount('#app');
//app.mount('#app')
