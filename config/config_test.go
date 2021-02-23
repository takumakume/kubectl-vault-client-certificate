package config

import (
	"fmt"
	"io/ioutil"
	"os"
	"reflect"
	"testing"

	"github.com/sergi/go-diff/diffmatchpatch"
)

func TestNewConfig(t *testing.T) {
	config := `---
vault:
- name: vault1
  address: https://localhost:8200
  tls:
    cert: /path/to/localhost.crt
    key: /path/to/localhost.key
    ca_cert: /path/to/ca.crt
- name: notls
  address: http://localhost:8200
vault_pki:
- name: pki1
  vault_server_name: vault1
  path: pki_int/issue/example-dot-local
  common_name: example.local
  ttl: 24h
`

	configFile, err := ioutil.TempFile("", "config")
	defer os.Remove(configFile.Name())
	if err != nil {
		panic(err)
	}
	_, err = configFile.WriteString(config)
	if err != nil {
		panic(err)
	}

	type args struct {
		path string
	}
	tests := []struct {
		name    string
		args    args
		want    *Config
		wantErr bool
	}{
		{
			name: "ok",
			args: args{
				path: configFile.Name(),
			},
			want: &Config{
				path: configFile.Name(),
				VaultServers: []VaultServerConfig{
					{
						Name:      "vault1",
						Address:   "https://localhost:8200",
						Namespace: "",
						TLS: VaultServerTLSConfig{
							Cert:       "/path/to/localhost.crt",
							Key:        "/path/to/localhost.key",
							CACert:     "/path/to/ca.crt",
							SkipVerify: false,
						},
					},
					{
						Name:      "notls",
						Address:   "http://localhost:8200",
						Namespace: "",
					},
				},
				VaultPKIs: []VaultPKIConfig{
					{
						Name:            "pki1",
						VaultServerName: "vault1",
						Path:            "pki_int/issue/example-dot-local",
						CommonName:      "example.local",
						TTL:             "24h",
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewConfig(tt.args.path)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewConfig() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				dmp := diffmatchpatch.New()
				t.Errorf("NewConfig() = %v", dmp.DiffPrettyText(dmp.DiffMain(fmt.Sprintf("%+v", got), fmt.Sprintf("%+v", tt.want), false)))
			}
		})
	}
}
