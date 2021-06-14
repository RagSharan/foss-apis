package entity

import "image"

type Country struct {
	CountryCode string      `json:"country-code" bson:"country-code"`
	CountryName string      `json:"country-name" bson:"country-name"`
	OffLang     string      `json:"off-lang" bson:"off-lang"`
	Flag        image.Image `json:"flag" bson:"flag"`
}
