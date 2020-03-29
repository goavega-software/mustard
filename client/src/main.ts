import Vue from "vue";
import App from "./App.vue";
import store from "./store/";
import VueSSE from "vue-sse";
Vue.config.productionTip = false;
Vue.use(VueSSE);
new Vue({
  render: h => h(App),
  store,
}).$mount("#app");
