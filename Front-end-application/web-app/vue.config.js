const { defineConfig } = require('@vue/cli-service')
module.exports = defineConfig({
  transpileDependencies: [
    'vuetify'
  ],
  devServer: {
    proxy:'http://42.192.21.119:8080'
  },
  pwa: {
    name: 'SAT301 Detection of New Coronary Pneumonia Disease',
    themeColor: '#4c89fe',
    msTileColor: '#4c89fe',
    manifestOptions: {
      start_url: '.',
      background_color: '#4c89fe'
    },
    workboxPluginMode: 'GenerateSW',
    workboxOptions: {
    }
  },
  // 名字和颜色涉及添加至桌面的应用名，及桌面进入的启动页面的样式
  lintOnSave: true
})
