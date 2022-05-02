import Vue from 'vue'
import App from './App.vue'
import vuetify from './plugins/vuetify'
import echarts from 'echarts'
import router from "./router";
import GlobalVariable from './GlobalVariable'
import axios from 'axios'
axios.defaults.baseURL = "http://42.192.21.119:8080"
Vue.prototype.$http = axios
Vue.prototype.GlobalVariable = GlobalVariable //挂载到Vue实例上面
Vue.prototype.$echarts = echarts
Vue.config.productionTip = false

new Vue({
  router,
  vuetify,
  render: h => h(App)
}).$mount('#app')
