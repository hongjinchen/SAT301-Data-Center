<script>
//首页
const HOME = 'https://www.yangshuaionline.com/'
//保存全局状态
import Vue from 'vue'
let State = new Vue({
  data() {
    return {
      isLogin: false,
    }
  },
})
export default {
  HOME,
  State,
}
</script>