<template>
  <v-app style="margin-top: 120px">
    <v-card style="margin: 0px 90px 30px 90px" color="grey lighten-3">
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
            <v-toolbar-title>User Information Table</v-toolbar-title>
            <v-spacer></v-spacer>
            <v-switch
              v-model="singleExpand"
              label="Single expand"
              class="mt-2"
            ></v-switch>
          </v-toolbar>
        </template>
        <template v-slot:expanded-item="{ headers }">
          <td :colspan="headers.length">
            <SonogramPage style="margin: 20px"></SonogramPage>
          </td>
        </template>
      </v-data-table>
    </v-card>

    <BottomFooter> </BottomFooter>
  </v-app>
</template>
<script>
import BottomFooter from "./BottomFooter";
import SonogramPage from "./SonogramPage";
export default {
  data() {
    return {
      expanded: [],
      singleExpand: false,
      dessertHeaders: [
        { text: "Location of User", value: "addressInfo" },
        { text: "Decteion result", value: "result" }
      ],
      desserts: [
        {
          addressInfo: "重庆市重庆市",
          result: "negative",
          audioFile: "../assets/test.mp3",
        },
        {
          addressInfo: "重庆市重庆市",
          result: "negative",
          audioFile: "../assets/test.mp3",
        },
        {
          addressInfo: "重庆市重庆市",
          result: "negative",
          audioFile: "../assets/test.mp3",
        },
        {
          addressInfo: "重庆市重庆市",
          result: "negative",
          audioFile: "../assets/test.mp3",
        },
        {
          addressInfo: "重庆市重庆市",
          result: "negative",
          audioFile: "../assets/test.mp3",
        },
        {
          addressInfo: "重庆市重庆市",
          result: "negative",
          audioFile: "../assets/test.mp3",
        },
        {
          addressInfo: "重庆市重庆市",
          result: "negative",
          audioFile: "../assets/test.mp3",
        },
        {
          addressInfo: "重庆市重庆市",
          result: "negative",
          audioFile: "../assets/test.mp3",
        },
        {
          addressInfo: "重庆市重庆市",
          result: "negative",
          audioFile: "../assets/test.mp3",
        },
        {
          addressInfo: "重庆市重庆市",
          result: "negative",
          audioFile: "../assets/test.mp3",
        },
      ],
    };
  },
  components: {
    BottomFooter,
    SonogramPage,
  },
};
</script>

<style scoped></style>
