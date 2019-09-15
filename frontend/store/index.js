
import Vue from 'vue'
import Vuex from 'vuex'
Vue.use(Vuex);

const store = () => new Vuex.Store({
  state: {
    captcha:null,
    token: "",
    login:""
  },
  mutations: {
    setCaptcha (state, data) {
      state.captcha = data;
    },
    setToken ( state, data) {
      state.token = data;
    },
    setLogin ( state, data) {
      state.login = data;
    }
  },
  actions: {
    async nuxtServerInit({ commit }, { req } ) {
      // const data = await this.$axios.$get()
    },
    acCaptcha( { commit }, data) {
      commit("setCaptcha", data)
    },
    acToken ({ commit }, data) {
      commit("setToken", data)
    },
    acLogin ({ commit }, data) {
      commit("setLogin", data)
    }
  }
});

export default store
