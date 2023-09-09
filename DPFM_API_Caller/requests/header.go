package requests

type Header struct {
	InspectionLot       int		`json:"InspectionLot"`
	IsMarkedForDeletion *bool	`json:"IsMarkedForDeletion"`
}
