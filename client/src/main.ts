import Vue from "vue";
import App from "./App.vue";
import store from "./store/";
// @ts-ignore
import VueSSE from "vue-sse";
import VueApexCharts from "vue-apexcharts";

Vue.config.productionTip = false;
Vue.use(VueSSE);
Vue.use(VueApexCharts);
Vue.component("apexchart", VueApexCharts);
new Vue({
  render: h => h(App),
  store
}).$mount("#app");
