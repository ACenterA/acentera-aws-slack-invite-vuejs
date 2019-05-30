import Vue from 'vue'
import Vuex from 'vuex'
import app from './modules/app'
import settings from './modules/settings'
import getters from './getters'
Vue.use(Vuex)

const store = new Vuex.Store({
  modules: {
    app,
    settings
  },
  getters
})
window.store = store

export default store
