package models

type StationDataResponse struct {
	Beginning      string  `json:"beginning"`
	End            string  `json:"end"`
	SeasonLength   int     `json:"season_length"`
	IceDays        int     `json:"ice_days"`
	FlTotal        float32 `json:"fl_total"`
	VolSum         float32 `json:"vol_sum"`
	MaxThickness   float32 `json:"max_thickness,omitempty"`
	MaxThickness2  float32 `json:"max_thickness2,omitempty"`
	Country        string  `json:"country"`
	CountryCode    string  `json:"country_code"`
	StationName    string  `json:"station_name"`
	StationCode    string  `json:"station_code"`
	RestrictedDays []int   `json:"restricted_days"`
}

type AnalysisDataResponse struct {
	AnalysisBasis string  `json:"analysis-basis"`
	Beginning     string  `json:"beginning"`
	End           string  `json:"end"`
	IceDays       float32 `json:"ice_days"`
	FlTotal       float32 `json:"fl_total"`
	VolSum        float32 `json:"vol_sum"`
	Country       string  `json:"country"`
	CountryCode   string  `json:"country_code"`
	StationName   string  `json:"station_name"`
	StationCode   string  `json:"station_code"`
}
