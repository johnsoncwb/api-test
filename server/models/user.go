package models

type User struct {
	ID      int         `json:"id,omitempty"`
	Name    string      `json:"name,omitempty"`
	Email   string      `json:"email,omitempty"`
	Phone   string      `json:"phone,omitempty"`
	Website string      `json:"website,omitempty"`
	Address UserAddress `json:"address,omitempty"`
	Company UserCompany `json:"company,omitempty"`
}

type UserAddress struct {
	Street  string  `json:"street,omitempty"`
	Suite   string  `json:"suite,omitempty"`
	Zipcode string  `json:"zipcode,omitempty"`
	Geo     UserGeo `json:"geo,omitempty"`
}

type UserGeo struct {
	Lat string `json:"lat,omitempty"`
	Lng string `json:"lng,omitempty"`
}

type UserCompany struct {
	Name        string `json:"name,omitempty"`
	CatchPhrase string `json:"catch_phrase,omitempty"`
	BS          string `json:"bs,omitempty"`
}
