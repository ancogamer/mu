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
		secret  string
		expDate int64
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
				secret:  "fiscaluno",
				expDate: 15000,
			},
			want:    "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyIjp7ImlkIjowLCJjcmVhdGVkQXQiOiIwMDAxLTAxLTAxVDAwOjAwOjAwWiIsInVwZGF0ZWRBdCI6IjAwMDEtMDEtMDFUMDA6MDA6MDBaIiwiZGVsZXRlZEF0IjpudWxsLCJmYWNlYm9va0lEIjoiMTIzNDUifSwiZXhwIjoxNTAwMCwiaXNzIjoibXUifQ.fPkpgTaNe3E7juy_efwGJRFm6fbpMYfasv65wLmTTww",
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
			got, err := u.GenerateToken(tt.args.secret, tt.args.expDate)
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

	tknTrue, _ := uMock.GenerateToken("fiscaluno", 95555550000)
	tknFalse, _ := uMock.GenerateToken("fiscaluno", 15000)

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
			name: "Validate JWT Token OK",
			fields: fields{
				CommonModelFields: CommonModelFields{},
				FacebookID:        "12345",
				Token:             tknTrue,
			},
			args: args{
				secret: "fiscaluno",
			},
			want:    true,
			wantErr: false,
		},
		{
			name: "Validate JWT Token NOT OK",
			fields: fields{
				CommonModelFields: CommonModelFields{},
				FacebookID:        "12345",
				Token:             tknFalse,
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
