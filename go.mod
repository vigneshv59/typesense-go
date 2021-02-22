module github.com/vigneshv59/typesense-go

go 1.15

require (
	github.com/deepmap/oapi-codegen v1.5.1
	github.com/go-chi/chi v4.0.2+incompatible // indirect
	github.com/golang/mock v1.4.4
	github.com/google/uuid v1.2.0
	github.com/jinzhu/copier v0.1.0
	github.com/kr/text v0.2.0 // indirect
	github.com/sony/gobreaker v0.4.1
	github.com/stretchr/testify v1.7.0
	github.com/testcontainers/testcontainers-go v0.9.0
	gopkg.in/check.v1 v1.0.0-20201130134442-10cb98267c6c // indirect
	gopkg.in/yaml.v3 v3.0.0-20200615113413-eeeca48fe776 // indirect
)

replace github.com/typesense/typesense-go/typesense => github.com/vigneshv59/typesense-go/typesense v0.2.1

replace github.com/typesense/typesense-go/typesense/api => github.com/vigneshv59/typesense-go/typesense/api v0.2.1
