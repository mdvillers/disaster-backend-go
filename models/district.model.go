package models

//District model
type District struct {
	DistrictName   string  `json:"districtName" gorm:"primaryKey;not null"`
	ProvinceNumber int     `json:"provinceNumber"`
	Latitude       float32 `json:"latitude"`
	Longitude      float32 `json:"longitude"`
}
