import Vue from 'vue';
import VueX from 'vuex';
import { eventType } from '../eventsink/';
Vue.use(VueX);
export default new VueX.Store({
  state: {
    textWidget: {},
    clockWidget: {},
    blogRoll: {},
    weather: {},
    comparisonWidget: {},
    slideshow: {}
  },
  mutations: {
    change(state, payload: eventType<object>) {
      state[payload.event] = payload.data;
    }
  },
  actions: {
    change({ commit }, payload) {
      commit('change', payload);
    }
  }
});
