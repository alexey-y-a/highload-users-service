package config

import "testing"

func TestGetServerAddr(t *testing.T) {
	tests := []struct {
		name string
		host string
		port int
		want string
	}{
		{
			name: "ipv4_any",
			host: "0.0.0.0",
			port: 8080,
			want: "0.0.0.0:8080",
		},
		{
			name: "ipv4_localhost",
			host: "127.0.0.1",
			port: 9090,
			want: "127.0.0.1:9090",
		},
		{
			name: "ipv6_localhost",
			host: "::1",
			port: 8080,
			want: "[::1]:8080",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			cfg := Config{
				env: &EnvSetting{
					ServerHost: test.host,
					ServerPort: test.port,
				},
			}

			got := cfg.GetServerAddr()
			if got != test.want {
				t.Fatalf("GetServerAddr() = %q, want %q", got, test.want)
			}
		})
	}
}
