module github.com/takumakume/kubectl-vault-client-certificate

go 1.15

require (
	github.com/BurntSushi/toml v0.3.1
	github.com/caarlos0/env v3.5.0+incompatible
	github.com/caarlos0/env/v6 v6.5.0
	github.com/goccy/go-yaml v1.8.8
	github.com/google/gofuzz v1.2.0 // indirect
	github.com/hashicorp/vault/api v1.0.4
	github.com/hashicorp/vault/sdk v0.1.13
	github.com/ilyakaznacheev/cleanenv v1.2.5
	github.com/imdario/mergo v0.3.11 // indirect
	github.com/jinzhu/configor v1.2.1
	github.com/olekukonko/tablewriter v0.0.5
	github.com/sergi/go-diff v1.1.0
	github.com/sirupsen/logrus v1.2.0
	github.com/spf13/cobra v1.1.3
	gopkg.in/yaml.v2 v2.4.0
	k8s.io/api v0.20.2 // indirect
	k8s.io/client-go v0.0.0-20210216172702-39da00799391
	k8s.io/klog v1.0.0 // indirect
	k8s.io/utils v0.0.0-20210111153108-fddb29f9d009 // indirect
)
