<template>
  <div>
    <v-row>
      <v-col cols="12" sm="6" md="8">
        <v-card color="white" elevation="1" class="record"></v-card>
      </v-col>
      <v-col cols="6" md="4">
        <v-card color="white" elevation="1" class="uploadVideo"></v-card>
      </v-col>
    </v-row>
  </div>
</template>

<script>
export default {
  name: "HomePage",

  components: {},
};
</script>
<style scoped>
.record {
  max-width: 950px;
  min-height: 600px;
}
.uploadVideo {
  max-width: 460px;
  min-height: 600px;
}
</style>
