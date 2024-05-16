package types

var WasteType = struct {
	FailedGetAll  string
	SuccessGetAll string
	FailedDelete  string
	SuccessDelete string
}{
	FailedGetAll:  "Failed to get all waste types",
	SuccessGetAll: "Successfully get all waste types",
	FailedDelete:  "Failed to delete waste type",
	SuccessDelete: "Successfully delete waste type",
}

// ID
var InvalidIDMessage string = "Invalid ID"
var MustUUIDValidError string = "The ID must be a valid UUID"
