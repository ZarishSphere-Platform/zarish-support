package models

import "time"

// Ticket represents a support ticket
type Ticket struct {
	BaseModel

	TicketNumber string `gorm:"size:50;uniqueIndex;not null" json:"ticket_number"`
	Subject      string `gorm:"size:500;not null" json:"subject"`
	Description  string `gorm:"type:text" json:"description"`
	Status       string `gorm:"size:50;default:'open';index" json:"status"` // open, in_progress, resolved, closed
	Priority     string `gorm:"size:50;default:'medium'" json:"priority"`   // low, medium, high, urgent
	Category     string `gorm:"size:100" json:"category"`

	ReporterID uint  `gorm:"index;not null" json:"reporter_id"`
	AssigneeID *uint `gorm:"index" json:"assignee_id,omitempty"`

	ResolvedAt *time.Time `json:"resolved_at,omitempty"`
	ClosedAt   *time.Time `json:"closed_at,omitempty"`
}

// Comment represents a comment on a ticket
type Comment struct {
	BaseModel

	TicketID   uint   `gorm:"index;not null" json:"ticket_id"`
	UserID     uint   `gorm:"index;not null" json:"user_id"`
	Content    string `gorm:"type:text;not null" json:"content"`
	IsInternal bool   `json:"is_internal"` // Internal notes vs public comments
}

// TableName overrides
func (Ticket) TableName() string {
	return "tickets"
}

func (Comment) TableName() string {
	return "ticket_comments"
}
