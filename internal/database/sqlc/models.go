// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0

package sqlc

import (
	"database/sql"
	"time"
)

type DesertFigure struct {
	ID          int64          `json:"id"`
	FullName    sql.NullString `json:"full_name"`
	FirstName   sql.NullString `json:"first_name"`
	LastName    sql.NullString `json:"last_name"`
	Type        int32          `json:"type"`
	DateOfBirth sql.NullTime   `json:"date_of_birth"`
	DateOfDeath sql.NullTime   `json:"date_of_death"`
	DateAdded   time.Time      `json:"date_added"`
	LastUpdated sql.NullTime   `json:"last_updated"`
	CreatedBy   int64          `json:"created_by"`
}

type Excerpt struct {
	ID             int64          `json:"id"`
	Body           string         `json:"body"`
	Type           int32          `json:"type"`
	ReferenceTitle sql.NullString `json:"reference_title"`
	ReferencePage  sql.NullInt32  `json:"reference_page"`
	ReferenceUrl   sql.NullString `json:"reference_url"`
	DesertFigure   int64          `json:"desert_figure"`
	DateAdded      time.Time      `json:"date_added"`
	LastUpdated    sql.NullTime   `json:"last_updated"`
	CreatedBy      int64          `json:"created_by"`
}

type ExcerptTag struct {
	ExcerptID int64 `json:"excerpt_id"`
	TagID     int64 `json:"tag_id"`
}

type Icon struct {
	ID           int64          `json:"id"`
	Url          string         `json:"url"`
	Description  sql.NullString `json:"description"`
	CreatedBy    int64          `json:"created_by"`
	DesertFigure int64          `json:"desert_figure"`
	DateAdded    time.Time      `json:"date_added"`
	LastUpdated  sql.NullTime   `json:"last_updated"`
}

type Tag struct {
	ID        int64     `json:"id"`
	Name      string    `json:"name"`
	DateAdded time.Time `json:"date_added"`
	CreatedBy int64     `json:"created_by"`
}

type User struct {
	ID            int64          `json:"id"`
	Name          string         `json:"name"`
	Email         string         `json:"email"`
	EmailVerified sql.NullBool   `json:"email_verified"`
	Image         sql.NullString `json:"image"`
}
