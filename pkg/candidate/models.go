package candidate

import (
	"candidate_service/pkg/commons"
	"time"
)

// AddressType represents address type of a candidate
type AddressType string

// Address represents address of a candidate
type Address struct {
	ID               string      `json:"id" gorm:"column:id,primary_key" validate:"required"`
	FormattedAddress string      `json:"formatted_address" gorm:"column:formatted_address" validate:"required"`
	Lat              string      `json:"lat" gorm:"column:lat" validate:"required"`
	Lng              string      `json:"lng" gorm:"column:lng" validate:"required"`
	City             string      `json:"city" gorm:"column:city" validate:"required"`
	Country          string      `json:"country" gorm:"column:country" validate:"required"`
	PinCode          string      `json:"pin_code" gorm:"column:pin_code" validate:"required"`
	Type             AddressType `json:"type" gorm:"column:type" validate:"required"`
	CandidateID      string
	CreatedAt        time.Time
	UpdatedAt        time.Time
	DeletedAt        *time.Time
}

// Candidate holds candidate details along with who columns
type Candidate struct {
	ID         string     `json:"id" gorm:"column:id,primary_key" validate:"required"`
	Name       string     `json:"name" gorm:"column:name" validate:"required"`
	Age        int        `json:"age" gorm:"column:age" validate:"gte=0,lte=130"`
	RollNumber string     `json:"roll_number" gorm:"column:roll_number" validate:"required"`
	Email      string     `json:"email" gorm:"column:email" validate:"required"`
	Mobile     string     `json:"mobile" gorm:"column:mobile" validate:"required"`
	Address    []*Address `json:"address" gorm:"ForeignKey:CandidateID,PRELOAD:true" validate:"required"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  *time.Time
}

// TableName to ensure table names are predictable
func (Candidate) TableName() string {
	return "candidates"
}

// TableName to ensure table names are predictable
func (Address) TableName() string {
	return "addresses"
}

// Response is standard response from candidate api
type Response struct {
	Data  Candidate        `json: "candidate"`
	Error commons.APIError `json: "error"`
}
