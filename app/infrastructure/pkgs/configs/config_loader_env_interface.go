package configs

import "time"

type ConfigEnvItf interface {
	ConfigEnvCommonItf
	ConfigEnvCustomItf
	ConfigMySQLEnv
	ConfigPostgreSQLEnv
}

type ConfigEnvCommonItf interface {
	GetServiceName() string
	GetMode() string
	GetPort() string
	GetUTCTimezone() string
	GetTimezone() string
}

type ConfigEnvCustomItf interface {
	GetGinMode() string
	GetSentryDNS() string
	GetSentryEnableTracing() bool
	GetSentryTracesSampleRate() float64
	GetSentryDebug() bool
	GetOtelExporterOTLPEndpoint() string
	GetLogLevelZap() string
}

type ConfigMySQLEnv interface {
	GetDBMySQLHost() string
	GetDBMySQLPort() int
	GetDBMySQLUserName() string
	GetDBMySQLPassword() string
	GetDBMySQLName() string
	GetDBMySQLCharset() string
	GetDBMySQLCollation() string
	GetDBMySQLParseTime() bool
	GetDBMySQLTimezone() string
	GetDBMySQLMaxIdleConn() int
	GetDBMySQLMaxOpenConn() int
	GetDBMySQLConnMaxLifetime() time.Duration
}

type ConfigPostgreSQLEnv interface {
	GetDBPostgreSQLHost() string
	GetDBPostgreSQLPort() int
	GetDBPostgreSQLUserName() string
	GetDBPostgreSQLPassword() string
	GetDBPostgreSQLName() string
	GetDBPostgreSQLCharset() string
	GetDBPostgreSQLCollation() string
	GetDBPostgreSQLParseTime() bool
	GetDBPostgreSQLTimezone() string
	GetDBPostgreSQLMaxIdleConn() int
	GetDBPostgreSQLMaxOpenConn() int
	GetDBPostgreSQLConnMaxLifetime() time.Duration
}
