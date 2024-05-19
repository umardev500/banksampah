package types

import "fmt"

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

// Wallet
var wallet = "wallet"
var Wallet = struct {
	FailedGetAll  string
	SuccessGetAll string
	FailedDelete  string
	SuccessDelete string
	SuccessUpdate string
	FaildUpdate   string
	SuccessCreate string
	FailedCreate  string
}{
	FailedGetAll:  fmt.Sprintf("Failed to get all %s", wallet),
	SuccessGetAll: fmt.Sprintf("Successfully get all %s", wallet),
	FailedDelete:  fmt.Sprintf("Failed to delete %s", wallet),
	SuccessDelete: fmt.Sprintf("Successfully delete %s", wallet),
	SuccessUpdate: fmt.Sprintf("Successfully update %s", wallet),
	FaildUpdate:   fmt.Sprintf("Failed to update %s", wallet),
	SuccessCreate: fmt.Sprintf("Successfully create %s", wallet),
	FailedCreate:  fmt.Sprintf("Failed to create %s", wallet),
}
