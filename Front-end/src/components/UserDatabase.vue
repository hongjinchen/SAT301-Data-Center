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
          </v-toolbar>
        </template>
        <template v-slot:expanded-item="{ headers }">
          <td :colspan="headers.length">
          <SonogramPage style="margin:20px"></SonogramPage>
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
  data () {
      return {
        expanded: [],
        singleExpand: true,
        dessertHeaders: [
          {
            text: 'No.',
            align: 'start',
            sortable: false,
            value: 'name',
          },
          { text: 'Location', value: 'addressInfo' },
          { text: 'Result', value: 'result' },
        ],
        desserts: [
          {
            addressInfo: '重庆市重庆市',
            name: 159,
            result:"negative",
            audioFile: 24,
          },
          {
            addressInfo: '重庆市重庆市',
            name: 237,
            result: "negative",
            audioFile: 37,
          },
          {
            addressInfo: '重庆市重庆市',
            name: 262,
            result: "positive",
            audioFile: 23,
          },
          {
            addressInfo: '重庆市重庆市',
            name: 305,
            result:"negative",
            audioFile: 67,
          },
          {
            addressInfo: '重庆市重庆市',
            name: 356,
            result: "positive",
            audioFile: 49,
          },
          {
            addressInfo: '重庆市重庆市',
            name: 375,
            result: "negative",
            audioFile: 94,
          },
          {
            addressInfo: '重庆市重庆市',
            name: 392,
            result:"negative",
            audioFile: 98,
          },
          {
            addressInfo: '重庆市重庆市',
            name: 408,
            result:"negative",
            audioFile: 87,
          },
          {
            addressInfo: '重庆市重庆市',
            name: 452,
            result:"negative",
            audioFile: 51,
          },
          {
            addressInfo: '重庆市重庆市',
            name: 518,
            result:"positive",
            audioFile: 65,
          },
        ],
      }
    },
  components: {
    BottomFooter,
    SonogramPage
  },
};
</script>

<style scoped></style>
