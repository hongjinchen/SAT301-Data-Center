<template>
  <v-app>
    <router-view></router-view>
    <NavBar> </NavBar>
  </v-app>
</template>

<script>
// import HelloWorld from "./components/HelloWorld";
import NavBar from "./components/NavBar";
// import WelcomePage from "./components/WelcomePage";
export default {
  name: "App",

  components: {
    // HelloWorld,
    // WelcomePage,
    NavBar,
  },

  data: () => ({
    //
  }),
};
</script>
