package config

import "os"

type Config struct {
	GRPCPort     string
	AMQPHost     string
	AMQPPort     string
	AMQPLogin    string
	AMQPPassword string
	AMQPVhost    string
	DBHost       string
	DBPort       string
	DBName       string
	DBUser       string
	DBPassword   string
}

func New() *Config {
	return &Config{
		GRPCPort:     os.Getenv("GRPC_PORT"),
		AMQPHost:     os.Getenv("AMQP_HOST"),
		AMQPPort:     os.Getenv("AMQP_PORT"),
		AMQPLogin:    os.Getenv("AMQP_LOGIN"),
		AMQPPassword: os.Getenv("AMQP_PASSWORD"),
		AMQPVhost:    os.Getenv("AMQP_VHOST"),
		DBHost:       os.Getenv("DB_HOST"),
		DBPort:       os.Getenv("DB_PORT"),
		DBName:       os.Getenv("DB_NAME"),
		DBUser:       os.Getenv("DB_USER"),
		DBPassword:   os.Getenv("DB_PASSWORD"),
	}
}
