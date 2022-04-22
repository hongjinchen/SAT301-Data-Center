<template>
      <div :class="{'voice_wrap':true,'large':duration>40,'middle':duration>20&&duration<41,'small':duration<21}">
    <div>
      <v-icon v-if="isPlaying" @click="playVoice"> mdi-pause </v-icon>
      <v-icon v-if="!isPlaying" @click="playVoice"> mdi-play </v-icon>
    </div>
    <div class="voice_time_line" id="waveform" ref="waveform" />
    <div class="voice-dialog-time">
      <p>{{ time }} </p>
    </div>
  </div>
</template>

<script>
import WaveSurfer from "wavesurfer.js";
import CursorPlugin from "wavesurfer.js/dist/plugin/wavesurfer.cursor.js"; //导入指针轴插件
export default {
  props: {
    rebot: {
      type: Object,
    },
    person: {
      type: Object,
    },
    url: {
      type: String,
      default: require("../assets/test.mp3"),
    },
    toStopVoice: {
      type: Boolean,
      default: false,
    },
    loadWave: {
      type: Boolean,
      default: true,
    },
  },
  data() {
    return {
      wavesurfer: null, // 波形图对象
      time: "00:00",
      isPlaying: false,
    };
  },
  mounted() {
    this.$nextTick(() => {
      if (this.loadWave) {
        this.loadVoice();
      }
    });
  },
  methods: {
    // 转化数字为字符串 如果小于0  返回 0x (01 02...)
    buZero(num) {
      return num > 9 ? num : "0" + num;
    },
    // 加载或者创建波形图
    loadVoice(bool) {
      if (this.wavesurfer) {
        if (bool) {
          this.time = "00:00";
          this.wavesurfer.load(this.url);
        }
      } else {
        this.wavesurfer = WaveSurfer.create({
          container: this.$refs.waveform,
          waveColor: "#65da74",
          progressColor: "purple",
          height: "60",
          backgroundColor: "#f7f3f4",
          scrollParent: true,
          plugins: [
            CursorPlugin.create({
              showTime: true,
              opacity: 1,
              customShowTimeStyle: {
                "background-color": "#000",
                color: "#fff",
                padding: "2px",
                "font-size": "10px",
              },
            }),
          ],
        });
        this.wavesurfer.load(this.url);
        this.wavesurfer.on("audioprocess", (e) => {
          const times = Math.ceil(e);
          this.time =
            this.buZero(Math.floor(times / 60)) + ":" + this.buZero(times % 60);
        });
        this.wavesurfer.on("finish", () => {
          this.isPlaying = false;
        });
      }
    },
    // 开始
    playVoice() {
      if (this.wavesurfer) {
        if (this.wavesurfer.isPlaying()) {
          this.isPlaying = false;
          this.wavesurfer.pause();
        } else {
          this.isPlaying = true;
          this.wavesurfer.play();
        }
      }
    },
  },
  watch: {
    loadWave: function () {
      this.playVoice();
    },
    url: function () {
      this.playVoice(true);
    },
    toStopVoice: function () {
      if (this.wavesurfer) {
        this.wavesurfer.pause();
      }
    },
  },
};
</script>
<style lang="scss" scoped>
.voice-dialog-content {
  display: flex;
  align-items: center;
}
#waveform {
  margin-left: 4px;
  margin-bottom: 4px;
  width: 98%;
  height: 60px;
}
.voice-dialog-time {
  font-size: 16px;
  line-height: 22px;
  color: #fff;
  margin-left: 8px;
}
.voice_wrap{
    background: #94BDEC;
    border-radius: 18px;
    height: 80px;
    display: flex;
    align-items: center;
    padding: 0 12px;
    box-sizing: border-box;
    img{
      width: 24px;
      height: 24px;
      margin-right: 4px;
    }
    .voice_time_line{
      height: 24px;
    }
    .duration{
      font-size: 16px;
      line-height: 22px;
      color: #fff;
      margin-left: 8px;
    }
  }
  .voice_wrap.large{
    width: 262px;
    .voice_time_line{
      width: 177px;
    }
  }
  .voice_wrap.middle{
    width: 203px;
    .voice_time_line{
      width: 118px;
    }
  }
  .voice_wrap.small{
    width: 144px;
    .voice_time_line{
      width: 59px;
    }
  }
</style>
