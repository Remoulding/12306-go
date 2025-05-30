module.exports = {
  devServer: {
    port: process.env.PORT || 12306,
    host: '0.0.0.0', // 允许容器外访问
    client: {
      overlay: false
    },
  },
  css: {
    loaderOptions: {
      less: {
        javascriptEnabled: true,
        modifyVars: {
          'border-radius-base': '4px',
          'card-radius': '4px'
        }
      }
    }
  }
}
