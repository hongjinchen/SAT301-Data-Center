<template>
  <v-app-bar
    app
    color="#6A76AB"
    src="https://picsum.photos/1920/1080?random"
    fade-img-on-scroll
    flat
  >
    <v-container class="py-0">
      <v-spacer></v-spacer>
      <v-tabs      grow>
        <v-tab @click="gotoHome">Statistics </v-tab>
        <v-tab @click="gotoDatabase">Database </v-tab>
      </v-tabs>
      <v-spacer></v-spacer>
    </v-container>
  </v-app-bar>
</template>
<script>
export default {
  data() {
    return {
      tabs: [
        { id: 1, name: "Statistics", route: `/` },
        { id: 2, name: "Database", route: `/userDatabase` },
      ],
    };
  },
  methods: {
    gotoHome() {
      this.$router.push({ path: "/" });
    },
    gotoDatabase() {
      this.$router.push({ path: "/userDatabase" });
    },
  },
};
</script>

<style scoped></style>
