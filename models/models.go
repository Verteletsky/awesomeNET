package models

type Profiles struct {
	Count   int       `json:"ul_count"`
	Profile []Profile `json:"ul"`
	Success bool      `json:"success"`
	Code    int       `json:"code"`
	Message string    `json:"message"`
}

type Profile struct {
	Inn     string `json:"inn"`
	Kpp     string `json:"ogrn"`
	Name    string `json:"name"`
	CeoName string `json:"ceo_name"`
}
