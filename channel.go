package digo

import "time"

type ChannelLink struct {
	ID uint `json:"id"`
	SimilarChannelID uint `json:"similar_channel_id"`
	Position uint `json:"position"`
}

type Channel struct {
	ChannelID uint `json:"channel_id"`
	AdChannels string `json:"ad_channels"`
	ChannelDirector string `json:"channel_director"`
	CreatedAt time.Time `json:"created_at"`
	Description string `json:"description"`
	ID uint `json:"id"`
	Key string `json:"key"`
	Name string `json:"name"`
	OldID string `json:"old_id"`
	Public bool `json:"public"`
	UpdatedAt time.Time `json:"updated_at"`
	AssetUrl string `json:"asset_url"`
	BannerUrl string `json:"banner_url"`
	SimilarChannels []ChannelLink `json:"similar_channels"`
	images struct {
		Default string `json:"default"`
		HorizontalBanner string `json:"horizontal_banner"`
		Compact string `json:"compact"`
	} `json:"images"`
}



