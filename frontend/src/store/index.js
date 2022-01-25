import Vue from 'vue';
import Vuex from 'vuex';
import axios from 'axios';

Vue.use(Vuex);

const baseURL = 'http://localhost:3000/v1/players/';
export default new Vuex.Store({
  state: {
    token: localStorage.getItem('token') || null,
  },
  getters: {
    loggedIn(state) {
      return state.token !== null;
    },
  },
  mutations: {
    setToken(state, token) {
      state.token = token;
    },
    destoyToken(state) {
      state.token = null;
    },
  },
  actions: {
    signIn(context, credentials) {
      return new Promise((resolve, reject) => {
        axios.post(`${baseURL}login`, {
          username: credentials.username,
          password: credentials.password,
        })
          .then((response) => {
            const { token } = response.data;

            localStorage.setItem('token', token);
            resolve(context.commit('setToken', token));
          })
          .catch((err) => {
            reject(err);
          });
      });
    },
    signOut(context) {
      localStorage.removeItem('token');
      context.commit('destoyToken');
    },
  },
  modules: {
  },
});
