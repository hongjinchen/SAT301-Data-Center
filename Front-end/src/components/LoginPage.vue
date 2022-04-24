<template>
  <div>
    <v-card>
      <v-card-title>Login</v-card-title>
    </v-card>
  </div>
</template>

<script>
export default {
  data() {
    return {};
  },
  methods: {
    changeState() {
      this.$store.state.isLogin = !this.$store.state.isLogin;
    },
  },
};
</script>

<style></style>
