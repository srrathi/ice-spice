package models

import (
	"log"

	"gorm.io/gorm"
)

type StationData struct {
	ID             uint     `gorm:"primary key;autoIncrement" json:"id"`
	Beginning      *string  `json:"beginning"`
	End            *string  `json:"end"`
	SeasonLength   *int     `json:"season_length"`
	IceDays        *int     `json:"ice_days"`
	FlTotal        *float32 `json:"fl_total"`
	VolSum         *float32 `json:"vol_sum"`
	MaxThickness   *float32 `json:"max_thickness,omitempty"`
	MaxThickness2  *float32 `json:"max_thickness2,omitempty"`
	Country        *string  `json:"country"`
	CountryCode    *string  `json:"country_code"`
	StationName    *string  `json:"station_name"`
	StationCode    *string  `json:"station_code"`
	RestrictedDays *int     `json:"restricted_days"`
}

type AnalysisData struct {
	AdId            uint     `gorm:"primary key;autoIncrement" json:"ad_id"`
	AnalysisBasis *string  `json:"analysis-basis"`
	Beginning     *string  `json:"beginning"`
	End           *string  `json:"end"`
	IceDays       *float32 `json:"ice_days"`
	FlTotal       *float32 `json:"fl_total"`
	VolSum        *float32 `json:"vol_sum"`
	Country       *string  `json:"country"`
	CountryCode   *string  `json:"country_code"`
	StationName   *string  `json:"station_name"`
	StationCode   *string  `json:"station_code"`
}

type RestrictedDays struct {
	RdId  uint `gorm:"primary key;autoIncrement" json:"-"`
	RD1 *int `json:"rd1"`
	RD2 *int `json:"rd2"`
	RD3 *int `json:"rd3"`
	RD4 *int `json:"rd4"`
	RD5 *int `json:"rd5"`
	RD6 *int `json:"rd6"`
	RD7 *int `json:"rd7"`
	RD8 *int `json:"rd8"`
	RD9 *int `json:"rd9"`
}

func MigrateDatabases(db *gorm.DB) error {
	err := db.AutoMigrate(&StationData{})
	if err != nil {
		log.Fatal(err)
		return err
	}

	err = db.AutoMigrate(&AnalysisData{})
	if err != nil {
		log.Fatal(err)
		return err
	}

	err = db.AutoMigrate(&RestrictedDays{})
	if err != nil {
		log.Fatal(err)
		return err
	}
	return err
}
