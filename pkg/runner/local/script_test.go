package local

import "testing"

func Test_getEndpoint(t *testing.T) {
	type args struct {
		base    string
		secrets map[string]string
	}
	tests := []struct {
		name string
		args args
		want string
	}{{
		name: "main",
		args: args{
			base:    "https://warp10.gra1.metrics.ovh.net/api/v0/exec",
			secrets: map[string]string{},
		},
		want: "https://warp10.gra1.metrics.ovh.net/api/v0/exec",
	}, {
		name: "with secret",
		args: args{
			base: "https://{{ host }}/api/v0/exec",
			secrets: map[string]string{
				"host": "warp10.gra1.metrics.ovh.net",
			},
		},
		want: "https://warp10.gra1.metrics.ovh.net/api/v0/exec",
	}, {
		name: "secret only",
		args: args{
			base: "{{ host }}",
			secrets: map[string]string{
				"host": "https://warp10.gra1.metrics.ovh.net/api/v0/exec",
			},
		},
		want: "https://warp10.gra1.metrics.ovh.net/api/v0/exec",
	}, {
		name: "broken",
		args: args{
			base: "https://{{ host/api/v0/exec",
			secrets: map[string]string{
				"host": "warp10.gra1.metrics.ovh.net",
			},
		},
		want: "https://warp10.gra1.metrics.ovh.net/api/v0/exec",
	}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getEndpoint(tt.args.base, tt.args.secrets); got != tt.want {
				t.Errorf("getEndpoint() = %v, want %v", got, tt.want)
			}
		})
	}
}
