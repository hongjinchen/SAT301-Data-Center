<template>
  <div>
    <div v-if="currentFile">
      <div>
        <v-progress-linear
          v-model="progress"
          color="light-blue"
          height="25"
          reactive
        >
          <strong>{{ progress }} %</strong>
        </v-progress-linear>
      </div>
    </div>
    <v-row no-gutters justify="center" align="center">
      <v-col cols="9">
        <v-file-input
          show-size
          style="margin: 10px 10px 10px 30px"
          label="File input"
          @change="selectFile"
        ></v-file-input>
      </v-col>
      <v-col cols="3" class="pl-2">
        <v-btn
          class="mx-2"
          fab
          dark
          small
          color="indigo lighten-3"
          @click="upload"
        >
          <v-icon dark> mdi-upload </v-icon>
        </v-btn>
      </v-col>
    </v-row>
    <v-alert v-if="message" border="left" color="blue-grey" dark>
      {{ message }}
    </v-alert>
    <v-card v-if="fileInfos.length > 0" class="mx-auto">
      <v-list>
        <v-subheader>List of Files</v-subheader>
        <v-list-item-group color="primary">
          <v-list-item v-for="(file, index) in fileInfos" :key="index">
            <a :href="file.url">{{ file.name }}</a>
          </v-list-item>
        </v-list-item-group>
      </v-list>
    </v-card>
  </div>
</template>
<script>
import UploadService from "../services/UploadFilesService";
export default {
  name: "upload-files",
  data() {
    return {
      currentFile: undefined,
      progress: 0,
      message: "",
      fileInfos: [],
    };
  },
  methods: {
    selectFile(file) {
      this.progress = 0;
      this.currentFile = file;
    },
    // 判断文件类型
    fileType(filePath) {
      //获取最后一个.的位置
      var index = filePath.lastIndexOf(".");
      //获取后缀
      var ext = filePath.substr(index + 1);

      //判断是否是视频类型
      if (
        ["mp4", "avi", "mov", "rmvb", "rm", "flv", "3gp"].indexOf(
          ext.toLowerCase()
        ) != -1
      ) {
        return "video";
      }
      //判断是否是图片类型
      if (
        [
          "png",
          "jpg",
          "jpeg",
          "bmp",
          "gif",
          "webp",
          "psd",
          "svg",
          "tiff",
        ].indexOf(ext.toLowerCase()) != -1
      ) {
        return "image";
      }
      //判断是否是音频类型
      if (
        ["cda", "wav", "mp3", "wmv", "flac", "aac"].indexOf(
          ext.toLowerCase()
        ) != -1
      ) {
        return "audio";
      }
      return "otherType";
    },
    // 上传
    upload() {
      if (!this.currentFile) {
        this.message = "Please select a file!";
        return;
      }
      this.message = "";
      if (this.fileType(this.currentFile.name) == "video") {
        setTimeout(() => console.log(3), 3000);
        this.$router.push({ path: "/ResultPage" });
        UploadService.upload(this.currentFile, (event) => {
          this.progress = Math.round((100 * event.loaded) / event.total);
        })
          .then((response) => {
            this.message = response.data.message;
          })
          .then((files) => {
            this.fileInfos = files.data;
            this.$router.push({ path: "/ResultPage" });
          })
          .catch(() => {
            this.progress = 0;
            this.message = "Could not upload the file!";
            this.currentFile = undefined;
          });
      } else {
        alert("Please upload a video file!");
      }
    },
  },
  mounted() {},
};
</script>
