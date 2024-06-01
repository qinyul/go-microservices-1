module barqi.com/user/auth

go 1.22.3

replace barqi.com/user/database => ../database

replace barqi.com/user/utils => ../utils

replace barqi.com/user/common => ../common

require (
	barqi.com/user/common v0.0.0-00010101000000-000000000000
	barqi.com/user/database v0.0.0-00010101000000-000000000000
	barqi.com/user/utils v0.0.0-00010101000000-000000000000
	github.com/sirupsen/logrus v1.9.3
	go.mongodb.org/mongo-driver v1.15.0
)

require (
	barqi.com/user/docs v0.0.0-00010101000000-000000000000 // indirect
	github.com/KyleBanks/depth v1.2.1 // indirect
	github.com/PuerkitoBio/purell v1.1.1 // indirect
	github.com/PuerkitoBio/urlesc v0.0.0-20170810143723-de5bf2ad4578 // indirect
	github.com/go-openapi/jsonpointer v0.19.5 // indirect
	github.com/go-openapi/jsonreference v0.19.6 // indirect
	github.com/go-openapi/spec v0.20.4 // indirect
	github.com/go-openapi/swag v0.19.15 // indirect
	github.com/golang-jwt/jwt/v5 v5.2.1 // indirect
	github.com/golang/snappy v0.0.1 // indirect
	github.com/josharian/intern v1.0.0 // indirect
	github.com/klauspost/compress v1.13.6 // indirect
	github.com/mailru/easyjson v0.7.6 // indirect
	github.com/montanaflynn/stats v0.0.0-20171201202039-1bf9dbcd8cbe // indirect
	github.com/natefinch/lumberjack v2.0.0+incompatible // indirect
	github.com/swaggo/swag v1.16.3 // indirect
	github.com/xdg-go/pbkdf2 v1.0.0 // indirect
	github.com/xdg-go/scram v1.1.2 // indirect
	github.com/xdg-go/stringprep v1.0.4 // indirect
	github.com/youmark/pkcs8 v0.0.0-20181117223130-1be2e3e5546d // indirect
	golang.org/x/crypto v0.23.0 // indirect
	golang.org/x/net v0.21.0 // indirect
	golang.org/x/sync v0.1.0 // indirect
	golang.org/x/sys v0.20.0 // indirect
	golang.org/x/text v0.15.0 // indirect
	golang.org/x/tools v0.7.0 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
)

replace barqi.com/user/docs => ../docs
