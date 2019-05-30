package main

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/acenterastatic/static"
	"github.com/aws/aws-lambda-go/events"
	"github.com/gin-gonic/gin"
)

var (
	Response404 = events.APIGatewayProxyResponse{
		Body: fmt.Sprintf("404\n"),
		Headers: map[string]string{
			"Content-Type": "text/html",
		},
		StatusCode: 404,
	}
	Response301 = events.APIGatewayProxyResponse{
		Body: fmt.Sprintf("redirect\n"),
		Headers: map[string]string{
			"Location":     "/",
			"Content-Type": "text/html",
		},
		StatusCode: 301,
	}
	responseError = events.APIGatewayProxyResponse{
		Body: "",
		Headers: map[string]string{
			"Content-Type": "application/javascript",
		},
		StatusCode: 500,
	}
)

func init() {
}

func ReturnResponse(apiGRresp events.APIGatewayProxyResponse, err error, c *gin.Context) {
	if err != nil {

		for pName, opts := range apiGRresp.Headers {
			c.Writer.Header().Set(pName, string(opts))
		}

		// if (apiGRresp != nil) {
		if apiGRresp.Body != "" {
			in := []byte(apiGRresp.Body)
			var raw map[string]interface{}
			json.Unmarshal(in, &raw)

			c.JSON(apiGRresp.StatusCode, raw)
		} else {

			in := []byte("{}")
			var raw map[string]interface{}
			json.Unmarshal(in, &raw)

			c.JSON(apiGRresp.StatusCode, raw)
		}
		// } else {
		//      c.JSON(501, gin.h{"status": "501", "message": "invalid request."})
		// }
	} else {

		in := []byte(apiGRresp.Body)
		var raw map[string]interface{}
		json.Unmarshal(in, &raw)
		c.JSON(apiGRresp.StatusCode, raw)
	}
}

func WebsitePublic(ctx context.Context, e events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	path := e.Path
	fileNamePath := strings.TrimPrefix(strings.Replace(path, "/api/static/static/", "/api/static/", 1), "/api/")
	fileNamePath = strings.TrimPrefix(fileNamePath, "/static/")
	data, err := static.Asset(fmt.Sprintf("dist%v", path))
	fmt.Println("Finding ", fmt.Sprintf("dist%v", path), "")
	if err != nil {
		// Asset was not found.
		fmt.Println("Finding ", fmt.Sprintf("dist/%v", fileNamePath))
		data, err = static.Asset(fmt.Sprintf("dist/%v", fileNamePath))
	}
	if err != nil {

		if strings.HasSuffix(fileNamePath, ".js") || strings.HasSuffix(fileNamePath, ".css") {
			return Response404, nil
		} else {
			fmt.Println("Finding as last resort ", fmt.Sprintf("dist/%v", "index.html"))
			data, err = static.Asset(fmt.Sprintf("dist/%v", "index.html"))
		}

	}
	extension := filepath.Ext(path)

	if extension == ".css" {
		return events.APIGatewayProxyResponse{
			Body: string(data),
			Headers: map[string]string{
				"Content-Type":  "text/css",
				"Cache-Control": "public, max-age=8600",
			},
			StatusCode: 200,
		}, nil
	} else if extension == ".js" {
		return events.APIGatewayProxyResponse{
			Body: string(data),
			Headers: map[string]string{
				"Content-Type":  "application/javascript",
				"Cache-Control": "public, max-age=8600",
			},
			StatusCode: 200,
		}, nil
	} else if extension == ".gif" {
		output := base64.StdEncoding.EncodeToString(data)
		return events.APIGatewayProxyResponse{
			Body: output,
			Headers: map[string]string{
				"Content-Type":  "image/gif",
				"Cache-Control": "public, max-age=8600",
			},
			IsBase64Encoded: true,
			StatusCode:      200,
		}, nil
	} else if extension == ".png" {
		output := base64.StdEncoding.EncodeToString(data)
		return events.APIGatewayProxyResponse{
			Body: fmt.Sprintf("%v", output),
			Headers: map[string]string{
				"Content-Type":  "image/png",
				"Cache-Control": "public, max-age=8600",
			},
			IsBase64Encoded: true,
			StatusCode:      200,
		}, nil
	} else if extension == ".jpg" || extension == ".jpeg" {
		output := base64.StdEncoding.EncodeToString(data)
		return events.APIGatewayProxyResponse{
			Body: output,
			Headers: map[string]string{
				"Content-Type":  "image/jpeg",
				"Cache-Control": "public, max-age=8600",
			},
			IsBase64Encoded: true,
			StatusCode:      200,
		}, nil
	} else if extension == ".svg" {
		output := base64.StdEncoding.EncodeToString(data)
		return events.APIGatewayProxyResponse{
			Body:            output,
			IsBase64Encoded: true,
			Headers: map[string]string{
				"Content-Type":  "image/svg+xml",
				"Cache-Control": "public, max-age=8600",
			},
			StatusCode: 200,
		}, nil
	} else if extension == ".woff2" {
		output := base64.StdEncoding.EncodeToString(data)
		return events.APIGatewayProxyResponse{
			Body:            output,
			IsBase64Encoded: true,
			Headers: map[string]string{
				"Content-Type":  "font/woff2",
				"Cache-Control": "public, max-age=8600",
			},
			StatusCode: 200,
		}, nil
	} else if extension == ".html" || extension == "" {
		// TODO: Detect language ??
		strTitle := os.Getenv("TITLE")
		if strTitle == "" {
			strTitle = os.Getenv("SITE_TITLE")
		}
		dataWithTitle := strings.Replace(string(data), "%%TITLE%%", strTitle, -1) //TODO: Customize name using CloudFormation Title
		if strTitle == "" {
			dataWithTitle = strings.Replace(string(data), "%%TITLE%%", "Serverless Portal", -1) //TODO: Customize name using CloudFormation Title
		}
		return events.APIGatewayProxyResponse{
			Body: dataWithTitle,
			Headers: map[string]string{
				"Content-Type":  "text/html",
				"Cache-Control": "public, must-revalidate, proxy-revalidate, max-age=0",
			},
			StatusCode: 200,
		}, nil
		// //fmt.Printf("File type: %s. MIME: %s\n", kind.Extension, kind.MIME.Value)
	} else {
		return Response404, nil
	}
}
