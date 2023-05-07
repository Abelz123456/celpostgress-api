package dto

type FilterPermissionPolicyUserDto struct {
	CompanyName string `json:"companyName"`
	Address     string `json:"address"`
}

type PagePermissionPolicyUserDto struct {
	Oid         string `json:"oid"`
	CompanyName string `json:"companyName"`
	Address     string `json:"address"`
}

type CreatePermissionPolicyUserDto struct {
	CompanyName   string `json:"companyName"`
	Address       string `json:"address"`
	ContactPerson string `json:"contactPerson"`
	PhoneNumber   string `json:"phoneNumber"`
	EmailName     string `json:"emailName"`
	Password      string `json:"password"`
	LevelUser     int    `json:"levelUser"`
	AndroidToken  string `json:"androidToken"`
	UserInserted  string `json:"userInserted"`
}

type UpdatePermissionPolicyUserDto struct {
	Oid           string `json:"oid"`
	CompanyName   string `json:"companyName"`
	Address       string `json:"address"`
	ContactPerson string `json:"contactPerson"`
	PhoneNumber   string `json:"phoneNumber"`
	EmailName     string `json:"emailName"`
	Password      string `json:"password"`
	LevelUser     int    `json:"levelUser"`
	AndroidToken  string `json:"androidToken"`
	LastUserId    string `json:"lastUserId"`
}
