package digo

import (
	"encoding/json"
	"fmt"
	"github.com/go-resty/resty"
	log "github.com/sirupsen/logrus"
)

type API struct {
	Caller resty.Client
}

func GenerateNewAPI(key string) *resty.Client {
	client := new(resty.Client)
	client.SetHeader("X-Session-Key", key)
	return client
}

func (api *API) GetChannels() []Channel {
	channels := make([]Channel, 0)
	channelResult, err := api.Caller.R().Post(
		fmt.Sprintf("%s/channels", APIRoot),
		)
	if err != nil {
		log.Error("Could ")
	}
	body := channelResult.Body()
	err = json.Unmarshal(body, &channels)
}
