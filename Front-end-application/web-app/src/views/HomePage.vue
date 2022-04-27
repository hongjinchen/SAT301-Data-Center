<template>
  <div>
    <v-row>
      <v-col cols="12" sm="6" md="8">
        <v-card color="blue-grey lighten-5" elevation="1" class="record">
          <div class="title">Record my voice</div>

          <div style="padding: 0px 20px 0px 20px">
            <v-btn
              block
              light
              depressed
              x-large
              v-if="isMicophone"
              color="indigo lighten-2"
              @click="onStartVoice"
            >
              Start/Restart
            </v-btn>

            <v-btn
              block
              depressed
              x-large
              v-else
              light
              color="indigo lighten-4"
              @click="onEndVoice"
            >
              Finish
            </v-btn>

            <v-btn
              block
              depressed
              x-large
              style="margin-top: 20px"
              v-if="isFinished"
              @click="onPlayAudio"
            >
              Play my audio
            </v-btn>

            <!-- 录音条 -->
            <div class="myrecord">
              <div class="record-play" v-show="isPlay">
                <audio id="audioVoice" controls autoplay></audio>
              </div>
            </div>
            <div
              style="text-align: center; display: flex; justify-content: center"
            >
              <v-btn
                v-if="isPlay"
                depressed
                x-large
                color="indigo lighten-3"
                style="margin-top: 20px"
                @click="uploadVideo"
                >Upload</v-btn
              >
            </div>
          </div>
        </v-card>
      </v-col>

      <v-col cols="12" sm="6" md="4">
        <v-card color="blue-grey lighten-5" elevation="1" class="uploadVideo">
          <div class="title">Uploading my record</div>
          <v-card class="input">
            <!-- <v-row>
              <v-col cols="9">
                <v-file-input
                  style="margin: 10px 10px 10px 30px"
                  small-chips
                  truncate-length="15"
                ></v-file-input>
              </v-col>
              <v-col cols="3">

              </v-col>
            </v-row> -->
            <UploadFiles></UploadFiles>
          </v-card>
          <v-divider></v-divider>
          <div class="word">
            1. Record coughs, numbers from 0 to 9, or words like
            "Ommmmmmmmm";<br />
            2. The recording duration is 12 seconds.
          </div>
          <div
            style="text-align: center; display: flex; justify-content: center"
          >
            <img style="margin-bottom: 20px" src="../assets/coughing.svg" />
          </div>
        </v-card>
      </v-col>
    </v-row>
  </div>
</template>

<script>
import Record from "../utils/record-sdk";
import UploadFiles from "../components/UploadFiles";
export default {
  name: "HomePage",
  data() {
    return {
      isMicophone: true,
      isMicophoneButton: "Start",
      isVoice: false,
      isFinished: false,
      isPlay: false,
      tipMsg: "录音",
      audio: "",
      isupload: false,
      recorder: new Record(),
    };
  },
  components: { UploadFiles },
  created() {
    this.getLocation();
  },
  methods: {
    recordVoice() {},
    // 开始录音
    onStartVoice() {
      this.onStopAudio();
      this.isFinished = false;
      this.isMicophone = !this.isMicophone;
      this.isPlay = false;

      if (this.isMicophone) {
        this.isMicophoneButton = "Finish";
      } else {
        this.isMicophoneButton = "Start";
      }

      this.recorder.startRecord({
        success: () => {
          this.isVoice = true;
        },
        error: (e) => {
          this.isVoice = false;
          this.$toast(e);
        },
      });
    },
    getLocation() {
     this.$store.state.cityCode = returnCitySN.cname;
      console.log("OK!");
      console.log(this.$store.state.cityCode);
      // console.log('<hr><br><h1> 老铁位置：'+JSON.stringify(returnCitySN)+'</h1>')
    },
    // 结束录音
    onEndVoice() {
      this.isFinished = true;
      this.isMicophone = !this.isMicophone;

      if (this.isMicophone) {
        this.isMicophoneButton = "Finish";
      } else {
        this.isMicophoneButton = "Start";
      }

      this.recorder.stopRecord({
        success: (res) => {
          this.isVoice = false;
          //此处可以获取音频源文件(res)，用于上传等操作
          console.log("音频源文件", res);
        },
        error: () => {
          this.isVoice = false;
        },
      });
    },

    // 播放录音
    onPlayAudio() {
      this.isVoice = false;
      this.isPlay = true;
      this.audio = document.getElementById("audioVoice");
      this.recorder.play(this.audio);
    },

    // 停止播放录音
    onStopAudio() {
      this.recorder.clear(this.audio);
    },

    // 上传录音
    uploadVideo() {
      console.log(this.audio);
      this.isupload = !this.isupload;
      if (this.isupload) {
        alert("UPLOAD COMPLETE!");
        // this.$router.push({ path: "/multiple_gift/" + this.$route.params.post_id });
        this.$router.push({ path: "/ResultPage" });
      }
    },
  },
};
</script>
<style scoped>
.record {
  margin: 30px;
  padding-top: 30px;
  max-width: 960px;
  min-height: 630px;
}
.uploadVideo {
  margin: 30px;
  padding-top: 30px;
  max-width: 600px;
  min-height: 600px;
}
.word {
  margin: 30px;
  font-size: 14px;
}
.title {
  margin: 0px 30px 30px 30px;
  text-align: center;
  display: flex;
  justify-content: center;
  /* margin-top margin-right margin-bottom margin-left */
  font-size: 30px;
}
.input {
  margin: 0px 30px 30px 30px;
}
.myrecord {
  margin-top: 20px;
  text-align: center;
  display: flex;
  justify-content: center;
  /* margin-top margin-right margin-bottom margin-left */
  font-size: 30px;
}
</style>
