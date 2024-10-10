package util

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConvertStringToDigitArray(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name    string
		args    args
		want    []int
		wantErr bool
	}{
		{
			name:    "case invalid digits string",
			args:    args{s: "123a456"},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "case valid digits string",
			args:    args{s: "1234567890"},
			want:    []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ConvertStringToDigitArray(tt.args.s)
			assert.Equal(t, tt.wantErr, err != nil)
			assert.Equal(t, tt.want, got)
		})
	}
}
