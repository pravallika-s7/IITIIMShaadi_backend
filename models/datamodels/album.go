package datamodels

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"gorm.io/gorm"
)

func Notifications(db *gorm.DB) {

	jsonBody := map[string]string{"token": "d53f61981a50b6e5baad02eec136db6c", "friend_user": "83441"}
	jsonValue, err := json.Marshal(jsonBody)

	response, err := http.Post("https://www.iitiimshaadi.com/apis/all_notifications.json", "application/json", bytes.NewBuffer(jsonValue))
	if err != nil {
		log.Fatalln(err)
	}

	defer response.Body.Close()
	bodyBytes, _ := ioutil.ReadAll(response.Body)

	var notification map[string]interface{}
	json.Unmarshal(bodyBytes, &notification)
	notify := notification["getAllNotificaitons"].([]interface{})
	//db.Create(&Notification{})
	for i := range notify {
		n := notify[i].(map[string]interface{})
		db.Create(&Notification{ToId: int(n["id"].(float64)),
			Type:   n["type"].(string),
			Msg:    n["message"].(string),
			Status: int(n["status"].(float64)),
		})
	}
}
