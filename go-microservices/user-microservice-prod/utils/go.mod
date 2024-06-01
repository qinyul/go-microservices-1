module barqi.com/user/utils

go 1.22.3

replace barqi.com/user/common => ../common

require (
	barqi.com/user/common v0.0.0-00010101000000-000000000000
	barqi.com/user/docs v0.0.0-00010101000000-000000000000
	github.com/golang-jwt/jwt/v5 v5.2.1
	github.com/sirupsen/logrus v1.9.3
	golang.org/x/crypto v0.23.0
)

require (
	github.com/KyleBanks/depth v1.2.1 // indirect
	github.com/PuerkitoBio/purell v1.1.1 // indirect
	github.com/PuerkitoBio/urlesc v0.0.0-20170810143723-de5bf2ad4578 // indirect
	github.com/go-openapi/jsonpointer v0.19.5 // indirect
	github.com/go-openapi/jsonreference v0.19.6 // indirect
	github.com/go-openapi/spec v0.20.4 // indirect
	github.com/go-openapi/swag v0.19.15 // indirect
	github.com/josharian/intern v1.0.0 // indirect
	github.com/mailru/easyjson v0.7.6 // indirect
	github.com/natefinch/lumberjack v2.0.0+incompatible // indirect
	github.com/swaggo/swag v1.16.3 // indirect
	golang.org/x/net v0.21.0 // indirect
	golang.org/x/sys v0.20.0 // indirect
	golang.org/x/text v0.15.0 // indirect
	golang.org/x/tools v0.7.0 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
)

replace barqi.com/user/docs => ../docs
