package user

import (
	"testing"
)

func TestUser_GenerateToken(t *testing.T) {
	type fields struct {
		CommonModelFields CommonModelFields
		FacebookID        string
		Token             string
	}
	type args struct {
		secret string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "Generate JWT",
			fields: fields{
				CommonModelFields: CommonModelFields{},
				FacebookID:        "12345",
			},
			args: args{
				secret: "fiscaluno",
			},
			want: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6MCwiY3JlYXRlZEF0IjoiMDAwMS0wMS0wMVQwMDowMDowMFoiLCJ1cGRhdGVkQXQiOiIwMDAxLTAxLTAxVDAwOjAwOjAwWiIsImRlbGV0ZWRBdCI6bnVsbCwiZmFjZWJvb2tJRCI6IjEyMzQ1IiwiZXhwIjoyNTAwMDAwMDAwLCJpc3MiOiJtdSJ9.lpkRuXhnKW-wpzEeqSvj1ew_r_MI25sgucGLsxHGKSg",
			// want:    "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyIjp7ImlkIjowLCJjcmVhdGVkQXQiOiIwMDAxLTAxLTAxVDAwOjAwOjAwWiIsInVwZGF0ZWRBdCI6IjAwMDEtMDEtMDFUMDA6MDA6MDBaIiwiZGVsZXRlZEF0IjpudWxsLCJmYWNlYm9va0lEIjoiMTIzNDUifSwiZXhwIjoxNTAwMCwiaXNzIjoibXUifQ.fPkpgTaNe3E7juy_efwGJRFm6fbpMYfasv65wLmTTww",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := User{
				CommonModelFields: tt.fields.CommonModelFields,
				FacebookID:        tt.fields.FacebookID,
				Token:             tt.fields.Token,
			}
			got, err := u.GenerateToken(tt.args.secret)
			if (err != nil) != tt.wantErr {
				t.Errorf("User.GenerateToken() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("User.GenerateToken() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUser_ValidateToken(t *testing.T) {

	uMock := User{
		CommonModelFields: CommonModelFields{},
		FacebookID:        "12345",
	}

	tkn, _ := uMock.GenerateToken("fiscaluno")

	type fields struct {
		CommonModelFields CommonModelFields
		FacebookID        string
		Token             string
	}
	type args struct {
		secret string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    bool
		wantErr bool
	}{
		{
			name: "Validate JWT",
			fields: fields{
				CommonModelFields: CommonModelFields{},
				FacebookID:        "12345",
				Token:             tkn,
			},
			args: args{
				secret: "fiscaluno",
			},
			want:    false,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := User{
				CommonModelFields: tt.fields.CommonModelFields,
				FacebookID:        tt.fields.FacebookID,
				Token:             tt.fields.Token,
			}
			got, err := u.ValidateToken(tt.args.secret)
			if (err != nil) != tt.wantErr {
				t.Errorf("User.ValidateToken() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("User.ValidateToken() = %v, want %v", got, tt.want)
			}
		})
	}
}
