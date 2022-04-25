import Vue from 'vue'
import App from './App.vue'
import vuetify from './plugins/vuetify'
import echarts from 'echarts'
import router from "./router";
import GlobalVariable from './GlobalVariable'

Vue.prototype.GlobalVariable = GlobalVariable //挂载到Vue实例上面
Vue.prototype.$echarts = echarts
Vue.config.productionTip = false

new Vue({
  router,
  vuetify,
  render: h => h(App)
}).$mount('#app')
