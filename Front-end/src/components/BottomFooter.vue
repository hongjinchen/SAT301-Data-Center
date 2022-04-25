<template>
  <v-footer dark padless>
    <v-card flat tile color="#6A76AB" class="white--text text-center">
      <v-card-text>
        <v-btn v-for="icon in icons" :key="icon" class="mx-4 white--text" icon>
          <v-icon size="24px">
            {{ icon }}
          </v-icon>
        </v-btn>
      </v-card-text>

      <v-card-text class="white--text pt-0">
        This website is the administrator system of the project "Application for
        Detection of New Coronary Pneumonia Disease by Breath Sound", which
        mainly includes two functions. Including preview statistics, and view
        the query history of all users. The first page is public, While the
        login the second feature requires logging in first, and only the unique
        administrator account can log in.
      </v-card-text>

      <v-divider></v-divider>

      <v-card-text class="white--text">
        {{ new Date().getFullYear() }} —
        <strong
          >SAT301 Application for Detection of New Coronary Pneumonia Disease by
          Breath Sound</strong
        >
      </v-card-text>
    </v-card>
  </v-footer>
</template>
<script>
export default {
  setup() {},
};
</script>

<style scoped></style>
