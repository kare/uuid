package uuid

import guuid "github.com/google/uuid"

// UUID wraps Google UUID type.
type UUID struct {
	guuid.UUID
}

// New returns a Random (Version 4) UUID.
func New() UUID {
	uid, err := guuid.NewRandom()
	id := guuid.Must(uid, err)
	return UUID{
		UUID: id,
	}
}

// NewString create a Random (Version 4) UUID conveniently as a string.
func NewString() string {
	id := New()
	s := id.String()
	return s
}

// Parse decodes s into a UUID or returns an error.
func Parse(s string) (UUID, error) {
	id, err := guuid.Parse(s)
	if err != nil {
		return UUID{}, err
	}
	parsed := UUID{
		UUID: id,
	}
	return parsed, nil
}
