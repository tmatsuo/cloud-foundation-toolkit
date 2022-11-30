package bpmetadata

import (
	"testing"
)

func TestGetBpRepoPath(t *testing.T) {
	tests := []struct {
		name    string
		path    string
		want    string
		wantErr bool
	}{
		{
			name:    "simple",
			path:    "testdata/bpmetadata/terraform-google-bp01",
			want:    "testdata/bpmetadata/terraform-google-bp01",
			wantErr: false,
		},
		{
			name:    "one level nested",
			path:    "testdata/bpmetadata/terraform-google-bp01/modules/bp01-01",
			want:    "testdata/bpmetadata/terraform-google-bp01",
			wantErr: false,
		},
		{
			name:    "two level nested",
			path:    "testdata/bpmetadata/terraform-google-bp01/modules/bp01-01/subbp01-01",
			want:    "testdata/bpmetadata/terraform-google-bp01",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := getBpRootPath(tt.path)
			if got != tt.want {
				t.Errorf("getBpRootPath() = %v, want %v", got, tt.want)
			}
		})
	}
}
