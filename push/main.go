// package function

// import (
// 	"context"
// 	"fmt"
// 	"net/http"

// 	firebase "firebase.google.com/go"

// 	"google.golang.org/api/option"
// )

// func Push(w http.ResponseWriter, r *http.Request) {
// 	opt := option.WithCredentialsFile("./serviceAccountKey.json")
// 	app, err := firebase.NewApp(context.Background(), nil, opt)
// 	if err != nil {
// 		return nil, fmt.Errorf("error initializing app: %v", err)
// 	}
// }

// thesis-4ef45

package main

import (
	"crypto/rand"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/go-gcm"
	"github.com/labstack/echo"
)

func main() {

	SendGMToClient2("Hello from GCM", "eHScHsD-hNU:APA91bHis1hmhOpoxGTSTd09Es_Y8XIMxKoTsytR7ounjtaEMxs1yLCD9jkAgid2vc3igLKjbmtS8HU8EpsWieAHyv0KJ8wU8BZ94PKwLqsMXVltGDkm1rdNO_wyA1dPBjWgfRHRduqK")

	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.POST("/api/v1/push-notif", SendGMToClient)
	e.Logger.Fatal(e.Start(":8080"))

	// opt := option.WithCredentialsFile("./serviceAccountKey.json")
	// app, err := firebase.NewApp(context.Background(), nil, opt)
	// if err != nil {
	// 	// return nil, fmt.Errorf("error initializing app: %v", err)
	// 	fmt.Printf("error initializing app: %v", err)
	// }
	// fmt.Println(app)

}

func SendGMToClient2(pushText string, pushToken string) {
	serverKey := "AAAAorSWiIM:APA91bGFfAnMlIt20vocPKeNkQc1qrblrUT6Q1AgAtY4ZyV4howzavhKrgtIBzFHi89i0b2Z62qcOy6xQsKpcNpl3MsX98UkkbbP51vNcz5LRtno5Dv737rOjXgUjxmjrvWGJk-5djVl"
	var msg gcm.HttpMessage
	data := map[string]interface{}{"message": pushText}
	regIDs := []string{pushToken}
	msg.RegistrationIds = regIDs
	msg.Data = data

	response, err := gcm.SendHttp(serverKey, msg)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		// fmt.Println("Response ", response)
		fmt.Println("Response ", response.Success)
		fmt.Println("MessageID ", response.MessageId)
		fmt.Println("Failure ", response.Failure)
		fmt.Println("Error ", response.Error)
		fmt.Println("Results ", response.Results)
	}
}

// SendGMToClient is a function that will push a message to client
func SendGMToClient(c echo.Context) error {
	serverKey := "AAAAorSWiIM:APA91bGFfAnMlIt20vocPKeNkQc1qrblrUT6Q1AgAtY4ZyV4howzavhKrgtIBzFHi89i0b2Z62qcOy6xQsKpcNpl3MsX98UkkbbP51vNcz5LRtno5Dv737rOjXgUjxmjrvWGJk-5djVl"
	var msg gcm.HttpMessage
	data := map[string]interface{}{"message": c.FormValue("message")}
	regIDs := []string{c.FormValue("client_token")}
	msg.RegistrationIds = regIDs
	msg.Data = data
	response, err := gcm.SendHttp(serverKey, msg)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		// fmt.Println("Response ", response)
		fmt.Println("Response ", response.Success)
		fmt.Println("MessageID ", response.MessageId)
		fmt.Println("Failure ", response.Failure)
		fmt.Println("Error ", response.Error)
		fmt.Println("Results ", response.Results)
	}
	// return c.JSON(http.StatusOK, gin.H{
	// 	"message":      c.FormValue("message"),
	// 	"client_token": c.FormValue("client_token"),
	// })

	t := time.Now()
	uuid, errUUID := newUUID()
	if errUUID != nil {
		fmt.Printf("error: %v\n", err)
	}

	if err != nil {
		return c.JSON(http.StatusBadRequest, gin.H{
			"requestID": uuid,
			"now":       t.Format("2006/01/02 15:04:05"),
			"code":      strconv.Itoa(http.StatusBadRequest) + "02",
			"message":   err.Error(),
			"data":      "[]",
		})
	} else {
		return c.JSON(http.StatusOK, gin.H{
			"requestID": uuid,
			"now":       t.Format("2006/01/02 15:04:05"),
			"code":      strconv.Itoa(http.StatusOK) + "01",
			"message":   response.Error,
			"data":      response.Results,
		})
	}
}

// newUUID generates a random UUID according to RFC 4122
func newUUID() (string, error) {
	uuid := make([]byte, 16)
	n, err := io.ReadFull(rand.Reader, uuid)
	if n != len(uuid) || err != nil {
		return "", err
	}
	// variant bits; see section 4.1.1
	uuid[8] = uuid[8]&^0xc0 | 0x80
	// version 4 (pseudo-random); see section 4.1.3
	uuid[6] = uuid[6]&^0xf0 | 0x40
	return fmt.Sprintf("%x-%x-%x-%x-%x", uuid[0:4], uuid[4:6], uuid[6:8], uuid[8:10], uuid[10:]), nil
}
