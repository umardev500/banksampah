package types

import "fmt"

var Waste = struct {
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
var FailedParseIDMessage string = "Failed to parse id"
var InvalidIDMessage string = "Invalid ID"
var MustUUIDValidError string = "The ID must be a valid UUID"

// Wallet
var wallet = "wallet"
var Wallet = struct {
	// find
	FailedGetAll  string
	SuccessGetAll string
	// Delete
	FailedDelete  string
	SuccessDelete string
	// Update
	SuccessUpdate string
	FailedUpdate  string
	// Create
	SuccessCreate string
	FailedCreate  string
	// Extension
	SuccessMoveBalance  string
	FaildMoveBalance    string
	OutOfBalance        string
	OutOfBalanceDetails string
}{
	FailedGetAll:        fmt.Sprintf("Failed to get all %s", wallet),
	SuccessGetAll:       fmt.Sprintf("Successfully get all %s", wallet),
	FailedDelete:        fmt.Sprintf("Failed to delete %s", wallet),
	SuccessDelete:       fmt.Sprintf("Successfully delete %s", wallet),
	SuccessUpdate:       fmt.Sprintf("Successfully update %s", wallet),
	FailedUpdate:        fmt.Sprintf("Failed to update %s", wallet),
	SuccessCreate:       fmt.Sprintf("Successfully create %s", wallet),
	FailedCreate:        fmt.Sprintf("Failed to create %s", wallet),
	SuccessMoveBalance:  "Successfully move balance",
	FaildMoveBalance:    "Failed to move balance",
	OutOfBalance:        "Out of balance",
	OutOfBalanceDetails: "Transfer may be failed because wallet is out of balance, or ensure it's wallet id is correct.",
}

var depo = "deposit"
var Deposit = struct {
	// find
	FailedGetAll  string
	SuccessGetAll string
	FailedGetOne  string
	SuccessGetOne string
	// Delete
	FailedDelete      string
	SuccessDelete     string
	FailedSoftDelete  string
	SuccessSoftDelete string
	// Update
	SuccessUpdate string
	FailedUpdate  string
	// Create
	SuccessCreate string
	FailedCreate  string
}{
	FailedGetAll:      fmt.Sprintf("Failed to get all %s", depo),
	SuccessGetAll:     fmt.Sprintf("Successfully get all %s", depo),
	FailedGetOne:      fmt.Sprintf("Failed to get %s by id", depo),
	SuccessGetOne:     fmt.Sprintf("Successfully get %s by id", depo),
	FailedDelete:      fmt.Sprintf("Failed to delete %s", depo),
	SuccessDelete:     fmt.Sprintf("Successfully delete %s", depo),
	FailedSoftDelete:  fmt.Sprintf("Failed to soft delete %s", depo),
	SuccessSoftDelete: fmt.Sprintf("Successfully soft delete %s", depo),
	SuccessUpdate:     fmt.Sprintf("Successfully update %s", depo),
	FailedUpdate:      fmt.Sprintf("Failed to update %s", depo),
	SuccessCreate:     fmt.Sprintf("Successfully create %s", depo),
	FailedCreate:      fmt.Sprintf("Failed to create %s", depo),
}
