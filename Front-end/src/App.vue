<template>
  <v-app>
    <router-view></router-view>
    <NavBar> </NavBar>
    <!-- <HelloWorld></HelloWorld> -->
  </v-app>
</template>

<script>
import NavBar from "./components/NavBar";
// import HelloWorld from "./components/HelloWorld"
export default {
  name: "App",

  components: {
    NavBar,
    // HelloWorld
  },

  data: () => ({
    //
  }),
};
</script>
