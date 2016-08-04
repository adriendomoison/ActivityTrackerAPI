package DTO

type Event struct {
	Id              int `json:"id"`
	Name            string `json:"name"`
	StartDate       string `json:"startDate"`
	EndDate         string `json:"endDate"`
	Address         string `json:"address"`
	Credit          float64 `json:"credit"`
	Constraints     []EventRegistrationConstraints `json:"constraints"`
	Details         string `json:"details"`
	MaxParticipants int `json:"maxParticipants"`
	Status          string `json:"status"`
	Type            string `json:"type"`
	Group           string `json:"group"`
	AttachedProgram int `json:"attachedProgram"`
}

type EventRegistrationConstraints int

const (
	CANT_BE_DROPPED EventRegistrationConstraints = iota
	ONE_BY_GROUPE
	ONE_BY_TYPE
)