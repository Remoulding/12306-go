module.exports = {
  devServer: {
    client: {
      overlay: false
    },
    // proxy: {
    //   '/api': {
    //     target: 'http://localhost:7878',
    //     changeOrigin: true,
    //     ws: true,
    //   }
    // }
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
