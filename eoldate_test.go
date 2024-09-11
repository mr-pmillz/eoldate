package eoldate

import (
	"testing"
	"time"
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
			got, _, _, err := c.IsSupportedSoftwareVersion(tt.args.softwareName, tt.args.version)
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

func TestCalculateTimeDifference(t *testing.T) {
	now := time.Now()
	type args struct {
		endDate time.Time
	}
	tests := []struct {
		name  string
		args  args
		want  int
		want1 int
		want2 int
	}{
		{
			name:  "One year ago",
			args:  args{endDate: now.AddDate(-1, 0, 0)},
			want:  1,
			want1: 0,
			want2: 0,
		},
		{
			name:  "Six months ago",
			args:  args{endDate: now.AddDate(0, -6, 0)},
			want:  0,
			want1: 6,
			want2: 0,
		},
		{
			name:  "One month and 15 days ago",
			args:  args{endDate: now.AddDate(0, -1, -15)},
			want:  0,
			want1: 1,
			want2: 15,
		},
		{
			name:  "One year and one month in the future",
			args:  args{endDate: now.AddDate(1, 1, 0)},
			want:  1,
			want1: 1,
			want2: 0,
		},
		{
			name:  "Six months in the future",
			args:  args{endDate: now.AddDate(0, 6, 0)},
			want:  0,
			want1: 6,
			want2: 0,
		},
		{
			name:  "15 days in the future",
			args:  args{endDate: now.AddDate(0, 0, 15)},
			want:  0,
			want1: 0,
			want2: 15,
		},
		{
			name:  "Current date",
			args:  args{endDate: now},
			want:  0,
			want1: 0,
			want2: 0,
		},
		{
			name:  "Two years, three months, and 10 days ago",
			args:  args{endDate: now.AddDate(-2, -3, -10)},
			want:  2,
			want1: 3,
			want2: 10,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, got2 := CalculateTimeDifference(tt.args.endDate)
			if got != tt.want {
				t.Errorf("CalculateTimeDifference() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("CalculateTimeDifference() got1 = %v, want %v", got1, tt.want1)
			}
			if got2 != tt.want2 {
				t.Errorf("CalculateTimeDifference() got2 = %v, want %v", got2, tt.want2)
			}
		})
	}
}
