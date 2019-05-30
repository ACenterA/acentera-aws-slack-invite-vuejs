<template>
  <div :style="[custom_style]" class="app-container bg">
    <div v-if="status === 'success'" class="title-container">
      <div class="title-container">
        <h3 class="title"> {{ invite_success }} </h3>
        <a v-if="slack_url" :href="slack_url">{{ slack_url }}</a>
      </div>
    </div>
    <div v-else>
      <div class="title-container">
        <h3 class="title"> {{ invite_title }} </h3>
      </div>
      <p> {{ invite_message }} </p>
      <el-form ref="slackInviteForm" :model="slackInviteForm" :rules="slackInviteRules" class="login-form" auto-complete="on" label-position="left" @submit.prevent.native="">
        <el-form-item prop="email">
          <span class="svg-container">
            <svg-icon icon-class="user" />
          </span>
          <el-input
            ref="email"
            v-model="slackInviteForm.email"
            placeholder="email@example.com"
            name="email"
            type="text"
            tabindex="1"
            auto-complete="on"
            @keyup.enter.native="handleSlackInvite"
          />
        </el-form-item>

        <div v-if="isLoaded()">
          <vue-recaptcha
            v-if="settings.recaptchaSiteKey !== ''"
            ref="recaptcha"
            :sitekey="settings.recaptchaSiteKey"
            size="invisible"
            @verify="onCaptchaVerified"
            @expired="onCaptchaExpired"/>
        </div>

        <el-button :loading="loading" :disabled="!isLoaded()" type="primary" style="width:100%;margin-bottom:30px;margin-top:15px;" @click.native.prevent="handleSlackInvite">{{ $t('slack.invite.submit' ) }}</el-button>

      </el-form>
    </div>
    <div class="powered_by">Made possible by <a href="https://acentera.com/">ACenterA</a></div>
  </div>
</template>

<script>
import { mapGetters } from 'vuex'
import VueRecaptcha from 'vue-recaptcha'
import { validateEmail } from '@/utils/validate'
import SvgIcon from '@/components/SvgIcon'

export default {
  name: 'Home',
  components: {
    'vue-recaptcha': VueRecaptcha,
    SvgIcon
  },
  data() {
    const validateEmailValidator = (rule, value, callback) => {
      console.error(validateEmail)
      if (!validateEmail(value)) {
        callback(new Error('Please enter a valid email address'))
      } else {
        callback()
      }
    }
    return {
      slackInviteForm: {
        email: ''
      },
      slackInviteRules: {
        email: [{ required: true, trigger: 'blur', validator: validateEmailValidator }]
      },
      status: 'pending',
      sucessfulServerResponse: '',
      serverError: '',
      loading: false
    }
  },
  computed: {
    ...mapGetters([
      'settings'
    ]),
    invite_title() {
      if (!this.isLoaded()) {
        return ''
      }
      if (this.settings.SlackInviteTitle) {
        return this.settings.SlackInviteTitle
      }
      return this.$t('slack.invite.title')
    },
    slack_url() {
      if (!this.isLoaded()) {
        return ''
      }
      if (this.settings.SlackId) {
        return 'https://' + this.settings.SlackId + '.slack.com/'
      }
      return ''
    },
    invite_success() {
      if (!this.isLoaded()) {
        return ''
      }
      if (this.settings.SlackInviteSuccess) {
        return this.settings.SlackInviteSuccess
      }
      return this.$t('slack.invite.success')
    },
    invite_message() {
      if (!this.isLoaded()) {
        return ''
      }
      if (this.settings.SlackInviteMessage) {
        return this.settings.SlackInviteMessage
      }
      return ''
    },
    custom_style() {
      var style = {}
      if (this.settings.HtmlTextColor) {
        style['color'] = this.settings.HtmlTextColor
      }
      if (this.settings.backgroundImage !== '') {
        style['background-image'] = 'url("' + this.settings.backgroundImage + '")'
      } else {
        style['background-image'] = 'url("/static/img/background.jpg")'
      }
      return style
    }
  },
  created() {
    this.$store.dispatch('GetSiteSettings').then(function(r) {
      // All good, wee got everything we needed
    }).catch((err) => {
      console.error(err)
      var errorMsg = 'slack.invite.network_error'
      if (err && err.data && err.data.message) {
        errorMsg = err.data.message.replace(/ /g, '.')
      }
      this.$notify({
        title: this.$t('slack.invite.config'),
        message: this.$t(errorMsg),
        type: 'error',
        duration: 2000
      })
    })
  },
  methods: {
    ...mapGetters([
      'isLoaded'
    ]),
    handleSlackInvite: function() {
      // this.status = "submitting";
      this.$refs.recaptcha.execute()
    },
    onCaptchaVerified: function(recaptchaToken) {
      const self = this
      self.status = 'submitting'
      self.$refs.recaptcha.reset()
      self.$store.dispatch('SlackSignup', {
        email: self.slackInviteForm.email,
        recaptchaToken: recaptchaToken
      }).then(function(response) {
        self.$notify({
          title: self.$t('slack.invite.title'),
          message: self.$t('slack.invite.success'),
          type: 'success',
          duration: 2000
        })
        self.status = 'success'
      }).catch((err) => {
        var responseBody = err.response
        if (!responseBody) {
          responseBody = err
        } else {
          responseBody = err.response.data || responseBody
        }
        responseBody = responseBody.data || responseBody
        console.error(responseBody)
        var msg = responseBody.error || responseBody.message || JSON.stringify(responseBody)
        if (responseBody.error) {
          msg = 'slack.error.' + msg
        }
        self.sttatus = 'error'
        if (msg === 'slack.error.already_invited') {
          self.status = 'success'
        }
        self.$notify({
          title: self.$t('slack.invite.title'),
          message: self.$t(msg),
          type: self.status,
          duration: 2000
        })
      })
    },
    onCaptchaExpired: function() {
      this.$refs.recaptcha.reset()
    }
  }
}
</script>

<style>
.bg {
  background: no-repeat center center fixed;
  -webkit-background-size: cover;
  -moz-background-size: cover;
  -o-background-size: cover;
  o-background-size: cover;
}
</style>
