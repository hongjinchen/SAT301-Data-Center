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
        {
          text: "No.",
          align: "start",
          sortable: false,
          value: "index",
        },
        { text: "Location", value: "addressInfo" },
        { text: "Result", value: "result" },
      ],
      desserts: [
        {
          addressInfo: "重庆市重庆市",
          index: 1,
          result: "negative",
          audioFile: 24,
        },
        {
          addressInfo: "重庆市重庆市",
          index: 2,
          result: "negative",
          audioFile: 37,
        },
        {
          addressInfo: "重庆市重庆市",
          index: 3,
          result: "positive",
          audioFile: 23,
        },
        {
          addressInfo: "广东省广州市",
          index: 4,
          result: "negative",
          audioFile: 67,
        },
        {
          addressInfo: "广东省广州市",
          index: 5,
          result: "positive",
          audioFile: 49,
        },
        {
          addressInfo: "广东省广州市",
          index: 6,
          result: "negative",
          audioFile: 94,
        },
        {
          addressInfo: "广东省广州市",
          index: 7,
          result: "negative",
          audioFile: 98,
        },
        {
          addressInfo: "广东省广州市",
          index: 8,
          result: "positive",
          audioFile: 87,
        },
        {
          addressInfo: "湖南省衡阳市",
          index: 9,
          result: "negative",
          audioFile: 51,
        },
        {
          addressInfo: "湖南省衡阳市",
          index: 10,
          result: "negative",
          audioFile: 65,
        },
        {
          addressInfo: "河南省郑州市",
          index: 11,
          result: "negative",
          audioFile: 65,
        },
        {
          addressInfo: "江苏省苏州市",
          index: 12,
          result: "negative",
          audioFile: 65,
        },
        {
          addressInfo: "江苏省苏州市",
          index: 13,
          result: "negative",
          audioFile: 65,
        },
        {
          addressInfo: "江苏省苏州市",
          index: 14,
          result: "negative",
          audioFile: 65,
        },
        {
          addressInfo: "江苏省苏州市",
          index: 15,
          result: "negative",
          audioFile: 65,
        },
      ],
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
        //           {
        //   addressInfo: "重庆市重庆市",
        //   index: 1,
        //   result: "negative",
        //   audioFile: 24,
        // },
        for(var item in res.data.data){
              console.log(item);
        }
          console.log(res.data);
        })
        .catch(() => {
          // alert("Failed to get  user information, please try again");
          console.log("获取用户信息失败");
        });
    },
  },
};
</script>

<style scoped></style>
