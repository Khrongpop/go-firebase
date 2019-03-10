package function

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type Result struct {
	Code    int
	Message string
}

type Body struct {
	Message string
}

type Noti struct {
	Title           string `json:title`
	Body            string `json:body`
	ClickAction     string `json:click_action`
	RegistrationIds string `json:regis_id`
}

const (
	serverKey = "AAAAorSWiIM:APA91bGFfAnMlIt20vocPKeNkQc1qrblrUT6Q1AgAtY4ZyV4howzavhKrgtIBzFHi89i0b2Z62qcOy6xQsKpcNpl3MsX98UkkbbP51vNcz5LRtno5Dv737rOjXgUjxmjrvWGJk-5djVl"
)

func Push(w http.ResponseWriter, r *http.Request) {

	data, _ := ioutil.ReadAll(r.Body)
	noti := Noti{}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	if err := json.Unmarshal(data, &noti); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(Result{
			500,
			"Internal Server Error",
		})
		return

	} else {
		json.NewEncoder(w).Encode(noti)
		return
	}
}
