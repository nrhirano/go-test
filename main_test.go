package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_triangle(t *testing.T) {
	type args struct {
		bottom int
		height int
	}
	tests := []struct {
		name    string
		args    args
		want    int
		wantErr bool
	}{
		{
			name:    "success",
			args:    args{bottom: 2, height: 3},
			want:    3,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := triangle(tt.args.bottom, tt.args.height)
			if tt.wantErr {
				assert.Error(t, err)
				return
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.want, got)
			}
		})
	}
}
