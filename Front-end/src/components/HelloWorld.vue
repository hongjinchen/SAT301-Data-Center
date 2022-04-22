<template>
  <v-app style="margin-top:60px" id="inspire">
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
export default {
  data: () => ({
    option: {},
    ChartList: [
      "Percentage of detection results",
      "Number of detection in China",
      "Number of detection for a single user",
    ],
    index: 0,
  }),
  components: {
    BottomFooter,
  },
  mounted() {
    this.ChangeChart(this.index);
  },
  methods: {
    ChangeChart(index) {
      let myChart = this.$echarts.init(document.getElementById("main"));
      if (index == 0) {
        this.option = {
          title: {
            text: "Percentage of detection results",
            subtext: "Fake Data",
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
              data: [
                { value: 484, name: "True" },
                { value: 300, name: "False" },
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
            text: "Number of users per day",
          },
          xAxis: {
            type: "category",
            data: ["Mon", "Tue", "Wed", "Thu", "Fri", "Sat", "Sun"],
          },
          yAxis: {
            type: "value",
          },
          series: [
            {
              data: [150, 230, 224, 218, 135, 147, 260],
              type: "line",
            },
          ],
        };
      } else if (index == 2) {
        this.option = {
          title: {
            text: "The number of detection for a single user",
            subtext: "Fake Data",
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
              name: "User number",
              type: "pie",
              radius: "50%",
              data: [
                { value: 1048, name: "1" },
                { value: 735, name: "2" },
                { value: 580, name: "3" },
                { value: 484, name: "4" },
                { value: 300, name: "5" },
                { value: 300, name: "over 5" },
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
      }
      myChart.setOption(this.option);
    },
    Refresh() {
      location.reload();
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
