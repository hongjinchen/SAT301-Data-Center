<template>
  <v-app style="margin-top: 120px">
    <v-card
      v-if="this.GlobalVariable.State.isLogin"
      style="margin: 0px 90px 30px 90px"
      color="grey lighten-3"
    >
      <v-data-table
        :headers="dessertHeaders"
        :items="desserts"
        :single-expand="singleExpand"
        :expanded.sync="expanded"
        item-key="name"
        show-expand
        class="elevation-1"
      >
        <template v-slot:top>
          <v-toolbar flat>
            <v-toolbar-title>Detect Information Table</v-toolbar-title>
            <v-spacer></v-spacer>
          </v-toolbar>
        </template>
        <template v-slot:expanded-item="{ headers }">
          <td :colspan="headers.length">
            <SonogramPage style="margin: 20px"></SonogramPage>
          </td>
        </template>
      </v-data-table>
    </v-card>
    <div v-else style="margin-bottom: 260px">
      <LoginPage></LoginPage>
    </div>
    <BottomFooter> </BottomFooter>
  </v-app>
</template>
<script>
import BottomFooter from "./BottomFooter";
import SonogramPage from "./SonogramPage";
import LoginPage from "./LoginPage";
import { getAll } from "@/router/api.js";
export default {
  data() {
    return {
      expanded: [],
      singleExpand: true,
      dessertHeaders: [
        { text: "Location", value: "addressInfo" },
        { text: "Result", value: "result" },
      ],
      desserts: [],
    };
  },
  components: {
    BottomFooter,
    SonogramPage,
    LoginPage,
  },
  created() {
    this.getUserInformation();
  },
  methods: {
    getUserInformation() {
      getAll()
        .then((res) => {
          this.desserts = res.data.data;
        })
        .catch(() => {
          alert("Failed to get  user information, please try again");
        });
    },
  },
};
</script>

<style scoped></style>
