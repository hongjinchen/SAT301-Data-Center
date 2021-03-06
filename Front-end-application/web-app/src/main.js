import Vue from 'vue'
import App from './App.vue'
import router from './router'
import store from './store'
import vuetify from './plugins/vuetify'
import axios from 'axios'
import './registerServiceWorker'
import GlobalVariable from './GlobalVariable'
axios.defaults.baseURL = "http://42.192.21.119:8080"
axios.defaults.withCredentials = true
Vue.prototype.$axios = axios //全局注册，使用方法为:this.$axios
Vue.prototype.GlobalVariable = GlobalVariable //挂载到Vue实例上面
Vue.config.productionTip = false

new Vue({
  router,
  store,
  vuetify,
  render: h => h(App)
}).$mount('#app')
