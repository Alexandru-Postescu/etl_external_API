package model

// Struct representation of the json fields
type Person struct {
	Firstname string `json:"first"`
	Lastname  string `json:"last"`
	Email     string `json:"email"`
	Address   string `json:"address"`
	Created   string `json:"created"`
	Balance   string `json:"balance"`
}

// Get Key returns the first letter of the first name of the struct
func (p Person) GetKey() string {
	// TODO: what if the Firstname is an empty string?
	return p.Firstname[0:1]
}
