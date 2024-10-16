package configs

import "github.com/spf13/viper"

// AMQPPORT=5672
// AMQPHOST=localhost
// AMQPUSER=guest
// AMQPPASS=guest
// AMQPQUEUE=job_queue

type Config struct {
	DBDriver   string `mapstructure:"DB_DRIVER"`
	DBHost     string `mapstructure:"DB_HOST"`
	DBPort     string `mapstructure:"DB_PORT"`
	DBUser     string `mapstructure:"DB_USER"`
	DBPassword string `mapstructure:"DB_PASSWORD"`
	DBName     string `mapstructure:"DB_NAME"`

	WebServerPort string `mapstructure:"WEB_SERVER_PORT"`

	AmqpPort     string `mapstructure:"AMQP_PORT"`
	AmqpHost     string `mapstructure:"AMQP_HOST"`
	AmqpUser     string `mapstructure:"AMQP_USER"`
	AmqpPassword string `mapstructure:"AMQP_PASSWORD"`
	AmqpQueue    string `mapstructure:"AMQP_QUEUE"`

	GRPCServerPort    string `mapstructure:"GRPC_SERVER_PORT"`
	GraphQLServerPort string `mapstructure:"GRAPHQL_SERVER_PORT"`
}

func LoadConfig(path string) (*Config, error) {
	var cfg *Config
	viper.SetConfigName("app_config")
	viper.SetConfigType("env")
	viper.AddConfigPath(path)
	viper.SetConfigFile(".env")
	viper.AutomaticEnv()
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
	err = viper.Unmarshal(&cfg)
	if err != nil {
		panic(err)
	}
	return cfg, err
}
