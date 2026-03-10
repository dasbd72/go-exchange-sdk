module github.com/dasbd72/go-exchange-sdk/okx

go 1.22.0

replace github.com/dasbd72/go-exchange-sdk/config => ../config

require (
	github.com/dasbd72/go-exchange-sdk/config v0.0.0
	github.com/google/go-cmp v0.7.0
)

require github.com/joho/godotenv v1.5.1 // indirect
