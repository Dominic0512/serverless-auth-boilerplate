package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"github.com/google/uuid"
)

// UserProvider holds the schema definition for the UserProvider entity.
type UserProvider struct {
	ent.Schema
}

// Fields of the UserProvider.
func (UserProvider) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("user_id", uuid.UUID{}),
		field.String("picture").Optional(),
		field.Enum("name").
			NamedValues(
				"Primary", "PPIMARY",
				"Facebook", "FACEBOOK",
				"Google", "GOOGLE",
			).Default("PPIMARY"),
		field.Time("createdAt").Default(time.Now),
		field.Time("updatedAt").Default(time.Now).UpdateDefault(time.Now),
	}
}

// Edges of the UserProvider.
func (UserProvider) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).
			Field("user_id").
			Ref("user_providers").
			Unique().
			Required(),
	}
}

// Indexes of the UserProvider.
func (UserProvider) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("user_id", "name").
			Unique(),
	}
}
