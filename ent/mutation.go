// Code generated by entc, DO NOT EDIT.

package ent

import (
	"annie-baby/ent/predicate"
	"annie-baby/ent/video"
	"context"
	"fmt"
	"sync"
	"time"

	"entgo.io/ent"
)

const (
	// Operation types.
	OpCreate    = ent.OpCreate
	OpDelete    = ent.OpDelete
	OpDeleteOne = ent.OpDeleteOne
	OpUpdate    = ent.OpUpdate
	OpUpdateOne = ent.OpUpdateOne

	// Node types.
	TypeVideo = "Video"
)

// VideoMutation represents an operation that mutates the Video nodes in the graph.
type VideoMutation struct {
	config
	op            Op
	typ           string
	id            *int
	link          *string
	_type         *string
	title         *string
	output_dir    *string
	proxy         *string
	status        *video.Status
	created_at    *time.Time
	updated_at    *time.Time
	clearedFields map[string]struct{}
	done          bool
	oldValue      func(context.Context) (*Video, error)
	predicates    []predicate.Video
}

var _ ent.Mutation = (*VideoMutation)(nil)

// videoOption allows management of the mutation configuration using functional options.
type videoOption func(*VideoMutation)

// newVideoMutation creates new mutation for the Video entity.
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

// withVideoID sets the ID field of the mutation.
func withVideoID(id int) videoOption {
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

// ID returns the ID value in the mutation. Note that the ID is only available
// if it was provided to the builder or after it was returned from the database.
func (m *VideoMutation) ID() (id int, exists bool) {
	if m.id == nil {
		return
	}
	return *m.id, true
}

// SetLink sets the "link" field.
func (m *VideoMutation) SetLink(s string) {
	m.link = &s
}

// Link returns the value of the "link" field in the mutation.
func (m *VideoMutation) Link() (r string, exists bool) {
	v := m.link
	if v == nil {
		return
	}
	return *v, true
}

// OldLink returns the old "link" field's value of the Video entity.
// If the Video object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *VideoMutation) OldLink(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldLink is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, fmt.Errorf("OldLink requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldLink: %w", err)
	}
	return oldValue.Link, nil
}

// ResetLink resets all changes to the "link" field.
func (m *VideoMutation) ResetLink() {
	m.link = nil
}

// SetType sets the "type" field.
func (m *VideoMutation) SetType(s string) {
	m._type = &s
}

// GetType returns the value of the "type" field in the mutation.
func (m *VideoMutation) GetType() (r string, exists bool) {
	v := m._type
	if v == nil {
		return
	}
	return *v, true
}

// OldType returns the old "type" field's value of the Video entity.
// If the Video object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *VideoMutation) OldType(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldType is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, fmt.Errorf("OldType requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldType: %w", err)
	}
	return oldValue.Type, nil
}

// ResetType resets all changes to the "type" field.
func (m *VideoMutation) ResetType() {
	m._type = nil
}

// SetTitle sets the "title" field.
func (m *VideoMutation) SetTitle(s string) {
	m.title = &s
}

// Title returns the value of the "title" field in the mutation.
func (m *VideoMutation) Title() (r string, exists bool) {
	v := m.title
	if v == nil {
		return
	}
	return *v, true
}

// OldTitle returns the old "title" field's value of the Video entity.
// If the Video object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *VideoMutation) OldTitle(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldTitle is only allowed on UpdateOne operations")
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

// ResetTitle resets all changes to the "title" field.
func (m *VideoMutation) ResetTitle() {
	m.title = nil
}

// SetOutputDir sets the "output_dir" field.
func (m *VideoMutation) SetOutputDir(s string) {
	m.output_dir = &s
}

// OutputDir returns the value of the "output_dir" field in the mutation.
func (m *VideoMutation) OutputDir() (r string, exists bool) {
	v := m.output_dir
	if v == nil {
		return
	}
	return *v, true
}

// OldOutputDir returns the old "output_dir" field's value of the Video entity.
// If the Video object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *VideoMutation) OldOutputDir(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldOutputDir is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, fmt.Errorf("OldOutputDir requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldOutputDir: %w", err)
	}
	return oldValue.OutputDir, nil
}

// ResetOutputDir resets all changes to the "output_dir" field.
func (m *VideoMutation) ResetOutputDir() {
	m.output_dir = nil
}

// SetProxy sets the "proxy" field.
func (m *VideoMutation) SetProxy(s string) {
	m.proxy = &s
}

// Proxy returns the value of the "proxy" field in the mutation.
func (m *VideoMutation) Proxy() (r string, exists bool) {
	v := m.proxy
	if v == nil {
		return
	}
	return *v, true
}

// OldProxy returns the old "proxy" field's value of the Video entity.
// If the Video object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *VideoMutation) OldProxy(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldProxy is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, fmt.Errorf("OldProxy requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldProxy: %w", err)
	}
	return oldValue.Proxy, nil
}

// ResetProxy resets all changes to the "proxy" field.
func (m *VideoMutation) ResetProxy() {
	m.proxy = nil
}

// SetStatus sets the "status" field.
func (m *VideoMutation) SetStatus(v video.Status) {
	m.status = &v
}

// Status returns the value of the "status" field in the mutation.
func (m *VideoMutation) Status() (r video.Status, exists bool) {
	v := m.status
	if v == nil {
		return
	}
	return *v, true
}

// OldStatus returns the old "status" field's value of the Video entity.
// If the Video object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *VideoMutation) OldStatus(ctx context.Context) (v video.Status, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldStatus is only allowed on UpdateOne operations")
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

// ResetStatus resets all changes to the "status" field.
func (m *VideoMutation) ResetStatus() {
	m.status = nil
}

// SetCreatedAt sets the "created_at" field.
func (m *VideoMutation) SetCreatedAt(t time.Time) {
	m.created_at = &t
}

// CreatedAt returns the value of the "created_at" field in the mutation.
func (m *VideoMutation) CreatedAt() (r time.Time, exists bool) {
	v := m.created_at
	if v == nil {
		return
	}
	return *v, true
}

// OldCreatedAt returns the old "created_at" field's value of the Video entity.
// If the Video object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *VideoMutation) OldCreatedAt(ctx context.Context) (v time.Time, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldCreatedAt is only allowed on UpdateOne operations")
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

// ResetCreatedAt resets all changes to the "created_at" field.
func (m *VideoMutation) ResetCreatedAt() {
	m.created_at = nil
}

// SetUpdatedAt sets the "updated_at" field.
func (m *VideoMutation) SetUpdatedAt(t time.Time) {
	m.updated_at = &t
}

// UpdatedAt returns the value of the "updated_at" field in the mutation.
func (m *VideoMutation) UpdatedAt() (r time.Time, exists bool) {
	v := m.updated_at
	if v == nil {
		return
	}
	return *v, true
}

// OldUpdatedAt returns the old "updated_at" field's value of the Video entity.
// If the Video object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *VideoMutation) OldUpdatedAt(ctx context.Context) (v time.Time, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldUpdatedAt is only allowed on UpdateOne operations")
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

// ResetUpdatedAt resets all changes to the "updated_at" field.
func (m *VideoMutation) ResetUpdatedAt() {
	m.updated_at = nil
}

// Where appends a list predicates to the VideoMutation builder.
func (m *VideoMutation) Where(ps ...predicate.Video) {
	m.predicates = append(m.predicates, ps...)
}

// Op returns the operation name.
func (m *VideoMutation) Op() Op {
	return m.op
}

// Type returns the node type of this mutation (Video).
func (m *VideoMutation) Type() string {
	return m.typ
}

// Fields returns all fields that were changed during this mutation. Note that in
// order to get all numeric fields that were incremented/decremented, call
// AddedFields().
func (m *VideoMutation) Fields() []string {
	fields := make([]string, 0, 8)
	if m.link != nil {
		fields = append(fields, video.FieldLink)
	}
	if m._type != nil {
		fields = append(fields, video.FieldType)
	}
	if m.title != nil {
		fields = append(fields, video.FieldTitle)
	}
	if m.output_dir != nil {
		fields = append(fields, video.FieldOutputDir)
	}
	if m.proxy != nil {
		fields = append(fields, video.FieldProxy)
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

// Field returns the value of a field with the given name. The second boolean
// return value indicates that this field was not set, or was not defined in the
// schema.
func (m *VideoMutation) Field(name string) (ent.Value, bool) {
	switch name {
	case video.FieldLink:
		return m.Link()
	case video.FieldType:
		return m.GetType()
	case video.FieldTitle:
		return m.Title()
	case video.FieldOutputDir:
		return m.OutputDir()
	case video.FieldProxy:
		return m.Proxy()
	case video.FieldStatus:
		return m.Status()
	case video.FieldCreatedAt:
		return m.CreatedAt()
	case video.FieldUpdatedAt:
		return m.UpdatedAt()
	}
	return nil, false
}

// OldField returns the old value of the field from the database. An error is
// returned if the mutation operation is not UpdateOne, or the query to the
// database failed.
func (m *VideoMutation) OldField(ctx context.Context, name string) (ent.Value, error) {
	switch name {
	case video.FieldLink:
		return m.OldLink(ctx)
	case video.FieldType:
		return m.OldType(ctx)
	case video.FieldTitle:
		return m.OldTitle(ctx)
	case video.FieldOutputDir:
		return m.OldOutputDir(ctx)
	case video.FieldProxy:
		return m.OldProxy(ctx)
	case video.FieldStatus:
		return m.OldStatus(ctx)
	case video.FieldCreatedAt:
		return m.OldCreatedAt(ctx)
	case video.FieldUpdatedAt:
		return m.OldUpdatedAt(ctx)
	}
	return nil, fmt.Errorf("unknown Video field %s", name)
}

// SetField sets the value of a field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *VideoMutation) SetField(name string, value ent.Value) error {
	switch name {
	case video.FieldLink:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetLink(v)
		return nil
	case video.FieldType:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetType(v)
		return nil
	case video.FieldTitle:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetTitle(v)
		return nil
	case video.FieldOutputDir:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetOutputDir(v)
		return nil
	case video.FieldProxy:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetProxy(v)
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

// AddedFields returns all numeric fields that were incremented/decremented during
// this mutation.
func (m *VideoMutation) AddedFields() []string {
	return nil
}

// AddedField returns the numeric value that was incremented/decremented on a field
// with the given name. The second boolean return value indicates that this field
// was not set, or was not defined in the schema.
func (m *VideoMutation) AddedField(name string) (ent.Value, bool) {
	return nil, false
}

// AddField adds the value to the field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *VideoMutation) AddField(name string, value ent.Value) error {
	switch name {
	}
	return fmt.Errorf("unknown Video numeric field %s", name)
}

// ClearedFields returns all nullable fields that were cleared during this
// mutation.
func (m *VideoMutation) ClearedFields() []string {
	return nil
}

// FieldCleared returns a boolean indicating if a field with the given name was
// cleared in this mutation.
func (m *VideoMutation) FieldCleared(name string) bool {
	_, ok := m.clearedFields[name]
	return ok
}

// ClearField clears the value of the field with the given name. It returns an
// error if the field is not defined in the schema.
func (m *VideoMutation) ClearField(name string) error {
	return fmt.Errorf("unknown Video nullable field %s", name)
}

// ResetField resets all changes in the mutation for the field with the given name.
// It returns an error if the field is not defined in the schema.
func (m *VideoMutation) ResetField(name string) error {
	switch name {
	case video.FieldLink:
		m.ResetLink()
		return nil
	case video.FieldType:
		m.ResetType()
		return nil
	case video.FieldTitle:
		m.ResetTitle()
		return nil
	case video.FieldOutputDir:
		m.ResetOutputDir()
		return nil
	case video.FieldProxy:
		m.ResetProxy()
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

// AddedEdges returns all edge names that were set/added in this mutation.
func (m *VideoMutation) AddedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// AddedIDs returns all IDs (to other nodes) that were added for the given edge
// name in this mutation.
func (m *VideoMutation) AddedIDs(name string) []ent.Value {
	return nil
}

// RemovedEdges returns all edge names that were removed in this mutation.
func (m *VideoMutation) RemovedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// RemovedIDs returns all IDs (to other nodes) that were removed for the edge with
// the given name in this mutation.
func (m *VideoMutation) RemovedIDs(name string) []ent.Value {
	return nil
}

// ClearedEdges returns all edge names that were cleared in this mutation.
func (m *VideoMutation) ClearedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// EdgeCleared returns a boolean which indicates if the edge with the given name
// was cleared in this mutation.
func (m *VideoMutation) EdgeCleared(name string) bool {
	return false
}

// ClearEdge clears the value of the edge with the given name. It returns an error
// if that edge is not defined in the schema.
func (m *VideoMutation) ClearEdge(name string) error {
	return fmt.Errorf("unknown Video unique edge %s", name)
}

// ResetEdge resets all changes to the edge with the given name in this mutation.
// It returns an error if the edge is not defined in the schema.
func (m *VideoMutation) ResetEdge(name string) error {
	return fmt.Errorf("unknown Video edge %s", name)
}
