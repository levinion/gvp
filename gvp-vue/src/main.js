import { createApp } from 'vue';
import ElementPlus from 'element-plus';
import Vuex from 'vuex';
import VueAxios from 'vue-axios';
import axios from 'axios';
import App from './App.vue';
import router from './router';
import store from './store';
import 'element-plus/dist/index.css';

createApp(App).use(store).use(ElementPlus).use(router)
  .use(VueAxios, axios)
  .use(Vuex)
  .mount('#app');
