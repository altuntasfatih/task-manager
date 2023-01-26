package models

type Repair struct {
	Id        string `json:"id"`
	Email     string `json:"email"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

func NewRepair(id, email, firsName, lastName string) *Repair {
	return &Repair{Id: id, Email: email, FirstName: firsName, LastName: lastName}
}
