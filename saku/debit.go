// Package database contains implementation of database service.
// Any database service should be implemented here.
package saku

import (
  	"time"
  	"github.com/condrowiyono/sakudompet/pass"
)

type Debit struct {
	Id          	uint   `gorm:"primary_key" json:"id"`
  	Name     		string `json:"name"`
  	IssuedBy 		string `json:"issued_by"`
  	Number     		string `json:"number"`
  	Balance     	string `json:"balance"`
  	PassId			uint 	`json:"pass_id"`
  	Pass 			*pass.Pass  	`gorm:"-" json:"pass"`
  	CreatedAt     	time.Time   `json:"created_at"`
  	UpdatedAt     	time.Time   `json:"updated_at"`
}

type DebitPrimaryField struct {
	Key 		string 		`json="key"`
	Value 		string 		`json="value"`
}

type DebitSecondaryField struct {
	Key 		string 		`json="key"`
	Value 		string 		`json="value"`
}

type DebitAuxiliaryField struct {
	Key 		string 		`json="key"`
	Value 		string 		`json="value"`
}