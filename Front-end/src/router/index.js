import Vue from 'vue'
import VueRouter from 'vue-router'
import HelloWorld  from '../components/HelloWorld.vue';
import UserDatabase  from '../components/UserDatabase.vue';
import WelcomePage from "../components/WelcomePage.vue"
import SonogramPage from "../components/SonogramPage.vue"

Vue.use(VueRouter)

const routes = [
  {    path: '/',
  name: WelcomePage,
  component: WelcomePage,},
  {
    path: '/helloWorld',
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