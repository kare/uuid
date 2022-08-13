package uuid_test

import (
	"regexp"
	"testing"

	"kkn.fi/uuid"
)

func TestUUID(t *testing.T) {
	id := uuid.New()
	re := "^[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}$"
	matched, err := regexp.Match(re, []byte(id.String()))
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if !matched {
		t.Errorf("generated id did not match expected format: '%v' len=%v", id.String(), len(id.String()))
	}
	t.Logf("%s len=%v", id.String(), len(id.String()))
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
