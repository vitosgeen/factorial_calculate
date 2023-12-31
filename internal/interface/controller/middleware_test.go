package controller

import (
	"fmt"
	"testing"

	"factorial_calculate/internal/domain/model"
)

func Test_validateDataRequest(t *testing.T) {
	tests := []struct {
		name    string
		dataReq model.DataRequest
		wantErr error
	}{
		{
			name: "valid data request",
			dataReq: model.DataRequest{
				A: 1,
				B: 2,
			},
			wantErr: nil,
		},
		{
			name: "negative A",
			dataReq: model.DataRequest{
				A: -1,
				B: 2,
			},
			wantErr: fmt.Errorf("a and b must be positive"),
		},
		{
			name: "negative B",
			dataReq: model.DataRequest{
				A: 1,
				B: -2,
			},
			wantErr: fmt.Errorf("a and b must be positive"),
		},
		{
			name: "negative A and B",
			dataReq: model.DataRequest{
				A: -1,
				B: -2,
			},
			wantErr: fmt.Errorf("a and b must be positive"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := validateDataRequest(tt.dataReq)
			if err != nil && err.Error() != tt.wantErr.Error() {
				t.Errorf("validateDataRequest() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
