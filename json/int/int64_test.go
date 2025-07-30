package jsonint_test

import (
	"encoding/json"
	"reflect"
	"testing"

	"github.com/hypershadow-io/contract/jsonint"
)

func TestInt64_MarshalJSON(t *testing.T) {
	tests := []struct {
		name    string
		value   jsonint.Int64
		want    []byte
		wantErr bool
	}{
		{
			name:    "correct",
			value:   10,
			want:    []byte(`"10"`),
			wantErr: false,
		},
		{
			name:    "correct negative",
			value:   -717,
			want:    []byte(`"-717"`),
			wantErr: false,
		},
		{
			name:    "correct negative 2",
			value:   -666,
			want:    []byte(`"-666"`),
			wantErr: false,
		},
		{
			name:    "correct 0",
			value:   -0,
			want:    []byte(`"0"`),
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := json.Marshal(tt.value)
			if (err != nil) != tt.wantErr {
				t.Errorf("MarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MarshalJSON() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInt64_UnmarshalJSON(t *testing.T) {
	tests := []struct {
		name    string
		value   jsonint.Int64
		data    []byte
		wantErr bool
	}{
		{
			name:    "correct",
			value:   10,
			data:    []byte(`"10"`),
			wantErr: false,
		},
		{
			name:    "correct",
			value:   10,
			data:    []byte(` 10`),
			wantErr: false,
		},
		{
			name:    "correct",
			value:   -33,
			data:    []byte(`"-33" `),
			wantErr: false,
		},
		{
			name:    "correct",
			value:   133,
			data:    []byte(`133`),
			wantErr: false,
		},
		{
			name:    "incorrect",
			value:   -33,
			data:    []byte(`-`),
			wantErr: true,
		},
		{
			name:    "incorrect",
			value:   -33,
			data:    []byte(`message`),
			wantErr: true,
		},
		{
			name:    "incorrect",
			value:   -33,
			data:    []byte(`{"u":1}`),
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var got jsonint.Int64
			if err := json.Unmarshal(tt.data, &got); (err != nil) != tt.wantErr {
				t.Errorf("UnmarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !tt.wantErr && !reflect.DeepEqual(got, tt.value) {
				t.Errorf("MarshalJSON() got = %v, want %v", got, tt.value)
			}
		})
	}
}
