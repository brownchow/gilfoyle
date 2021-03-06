// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/dreamvo/gilfoyle/ent/media"
	"github.com/dreamvo/gilfoyle/ent/video"
	"github.com/google/uuid"

	"github.com/facebookincubator/ent"
)

const (
	// Operation types.
	OpCreate    = ent.OpCreate
	OpDelete    = ent.OpDelete
	OpDeleteOne = ent.OpDeleteOne
	OpUpdate    = ent.OpUpdate
	OpUpdateOne = ent.OpUpdateOne

	// Node types.
	TypeMedia = "Media"
	TypeVideo = "Video"
)

// MediaMutation represents an operation that mutate the MediaSlice
// nodes in the graph.
type MediaMutation struct {
	config
	op            Op
	typ           string
	id            *uuid.UUID
	title         *string
	status        *media.Status
	created_at    *time.Time
	updated_at    *time.Time
	clearedFields map[string]struct{}
	done          bool
	oldValue      func(context.Context) (*Media, error)
}

var _ ent.Mutation = (*MediaMutation)(nil)

// mediaOption allows to manage the mutation configuration using functional options.
type mediaOption func(*MediaMutation)

// newMediaMutation creates new mutation for $n.Name.
func newMediaMutation(c config, op Op, opts ...mediaOption) *MediaMutation {
	m := &MediaMutation{
		config:        c,
		op:            op,
		typ:           TypeMedia,
		clearedFields: make(map[string]struct{}),
	}
	for _, opt := range opts {
		opt(m)
	}
	return m
}

// withMediaID sets the id field of the mutation.
func withMediaID(id uuid.UUID) mediaOption {
	return func(m *MediaMutation) {
		var (
			err   error
			once  sync.Once
			value *Media
		)
		m.oldValue = func(ctx context.Context) (*Media, error) {
			once.Do(func() {
				if m.done {
					err = fmt.Errorf("querying old values post mutation is not allowed")
				} else {
					value, err = m.Client().Media.Get(ctx, id)
				}
			})
			return value, err
		}
		m.id = &id
	}
}

// withMedia sets the old Media of the mutation.
func withMedia(node *Media) mediaOption {
	return func(m *MediaMutation) {
		m.oldValue = func(context.Context) (*Media, error) {
			return node, nil
		}
		m.id = &node.ID
	}
}

// Client returns a new `ent.Client` from the mutation. If the mutation was
// executed in a transaction (ent.Tx), a transactional client is returned.
func (m MediaMutation) Client() *Client {
	client := &Client{config: m.config}
	client.init()
	return client
}

// Tx returns an `ent.Tx` for mutations that were executed in transactions;
// it returns an error otherwise.
func (m MediaMutation) Tx() (*Tx, error) {
	if _, ok := m.driver.(*txDriver); !ok {
		return nil, fmt.Errorf("ent: mutation is not running in a transaction")
	}
	tx := &Tx{config: m.config}
	tx.init()
	return tx, nil
}

// SetID sets the value of the id field. Note that, this
// operation is accepted only on Media creation.
func (m *MediaMutation) SetID(id uuid.UUID) {
	m.id = &id
}

// ID returns the id value in the mutation. Note that, the id
// is available only if it was provided to the builder.
func (m *MediaMutation) ID() (id uuid.UUID, exists bool) {
	if m.id == nil {
		return
	}
	return *m.id, true
}

// SetTitle sets the title field.
func (m *MediaMutation) SetTitle(s string) {
	m.title = &s
}

// Title returns the title value in the mutation.
func (m *MediaMutation) Title() (r string, exists bool) {
	v := m.title
	if v == nil {
		return
	}
	return *v, true
}

// OldTitle returns the old title value of the Media.
// If the Media object wasn't provided to the builder, the object is fetched
// from the database.
// An error is returned if the mutation operation is not UpdateOne, or database query fails.
func (m *MediaMutation) OldTitle(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldTitle is allowed only on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, fmt.Errorf("OldTitle requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldTitle: %w", err)
	}
	return oldValue.Title, nil
}

// ResetTitle reset all changes of the "title" field.
func (m *MediaMutation) ResetTitle() {
	m.title = nil
}

// SetStatus sets the status field.
func (m *MediaMutation) SetStatus(value media.Status) {
	m.status = &value
}

// Status returns the status value in the mutation.
func (m *MediaMutation) Status() (r media.Status, exists bool) {
	v := m.status
	if v == nil {
		return
	}
	return *v, true
}

// OldStatus returns the old status value of the Media.
// If the Media object wasn't provided to the builder, the object is fetched
// from the database.
// An error is returned if the mutation operation is not UpdateOne, or database query fails.
func (m *MediaMutation) OldStatus(ctx context.Context) (v media.Status, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldStatus is allowed only on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, fmt.Errorf("OldStatus requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldStatus: %w", err)
	}
	return oldValue.Status, nil
}

// ResetStatus reset all changes of the "status" field.
func (m *MediaMutation) ResetStatus() {
	m.status = nil
}

// SetCreatedAt sets the created_at field.
func (m *MediaMutation) SetCreatedAt(t time.Time) {
	m.created_at = &t
}

// CreatedAt returns the created_at value in the mutation.
func (m *MediaMutation) CreatedAt() (r time.Time, exists bool) {
	v := m.created_at
	if v == nil {
		return
	}
	return *v, true
}

// OldCreatedAt returns the old created_at value of the Media.
// If the Media object wasn't provided to the builder, the object is fetched
// from the database.
// An error is returned if the mutation operation is not UpdateOne, or database query fails.
func (m *MediaMutation) OldCreatedAt(ctx context.Context) (v time.Time, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldCreatedAt is allowed only on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, fmt.Errorf("OldCreatedAt requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldCreatedAt: %w", err)
	}
	return oldValue.CreatedAt, nil
}

// ResetCreatedAt reset all changes of the "created_at" field.
func (m *MediaMutation) ResetCreatedAt() {
	m.created_at = nil
}

// SetUpdatedAt sets the updated_at field.
func (m *MediaMutation) SetUpdatedAt(t time.Time) {
	m.updated_at = &t
}

// UpdatedAt returns the updated_at value in the mutation.
func (m *MediaMutation) UpdatedAt() (r time.Time, exists bool) {
	v := m.updated_at
	if v == nil {
		return
	}
	return *v, true
}

// OldUpdatedAt returns the old updated_at value of the Media.
// If the Media object wasn't provided to the builder, the object is fetched
// from the database.
// An error is returned if the mutation operation is not UpdateOne, or database query fails.
func (m *MediaMutation) OldUpdatedAt(ctx context.Context) (v time.Time, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldUpdatedAt is allowed only on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, fmt.Errorf("OldUpdatedAt requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldUpdatedAt: %w", err)
	}
	return oldValue.UpdatedAt, nil
}

// ResetUpdatedAt reset all changes of the "updated_at" field.
func (m *MediaMutation) ResetUpdatedAt() {
	m.updated_at = nil
}

// Op returns the operation name.
func (m *MediaMutation) Op() Op {
	return m.op
}

// Type returns the node type of this mutation (Media).
func (m *MediaMutation) Type() string {
	return m.typ
}

// Fields returns all fields that were changed during
// this mutation. Note that, in order to get all numeric
// fields that were in/decremented, call AddedFields().
func (m *MediaMutation) Fields() []string {
	fields := make([]string, 0, 4)
	if m.title != nil {
		fields = append(fields, media.FieldTitle)
	}
	if m.status != nil {
		fields = append(fields, media.FieldStatus)
	}
	if m.created_at != nil {
		fields = append(fields, media.FieldCreatedAt)
	}
	if m.updated_at != nil {
		fields = append(fields, media.FieldUpdatedAt)
	}
	return fields
}

// Field returns the value of a field with the given name.
// The second boolean value indicates that this field was
// not set, or was not define in the schema.
func (m *MediaMutation) Field(name string) (ent.Value, bool) {
	switch name {
	case media.FieldTitle:
		return m.Title()
	case media.FieldStatus:
		return m.Status()
	case media.FieldCreatedAt:
		return m.CreatedAt()
	case media.FieldUpdatedAt:
		return m.UpdatedAt()
	}
	return nil, false
}

// OldField returns the old value of the field from the database.
// An error is returned if the mutation operation is not UpdateOne,
// or the query to the database was failed.
func (m *MediaMutation) OldField(ctx context.Context, name string) (ent.Value, error) {
	switch name {
	case media.FieldTitle:
		return m.OldTitle(ctx)
	case media.FieldStatus:
		return m.OldStatus(ctx)
	case media.FieldCreatedAt:
		return m.OldCreatedAt(ctx)
	case media.FieldUpdatedAt:
		return m.OldUpdatedAt(ctx)
	}
	return nil, fmt.Errorf("unknown Media field %s", name)
}

// SetField sets the value for the given name. It returns an
// error if the field is not defined in the schema, or if the
// type mismatch the field type.
func (m *MediaMutation) SetField(name string, value ent.Value) error {
	switch name {
	case media.FieldTitle:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetTitle(v)
		return nil
	case media.FieldStatus:
		v, ok := value.(media.Status)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetStatus(v)
		return nil
	case media.FieldCreatedAt:
		v, ok := value.(time.Time)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetCreatedAt(v)
		return nil
	case media.FieldUpdatedAt:
		v, ok := value.(time.Time)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetUpdatedAt(v)
		return nil
	}
	return fmt.Errorf("unknown Media field %s", name)
}

// AddedFields returns all numeric fields that were incremented
// or decremented during this mutation.
func (m *MediaMutation) AddedFields() []string {
	return nil
}

// AddedField returns the numeric value that was in/decremented
// from a field with the given name. The second value indicates
// that this field was not set, or was not define in the schema.
func (m *MediaMutation) AddedField(name string) (ent.Value, bool) {
	return nil, false
}

// AddField adds the value for the given name. It returns an
// error if the field is not defined in the schema, or if the
// type mismatch the field type.
func (m *MediaMutation) AddField(name string, value ent.Value) error {
	switch name {
	}
	return fmt.Errorf("unknown Media numeric field %s", name)
}

// ClearedFields returns all nullable fields that were cleared
// during this mutation.
func (m *MediaMutation) ClearedFields() []string {
	return nil
}

// FieldCleared returns a boolean indicates if this field was
// cleared in this mutation.
func (m *MediaMutation) FieldCleared(name string) bool {
	_, ok := m.clearedFields[name]
	return ok
}

// ClearField clears the value for the given name. It returns an
// error if the field is not defined in the schema.
func (m *MediaMutation) ClearField(name string) error {
	return fmt.Errorf("unknown Media nullable field %s", name)
}

// ResetField resets all changes in the mutation regarding the
// given field name. It returns an error if the field is not
// defined in the schema.
func (m *MediaMutation) ResetField(name string) error {
	switch name {
	case media.FieldTitle:
		m.ResetTitle()
		return nil
	case media.FieldStatus:
		m.ResetStatus()
		return nil
	case media.FieldCreatedAt:
		m.ResetCreatedAt()
		return nil
	case media.FieldUpdatedAt:
		m.ResetUpdatedAt()
		return nil
	}
	return fmt.Errorf("unknown Media field %s", name)
}

// AddedEdges returns all edge names that were set/added in this
// mutation.
func (m *MediaMutation) AddedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// AddedIDs returns all ids (to other nodes) that were added for
// the given edge name.
func (m *MediaMutation) AddedIDs(name string) []ent.Value {
	return nil
}

// RemovedEdges returns all edge names that were removed in this
// mutation.
func (m *MediaMutation) RemovedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// RemovedIDs returns all ids (to other nodes) that were removed for
// the given edge name.
func (m *MediaMutation) RemovedIDs(name string) []ent.Value {
	return nil
}

// ClearedEdges returns all edge names that were cleared in this
// mutation.
func (m *MediaMutation) ClearedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// EdgeCleared returns a boolean indicates if this edge was
// cleared in this mutation.
func (m *MediaMutation) EdgeCleared(name string) bool {
	return false
}

// ClearEdge clears the value for the given name. It returns an
// error if the edge name is not defined in the schema.
func (m *MediaMutation) ClearEdge(name string) error {
	return fmt.Errorf("unknown Media unique edge %s", name)
}

// ResetEdge resets all changes in the mutation regarding the
// given edge name. It returns an error if the edge is not
// defined in the schema.
func (m *MediaMutation) ResetEdge(name string) error {
	return fmt.Errorf("unknown Media edge %s", name)
}

// VideoMutation represents an operation that mutate the Videos
// nodes in the graph.
type VideoMutation struct {
	config
	op            Op
	typ           string
	id            *uuid.UUID
	title         *string
	status        *video.Status
	created_at    *time.Time
	updated_at    *time.Time
	clearedFields map[string]struct{}
	done          bool
	oldValue      func(context.Context) (*Video, error)
}

var _ ent.Mutation = (*VideoMutation)(nil)

// videoOption allows to manage the mutation configuration using functional options.
type videoOption func(*VideoMutation)

// newVideoMutation creates new mutation for $n.Name.
func newVideoMutation(c config, op Op, opts ...videoOption) *VideoMutation {
	m := &VideoMutation{
		config:        c,
		op:            op,
		typ:           TypeVideo,
		clearedFields: make(map[string]struct{}),
	}
	for _, opt := range opts {
		opt(m)
	}
	return m
}

// withVideoID sets the id field of the mutation.
func withVideoID(id uuid.UUID) videoOption {
	return func(m *VideoMutation) {
		var (
			err   error
			once  sync.Once
			value *Video
		)
		m.oldValue = func(ctx context.Context) (*Video, error) {
			once.Do(func() {
				if m.done {
					err = fmt.Errorf("querying old values post mutation is not allowed")
				} else {
					value, err = m.Client().Video.Get(ctx, id)
				}
			})
			return value, err
		}
		m.id = &id
	}
}

// withVideo sets the old Video of the mutation.
func withVideo(node *Video) videoOption {
	return func(m *VideoMutation) {
		m.oldValue = func(context.Context) (*Video, error) {
			return node, nil
		}
		m.id = &node.ID
	}
}

// Client returns a new `ent.Client` from the mutation. If the mutation was
// executed in a transaction (ent.Tx), a transactional client is returned.
func (m VideoMutation) Client() *Client {
	client := &Client{config: m.config}
	client.init()
	return client
}

// Tx returns an `ent.Tx` for mutations that were executed in transactions;
// it returns an error otherwise.
func (m VideoMutation) Tx() (*Tx, error) {
	if _, ok := m.driver.(*txDriver); !ok {
		return nil, fmt.Errorf("ent: mutation is not running in a transaction")
	}
	tx := &Tx{config: m.config}
	tx.init()
	return tx, nil
}

// SetID sets the value of the id field. Note that, this
// operation is accepted only on Video creation.
func (m *VideoMutation) SetID(id uuid.UUID) {
	m.id = &id
}

// ID returns the id value in the mutation. Note that, the id
// is available only if it was provided to the builder.
func (m *VideoMutation) ID() (id uuid.UUID, exists bool) {
	if m.id == nil {
		return
	}
	return *m.id, true
}

// SetTitle sets the title field.
func (m *VideoMutation) SetTitle(s string) {
	m.title = &s
}

// Title returns the title value in the mutation.
func (m *VideoMutation) Title() (r string, exists bool) {
	v := m.title
	if v == nil {
		return
	}
	return *v, true
}

// OldTitle returns the old title value of the Video.
// If the Video object wasn't provided to the builder, the object is fetched
// from the database.
// An error is returned if the mutation operation is not UpdateOne, or database query fails.
func (m *VideoMutation) OldTitle(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldTitle is allowed only on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, fmt.Errorf("OldTitle requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldTitle: %w", err)
	}
	return oldValue.Title, nil
}

// ResetTitle reset all changes of the "title" field.
func (m *VideoMutation) ResetTitle() {
	m.title = nil
}

// SetStatus sets the status field.
func (m *VideoMutation) SetStatus(v video.Status) {
	m.status = &v
}

// Status returns the status value in the mutation.
func (m *VideoMutation) Status() (r video.Status, exists bool) {
	v := m.status
	if v == nil {
		return
	}
	return *v, true
}

// OldStatus returns the old status value of the Video.
// If the Video object wasn't provided to the builder, the object is fetched
// from the database.
// An error is returned if the mutation operation is not UpdateOne, or database query fails.
func (m *VideoMutation) OldStatus(ctx context.Context) (v video.Status, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldStatus is allowed only on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, fmt.Errorf("OldStatus requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldStatus: %w", err)
	}
	return oldValue.Status, nil
}

// ResetStatus reset all changes of the "status" field.
func (m *VideoMutation) ResetStatus() {
	m.status = nil
}

// SetCreatedAt sets the created_at field.
func (m *VideoMutation) SetCreatedAt(t time.Time) {
	m.created_at = &t
}

// CreatedAt returns the created_at value in the mutation.
func (m *VideoMutation) CreatedAt() (r time.Time, exists bool) {
	v := m.created_at
	if v == nil {
		return
	}
	return *v, true
}

// OldCreatedAt returns the old created_at value of the Video.
// If the Video object wasn't provided to the builder, the object is fetched
// from the database.
// An error is returned if the mutation operation is not UpdateOne, or database query fails.
func (m *VideoMutation) OldCreatedAt(ctx context.Context) (v time.Time, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldCreatedAt is allowed only on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, fmt.Errorf("OldCreatedAt requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldCreatedAt: %w", err)
	}
	return oldValue.CreatedAt, nil
}

// ResetCreatedAt reset all changes of the "created_at" field.
func (m *VideoMutation) ResetCreatedAt() {
	m.created_at = nil
}

// SetUpdatedAt sets the updated_at field.
func (m *VideoMutation) SetUpdatedAt(t time.Time) {
	m.updated_at = &t
}

// UpdatedAt returns the updated_at value in the mutation.
func (m *VideoMutation) UpdatedAt() (r time.Time, exists bool) {
	v := m.updated_at
	if v == nil {
		return
	}
	return *v, true
}

// OldUpdatedAt returns the old updated_at value of the Video.
// If the Video object wasn't provided to the builder, the object is fetched
// from the database.
// An error is returned if the mutation operation is not UpdateOne, or database query fails.
func (m *VideoMutation) OldUpdatedAt(ctx context.Context) (v time.Time, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldUpdatedAt is allowed only on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, fmt.Errorf("OldUpdatedAt requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldUpdatedAt: %w", err)
	}
	return oldValue.UpdatedAt, nil
}

// ResetUpdatedAt reset all changes of the "updated_at" field.
func (m *VideoMutation) ResetUpdatedAt() {
	m.updated_at = nil
}

// Op returns the operation name.
func (m *VideoMutation) Op() Op {
	return m.op
}

// Type returns the node type of this mutation (Video).
func (m *VideoMutation) Type() string {
	return m.typ
}

// Fields returns all fields that were changed during
// this mutation. Note that, in order to get all numeric
// fields that were in/decremented, call AddedFields().
func (m *VideoMutation) Fields() []string {
	fields := make([]string, 0, 4)
	if m.title != nil {
		fields = append(fields, video.FieldTitle)
	}
	if m.status != nil {
		fields = append(fields, video.FieldStatus)
	}
	if m.created_at != nil {
		fields = append(fields, video.FieldCreatedAt)
	}
	if m.updated_at != nil {
		fields = append(fields, video.FieldUpdatedAt)
	}
	return fields
}

// Field returns the value of a field with the given name.
// The second boolean value indicates that this field was
// not set, or was not define in the schema.
func (m *VideoMutation) Field(name string) (ent.Value, bool) {
	switch name {
	case video.FieldTitle:
		return m.Title()
	case video.FieldStatus:
		return m.Status()
	case video.FieldCreatedAt:
		return m.CreatedAt()
	case video.FieldUpdatedAt:
		return m.UpdatedAt()
	}
	return nil, false
}

// OldField returns the old value of the field from the database.
// An error is returned if the mutation operation is not UpdateOne,
// or the query to the database was failed.
func (m *VideoMutation) OldField(ctx context.Context, name string) (ent.Value, error) {
	switch name {
	case video.FieldTitle:
		return m.OldTitle(ctx)
	case video.FieldStatus:
		return m.OldStatus(ctx)
	case video.FieldCreatedAt:
		return m.OldCreatedAt(ctx)
	case video.FieldUpdatedAt:
		return m.OldUpdatedAt(ctx)
	}
	return nil, fmt.Errorf("unknown Video field %s", name)
}

// SetField sets the value for the given name. It returns an
// error if the field is not defined in the schema, or if the
// type mismatch the field type.
func (m *VideoMutation) SetField(name string, value ent.Value) error {
	switch name {
	case video.FieldTitle:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetTitle(v)
		return nil
	case video.FieldStatus:
		v, ok := value.(video.Status)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetStatus(v)
		return nil
	case video.FieldCreatedAt:
		v, ok := value.(time.Time)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetCreatedAt(v)
		return nil
	case video.FieldUpdatedAt:
		v, ok := value.(time.Time)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetUpdatedAt(v)
		return nil
	}
	return fmt.Errorf("unknown Video field %s", name)
}

// AddedFields returns all numeric fields that were incremented
// or decremented during this mutation.
func (m *VideoMutation) AddedFields() []string {
	return nil
}

// AddedField returns the numeric value that was in/decremented
// from a field with the given name. The second value indicates
// that this field was not set, or was not define in the schema.
func (m *VideoMutation) AddedField(name string) (ent.Value, bool) {
	return nil, false
}

// AddField adds the value for the given name. It returns an
// error if the field is not defined in the schema, or if the
// type mismatch the field type.
func (m *VideoMutation) AddField(name string, value ent.Value) error {
	switch name {
	}
	return fmt.Errorf("unknown Video numeric field %s", name)
}

// ClearedFields returns all nullable fields that were cleared
// during this mutation.
func (m *VideoMutation) ClearedFields() []string {
	return nil
}

// FieldCleared returns a boolean indicates if this field was
// cleared in this mutation.
func (m *VideoMutation) FieldCleared(name string) bool {
	_, ok := m.clearedFields[name]
	return ok
}

// ClearField clears the value for the given name. It returns an
// error if the field is not defined in the schema.
func (m *VideoMutation) ClearField(name string) error {
	return fmt.Errorf("unknown Video nullable field %s", name)
}

// ResetField resets all changes in the mutation regarding the
// given field name. It returns an error if the field is not
// defined in the schema.
func (m *VideoMutation) ResetField(name string) error {
	switch name {
	case video.FieldTitle:
		m.ResetTitle()
		return nil
	case video.FieldStatus:
		m.ResetStatus()
		return nil
	case video.FieldCreatedAt:
		m.ResetCreatedAt()
		return nil
	case video.FieldUpdatedAt:
		m.ResetUpdatedAt()
		return nil
	}
	return fmt.Errorf("unknown Video field %s", name)
}

// AddedEdges returns all edge names that were set/added in this
// mutation.
func (m *VideoMutation) AddedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// AddedIDs returns all ids (to other nodes) that were added for
// the given edge name.
func (m *VideoMutation) AddedIDs(name string) []ent.Value {
	return nil
}

// RemovedEdges returns all edge names that were removed in this
// mutation.
func (m *VideoMutation) RemovedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// RemovedIDs returns all ids (to other nodes) that were removed for
// the given edge name.
func (m *VideoMutation) RemovedIDs(name string) []ent.Value {
	return nil
}

// ClearedEdges returns all edge names that were cleared in this
// mutation.
func (m *VideoMutation) ClearedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// EdgeCleared returns a boolean indicates if this edge was
// cleared in this mutation.
func (m *VideoMutation) EdgeCleared(name string) bool {
	return false
}

// ClearEdge clears the value for the given name. It returns an
// error if the edge name is not defined in the schema.
func (m *VideoMutation) ClearEdge(name string) error {
	return fmt.Errorf("unknown Video unique edge %s", name)
}

// ResetEdge resets all changes in the mutation regarding the
// given edge name. It returns an error if the edge is not
// defined in the schema.
func (m *VideoMutation) ResetEdge(name string) error {
	return fmt.Errorf("unknown Video edge %s", name)
}
