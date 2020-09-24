package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type (
	//User registration table
	User struct {
		gorm.Model
		Username string `json:"username"`
		Fullname string `json:"name"`
		Email    string `json:"email"`
		//Password string    `json:"password"`
		Gender string `json:"gender"`
		Mobile string `json:"mobile_no"`
		//DOB    time.Time.Date() `json:"birth_date"`
	}
	//Profile of user
	Profile struct {
		ProfileFor string `json:"profile_created_for"`
		//Image
		Role       string `json:"role"`
		Religion   string `json:"religion"`
		Caste      string `json:"caste"`
		MotherTng  string `json:"mother_tounge"`
		AltMobile  string `json:"alternate_no"`
		MaritalSts string `json:"marital_status"`
		Interest   string `json:"interest"`
		Diet       string `json:"diet"`
		Height     string `json:"height"`
		Drink      string `json:"drink"`
		Smoke      string `json:"smoke"`
		HealthIss  string `json:"health_issue"`
		AboutMe    string `json:"about_me"`
	}
	//Family details of user
	Family struct {
		FatherName string `json:"father_name"`
		FatherOccp string `json:"father_occupation"`
		MotherName string `json:"mother_name"`
		MotherOccp string `json:"mother_occupation"`
		Brother    string `json:"brother"`
		Sister     string `json:"sister"`
	}
	//Education details of user
	Education struct {
		HighestQual string `json:"highest_education"`
		PostGradClg string `json:"post_graduation_college"`
		PostGradYr  int    `json:"post_graduation_year"`
		PostGrad    string `json:"post_graduation"`
		GradClg     string `json:"graduation_college"`
		GradYr      int    `json:"graduation_year"`
		Grad        string `json:"graduation"`
		SchoolYr    int    `json:"schooling_year"`
		School      string `json:"schooling"`
	}
	//Job details of user
	Job struct {
		Company  string `json:"name_of_company"`
		JobTitle string `json:"working_as"`
		JobLoc   string `json:"job_location"`
		LinkedIn string `json:"linked_in"`
		IncomeAn string `json:"annual_income"`
	}
	//Address details of user
	Address struct {
		PermAddr  string `json:"pemanent_address"`
		PermCity  string `json:"permanent_city"`
		PermState string `json:"permanent_state"`
		PermCntry string `json:"permanent_country"`
		PermZipCd string `json:"permanent_zipcode"`
		CurrAddr  string `json:"current_address"`
		CurrCity  string `json:"current_city"`
		CurrState string `json:"current_state"`
		CurrCntry string `json:"current_country"`
		CurrZipCd string `json:"current_zipcode"`
	}
	//Other details regarding user
	Other struct {
		MPrivacy int    `json:"mprivacy"`
		BrideChe string `json:"choice_of_bride"`
		GroomChe string `json:"choice_of_groom"`
		PrflCmpl int    `json:"profile_complition"`
	}
	//PartnerChoice preferred by user
	PartnerChoice struct {
		PartnerChe string `json:"choice_of_partner"`
		PartnerEdu string `json:"prefered_partner_education"`
		PartnerRlg string `json:"prefered_partner_religion"`
		PartnerCst string `json:"prefered_partner_caste"`
		PartnerCtr string `json:"prefered_partner_country"`
		PartMinAge int    `json:"prefered_partner_min_age"`
		PartMaxAge int    `json:"prefered_partner_max_age"`
		PartHtMax  string `json:"prefered_partner_height_max"`
		PartHtMin  string `json:"prefered_partner_height_min"`
		PartMartSt bool   `json:"prefered_partner_marital_status"`
	}
	//EmailData of user
	EmailData struct {
		EmailIdNo   int       `json:"id"`
		UserNick    string    `json:"user_nickname"`
		EmailStTm   time.Time `json:"email_sent_time"`
		DocName     string    `json:"document_name"`
		DocVerified int8      `json:"document_verified"`
		UserUrl     string    `json:"user_url"`
		//CreatedDate
		FbProfileId  string `json:"facebook_profileid"`
		FbStatus     int8   `json:"facebook_status"`
		GgProfileId  string `json:"google_profileid"`
		GgStatus     int8   `json:"google_status"`
		UserActKey   string `json:"user_activation_key"`
		AccStatus    int8   `json:"account_status"`
		SendMsg      int8   `json:"send_message"`
		SendReq      int8   `json:"send_request"`
		Shortlistd   int    `json:"shortlisted"`
		Favourite    int8   `json:"favourate"`
		DispName     string `json:"display_name"`
		PrflEmlSent  int8   `json:"profile_email_sent"`
		EmailSentSt  int8   `json:"emailsentstatus"`
		EmailSentSt1 int8   `json:"emailsentstatus1"`
		DeactEmail   int8   `json:"deactivatemail"`
		Pemail       int8   `json:"pemail"`
	}
	//VerificationData for user
	VerificationData struct {
		EmailStatus   string `json:"emailStatus"`
		MobileStatus  int8   `json:"mobileStatus"`
		BiodataStatus int8   `json:"biodata_status"`
		IdProofStatus int8   `json:"identity_proof_verified"`
		//DocVerified  int8 	`json:"document_verified"`
	}
)

func main() {
	db, err := gorm.Open(sqlite.Open("UserProfile.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	db.AutoMigrate(&User{}, &Profile{}, &Family{},
		&Education{}, &Job{}, &Address{}, &Other{},
		&PartnerChoice{}, &EmailData{}, &VerificationData{})

	jsonBody := map[string]string{"token": "12931e786a9da517f10c52880c5711eb", "userId": "9798"}
	jsonValue, err := json.Marshal(jsonBody)

	response, err := http.Post("https://www.iitiimshaadi.com/apis/my_profile.json", "application/json", bytes.NewBuffer(jsonValue))
	if err != nil {
		log.Fatalln(err)
	}

	defer response.Body.Close()
	bodyBytes, _ := ioutil.ReadAll(response.Body)

	//Populating data in each table
	userTable(db, bodyBytes)
	profileTable(db, bodyBytes)
	familyTable(db, bodyBytes)
	educationTable(db, bodyBytes)
	jobTable(db, bodyBytes)
	addressTable(db, bodyBytes)
	otherTable(db, bodyBytes)
	partnerChoiceTable(db, bodyBytes)
	emailDataTable(db, bodyBytes)
	verificationDataTable(db, bodyBytes)

}

func userTable(db *gorm.DB, body []byte) {

	var userData map[string]interface{}
	json.Unmarshal(body, &userData)
	user := userData["basicData"].(map[string]interface{})

	db.Create(&User{Username: user["username"].(string),
		Fullname: user["name"].(string),
		Email:    userData["emailData"].(map[string]interface{})["email"].(string),
		Gender:   user["gender"].(string),
		Mobile:   user["mobile_no"].(string),
		//DOB:      user["birth_data"].(string),
	})
}

func profileTable(db *gorm.DB, body []byte) {

	var profileData map[string]interface{}
	json.Unmarshal(body, &profileData)
	profile := profileData["basicData"].(map[string]interface{})

	db.Create(&Profile{ProfileFor: profile["profile_created_for"].(string),
		Religion:   profile["religion"].(string),
		Caste:      profile["caste"].(string),
		MotherTng:  profile["mother_tounge"].(string),
		AltMobile:  profile["alternate_no"].(string),
		MaritalSts: profile["marital_status"].(string),
		Interest:   profile["interest"].(string),
		Diet:       profile["diet"].(string),
		Height:     profile["height"].(string),
		Drink:      profile["drink"].(string),
		Smoke:      profile["smoke"].(string),
		HealthIss:  profile["health_issue"].(string),
		AboutMe:    profile["about_me"].(string),
	})
}

func familyTable(db *gorm.DB, body []byte) {

	var familyData map[string]interface{}
	json.Unmarshal(body, &familyData)
	family := familyData["basicData"].(map[string]interface{})

	db.Create(&Family{FatherName: family["father_name"].(string),
		FatherOccp: family["father_occupation"].(string),
		MotherName: family["mother_name"].(string),
		MotherOccp: family["mother_occupation"].(string),
		Brother:    family["brother"].(string),
		Sister:     family["sister"].(string),
	})
}

func educationTable(db *gorm.DB, body []byte) {

	var educationData map[string]interface{}
	json.Unmarshal(body, &educationData)
	education := educationData["basicData"].(map[string]interface{})

	db.Create(&Education{HighestQual: education["highest_education"].(string),
		PostGradClg: education["post_graduation_college"].(string),
		PostGradYr:  int(education["post_graduation_year"].(float64)),
		PostGrad:    education["post_graduation"].(string),
		GradClg:     education["graduation_college"].(string),
		GradYr:      int(education["graduation_year"].(float64)),
		Grad:        education["graduation"].(string),
		SchoolYr:    int(education["schooling_year"].(float64)),
		School:      education["schooling"].(string),
	})
}

func jobTable(db *gorm.DB, body []byte) {

	var jobData map[string]interface{}
	json.Unmarshal(body, &jobData)
	job := jobData["basicData"].(map[string]interface{})

	db.Create(&Job{Company: job["name_of_company"].(string),
		JobTitle: job["working_as"].(string),
		JobLoc:   job["job_location"].(string),
		LinkedIn: job["linked_in"].(string),
		IncomeAn: job["annual_income"].(string),
	})
}

func addressTable(db *gorm.DB, body []byte) {

	var addressData map[string]interface{}
	json.Unmarshal(body, &addressData)
	address := addressData["basicData"].(map[string]interface{})

	db.Create(&Address{PermAddr: address["permanent_address"].(string),
		PermCity:  address["permanent_city"].(string),
		PermState: address["permanent_state"].(string),
		PermCntry: address["permanent_country"].(string),
		PermZipCd: address["permanent_zipcode"].(string),
		CurrAddr:  address["current_address"].(string),
		CurrCity:  address["current_city"].(string),
		CurrState: address["current_state"].(string),
		CurrCntry: address["current_country"].(string),
		CurrZipCd: address["current_zipcode"].(string),
	})
}

func otherTable(db *gorm.DB, body []byte) {

	var otherData map[string]interface{}
	json.Unmarshal(body, &otherData)
	other := otherData["basicData"].(map[string]interface{})

	db.Create(&Other{MPrivacy: int(other["mprivacy"].(float64)),
		BrideChe: other["choice_of_bride"].(string),
		GroomChe: other["choice_of_groom"].(string),
		PrflCmpl: int(other["profile_complition"].(float64)),
	})
}

func partnerChoiceTable(db *gorm.DB, body []byte) {

	var partnerChoiceData map[string]interface{}
	json.Unmarshal(body, &partnerChoiceData)
	partnerChoice := partnerChoiceData["partnerBasicData"].(map[string]interface{})

	db.Create(&PartnerChoice{PartnerChe: partnerChoice["choice_of_partner"].(string),
		PartnerEdu: partnerChoice["prefered_partner_education"].(string),
		PartnerRlg: partnerChoice["prefered_partner_religion"].(string),
		PartnerCst: partnerChoice["prefered_partner_caste"].(string),
		PartnerCtr: partnerChoice["prefered_partner_country"].(string),
		PartMinAge: int(partnerChoice["prefered_partner_min_age"].(float64)),
		PartMaxAge: int(partnerChoice["prefered_partner_max_age"].(float64)),
		PartHtMax:  partnerChoice["prefered_partner_height_max"].(string),
		PartHtMin:  partnerChoice["prefered_partner_height_min"].(string),
		PartMartSt: partnerChoice["prefered_partner_marital_status"].(bool),
	})
}

func emailDataTable(db *gorm.DB, body []byte) {

	var emailData map[string]interface{}
	json.Unmarshal(body, &emailData)
	email := emailData["emailData"].(map[string]interface{})

	db.Create(&EmailData{EmailIdNo: int(email["id"].(float64)),
		UserNick: email["user_nickname"].(string),
		//EmailStTm:   email["email_sent_time"].(time.Time),
		DocName:     email["document_name"].(string),
		DocVerified: int8(email["document_verified"].(float64)),
		UserUrl:     email["user_url"].(string),
		//CreatedDate
		FbProfileId:  email["facebook_profileid"].(string),
		FbStatus:     int8(email["facebook_status"].(float64)),
		GgProfileId:  email["google_profileid"].(string),
		GgStatus:     int8(email["google_status"].(float64)),
		UserActKey:   email["user_activation_key"].(string),
		AccStatus:    int8(email["account_status"].(float64)),
		SendMsg:      int8(email["send_message"].(float64)),
		SendReq:      int8(email["send_request"].(float64)),
		Shortlistd:   int(email["shortlisted"].(float64)),
		Favourite:    int8(email["favourate"].(float64)),
		DispName:     email["display_name"].(string),
		PrflEmlSent:  int8(email["profile_email_sent"].(float64)),
		EmailSentSt:  int8(email["emailsentstatus"].(float64)),
		EmailSentSt1: int8(email["emailsentstatus1"].(float64)),
		DeactEmail:   int8(email["deactivatemail"].(float64)),
		Pemail:       int8(email["pemail"].(float64)),
	})
}

func verificationDataTable(db *gorm.DB, body []byte) {

	var verificationData map[string]interface{}
	json.Unmarshal(body, &verificationData)
	verify := verificationData["verificationData"].(map[string]interface{})

	db.Create(&VerificationData{EmailStatus: verify["emailStatus"].(string),
		MobileStatus:  int8(verify["mobileStatus"].(float64)),
		BiodataStatus: int8(verify["biodata_status"].(float64)),
		IdProofStatus: int8(verify["identity_proof_verified"].(float64)),
		//DocVerified: int8(verify["document_verified"].(float64)),
	})
}