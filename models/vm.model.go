package models

//VM model
type VM struct {
	VMID       uint   `gorm:"primaryKey;autoIncrement;not null"`
	Name       string `json:"name"`
	DistrictID int    `json:"districtId"`
}
