// const PurgecssPlugin = require('purgecss-webpack-plugin');
// const glob = require('glob-all');
const path = require('path');
// const VueLoaderPlugin = require('vue-loader/lib/plugin')

module.exports = {
  // resolve: {
  //   alias: {
  //     'vue$': 'vue/dist/vue.esm.js' // 'vue/dist/vue.common.js' for webpack 1
  //   }
  // },
  chainWebpack: config => {
    config.module
        .rule('vue')
        .use('vue-loader')
        .loader('vue-loader')
        .tap(options => {
          options['transformAssetUrls'] = {
            img: 'src',
            image: 'xlink:href',
            'b-carousel-slide': 'img-src',
            'b-embed': 'src'
          }

          return options
        })
  },

  configureWebpack: {
    resolve: {
      alias: {
        'vue$': 'vue/dist/vue.esm.js' // 'vue/dist/vue.common.js' for webpack 1
      }
    },
    // plugins: [
    //   new VueLoaderPlugin()
    // ]
  },
  // plugins: [
  //   // убедитесь что подключили плагин!
  //   new VueLoaderPlugin()
  // ],
  css: {
    loaderOptions: {
      scss: {
        prependData: '@import "@/styles/_variables.scss";'
      }
    }
  },

  pwa: {
    name: 'Acme',
    themeColor: '#5f4884',
    msTileColor: '#5f4884'
  }
};
