package client

import (
	"fmt"
	"testing"

	"github.com/mtslzr/pokeapi-go/structs"
	"github.com/stretchr/testify/assert"
)

func Test_pokeClient_FetchPokemon(t *testing.T) {
	type args struct {
		ID string
	}
	tests := []struct {
		name    string
		args    args
		want    *structs.Pokemon
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name:    "Fetch valid pokemon by ID",
			args:    args{ID: "25"},                    // Assuming 25 is a valid ID for Pikachu
			want:    &structs.Pokemon{Name: "pikachu"}, // Fill in the expected details
			wantErr: assert.NoError,
		},
		{
			name:    "Fetch valid pokemon by name",
			args:    args{ID: "pikachu"},               // Assuming the function also accepts names
			want:    &structs.Pokemon{Name: "pikachu"}, // Fill in the expected details
			wantErr: assert.NoError,
		},
		{
			name:    "Fetch with invalid ID",
			args:    args{ID: "invalid"},
			want:    nil,
			wantErr: assert.Error,
		},
		{
			name:    "Fetch with empty ID",
			args:    args{ID: ""},
			want:    nil,
			wantErr: assert.NoError,
		},
		{
			name:    "Fetch with boundary ID",
			args:    args{ID: "893"},                  // Assuming 893 is the highest valid ID at the time
			want:    &structs.Pokemon{Name: "zarude"}, // Fill in the expected details
			wantErr: assert.NoError,
		},
		{
			name:    "Fetch with non-numeric ID",
			args:    args{ID: "pika!"},
			want:    nil,
			wantErr: assert.Error,
		},
		{
			name:    "Fetch with non-existent numeric ID",
			args:    args{ID: "9999"}, // Assuming 9999 is not a valid ID
			want:    nil,
			wantErr: assert.Error,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := pokeClient{}
			got, err := p.FetchPokemon(tt.args.ID)
			if !tt.wantErr(t, err, fmt.Sprintf("FetchPokemon(%v)", tt.args.ID)) {
				return
			}
			if tt.want != nil {
				assert.Equalf(t, tt.want.Name, got.Name, "FetchPokemon(%v)", tt.args.ID)
			}
		})
	}
}
