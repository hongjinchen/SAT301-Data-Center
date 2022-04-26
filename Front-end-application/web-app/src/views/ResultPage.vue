<template>
  <div style="margin: 20px 240px 20px 240px">
    <v-card elevation="2">
      <v-card-title>Your test result is:</v-card-title>
      <v-card-subtitle>{{ result }}</v-card-subtitle>
      <div style="text-align: center; display: flex; justify-content: center">
        <img
          v-if="isNegative"
          style="height: 300px; width: auto"
          src="../assets/congratulation.png"
        />
        <img
          v-else
          style="height: 300px; width: auto"
          src="../assets/coughing.svg"
        />
      </div>
      <v-card-actions @click="Gohome">Back to home page</v-card-actions>
    </v-card>
  </div>
</template>

<script>
export default {
  data() {
    return {
      result: "Negative",
      isNegative: false,
      loading: true,
    };
  },
  mounted() {
    this.start();
    this.getResult();
  },
  methods: {
    Gohome() {
      this.$router.push({ path: "/" });
    },
    start() {
      var self = this;
      setTimeout(
        function () {
          self.loading = !self.loading; //改为self
        }.bind(this),
        2000
      );
    },
    getResult() {
      this.isNegative = true;
      if (this.isNegative) {
        this.result = "Negative";
      } else {
        this.result = "Positive";
      }
    },
  },
};
</script>

<style></style>
