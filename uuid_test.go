package uuid_test

import (
	"regexp"
	"testing"

	"kkn.fi/uuid"
)

var validUUIDRegexp = "^[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}$"

func TestNew(t *testing.T) {
	id := uuid.New()
	b := []byte(id.String())
	matched, err := regexp.Match(validUUIDRegexp, b)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if !matched {
		t.Errorf("generated id did not match expected format: '%v' len=%v", id.String(), len(id.String()))
	}
}

func TestNewString(t *testing.T) {
	id := uuid.NewString()
	matched, err := regexp.Match(validUUIDRegexp, []byte(id))
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if !matched {
		t.Errorf("generated id did not match expected format: '%s' len=%v", id, len(id))
	}
}

func TestParse(t *testing.T) {
	validID := "ca12c697-468c-45e3-88e6-071614dbe7d4"
	tests := []struct {
		name    string
		input   string
		wantErr bool
	}{
		{
			"valid",
			validID,
			false,
		},
		{
			"not valid",
			validID[1:],
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt := tt
			_, err := uuid.Parse(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("Parse() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

// tmpu avoids compiler optimisations in benchmarks
var tmpu uuid.UUID

func BenchmarkUUID(b *testing.B) {
	b.ReportAllocs()
	for n := 0; n < b.N; n++ {
		tmpu = uuid.New()
	}
}
