package luhn

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsValidLuhn(t *testing.T) {
	type args struct {
		digits []int
	}
	tests := []struct {
		name    string
		args    args
		want    bool
		wantErr bool
	}{
		{
			name:    "case normal credit card num",
			args:    args{digits: []int{5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 4, 4, 4, 4}},
			want:    true,
			wantErr: false,
		},
		{
			name:    "case invalid credit card",
			args:    args{digits: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0, 1, 2, 3, 4, 5, 6}},
			want:    false,
			wantErr: false,
		},
		{
			name:    "case invalid num of digits",
			args:    args{digits: []int{1, 2}},
			want:    false,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			valid, err := IsValidLuhn(tt.args.digits)
			if tt.wantErr {
				assert.True(t, err != nil)
			} else {
				assert.Equal(t, tt.want, valid)
			}
		})
	}
}
