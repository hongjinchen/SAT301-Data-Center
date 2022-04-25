<template>
  <v-app>
    <NavBar> </NavBar>
        <v-main>
      <router-view />
    </v-main>
    <!-- <HelloWorld></HelloWorld> -->
    <!-- <BottomFooter></BottomFooter> -->
  </v-app>
</template>

<script>
import NavBar from "./components/NavBar";
// import BottomFooter from "./components/BottomFooter";
// import HelloWorld from "./components/HelloWorld"
export default {
  name: "App",

  components: {
    NavBar,
    // BottomFooter
    // HelloWorld
  },

  data: () => ({
    //
  }),
};
</script>
