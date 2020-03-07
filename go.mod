module go-gin-dcard

go 1.13

require (
	github.com/gin-gonic/gin v1.5.0
	github.com/go-ini/ini v1.52.0
	github.com/go-redis/redis/v7 v7.2.0
	github.com/smartystreets/goconvey v1.6.4 // indirect
	gopkg.in/ini.v1 v1.52.0 // indirect
)

replace (
	go-gin-dcard/conf => /Users/jensonsu/MyProject/go-gin-dcard/conf
	go-gin-dcard/database => /Users/jensonsu/MyProject/go-gin-dcard/database
	go-gin-dcard/middleware => /Users/jensonsu/MyProject/go-gin-dcard/middleware
	go-gin-dcard/models => /Users/jensonsu/MyProject/go-gin-dcard/models
	go-gin-dcard/pkg => /Users/jensonsu/MyProject/go-gin-dcard/pkg
	go-gin-dcard/router => /Users/jensonsu/MyProject/go-gin-dcard/router
)
