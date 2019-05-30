export default {
  Missing: {
    Authentication: {
      Token: 'Api Error'
    }
  },
  slack: {
    error: {
      already_invited: 'An invite has already been sent.',
      channel_not_found: 'Invite channel was not found',
      not_allowed_token_type: 'Token is not allowed',
      invite_limit_reached: 'Maximum invite limit reached',
      requires_one_channel: 'One channel is required.',
      user_disabled: 'The email address used is disabled for this team',
    },
    invite: {
      config: 'Slack Invite Config',
      title: 'Slack Invite',
      success: 'An email has just been sent to your e-mail address. Thank you',
      submit: 'Submit'
    }
  }
}
