package eoldate

import (
	"testing"
)

func TestClient_IsSupportedSoftwareVersion(t *testing.T) {
	c := NewClient()
	type args struct {
		softwareName string
		version      string
	}
	tests := []struct {
		name    string
		args    args
		want    bool
		wantErr bool
	}{
		{name: "test IsSupportedSoftwareVersion", args: args{
			softwareName: "dotnetfx",
			version:      "4.0.30319",
		}, want: false, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, _, err := c.IsSupportedSoftwareVersion(tt.args.softwareName, tt.args.version)
			if (err != nil) != tt.wantErr {
				t.Errorf("IsSupportedSoftwareVersion() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("IsSupportedSoftwareVersion() got = %v, want %v", got, tt.want)
			}
		})
	}
}
