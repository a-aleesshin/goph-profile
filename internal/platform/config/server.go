package config

type Server struct {
	Address string `default:":9000" envconfig:"SERVER_ADDRESS" validate:"required"`
}
