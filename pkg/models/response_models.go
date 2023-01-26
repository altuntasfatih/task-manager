package models

type GetRepairsResponse struct {
	Repairs []*Repair
}
type GetRepairResponse struct {
	*Repair
}
