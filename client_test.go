package amazon_scraper

import (
	"context"
	"testing"
)

func TestClient_GetAppInfo(t *testing.T) {
	cli := New()

	type args struct {
		ctx  context.Context
		asin string
	}
	tests := []struct {
		name    string
		args    args
		wantErr error
	}{
		{
			name: "success",
			args: args{
				asin: "B0937HJ76H",
			},
		},
		{
			name: "not found",
			args: args{
				asin: "XXXXXXXXXX",
			},
			wantErr: ErrNotFound,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := cli.GetAppInfo(tt.args.asin)
			if (err != nil && tt.wantErr == nil) || (err == nil && tt.wantErr != nil) {
				t.Errorf("GetAppInfo() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if tt.wantErr == nil {
				if got.Icon == "" {
					t.Errorf("%s icon is empty", tt.args.asin)
				}
			}
		})
	}
}
