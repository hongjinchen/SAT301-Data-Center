<template>
  <v-app>
    <NavigationBar></NavigationBar>
    <v-main>
      <router-view />
    </v-main>
    <FooterPart></FooterPart>
  </v-app>
</template>

<script>
import NavigationBar from "./components/NavigationBar.vue";
import FooterPart from "./components/FooterPart.vue";

export default {
  name: "App",

  data: () => ({
    //
  }),
  components: {
    NavigationBar,
    FooterPart,
  },
};
</script>

<style scoped></style>
