import store from '@/store'
import { getSiteSettings } from '@/api/app.js'

const settings = {
  state: {
    isLoaded: false,
    allowRegister: false,
    recaptchaSecret: '',
    recaptchaSiteKey: '', 
    inviteToken: '',
    slackToken: '',
    backgroundImage: '',
    SlackInviteMessage: '',
    SlackInviteSuccess: '',
    HtmlTextColor: 'black',
    SlackInviteTitle: '',
    SlackId: '',
    communityName: '',
    title: '',
    stage: '',
    slackUrl: ''
  },
  getters: {
    settings(state) {
      return state
    },
    isLoaded(state) {
      return state.isLoaded === true
    }
  },
  mutations: {
    SET_RECAPCHA_SECRET: (state, recaptchaKey) => {
      state.recaptchaSecret = recaptchaSecret
    },
    SET_RECAPCHA_SITE_KEY: (state, recaptchaSiteKey) => {
      state.recaptchaSiteKey = recaptchaSiteKey
    },
    SET_INVITE_TOKEN: (state, inviteToken) => {
      state.inviteToken = inviteToken
    },
    SET_SLACK_TOKEN: (state, slackToken) => {
      state.slackToken = slackToken
    },
    SET_LOADED: (state, val) => {
      state.isLoaded = val
    },
    SET_SLACK_URL: (state, slackUrl) => {
      state.slackUrl = slackUrl
    },
    SET_COMMUNITY_NAME: (state, communityName) => {
      state.communityName = communityName
    },
    SET_SITE_TITLE: (state, title) => {
      state.title = title
    },
    SET_SLACK_INVITE_SUCCESS: (state, msg) => {
      state.SlackInviteSuccess = msg
    },
    SET_SLACK_INVITE_MESSAGE: (state, msg) => {
      state.SlackInviteMessage = msg
    },
    SET_BACKGROUND_IMAGE:(state, backgroundImage) => {
      state.backgroundImage = backgroundImage
    },
    SET_HTML_TEXT_COLOR: (state, color) => {
      state.HtmlTextColor = color
    },
    SET_SLACK_ID: (state, slackId) => {
      state.SlackId = slackId
    },
    SET_SLACK_INVITE_TITLE: (state, title) => {
      state.SlackInviteTitle = title
    },
  },
  actions: {
    UpdateSiteSettings({ commit, state }, input) {
      var data = input.data || input
      commit('SET_RECAPCHA_SITE_KEY', data.RecaptchaSiteKey)
      commit('SET_SLACK_URL', data.SlackUrl)
      commit('SET_SLACK_ID', data.SlackId)
      console.error(data)
      commit('SET_SLACK_INVITE_TITLE', data.SlackInviteTitle)
      commit('SET_SLACK_INVITE_MESSAGE', data.SlackInviteMessage)
      commit('SET_SLACK_INVITE_SUCCESS', data.SlackInviteSuccess)
      commit('SET_BACKGROUND_IMAGE', data.BackgroundImage)
      commit('SET_HTML_TEXT_COLOR', data.HtmlTextColor)
      
      // commit('SET_COMMUNITY_NAME', data.communityName)
      // commit('SET_SITE_TITLE', data.SiteTitle)
      // TODO: commit('SET_LOCALE', data.locale)
      commit('SET_LOADED', true) // FORCE LOADED
    },
    GetSiteSettings({ commit, state }) {
      return new Promise((resolve, reject) => {
        getSiteSettings().then(function(response) {
          if (state.isLoaded || (window.preventLoop === true)) { // preventLoop not really works when using Uglify and chunks ?
            resolve('loaded')
          } else {
            window.preventLoop = true
            if (response && response.data) {
              store.dispatch('UpdateSiteSettings', response.data).then(() => {
                resolve(state)
              })
            } else {
              reject(response)
            }
          }
        }).catch((er) => {
          reject(er)
        })
      })
    }
  }
}

export default settings
