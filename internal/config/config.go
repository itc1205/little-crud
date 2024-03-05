package config

type Config struct {
	Port int `env:"SERVER_PORT" envDefault:"12005"`
	PostgresConfig
	NatsConfig
	RedisConfig
}

type PostgresConfig struct {
	PgPort   string `env:"PG_PORT" envDefault:"5458"`
	PgHost   string `env:"PG_HOST" envDefault:"0.0.0.0"`
	PgDBName string `env:"PG_DB_NAME" envDefault:"db"`
	PgUser   string `env:"PG_USER" envDefault:"db"`
	PgPwd    string `env:"PG_PWD" envDefault:"db"`
}

type ClickHouseConfig struct {
	ChPort   string `env:"CH_PORT" envDefault:"8123"`
	ChHost   string `env:"CH_HOST" envDefault:"0.0.0.0"`
	ChDBName string `env:"CH_DB_NAME" envDefault:"db"`
	ChUser   string `env:"CH_USER" envDefault:"db"`
	ChPwd    string `env:"CH_PWD" envDefault:"db"`
}

type NatsConfig struct {
	NPort string `env:"N_PORT" envDefault:"4222"`
	NHost string `env:"N_HOST" envDefault:"nats"`
	NSubj string `env:"N_Subj" envDefault:"logs"`
}

type RedisConfig struct {
	RPort string `env:"R_PORT" envDefault:"6379"`
	RHost string `env:"R_HOST" envDefault:"0.0.0.0"`
}
