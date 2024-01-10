package configs

import "time"

type configEnv struct {
	ServiceName                 string        `mapstructure:"SERVICE_NAME"`
	Mode                        string        `mapstructure:"MODE"`
	Port                        string        `mapstructure:"PORT"`
	UtcTimezone                 string        `mapstructure:"UTC_TIMEZONE"`
	Timezone                    string        `mapstructure:"TIMEZONE"`
	DBMySQLHost                 string        `mapstructure:"DATABASE_MYSQL_HOST"`
	DBMySQLPort                 int           `mapstructure:"DATABASE_MYSQL_PORT"`
	DBMySQLUserName             string        `mapstructure:"DATABASE_MYSQL_USER_NAME"`
	DBMySQLPassword             string        `mapstructure:"DATABASE_MYSQL_PASSWORD"`
	DBMySQLName                 string        `mapstructure:"DATABASE_MYSQL_NAME"`
	DBMySQLCharset              string        `mapstructure:"DATABASE_MYSQL_CHARSET"`
	DBMySQLCollation            string        `mapstructure:"DATABASE_MYSQL_COLLATION"`
	DBMySQLParseTime            bool          `mapstructure:"DATABASE_MYSQL_PARSE_TIME"`
	DBMySQLTimezone             string        `mapstructure:"DATABASE_MYSQL_TIMEZONE"`
	DBMySQLMaxIdleConn          int           `mapstructure:"DATABASE_MYSQL_MAX_IDLE_CONN"`
	DBMySQLMaxOpenConn          int           `mapstructure:"DATABASE_MYSQL_MAX_OPEN_CONN"`
	DBMySQLConnMaxLifetime      time.Duration `mapstructure:"DATABASE_MYSQL_CONN_MAX_LIFETIME"`
	DBPostgreSQLHost            string        `mapstructure:"DATABASE_POSTGRESQL_HOST"`
	DBPostgreSQLPort            int           `mapstructure:"DATABASE_POSTGRESQL_PORT"`
	DBPostgreSQLUserName        string        `mapstructure:"DATABASE_POSTGRESQL_USER_NAME"`
	DBPostgreSQLPassword        string        `mapstructure:"DATABASE_POSTGRESQL_PASSWORD"`
	DBPostgreSQLName            string        `mapstructure:"DATABASE_POSTGRESQL_NAME"`
	DBPostgreSQLCharset         string        `mapstructure:"DATABASE_POSTGRESQL_CHARSET"`
	DBPostgreSQLCollation       string        `mapstructure:"DATABASE_POSTGRESQL_COLLATION"`
	DBPostgreSQLParseTime       bool          `mapstructure:"DATABASE_POSTGRESQL_PARSE_TIME"`
	DBPostgreSQLTimezone        string        `mapstructure:"DATABASE_POSTGRESQL_TIMEZONE"`
	DBPostgreSQLMaxIdleConn     int           `mapstructure:"DATABASE_POSTGRESQL_MAX_IDLE_CONN"`
	DBPostgreSQLMaxOpenConn     int           `mapstructure:"DATABASE_POSTGRESQL_MAX_OPEN_CONN"`
	DBPostgreSQLConnMaxLifetime time.Duration `mapstructure:"DATABASE_POSTGRESQL_CONN_MAX_LIFETIME"`
	GinMode                     string        `mapstructure:"GIN_MODE"`
	SentryDNS                   string        `mapstructure:"SENTRY_DSN"`
	SentryEnableTracing         bool          `mapstructure:"SENTRY_ENABLE_TRACING"`
	SentryTracesSampleRate      float64       `mapstructure:"SENTRY_TRACES_SAMPLE_RATE"`
	SentryDebug                 bool          `mapstructure:"SENTRY_DEBUG"`
	OtelExporterOTLPEndpoint    string        `mapstructure:"OTEL_EXPORTER_OTLP_ENDPOINT"`
	LogLevelZap                 string        `mapstructure:"LOG_LEVEL_ZAP`
}
