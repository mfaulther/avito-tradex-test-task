package model

type Statistics struct {
	Time   StatDate `json:"time"`
	Views  int
	Clicks int
	Cost   float32
	Cpc    float32
	Cpm    float32
}
