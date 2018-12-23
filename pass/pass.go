// Package database contains implementation of database service.
// Any database service should be implemented here.
package pass

import (
	"time"
)

type Pass struct {
	Id 							uint   		 	`gorm:"primary_key" json:"id"`
	Type 							uint   		 	`json:"type"`
	Subtype 							uint   		 	`json:"subtype"`
	Logo 						string 		  `json:"logo"`
	LogoText 				string 		  `json:"logo_text"`
	HeaderFields 		string 			`json:"header_fields"`
	PrimaryField 		string			`gorm:"type:json" json:"primary_fields"`
	SecondaryField 	string			`gorm:"type:json" json:"secondary_fields"`
	AuxiliaryField 	string			`gorm:"type:json" json:"auxiliary_fields"`
	Background 			string			`json:"background"`
	Thumbnail 			string			`json:"thumbnail"`
	BarcodeType 		string			`json:"barcode_type"`
	BarcodeMessage 	string			`json:"barcode_message"`
	CreatedAt 			time.Time   `json:"created_at"`
	UpdatedAt 			time.Time   `json:"updated_at"`
}
