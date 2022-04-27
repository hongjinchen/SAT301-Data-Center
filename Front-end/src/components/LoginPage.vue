<template>
  <div>
    <v-card style="margin: 30px 300px 30px 300px">
      <v-card-title>Login</v-card-title>
      <div style="margin: 30px 120px 30px 120px">
        <v-text-field
          label="User Name"
          :rules="rules"
          v-model="username"
          hide-details="auto"
        ></v-text-field>
        <v-text-field
          v-model="password"
          type="password"
          label="Password"
        ></v-text-field>
      </div>
      <div style="text-align: center; display: flex; justify-content: center">
        <v-btn style="margin-bottom: 30px" @click="changeState">
          Continue
        </v-btn>
      </div>
    </v-card>
  </div>
</template>

<script>
export default {
  data() {
    return {
      username: "",
      password: "",
    };
  },
  methods: {
    changeState() {
      if (this.username == "admin" && this.password == "123456") {
        this.GlobalVariable.State.isLogin = true;
        alert("Login successfully!");
      } else {
        alert(
          "Incorrect username or password, please check your username and password!"
        );
      }
    },
  },
};
</script>

<style></style>
