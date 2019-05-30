package main

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"encoding/json"
	"io/ioutil"
	"net/url"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/awslabs/aws-lambda-go-api-proxy/gin"
	"github.com/gin-gonic/gin"
)

var ginLambda *ginadapter.GinLambda

func init() {
}

type SiteSettings struct {
	ID                 int    `db:"id, json:"Id"`
	Title              string `db:"title, json:"Title"`
	SlackUrl           string `db:"slack_url, json:"SlackUrl"`
	SlackToken         string `db:"slack_token, json:"SlackToken"`
	SlackInviteToken   string `db:"slack_invite_token, json:"SlackInviteToken"`
	SlackInviteTitle   string `db:"slack_invite_title, json:"SlackInviteTitle"`
	SlackInviteMessage string `db:"slack_invite_message, json:"SlackInviteMessage"`
	SlackId            string `db:"slack_id, json:"SlackId"`
	BackgroundImage    string `db:"background_image, json:"BackgroundImage"`
	HtmlTextColor      string `db:"html_text_color, json:"HmlTextColor"`
	RecaptchaSiteKey   string `db:"recaptcha_site, json:"RecaptchaSiteKey"`
	CommunityName      string `db:"recaptcha_site, json:"CommunityName"`
}

type InviteData struct {
	Email          string `json:"email"`
	RecaptchaToken string `json:"recaptchaToken"`
}

type RecaptchaResponse struct {
	Success     bool      `json:"success"`
	Score       float64   `json:"score"`
	Action      string    `json:"action"`
	ChallengeTS time.Time `json:"challenge_ts"`
	Hostname    string    `json:"hostname"`
	ErrorCodes  []string  `json:"error-codes"`
}

type SlackResponse struct {
	Ok      bool   `json:"ok"`
	Error   string `json:"error"`
	Warning string `json:"warning"`
}

const recaptchaServerName = "https://www.google.com/recaptcha/api/siteverify"

// Check uses the client ip address and the challenge code
func ReCaptchaCheck(remoteip, response string) (r RecaptchaResponse, err error) {

	recaptchaPrivateKey := os.Getenv("RECAPTCHA_SECRET")
	fmt.Println("KEY IS : ", recaptchaPrivateKey)
	resp, err := http.PostForm(recaptchaServerName, url.Values{"secret": {recaptchaPrivateKey}, "remoteip": {remoteip}, "response": {response}})
	if err != nil {
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}
	err = json.Unmarshal(body, &r)
	if err != nil {
		fmt.Println("Read error: got invalid JSON:")
		fmt.Println(err)
		return
	}
	return
}

// Slack invite email
func InviteUser(email string) (r SlackResponse, err error) {
	// SlackResponse

	slackInviteUrl := fmt.Sprintf("https://%s.slack.com/api/users.admin.invite", os.Getenv("SLACK_TEAM"))
	resp, err := http.PostForm(slackInviteUrl, url.Values{"email": {email}, "token": {os.Getenv("SLACK_OAUTH_TOKEN")}, "channels": {os.Getenv("SLACK_CHANNEL")}, "set_active": {"true"}})
	if err != nil {
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}
	err = json.Unmarshal(body, &r)
	if err != nil {
		fmt.Println("Read error: got invalid JSON:")
		fmt.Println(err)
		return
	}
	return
}

func getSiteSettings(c *gin.Context) {
	_siteSettings := &SiteSettings{
		ID:                 1,
		RecaptchaSiteKey:   os.Getenv("RECAPTCHA_SITE_KEY"),
		SlackInviteTitle:   os.Getenv("SLACK_INVITE_TITLE"),
		SlackInviteMessage: os.Getenv("SLACK_INVITE_MESSAGE"),
		BackgroundImage:    os.Getenv("BACKGROUND_IMAGE"),
		HtmlTextColor:      os.Getenv("HTML_TEXT_COLOR"),
		SlackId:            os.Getenv("SLACK_TEAM"),
	}
	c.JSON(http.StatusOK, _siteSettings)
}
func postInvite(c *gin.Context) {

	var err error
	c.Header("Content-Type", "application/json; charset=utf-8")

	p := InviteData{}
	if err = c.BindJSON(&p); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":  "json decoding : " + err.Error(),
			"status": http.StatusBadRequest,
		})
		return
	}
	resp, err := ReCaptchaCheck(c.ClientIP(), p.RecaptchaToken)
	if err != nil || !resp.Success {
		c.JSON(http.StatusForbidden, resp)
		return
	}

	respSlack, errSlack := InviteUser(p.Email)
	if errSlack != nil || !respSlack.Ok {
		c.JSON(http.StatusForbidden, respSlack)
		return
	}
	// Ok at this point we can submit to the slack api.
	c.JSON(http.StatusOK, respSlack)
	return
}

func Handler(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	if ginLambda == nil {
		// stdout and stderr are sent to AWS CloudWatch Logs
		r := gin.Default()

		// REST API Queries here
		r.GET("/api/settings", getSiteSettings)
		r.POST("/api/invite", postInvite)
		ginLambda = ginadapter.New(r)
	}

	// If no name is provided in the HTTP request body, throw an error
	return ginLambda.Proxy(req)
}

func main() {
	if os.Getenv("TYPE") == "WEBSITE" {
		lambda.Start(WebsitePublic)
	} else {
		lambda.Start(Handler)
	}
}
