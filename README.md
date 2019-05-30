ACenterA Slack Invite Automation
------------

Vue.JS Slack Invite Project

If this app has been deployed fully using the [serverless repository link](https://serverlessrepo.aws.amazon.com/applications/arn:aws:serverlessrepo:us-east-1:356769441913:applications~slack-invite), once the cloudformation stack has been created, you should access the web application from the *WebsiteUrl cloudformation outputs*.

# ENV File Informations

    $ cat .env

| Environment   | Value          | Info  |
| ------------- |---------------| -----:|
| SITE_TITLE | Slack Invite - Local | Vue.JS Site Title |
| ENV_CONFIG | local | Specify local to use local graphql, anything else would use values from the DB |
| SLACK_OAUTH_TOKEN | xxxxx | Your access token for Slack. (see [Issue token](#issue-token)) |
| RECAPTCHA_SITE_KEY | xxxxx | ReCaptcha Site Key |
| RECAPTCHA_SECRET | xxxxx | ReCaptcha Secret Key |
| SLACK_OAUTH_TOKEN | xox.... | Slack Admin Invite OAuth Token |
| SLACK_TEAM | xxxxx | Slack Team Id / Name |
| SLACK_CHANNEL | | (Optional) Slack Channel ID to invite user |
| SLACK_INVITE_TITLE | Automatic invites to the Public Slack channel! | H1 Heading |
| SLACK_INVITE_MESSAGE | Want to help? Enter your email address and join us. | Text under the H1 Heading |
| HTML_TEXT_COLOR | black | HTML Text color to fit with the background |

# How to launch a local Development environment

Using Makefile

    $ make dev

Using Docker Compose only

    $ docker-compose -f docker-compose.yml up -d --build --force-recreate


# Endpoints

| Endpoint Name  | Value          | Info  |
| ------------- |---------------| -----:|
| Vue.JS App | http://127.0.0.1/ | Vue.JS App |
| API Gateway Proxy | http://127.0.0.1:2000/ | API Gateway Proxy (acentera core or plugin routing) |

# Launching using AWS Serverless Repo

Once you have deployed the serverless repo through the Serverless Repository, you should have a Cloudfront distributin created.

You may visit the *aws cloudformation* and you will find in the the stack output the cloudfront WebsiteUrl.


## Issue token
**You should generate the token in admin user, not owner.** If you generate the token in owner user, a `missing_scope` error may occur.

### OAuth tokens

1. Visit <https://api.slack.com/apps> and click Create New App.

2. Click "Permissions".

3. In "OAuth & Permissions" page, select `admin` scope under "Permission Scopes" menu and save changes.

4. Click "Install App to Workspace".

5. Return to your app Basic informations

Set `SlackOauthToken` or `SLACK_OAUTH_TOKEN`


## reCAPTCHA
Register a new site in [Google reCAPTHCA](https://www.google.com/recaptcha/)
as reCAPTCHA v2 type [Admin](https://www.google.com/recaptcha/admin/).

Set "Site key" as `RecaptchaSiteKey` or `RECAPTCHA_SITE_KEY`,
and "Secret key" as `RecaptchaSecret` or `RECAPTCHA_SECRET`.


# License
The software is licensed under the MIT license.

# Attribution
By using this UI, we would like that any application which incorporates it shall prominently display the message “Made with ACenterA” in a legible manner in the footer of the admin console. This message must open a link to acentera.com when clicked or touched.
