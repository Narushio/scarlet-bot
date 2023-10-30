package model

import "github.com/Narushio/scarlet-bot/helper/image"

type BoostInfo struct {
	Description string          `json:"description,omitempty"`
	Start       image.Base64Src `json:"start,omitempty"`
	End         image.Base64Src `json:"end,omitempty"`
}
