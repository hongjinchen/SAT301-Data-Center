import Vue from 'vue'
import VueRouter from 'vue-router'
import HelloWorld  from '../components/StatisticsPage.vue';
import UserDatabase  from '../components/UserDatabase.vue';
import SonogramPage from "../components/SonogramPage.vue"
import LoginPage from "../components/LoginPage.vue"
Vue.use(VueRouter)

const routes = [
  
  {
    path: '/LoginPage',
    name: LoginPage,
    component: LoginPage,
  },
  {
    path: '/',
    name: HelloWorld,
    component: HelloWorld,
  },
  {
    path: '/userDatabase',
    name: UserDatabase,
    component: UserDatabase,
  },
  {
    path: '/sonogramPage',
    name: SonogramPage,
    component: SonogramPage,
  }

]
const router = new VueRouter({
  mode: 'history',
  base: process.env.BASE_URL,
  routes
})

export default router