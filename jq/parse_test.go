package jq

import (
	"reflect"
	"testing"
)

func TestParseJQ(t *testing.T) {
	tests := []struct {
		name  string
		query string
		input []byte
		want  []byte
	}{
		{
			name:  "simple person name",
			query: ".person.name",
			input: []byte(`{
				"person": {
					"name": "John",
					"age": 30,
					"city": "New York"
				}
			}`),
			want: []byte(`["John"]`),
		},
		{
			name:  "filter string",
			query: `.[] | select(.Priority == "Urgent") | .Id`,
			input: []byte(`
			[
				{
				  "Id": "1",
				  "Priority": "Low"
				},
				{
				  "Id": "2",
				  "Priority": "Urgent"
				},
				{
				  "Id": "3",
				  "Priority": "Urgent"
				}
			  ]
			`),
			want: []byte(`["2","3"]`),
		},
		{
			name:  "filter array",
			query: `.[] | select(.Tags | index("urgent") > 0) | .Id`,
			input: []byte(`
			[
				{
				  "Id": "1",
				  "Tags": ["simple", "low"]
				},
				{
				  "Id": "2",
				  "Tags": ["simple", "urgent"]
				},
				{
				  "Id": "3",
				  "Tags": ["simple", "urgent"]
				}
			  ]
			`),
			want: []byte(`["2","3"]`),
		},
		{
			name:  "filter array one found",
			query: `.[] | select(.Tags | index("low") > 0) | .Id`,
			input: []byte(`
			[
				{
				  "Id": "1",
				  "Tags": ["simple", "low"]
				},
				{
				  "Id": "2",
				  "Tags": ["simple", "urgent"]
				},
				{
				  "Id": "3",
				  "Tags": ["simple", "urgent"]
				}
			  ]
			`),
			want: []byte(`["1"]`),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseJQ(tt.input, tt.query)
			if err != nil {
				t.Fatal(err)
			} else if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ParseJQ() = %s, want %s", got, tt.want)
			}
		})
	}
}
