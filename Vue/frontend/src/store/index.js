import Vue from "vue";
import Vuex from "vuex";

Vue.use(Vuex);

export default new Vuex.Store({
  state: { role: null },
  mutations: {
    // for login
    setRole(state, role) {
      state.role = role;
    },

    // for logout
    clearRole(state) {
      state.role = null;
    },
  },
  actions: {},
  modules: {},
});
