// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"github.com/dreamvo/gilfoyle/ent/media"
	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/google/uuid"
)

// Media is the model entity for the Media schema.
type Media struct {
	config `json:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// Title holds the value of the "title" field.
	Title string `json:"title,omitempty"`
	// Status holds the value of the "status" field.
	Status media.Status `json:"status,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Media) scanValues() []interface{} {
	return []interface{}{
		&uuid.UUID{},      // id
		&sql.NullString{}, // title
		&sql.NullString{}, // status
		&sql.NullTime{},   // created_at
		&sql.NullTime{},   // updated_at
	}
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Media fields.
func (m *Media) assignValues(values ...interface{}) error {
	if m, n := len(values), len(media.Columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	if value, ok := values[0].(*uuid.UUID); !ok {
		return fmt.Errorf("unexpected type %T for field id", values[0])
	} else if value != nil {
		m.ID = *value
	}
	values = values[1:]
	if value, ok := values[0].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field title", values[0])
	} else if value.Valid {
		m.Title = value.String
	}
	if value, ok := values[1].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field status", values[1])
	} else if value.Valid {
		m.Status = media.Status(value.String)
	}
	if value, ok := values[2].(*sql.NullTime); !ok {
		return fmt.Errorf("unexpected type %T for field created_at", values[2])
	} else if value.Valid {
		m.CreatedAt = value.Time
	}
	if value, ok := values[3].(*sql.NullTime); !ok {
		return fmt.Errorf("unexpected type %T for field updated_at", values[3])
	} else if value.Valid {
		m.UpdatedAt = value.Time
	}
	return nil
}

// Update returns a builder for updating this Media.
// Note that, you need to call Media.Unwrap() before calling this method, if this Media
// was returned from a transaction, and the transaction was committed or rolled back.
func (m *Media) Update() *MediaUpdateOne {
	return (&MediaClient{config: m.config}).UpdateOne(m)
}

// Unwrap unwraps the entity that was returned from a transaction after it was closed,
// so that all next queries will be executed through the driver which created the transaction.
func (m *Media) Unwrap() *Media {
	tx, ok := m.config.driver.(*txDriver)
	if !ok {
		panic("ent: Media is not a transactional entity")
	}
	m.config.driver = tx.drv
	return m
}

// String implements the fmt.Stringer.
func (m *Media) String() string {
	var builder strings.Builder
	builder.WriteString("Media(")
	builder.WriteString(fmt.Sprintf("id=%v", m.ID))
	builder.WriteString(", title=")
	builder.WriteString(m.Title)
	builder.WriteString(", status=")
	builder.WriteString(fmt.Sprintf("%v", m.Status))
	builder.WriteString(", created_at=")
	builder.WriteString(m.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", updated_at=")
	builder.WriteString(m.UpdatedAt.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// MediaSlice is a parsable slice of Media.
type MediaSlice []*Media

func (m MediaSlice) config(cfg config) {
	for _i := range m {
		m[_i].config = cfg
	}
}
