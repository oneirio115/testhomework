package schema

import "github.com/facebookincubator/ent"

// Testing holds the schema definition for the Testing entity.
type Testing struct {
	ent.Schema
}

// Fields of the Testing.
func (Testing) Fields() []ent.Field {
	return nil
}

// Edges of the Testing.
func (Testing) Edges() []ent.Edge {
	return nil
}
