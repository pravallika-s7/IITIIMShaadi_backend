package datamodels

import (
	"time"

	"gorm.io/gorm"
)

//const layout = "1989-10-20T00:00:00+0000"

type (
	//User registration table
	User struct {
		gorm.Model
		Username string `json:"username"`
		Fullname string `json:"name"`
		Email    string `json:"email"`
		Password string `json:"password"`
		Gender   string `json:"gender"`
		Mobile   string `json:"mobile_no"`
		DOB      string `json:"birth_date"`
		UserId   int    `gorm:"primaryKey",json:"id"`

		Friends []Friend `gorm:"foreignkey:UserID;references:UserId"`

		SentChats     []SentChat     `gorm:"foreignkey:UserID;references:UserId"`
		ReceivedChats []ReceivedChat `gorm:"foreignkey:UserID;references:UserId"`
		Subscriptions []Subscription `gorm:"foreignkey:UserID;references:UserId"`

		Albums         []Album         `gorm:"foreignkey:UserID;references:UserId"`
		FolderListings []FolderListing `gorm:"foreignkey:UserID;references:UserId"`
		Contacts       []Contact       `gorm:"foreignkey:UserID;references:UserId"`
	}
	//Profile of user
	Profile struct {
		gorm.Model
		ProfileFor string     `json:"profile_created_for"`
		Image      []byte     `json:"profile_image"`
		Role       string     `json:"role"`
		Religion   string     `json:"religion"`
		Caste      string     `json:"caste"`
		MotherTng  string     `json:"mother_tounge"`
		AltMobile  string     `json:"alternate_no"`
		MaritalSts string     `json:"marital_status"`
		Interests  []Interest `json:"interest"`
		Diet       string     `json:"diet"`
		Height     string     `json:"height"`
		Drink      string     `json:"drink"`
		Smoke      string     `json:"smoke"`
		HealthIss  string     `json:"health_issue"`
		AboutMe    string     `json:"about_me"`
		UserID     int
	}
	Interest struct {
		gorm.Model
		Intr      string
		ProfileID uint
	}
	//Family details of user
	Family struct {
		FatherName string `json:"father_name"`
		FatherOccp string `json:"father_occupation"`
		MotherName string `json:"mother_name"`
		MotherOccp string `json:"mother_occupation"`
		Brother    string `json:"brother"`
		Sister     string `json:"sister"`
		UserID     int
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
		UserID      int
	}
	//Job details of user
	Job struct {
		Company  string `json:"name_of_company"`
		JobTitle string `json:"working_as"`
		JobLoc   string `json:"job_location"`
		LinkedIn string `json:"linked_in"`
		IncomeAn string `json:"annual_income"`
		UserID   int
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
		UserID    int
	}
	//Other details regarding user
	Other struct {
		MPrivacy int    `json:"mprivacy"`
		BrideChe string `json:"choice_of_bride"`
		GroomChe string `json:"choice_of_groom"`
		PrflCmpl int    `json:"profile_complition"`
		UserID   int
	}
	//PartnerChoice preferred by user
	PartnerChoice struct {
		gorm.Model
		PartnerChe  string       `json:"choice_of_partner"`
		PartnerEdu  string       `json:"prefered_partner_education"`
		PartnerRlg  string       `json:"prefered_partner_religion"`
		PartnerCst  string       `json:"prefered_partner_caste"`
		PartnerCtr  string       `json:"prefered_partner_country"`
		PartMinAge  int          `json:"prefered_partner_min_age"`
		PartMaxAge  int          `json:"prefered_partner_max_age"`
		PartHtMax   string       `json:"prefered_partner_height_max"`
		PartHtMin   string       `json:"prefered_partner_height_min"`
		PartMartSts []PartMartSt `json:"prefered_partner_marital_status"`
		UserID      int
	}
	PartMartSt struct {
		gorm.Model
		MarSt           string
		PartnerChoiceID uint
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
		UserID       int
	}
	//VerificationData for user
	VerificationData struct {
		EmailStatus   string `json:"emailStatus"`
		MobileStatus  int8   `json:"mobileStatus"`
		BiodataStatus int8   `json:"biodata_status"`
		IdProofStatus int8   `json:"identity_proof_verified"`
		//DocVerified  int8 	`json:"document_verified"`
		UserID int
	}
	//Friend of user
	Friend struct {
		gorm.Model
		FriendId     int `json:"id"`
		ShortListed  bool
		RequestFrnd  bool
		InviteedFrnd bool
		UserID       int
	}
	//SentChat from user to others
	SentChat struct {
		ChatId        int       `json:"id"`
		ToUserId      int       `json:"to_user_id"`
		Message       string    `json:"message"`
		ChatType      string    `json:"chat_type"`
		SentTime      time.Time `json:"sent"`
		Recd          uint      `json:"recd"`
		LatestMsg     string    `json:"latestMessage"`
		LatestMsgDate time.Time `json:"latestDate"`
		MsgStatus     string    `json:"messageStatus"`
		UserID        int
	}
	//ReceivedChat from others
	ReceivedChat struct {
		ChatId        int       `json:"id"`
		FromUserId    int       `json:"from_user_id"`
		Message       string    `json:"message"`
		ChatType      string    `json:"chat_type"`
		SentTime      time.Time `json:"sent"`
		Recd          uint      `json:"recd"`
		LatestMsg     string    `json:"latestMessage"`
		LatestMsgDate time.Time `json:"latestDate"`
		MsgStatus     string    `json:"messageStatus"`
		UserID        int
	}
	//Subscription of user
	Subscription struct {
		SubId     int       `json:"ID"`
		TransId   string    `json:"transaction_id"`
		PayMode   string    `json:"payment_mode"`
		Amount    string    `json:"amount"`
		PayDate   time.Time `json:"payment_date"`
		ExpDate   time.Time `json:"exp_date"`
		Duration  string    `json:"duration"`
		LastRenew time.Time `json:"last_renewed_on"`
		NextRenew time.Time `json:"next_renew_on"`
		UserID    int
	}
	//Album of user
	Album struct {
		AlbumId       int       `json:"id"`
		OwnerType     string    `json:"owner_type"`
		DateUpload    time.Time `json:"date_uploaded"`
		Title         string    `json:"title"`
		Description   string    `json:"description"`
		Privacy       int       `json:"privacy"`
		PicOrgUrl     string    `json:"pic_org_url"`
		PicOrgPath    string    `json:"pic_org_path"`
		PicMidUrl     string    `json:"pic_mid_url"`
		PicMidPath    string    `json:"pic_mid_path"`
		PicThumbUrl   string    `json:"pic_thumb_url"`
		PicThumbPath  string    `json:"pic_thumb_path"`
		ImportSts     int       `json:"import_status"`
		OldActivityId int       `json:"old_activity_id"`
		NewActivityId int       `json:"new_activity_id"`
		Favoutite     bool      `json:"favorites"`
		UserID        int
	}
	//Contact bu user
	Contact struct {
		gorm.Model
		Viewed      int `json:"viewedContacts"`
		Denied      int `json:"deniedUsersIds"`
		RequestSent int `json:"allRequestSentIds"`
		DeniedByMe  int `json:"deniedByMeUsersIds"`
		UserID      int
	}
	//FolderListing of user
	FolderListing struct {
		ListingId int    `json:"id"`
		MetaKey   string `json:"meta_key"`
		MetaValue string `json:"meta_value"`
		FrndsIds  int    `json:"friends_ids"`
		//DateAdded time.Time `json:"date_added"`
		UserID int
	}
	//SuccessStory of org
	SuccessStory struct {
		gorm.Model
		Content string `json:"content"`
		Title   string `json:"title"`
		//DateAdded time.Time `json:"date_added"`
	}
	//MediaCoverage of org
	MediaCoverage struct {
		gorm.Model
		Title     string `json:"title"`
		Link      string `json:"link"`
		Image     string `json:"image"`
		SortOrder int    `json:"sort_order"`
		//DateAdded time.Time `json:"date_added"`
	}
	//Notification
	Notification struct {
		gorm.Model
		ToId   int    `json:"id"`
		Type   string `json:"type"`
		Msg    string `json:"message"`
		Status int    `json:"status"`
		//DateAdded time.Time `json:"date_added"`
	}
)
