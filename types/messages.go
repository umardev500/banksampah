package types

var WasteType = struct {
	FailedGetAll  string
	SuccessGetAll string
	FailedDelete  string
	SuccessDelete string
	SuccessUpdate string
	FaildUpdate   string
	SuccessCreate string
	FailedCreate  string
}{
	FailedGetAll:  "Failed to get all waste types",
	SuccessGetAll: "Successfully get all waste types",
	FailedDelete:  "Failed to delete waste type",
	SuccessDelete: "Successfully delete waste type",
	SuccessUpdate: "Successfully update waste type",
	FaildUpdate:   "Failed to update waste type",
	SuccessCreate: "Successfully create waste type",
	FailedCreate:  "Failed to create waste type",
}

// ID
var InvalidIDMessage string = "Invalid ID"
var MustUUIDValidError string = "The ID must be a valid UUID"
