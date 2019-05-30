import Vue from 'vue'
import Vuetify from 'vuetify'

// # Importing jQuery for plugin load... temp addition
window.$ = require('jquery')
window.JQuery = require('jquery')
window._ = require('underscore')
window.select = require('select.js')

const fixConsoleLog = {
  log: function() {

  },
  error: function() {

  }
}

Vue.use(Vuetify, {
  // theme: {
  //   primary: colors.indigo.base, // #E53935
  //   secondary: colors.indigo.lighten4, // #FFCDD2
  //   accent: colors.indigo.base // #3F51B5
  // },
  options: {
    themeVariations: ['primary', 'secondary', 'accent'],
    extra: {
      mainToolbar: {
        color: 'primary'
      },
      sideToolbar: {
      },
      sideNav: 'primary',
      mainNav: 'primary lighten-1',
      bodyBg: ''
    }
  }
})
window.console = window.console || fixConsoleLog

import Cookies from 'js-cookie'

import 'normalize.css/normalize.css' // A modern alternative to CSS resets

import Element from 'element-ui'
import 'element-ui/lib/theme-chalk/index.css'

import '@/styles/index.scss' // global css

import './permission' // permission control

import App from './App'
import router from './router'
import store from './store'

//global registration
import i18n from './lang' // Internationalization

// Import Amplify aws
import Amplify from 'aws-amplify'
window.Amplify = Amplify

Vue.use(Element, {
  size: Cookies.get('size') || 'medium', // set element-ui default size
  i18n: (key, value) => i18n.t(key, value)
})

Vue.config.productionTip = false

const app = new Vue({
  el: '#app',
  router,
  store,
  i18n,
  render: h => h(App)
})
window.app = app
