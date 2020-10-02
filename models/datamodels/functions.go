package datamodels

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"image"
	"log"
	"strings"

	"gorm.io/gorm"
)

const layout = "1989-10-20T00:00:00+0000"

func (prof *Profile) AddItem(intr Interest) []Interest {
	prof.Interests = append(prof.Interests, intr)
	return prof.Interests
}

func (pc *PartnerChoice) AddItem(marSt PartMartSt) []PartMartSt {
	pc.PartMartSts = append(pc.PartMartSts, marSt)
	return pc.PartMartSts
}

func UserTable(db *gorm.DB, body []byte) {

	var userData map[string]interface{}
	json.Unmarshal(body, &userData)
	user := userData["basicData"].(map[string]interface{})
	//fmt.Print(user)
	/*
		str, ok := user["birth_date"].(string)
		if !ok {
			fmt.Printf("ERROR: not a string -> %#v\n", user["birth_date"])
		}
		fmt.Printf(str)
	*/
	db.Create(&User{Username: user["username"].(string),
		Fullname: user["name"].(string),
		Email:    userData["emailData"].(map[string]interface{})["email"].(string),
		//Password: user["password"].(string), //HashPassword can also be used
		Gender: user["gender"].(string),
		Mobile: user["mobile_no"].(string),
		DOB:    user["birth_date"].(string),
		UserId: int(userData["emailData"].(map[string]interface{})["id"].(float64)),
	})
}

func ProfileTable(db *gorm.DB, body []byte) {

	var profileData map[string]interface{}
	json.Unmarshal(body, &profileData)
	profile := profileData["basicData"].(map[string]interface{})
	//fmt.Print(profile)
	interests := []Interest{}
	prof := Profile{Interests: interests}

	for i := 0; i < len(profile["interest"].([]interface{})); i++ {
		a := profile["interest"].([]interface{})[i]
		item := Interest{Intr: a.(string)}
		prof.AddItem(item)
	}

	Image, ok := profile["profile_image"].(string)
	if !ok {
		fmt.Printf("ERROR: not a string -> %#v\n", profile["profile_image"])
	}
	//fmt.Printf(image)
	/*i := strings.Index(image, ",")
	if i < 0 {
		log.Fatal("no comma")
	}
	dec := base64.NewDecoder(base64.StdEncoding, strings.NewReader(image[i+1:]))
	fmt.Print(dec)
	*/

	reader := base64.NewDecoder(base64.StdEncoding, strings.NewReader(Image))
	config, format, err := image.DecodeConfig(reader)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Width:", config.Width, "Height:", config.Height, "Format:", format)

	db.Create(&Profile{ProfileFor: profile["profile_created_for"].(string),
		//Image:
		Religion:   profile["religion"].(string),
		Caste:      profile["caste"].(string),
		MotherTng:  profile["mother_tounge"].(string),
		AltMobile:  profile["alternate_no"].(string),
		MaritalSts: profile["marital_status"].(string),
		Interests:  prof.Interests,
		Diet:       profile["diet"].(string),
		Height:     profile["height"].(string),
		Drink:      profile["drink"].(string),
		Smoke:      profile["smoke"].(string),
		HealthIss:  profile["health_issue"].(string),
		AboutMe:    profile["about_me"].(string),
		UserID:     int(profileData["emailData"].(map[string]interface{})["id"].(float64)),
	})
}

func FamilyTable(db *gorm.DB, body []byte) {

	var familyData map[string]interface{}
	json.Unmarshal(body, &familyData)
	family := familyData["basicData"].(map[string]interface{})

	db.Create(&Family{FatherName: family["father_name"].(string),
		FatherOccp: family["father_occupation"].(string),
		MotherName: family["mother_name"].(string),
		MotherOccp: family["mother_occupation"].(string),
		Brother:    family["brother"].(string),
		Sister:     family["sister"].(string),
		UserID:     int(familyData["emailData"].(map[string]interface{})["id"].(float64)),
	})
}

func EducationTable(db *gorm.DB, body []byte) {

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
		UserID:      int(educationData["emailData"].(map[string]interface{})["id"].(float64)),
	})
}

func JobTable(db *gorm.DB, body []byte) {

	var jobData map[string]interface{}
	json.Unmarshal(body, &jobData)
	job := jobData["basicData"].(map[string]interface{})

	db.Create(&Job{Company: job["name_of_company"].(string),
		JobTitle: job["working_as"].(string),
		JobLoc:   job["job_location"].(string),
		LinkedIn: job["linked_in"].(string),
		IncomeAn: job["annual_income"].(string),
		UserID:   int(jobData["emailData"].(map[string]interface{})["id"].(float64)),
	})
}

func AddressTable(db *gorm.DB, body []byte) {

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
		UserID:    int(addressData["emailData"].(map[string]interface{})["id"].(float64)),
	})
}

func OtherTable(db *gorm.DB, body []byte) {

	var otherData map[string]interface{}
	json.Unmarshal(body, &otherData)
	other := otherData["basicData"].(map[string]interface{})

	var bride string
	if other["choice_of_bride"] == nil {
		bride = "null"
	} else {
		bride = other["choice_of_bride"].(string)
	}

	var groom string
	if other["choice_of_groom"] == nil {
		groom = "null"
	} else {
		groom = other["choice_of_groom"].(string)
	}

	db.Create(&Other{MPrivacy: int(other["mprivacy"].(float64)),
		BrideChe: bride,
		GroomChe: groom,
		PrflCmpl: int(other["profile_complition"].(float64)),
		UserID:   int(otherData["emailData"].(map[string]interface{})["id"].(float64)),
	})
}

func PartnerChoiceTable(db *gorm.DB, body []byte) {

	var partnerChoiceData map[string]interface{}
	json.Unmarshal(body, &partnerChoiceData)
	partnerChoice := partnerChoiceData["partnerBasicData"].(map[string]interface{})

	Ms := []PartMartSt{}
	pc := PartnerChoice{PartMartSts: Ms}

	for i := 0; i < len(partnerChoice["prefered_partner_marital_status"].([]interface{})); i++ {
		b := partnerChoice["prefered_partner_marital_status"].([]interface{})[i]
		married := PartMartSt{MarSt: b.(string)}
		pc.AddItem(married)
	}

	var choice string
	if partnerChoice["choice_of_partner"] == nil {
		choice = "null"
	} else {
		choice = partnerChoice["choice_of_partner"].(string)
	}

	db.Create(&PartnerChoice{PartnerChe: choice,
		PartnerEdu:  partnerChoice["prefered_partner_education"].(string),
		PartnerRlg:  partnerChoice["prefered_partner_religion"].(string),
		PartnerCst:  partnerChoice["prefered_partner_caste"].(string),
		PartnerCtr:  partnerChoice["prefered_partner_country"].(string),
		PartMinAge:  int(partnerChoice["prefered_partner_min_age"].(float64)),
		PartMaxAge:  int(partnerChoice["prefered_partner_max_age"].(float64)),
		PartHtMax:   partnerChoice["prefered_partner_height_max"].(string),
		PartHtMin:   partnerChoice["prefered_partner_height_min"].(string),
		PartMartSts: pc.PartMartSts,
		UserID:      int(partnerChoiceData["emailData"].(map[string]interface{})["id"].(float64)),
	})
}

func EmailDataTable(db *gorm.DB, body []byte) {

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
		UserID:       int(email["id"].(float64)),
	})
}

func VerificationDataTable(db *gorm.DB, body []byte) {

	var verificationData map[string]interface{}
	json.Unmarshal(body, &verificationData)
	verify := verificationData["verificationData"].(map[string]interface{})

	db.Create(&VerificationData{EmailStatus: verify["emailStatus"].(string),
		MobileStatus:  int8(verify["mobileStatus"].(float64)),
		BiodataStatus: int8(verify["biodata_status"].(float64)),
		IdProofStatus: int8(verify["identity_proof_verified"].(float64)),
		//DocVerified: int8(verify["document_verified"].(float64)),
		UserID: int(verificationData["emailData"].(map[string]interface{})["id"].(float64)),
	})
}
