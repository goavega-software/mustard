import Vue from "vue";
import VueX from "vuex";
import { eventType } from "../eventsink/";
Vue.use(VueX);
interface State {
  [key: string]: {
    data: Record<string, unknown>;
  };
}
const getStoreItem = <T>(state: State, stateName?: string): T | undefined => {
  if (!stateName) {
    return;
  }
  const module = state[stateName];
  if (module) {
    return module.data as T;
  }
};
export default new VueX.Store({
  mutations: {
    change(state: State, payload: eventType<Record<string, unknown>>) {
      if (!state[payload.event]) {
        return;
      }
      state[payload.event]["data"] = payload.data;
    }
  },
  actions: {
    change({ commit }, payload) {
      commit("change", payload);
    }
  }
});

export { getStoreItem, State };
