package main

import (
	"encoding/json"
	"fmt"
	
	"log"
	"net/http"
	"os"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/gmail/v1"

	"github.com/gin-gonic/gin"
)

type Credentials struct {
	Token        string `json:"token"`
	ClientID     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
}

func main() {
	router := gin.Default()
	var credentials Credentials
	// Load the credentials from file.

	creds, err := os.ReadFile("credientials.json")
	if err != nil {
		log.Fatalf("Unable to load credentials from file: %v", err)
	}

	if err := json.Unmarshal(creds, &credentials); err != nil {
		log.Fatal("Some error occuerd")
	}

	client := &oauth2.Config{
		ClientID:     credentials.ClientID,
		ClientSecret: credentials.ClientSecret,
		Scopes: []string{
			gmail.GmailReadonlyScope,
		},
		Endpoint: google.Endpoint,
	}

	router.GET("/emails", func(c *gin.Context) {

		fileContents, err := os.ReadFile("data.json")
		if err != nil {
			c.JSON(http.StatusInternalServerError, err)
		}

		var data map[string]string
		err = json.Unmarshal(fileContents, &data)
		if err != nil {
			c.JSON(http.StatusInternalServerError, err)
		}

		some, err := client.Exchange(c, "replace with server auth Code") //TODO: Replace with server auth code
		if err != nil {
			c.JSON(http.StatusInternalServerError, err)
		}
		fmt.Print(some)

		dt, err := json.Marshal(some)
		print(dt)
		if err != nil {
			c.JSON(http.StatusInternalServerError, err)
		}

		c.JSON(http.StatusOK, gin.H{"code": dt})
		// gmailService, err := gmail.NewService(c, option.WithTokenSource(config.TokenSource(c, &oauth2.Token{AccessToken: data["accessToken"]})))
		// if err != nil {
		// 	c.JSON(http.StatusInternalServerError, err)
		// 	return
		// }

		// messages, err := gmailService.Users.Messages.List("me").Do()
		// if err != nil {
		// 	c.JSON(http.StatusInternalServerError, err)
		// 	return
		// }

		// var snippets []string

		// for index, message := range messages.Messages {
		// 	if index == 5 {
		// 		break
		// 	}
		// 	msg, err := gmailService.Users.Messages.Get("me", message.Id).Do()
		// 	if err != nil {
		// 		c.AbortWithError(http.StatusInternalServerError, err)
		// 		return
		// 	}
		// 	snippets = append(snippets, msg.Snippet)
		// }

		// // Return the list of snippets as a JSON response.
		// c.JSON(http.StatusOK, gin.H{"snippets": snippets})
	})

	// router.GET("/login", func(c *gin.Context) {
	// 	// Get the access token from the request header.

	// 	data := map[string]string{"accessToken": token["accessToken"]}

	// 	file, err := json.MarshalIndent(data, "", " ")
	// 	if err != nil {
	// 		log.Fatal("Unable to store access token", err)
	// 	}

	// 	os.WriteFile("data.json", file, 0644)

	// 	c.JSON(http.StatusOK, gin.H{"message": "Authentication successfull"})

	// })

	// Start the server.
	if err := router.Run(":8000"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
