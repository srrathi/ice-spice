package models

// 			   "Beginning": "03.Jan.2003",
//             "End": "18.Mar.2003",
//             "Season-length": "75",
//             "Ice-days": "55",
//             "Fl-total": "35.7",
//             "Vol-sum": "313.9",
//             "Max-thickness": "NaN",
//             "Max-thickness2": "22",
//             "Country": "FINLAND",
//             "Country-code": "FI",
//             "Station-name": "Station: See außerhalb Nyhamn u. Marhällan",
//             "Restricted-days": [
//                 "53",
//                 "41",
//                 "16",
//                 "6",
//                 "0",
//                 "0",
//                 "0",
//                 "0",
//                 "0"
//             ],
//             "Station-code": "STSEAUNYUMA"

type StationData struct {
	ID             uint    `gorm:"primary key;autoIncrement" json:"id"`
	Beginning      string  `json:"beginning"`
	End            string  `json:"end"`
	SeasonLength   int     `json:"season-length"`
	IceDays        int     `json:"ice-days"`
	FlTotal        float32 `json:"fl-total"`
	VolSum         float32 `json:"vol-sum"`
	MaxThickness   float32 `json:"max-thickness,omitempty"`
	MaxThickness2  float32 `json:"max-thickness2,omitempty"`
	Country        string  `json:"country"`
	CountryCode    string  `json:"country-code"`
	StationName    string  `json:"station-name"`
	StationCode    string  `json:"station-code"`
	RestrictedDays uint    `json:"restricted-days"`
}


// "Analysis-basis": "1.quartile",
// "Beginning-ice": "29.Nov",
// "End-ice": "13.May",
// "Ice-days": "138.0",
// "Fl-total": "134.5",
// "Vol-sum": "5780.0",
// "Station-name": "Station: Röyttä – Etukari",
// "Station-code": "STRÖET",
// "Country": "FINLAND",
// "Country-code": "FI"

type AnalysisData struct{
	ID uint `gorm:"primary key;autoIncrement" json:"ad_id"`
	AnalysisBasis string `json:"analysis-basis"`
	Beginning      string  `json:"beginning"`
	End            string  `json:"end"`
	IceDays        int     `json:"ice-days"`
	FlTotal        float32 `json:"fl-total"`
	VolSum         float32 `json:"vol-sum"`
	Country        string  `json:"country"`
	CountryCode    string  `json:"country-code"`
	StationName    string  `json:"station-name"`
	StationCode    string  `json:"station-code"`
}

