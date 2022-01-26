module SnakeGame/backend

// +heroku goVersion go1.17
go 1.17

require (
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/go-chi/chi v1.5.4
	github.com/go-chi/cors v1.2.0
	github.com/go-chi/jwtauth v1.2.0
	github.com/jackc/pgx/v4 v4.14.1
	golang.org/x/crypto v0.0.0-20220112180741-5e0467b6c7ce
)

require gopkg.in/yaml.v3 v3.0.0-20210107192922-496545a6307b // indirect

require (
	github.com/decred/dcrd/dcrec/secp256k1/v4 v4.0.1 // indirect
	github.com/goccy/go-json v0.9.3 // indirect
	github.com/lestrrat-go/jwx v1.2.17 // indirect
	github.com/stretchr/testify v1.7.0
)
