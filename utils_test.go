package ginauth

import "testing"

func Test_extractTokenFromAuthHeader(t *testing.T) {
	type args struct {
		val string
	}
	tests := []struct {
		name      string
		args      args
		wantToken string
		wantOk    bool
	}{
		{
			name: "Get Bearer auth token",
			args: args{
				val: "bearer dGVzdHRva2VuCg==",
			},
			wantToken: "dGVzdHRva2VuCg==",
			wantOk:    true,
		},
		{
			name: "NG blank",
			args: args{
				val: "",
			},
			wantToken: "",
			wantOk:    false,
		},
		{
			name: "NG invalid string pattern 1",
			args: args{
				val: "dGVzdHRva2VuCg==",
			},
			wantToken: "",
			wantOk:    false,
		},
		{
			name: "NG invalid string pattern 2",
			args: args{
				val: "bearer dGVzdHRva 2VuCg==",
			},
			wantToken: "",
			wantOk:    false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotToken, gotOk := extractTokenFromAuthHeader(tt.args.val)
			if gotToken != tt.wantToken {
				t.Errorf("extractTokenFromAuthHeader() gotToken = %v, want %v", gotToken, tt.wantToken)
			}
			if gotOk != tt.wantOk {
				t.Errorf("extractTokenFromAuthHeader() gotOk = %v, want %v", gotOk, tt.wantOk)
			}
		})
	}
}
