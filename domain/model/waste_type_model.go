package model

import (
	"time"

	"github.com/google/uuid"
)

type WasteType struct {
	ID          uuid.UUID  `json:"id"`
	Name        string     `json:"name"`
	Point       float64    `json:"point"`
	Description string     `json:"description"`
	VersionID   uuid.UUID  `json:"version_id"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   *time.Time `json:"updated_at,omitempty"`
	DeletedAt   *time.Time `json:"deleted_at,omitempty"`
	CreatedBy   *string    `json:"created_by,omitempty"`
	UpdatedBy   *string    `json:"updated_by,omitempty"`
	DeletedBy   *string    `json:"deleted_by,omitempty"`
}

type FindWasteTypeResponse struct {
	Total      int
	WasteTypes []WasteType
}

type WasteTypeUpdateWithVersionRequest struct {
	SOURCEID    string  `json:"-"`
	VERSIONID   string  `json:"-"`
	Name        string  `json:"name" db:"name"`
	Point       float64 `json:"point" db:"point"`
	Description string  `json:"description" db:"description"`
	UpdatedBy   string  `json:"created_by" db:"created_by"`
}

type WasteTypeCreateWithVersion struct {
	SOURCEID    string  `json:"-"`
	VERSIONID   string  `json:"-"`
	Name        string  `json:"name" db:"name"`
	Point       float64 `json:"point" db:"point"`
	Description string  `json:"description" db:"description"`
	CreatedBy   string  `json:"created_by" db:"created_by"`
}
