package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New),
		field.Enum("state").NamedValues("Pending", "PENDING", "Verified", "VERIFIED").Default("VERIFIED"),
		field.Enum("role").NamedValues("User", "USER", "Admin", "ADMIN").Default("USER"),
		field.String("name").NotEmpty(),
		field.String("email").Unique(),
		field.String("password").Optional().Nillable(),
		field.Time("createdAt").Default(time.Now),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("user_providers", UserProvider.Type),
	}
}
