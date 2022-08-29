package models

type Profiles struct {
	UlCount   int       `json:"ul_count"`
	IpCount   int       `json:"ip_count"`
	ProfileUl []Profile `json:"ul"`
	ProfileIp []Profile `json:"ip"`
	Success   bool      `json:"success"`
	Code      int       `json:"code"`
	Message   string    `json:"message"`
}

type Profile struct {
	Inn     string `json:"inn"`
	Kpp     string `json:"ogrn"`
	KppIP   string `json:"ogrnip"`
	Name    string `json:"name"`
	CeoName string `json:"ceo_name"`
}
