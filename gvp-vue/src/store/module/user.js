import storageService from '@/service/storageService';
import userService from '@/service/userService';

const userModule = {
  namespaced: true,
  state: {
    token: storageService.get(storageService.USER_TOKEN),
    userInfo: storageService.get(storageService.USER_INFO) ? JSON.parse(storageService.get(storageService.USER_INFO)) : null, //eslint-disable-line
  },
  mutations: {
    SET_TOKEN(state, token) {
      storageService.set(storageService.USER_TOKEN, token);
      state.token = token;
    },
    SET_USERINFO(state, userInfo) {
      storageService.set(storageService.USER_INFO, JSON.stringify(userInfo));
      state.userInfo = userInfo;
    },
  },
  actions: {
    register(context, { name, telephone, password }) {
      return new Promise((resolve, reject) => {
        userService.register({ name, telephone, password }).then((res) => {
          context.commit('SET_TOKEN', res.data.data.token);
          return userService.info();
        }).then((res) => {
          context.commit('SET_USERINFO', res.data.data.user);
          resolve(res);
        }).catch((error) => {
          reject(error);
        });
      });
    },
    login(context, { telephone, password }) {
      return new Promise((resolve, reject) => {
        userService.login({ telephone, password }).then((res) => {
          context.commit('SET_TOKEN', res.data.data.token);
          return userService.info();
        }).then((res) => {
          context.commit('SET_USERINFO', res.data.data.user);
          resolve(res);
        }).catch((error) => {
          reject(error);
        });
      });
    },
    logout({ commit }) {
      commit('SET_TOKEN', '');
      commit('SET_USERINFO', '');
    },
  },
};

export default {
  ...userModule,
};
