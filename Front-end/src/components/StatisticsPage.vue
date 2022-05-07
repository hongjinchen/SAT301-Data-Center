<template>
  <v-app style="margin-top: 60px" id="inspire">
    <v-card style="margin: 60px 60px 60px 60px" color="grey lighten-3">
      <v-container>
        <v-row>
          <v-col cols="4">
            <v-sheet rounded="lg">
              <v-list color="transparent">
                <v-list-item
                  v-for="(name, index) in ChartList"
                  :key="name"
                  link
                  @click="ChangeChart(index)"
                >
                  <v-list-item-content>
                    <v-list-item-title> {{ name }} </v-list-item-title>
                  </v-list-item-content>
                </v-list-item>

                <v-divider class="my-2"></v-divider>

                <v-list-item link color="grey lighten-4" @click="Refresh">
                  <v-list-item-content>
                    <v-list-item-title> Refresh </v-list-item-title>
                  </v-list-item-content>
                </v-list-item>
              </v-list>
            </v-sheet>
          </v-col>

          <v-col>
            <v-sheet rounded="lg" style="padding: 40px">
              <div
                id="main"
                class="main"
                style="
                  min-width: 400px;
                  min-height: 300px;
                  width: 600px;
                  height: 450px;
                "
              ></div>
            </v-sheet>
          </v-col>
        </v-row>
      </v-container>
    </v-card>
    <BottomFooter> </BottomFooter>
  </v-app>
</template>

<script>
import BottomFooter from "./BottomFooter";
import {
  getPercentageDetectionResults,
  getNumberDetectionByRegion,
} from "@/router/api.js";

export default {
  data: () => ({
    option: {},
    piehChartData: [],
    linechartCatgory: [],
    linechartData: [],
    ChartList: [
      "Percentage of detection results",
      "Number of detection in China",
      // "Number of detection for a single user",
    ],
    index: 0,
  }),
  components: {
    BottomFooter,
  },
  mounted() {
    // this.getStatistics();
    this.ChangeChart(this.index);
  },
  created() {
    // this.ChangeChart(this.index);
  },

  methods: {
    ChangeChart(index) {
      let myChart = this.$echarts.init(document.getElementById("main"));
      if (index == 0) {
        this.option = {
          title: {
            text: "Percentage of detection results",
            left: "center",
          },
          tooltip: {
            trigger: "item",
          },
          legend: {
            orient: "vertical",
            left: "left",
          },
          series: [
            {
              name: "Access From",
              type: "pie",
              radius: "50%",
              // data: this.piehChartData,
              data:[
                { value: 3, name: "True" },
                { value: 12, name: "False" },
              ],
              emphasis: {
                itemStyle: {
                  shadowBlur: 10,
                  shadowOffsetX: 0,
                  shadowColor: "rgba(0, 0, 0, 0.5)",
                },
              },
            },
          ],
        };
      } else if (index == 1) {
        this.option = {
          title: {
            text: "Number of decetion per region",
          },
          xAxis: {
            type: "category",
            // data: this.linechartCatgory,
             data: ["重庆市重庆市", "广东省广州市", "湖南省衡阳市", "河南省郑州市", "江苏省苏州市"],
          },
          yAxis: {
            type: "value",
          },
          series: [
            {
              // data: this.linechartData,
              data:[3,5,2,1,4],
              type: "line",
            },
          ],
        };
      }
      myChart.setOption(this.option);
    },
    Refresh() {
      location.reload();
    },
    getStatistics() {
      getPercentageDetectionResults()
        .then((res) => {
          // data: {TrueNumber: 0, FalseNumber: 1}
          this.piehChartData = [
            { value: res.data.data.TrueNumber, name: "True" },
            { value: res.data.data.FalseNumber, name: "False" },
          ];
          console.log(this.res.data.data);
          console.log("getPercentageDetectionResults");
        })
        .catch(() => {
          alert("Failed to get Percentage of Detection Results");
          console.log("Failed to get Percentage of Detection Results");
        });
      getNumberDetectionByRegion()
        .then((res) => {
          for (var item in res.data.data) {
            this.linechartCatgory.push(item);
            this.linechartData.push(res.data.data[item]);
          }
          console.log(res.data.data);
        })
        .catch(() => {
          alert("Failed to get Number of Detection By Region");
          console.log("Failed to get Number of Detection By Region");
        });
    },
  },
};
</script>
<style scoped>
.main {
  text-align: center;
  margin: auto;
  position: relative;
  top: 0;
  left: 0;
  right: 0;
  bottom: 0;
}
.img-logo {
  width: 100%;
  height: 100%;
  background: url("../assets/svgexport-5.png") no-repeat center/contain;
}
</style>
