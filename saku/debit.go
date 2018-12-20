// Package database contains implementation of database service.
// Any database service should be implemented here.
package saku

import (
  	"time"
)

type Debit struct {
	Id          	uint   `gorm:"primary_key" json:"id"`
  	Name     		string `json:"name"`
  	IssuedBy 		string `json:"issued_by"`
  	Number     		string `json:"number"`
  	CreatedAt     	time.Time   `json:"created_at"`
  	UpdatedAt     	time.Time   `json:"updated_at"`
}