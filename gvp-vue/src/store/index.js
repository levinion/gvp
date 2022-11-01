import { createStore } from 'vuex';
import userModule from './module/user';

export default createStore({
  strict: process.env.NODE_ENV !== 'production',
  state: {
  },
  getters: {
  },
  mutations: {
  },
  actions: {
  },
  modules: {
    userModule,
  },
});
