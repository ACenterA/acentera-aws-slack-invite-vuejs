import { sendSlackInvite } from '@/api/app.js'

const app = {
  state: {
    customClass: ''
  },
  getters: {
  },
  mutations: {
  },
  actions: {
    SlackSignup({ commit, state }, data) {
      return new Promise((resolve, reject) => {
        sendSlackInvite(data).then(function(response) {
          if (response && response.data) {
            resolve(response)
          } else {
            reject(rsponse)
          }
        }).catch((err) => {
          reject(err)
        })
      })
    }
  }
}

export default app
