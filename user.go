package digo

import (
	"encoding/json"
	"fmt"
	"github.com/go-resty/resty"
	"log"
	"time"
)

// UserData from the API
type UserData struct {
	AudioToken string `json:"audio_token"`
	CurrentLoginTime time.Time `json:"current_login_at"`
	NetworkID uint `json:"network_id"`
	MemberID uint `json:"member_id"`
	Key string `json:"key"`
	ID uint `json:"id"'`
	SSOToken string `json:"sso_token"`
	SSOTokenExpiration time.Time `json:"sso_token_expires_at"`
	CurrentClientDevice DeviceData `json:"current_client_device"`
	Member MemberData `json:"member"`
}

type DeviceData struct {
	AppVersion string `json:"app_version"`
	ClientDeviceCategoryID uint `json:"client_device_category_id"`
	Description string `json:"description"`
	ID uint `json:"id"`
	Key string `json:"key"`
	Name string `json:"name"`
	OperatingSystem string `json:"operating_system"`
	OSVersion string `json:"os_version"`
}

type MemberData struct {
	Active bool `json:"active"`
	APIKey string `json:"api_key"`
	APIKeyCreated time.Time `json:"api_key_created_at"`
	Banned bool `json:"banned"`
	BillingCountryID uint `json:"billing_country_id"`
	Bio string `json:"bio"`
	ConfirmationToken string `json:"confirmation_token"`
	Email string `json:"email"`
	FailedLoginCount uint `json:"failed_login_count"`
	FirstName string `json:"first_name"`
	ID uint `json:"id"`
	LastListenedAt time.Time `json:"last_listened_at"`
	LastLoginAttempt time.Time `json:"last_login_attempt"`
	ListenKey string `json:"listen_key"`
	Locale string `json:"locale"`
	NetworkID uint `json:"network_id"`
	Timezone string `json:"timezone"`
	UpdatedAt time.Time `json:"updated_at"`
	NetworkFavoriteChannels []ChannelLink `json:"network_favorite_channels"`
	Activated bool `json:"activated"`
	UserType string `json:"user_type"`
	HasPurchases bool `json:"has_purchases"`
	Subscriptions []SubscriptionData `json:"subscriptions"`
}

type SubscriptionData struct {
	AutoRenew bool `json:"auto_renew"`
	CreatedAt time.Time `json:"created_at"`
	Starts string `json:"starts_on"`
	NetworkID uint `json:"network_id"`
	MemberID uint `json:"member_id"`
	ID uint `json:"id"`
	Expires string `json:"expires_on"`
	Services []Service `json:"services"`
}

type Service struct {
	ID uint `json:"id"`
	Key string `json:"key"`
	Name string `json:"name"`
}

// Gets user data from the API and returns a parsed UserData
func GetUserData(username string, password string) *UserData {
	userData := new(UserData)
	formData := make(map[string]string)
	formData["member_session[username]"] = username
	formData["member_session[password]"] = password
	resp, err := resty.R().
		SetFormData(formData).
		SetBasicAuth("streams", "diradio").
		Post(
			fmt.Sprintf("%s/member_sessions", APIRoot),
		)
	if err != nil {
		log.Fatal("Could not complete API call to get user data")
	}
	body := resp.Body()
	fmt.Println(string(body))
	err = json.Unmarshal(body, &userData)
	if err != nil {
		fmt.Println(err.Error())
		log.Fatal("Could not parse JSON into UserData")
	}
	return userData
}
