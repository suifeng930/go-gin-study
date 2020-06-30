module github.com/mahongcheng/go-gin-study

go 1.14

require (
	github.com/gin-gonic/gin v1.6.3 // indirect
	github.com/go-ini/ini v1.57.0
	github.com/go-playground/validator/v10 v10.3.0 // indirect
	github.com/golang/protobuf v1.4.2 // indirect
	github.com/jinzhu/gorm v1.9.14 // indirect
	github.com/json-iterator/go v1.1.10 // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.1 // indirect
	github.com/unknwon/com v1.0.1 // indirect
	golang.org/x/sys v0.0.0-20200625212154-ddb9806d33ae // indirect
	google.golang.org/protobuf v1.25.0 // indirect
	gopkg.in/yaml.v2 v2.3.0 // indirect
)

replace (
	github.com/mahongcheng/go-gin-study/conf => ./go-gin-study/conf
	github.com/mahongcheng/go-gin-study/middleware => ./go-gin-study/middleware
	github.com/mahongcheng/go-gin-study/models => ./go-gin-study/models
	github.com/mahongcheng/go-gin-study/pkg/e => ./go-gin-study/pkg/e
	github.com/mahongcheng/go-gin-study/pkg/setting => ./go-gin-study/pkg/setting
	github.com/mahongcheng/go-gin-study/pkg/util => ./go-gin-study/pkg/util
	github.com/mahongcheng/go-gin-study/routers => ./go-gin-study/routers
)
