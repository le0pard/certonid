module github.com/certonid/certonid/cli

go 1.12

require (
	github.com/ScaleFT/sshkeys v0.0.0-20200327173127-6142f742bca5
	github.com/aws/aws-sdk-go-v2/config v1.15.14 // indirect
	github.com/aws/aws-sdk-go-v2/service/kms v1.18.0 // indirect
	github.com/aws/aws-sdk-go-v2/service/lambda v1.23.4 // indirect
	github.com/certonid/certonid v0.0.0-20220722152623-ecd67815e327
	github.com/go-playground/universal-translator v0.18.0 // indirect
	github.com/leodido/go-urn v1.2.1 // indirect
	github.com/mitchellh/go-homedir v1.1.0
	github.com/pelletier/go-toml/v2 v2.0.2 // indirect
	github.com/rs/zerolog v1.27.0
	github.com/sethvargo/go-password v0.2.0 // indirect
	github.com/spf13/afero v1.9.2 // indirect
	github.com/spf13/cobra v1.5.0
	github.com/spf13/viper v1.12.0
	github.com/subosito/gotenv v1.4.0 // indirect
	golang.org/x/crypto v0.0.0-20220722155217-630584e8d5aa
	golang.org/x/sys v0.0.0-20220728004956-3c1f35247d10 // indirect
	golang.org/x/term v0.0.0-20220722155259-a9ba230a4035 // indirect
	gopkg.in/go-playground/assert.v1 v1.2.1 // indirect
	gopkg.in/go-playground/validator.v9 v9.31.0 // indirect
	gopkg.in/ini.v1 v1.66.6 // indirect
)

replace (
	github.com/certonid/certonid/adapters => ./../adapters
	github.com/certonid/certonid/proto => ./../proto
	github.com/certonid/certonid/utils => ./../utils
)
