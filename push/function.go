package function

import (
	"encoding/json"
	"net/http"

	"github.com/google/go-gcm"
)

type Result struct {
	Code    int
	Message string
}

const (
	serverKey = "AAAAorSWiIM:APA91bGFfAnMlIt20vocPKeNkQc1qrblrUT6Q1AgAtY4ZyV4howzavhKrgtIBzFHi89i0b2Z62qcOy6xQsKpcNpl3MsX98UkkbbP51vNcz5LRtno5Dv737rOjXgUjxmjrvWGJk-5djVl"
)

func Push(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	notification := gcm.Notification{
		Title:       r.FormValue("title"),
		Body:        r.FormValue("body"),
		ClickAction: r.FormValue("clickAction"),
	}
	msg := gcm.HttpMessage{
		Data:            map[string]interface{}{"message": r.FormValue("message")},
		RegistrationIds: []string{r.FormValue("client_token")},
		Notification:    &notification,
	}
	_, err := gcm.SendHttp(serverKey, msg)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(Result{
			500,
			"Internal Server Error",
		})
		return
	} else {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(Result{
			200,
			"Push Message Success",
		})
		return
	}

}
