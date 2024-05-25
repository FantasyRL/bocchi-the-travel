import { createStore } from 'vuex';

export default createStore({
  state: {
    isLoading: false
  },
  mutations: {
    setLoading(state, payload) {
      state.isLoading = payload;
    }
  }
});
