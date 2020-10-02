package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"models/datamodels"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {

	db, err := gorm.Open(sqlite.Open("UserProfile.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	db.AutoMigrate(&datamodels.User{}, &datamodels.Profile{}, &datamodels.Interest{}, &datamodels.Family{},
		&datamodels.Education{}, &datamodels.Job{}, &datamodels.Address{}, &datamodels.Other{},
		&datamodels.PartnerChoice{}, &datamodels.PartMartSt{}, &datamodels.EmailData{}, &datamodels.VerificationData{}, &datamodels.Friend{},
		&datamodels.SentChat{}, &datamodels.ReceivedChat{}, &datamodels.Subscription{}, &datamodels.Album{}, &datamodels.Contact{},
		&datamodels.FolderListing{}, &datamodels.SuccessStory{}, &datamodels.MediaCoverage{}, &datamodels.Notification{})

	jsonBody := map[string]string{"token": "d53f61981a50b6e5baad02eec136db6c", "userId": "93714"}
	jsonValue, err := json.Marshal(jsonBody)

	response, err := http.Post("https://www.iitiimshaadi.com/apis/my_profile.json", "application/json", bytes.NewBuffer(jsonValue))

	if err != nil {
		log.Fatalln(err)
	}

	defer response.Body.Close()
	bodyBytes, _ := ioutil.ReadAll(response.Body)

	//Populating data in each table
	datamodels.UserTable(db, bodyBytes)
	datamodels.ProfileTable(db, bodyBytes)
	datamodels.FamilyTable(db, bodyBytes)
	datamodels.EducationTable(db, bodyBytes)
	datamodels.JobTable(db, bodyBytes)
	datamodels.AddressTable(db, bodyBytes)
	datamodels.OtherTable(db, bodyBytes)
	datamodels.PartnerChoiceTable(db, bodyBytes)
	datamodels.EmailDataTable(db, bodyBytes)
	datamodels.VerificationDataTable(db, bodyBytes)
	datamodels.Notifications(db)

}
