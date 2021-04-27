package domain

import (
	"testing"
)

func TestNewAddress(t *testing.T) {
	type args struct {
		rawurl string
	}
	tests := []struct {
		name    string
		args    args
		wantA   string
		wantErr bool
	}{
		{name: "Empty address", args: args{rawurl: ""}, wantErr: true},
		{name: "Check domain with http scheme", args: args{rawurl: "http://google.com"}, wantA: "http://google.com", wantErr: false},
		{name: "Check domain without scheme", args: args{rawurl: "google.com"}, wantA: "http://google.com", wantErr: false},
		{name: "ComplexDomain", args: args{rawurl: "http://reddit.com/r/notfunny"}, wantA: "http://reddit.com/r/notfunny", wantErr: false},
		{name: "ComplexDomain without scheme", args: args{rawurl: "reddit.com/r/notfunny"}, wantA: "http://reddit.com/r/notfunny", wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotA, err := NewAddress(tt.args.rawurl)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewAddress() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if err != nil {
				return
			}
			if gotA.u == nil {
				t.Error("URL is nil")
				return
			}
			if tt.wantA != gotA.u.String() {
				t.Errorf("Url cannot matched expected = %v Parsed = %v", tt.wantA, gotA.u.String())
			}
		})
	}
}
