import Vue from 'vue'
import App from './App.vue'
import vuetify from './plugins/vuetify'
import echarts from 'echarts'
import router from "./router";
import store from "./store";

Vue.prototype.$echarts = echarts
Vue.config.productionTip = false

new Vue({
  router,
  vuetify,
  store,
  render: h => h(App)
}).$mount('#app')
