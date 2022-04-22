<template>
  <div :class="{'voice_wrap':true,'large':duration>40,'middle':duration>20&&duration<41,'small':duration<21}">
    <v-icon v-if="!isplay" alt="" @click="play">mdi-play</v-icon>
    <v-icon v-else  alt="" @click="pause" >mdi-pause</v-icon>
    <div class="voice_time_line" :id="'wave'+id"></div>
    <div class="duration">{{duration}}"</div>
  </div>
</template>

<script>
//引入wavesurfer.js
import WaveSurfer from 'wavesurfer.js' //导入wavesurfer.js
export default {
  props:{
    id:Number,
    url:String,
    duration:Number
  },
  data(){
    return {
      isplay:false,
      wavesurfer: null,
    }
  },
  mounted(){
    let that = this;
    this.$nextTick(()=>{
      that.wavesurfer = WaveSurfer.create({
        height:24,
		container: '#wave' + that.id,
        cursorColor:'transparent',
        maxCanvasWidth:60,
        progressColor:'rgba(255,255,255,1)',
        waveColor:'rgba(255,255,255,0.4)',
        barGap:1.8,
        barWidth:1.2,
        barHeight:16,
        barMinHeight:0.5
      });
      // 特别提醒：如果是本地文件此处需要使用require(相对路径)，否则会报错
      that.wavesurfer.load(that.url);
      that.wavesurfer.on('finish', function () {
// eslint-disable-next-line no-mixed-spaces-and-tabs
		  console.log('播放完毕：finish');
        that.isplay = false;
        that.wavesurfer.stop();
      });
    })
  },
  methods:{
    play(){
      this.isplay = true;
      this.wavesurfer.play();
    },
    pause(){
      this.isplay = false;
      this.wavesurfer.pause();
    }
  }
}
</script>

<style lang="scss">
  .voice_wrap{
    background: #94BDEC;
    border-radius: 18px;
    height: 36px;
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

