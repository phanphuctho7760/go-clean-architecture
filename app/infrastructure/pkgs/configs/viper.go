package configs

import (
	"log"
	"time"

	_ "github.com/joho/godotenv/autoload"
	"github.com/spf13/viper"

	"github.com/phanphuctho7760/go-clean-architecture/app/infrastructure/pkgs/colorconst"
	"github.com/phanphuctho7760/go-clean-architecture/app/utils/helpers"
)

const configFileName = ".env"

type viperConfigLoaderEnv struct {
	configEnv configEnv
}

func newViperConfigLoaderEnv(configPath string) ConfigEnvItf {
	var v viperConfigLoaderEnv
	viper.AddConfigPath(configPath)
	viper.SetConfigFile(configFileName)
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			// Config file was found but another error was produced
			log.Printf("%sError to read config. Error: %s at %s%s\n", colorconst.ANSIColorRed, err.Error(), helpers.GetCallerLocationSkip(1), colorconst.ANSIColorWhite)
			return nil
		}
		// Config file not found; ignore error if desired
		log.Printf("%sNot found config file. Error: %s at %s%s\n", colorconst.ANSIColorRed, err.Error(), helpers.GetCallerLocationSkip(1), colorconst.ANSIColorWhite)
		return nil
	}

	err := viper.Unmarshal(&v.configEnv)
	if err != nil {
		log.Printf("%sFailed to unmarshal config. Error: %s at %s%s\n", colorconst.ANSIColorRed, err.Error(), helpers.GetCallerLocationSkip(1), colorconst.ANSIColorWhite)
		return nil
	}
	return &v
}

func (receiver viperConfigLoaderEnv) GetServiceName() string {
	return receiver.configEnv.ServiceName
}

func (receiver viperConfigLoaderEnv) GetMode() string {
	return receiver.configEnv.Mode
}

func (receiver viperConfigLoaderEnv) GetPort() string {
	return receiver.configEnv.Port
}

func (receiver viperConfigLoaderEnv) GetUTCTimezone() string {
	return receiver.configEnv.UtcTimezone
}

func (receiver viperConfigLoaderEnv) GetTimezone() string {
	return receiver.configEnv.Timezone
}

func (receiver viperConfigLoaderEnv) GetGinMode() string {
	return receiver.configEnv.GinMode
}

func (receiver viperConfigLoaderEnv) GetSentryDNS() string {
	return receiver.configEnv.SentryDNS
}

func (receiver viperConfigLoaderEnv) GetSentryEnableTracing() bool {
	return receiver.configEnv.SentryEnableTracing
}

func (receiver viperConfigLoaderEnv) GetSentryTracesSampleRate() float64 {
	return receiver.configEnv.SentryTracesSampleRate
}

func (receiver viperConfigLoaderEnv) GetSentryDebug() bool {
	return receiver.configEnv.SentryDebug
}

func (receiver viperConfigLoaderEnv) GetOtelExporterOTLPEndpoint() string {
	return receiver.configEnv.OtelExporterOTLPEndpoint
}

func (receiver viperConfigLoaderEnv) GetLogLevelZap() string {
	return receiver.configEnv.LogLevelZap
}

// mysql implement

func (receiver viperConfigLoaderEnv) GetDBMySQLHost() string {
	return receiver.configEnv.DBMySQLHost
}

func (receiver viperConfigLoaderEnv) GetDBMySQLPort() int {
	return receiver.configEnv.DBMySQLPort
}

func (receiver viperConfigLoaderEnv) GetDBMySQLUserName() string {
	return receiver.configEnv.DBMySQLUserName
}

func (receiver viperConfigLoaderEnv) GetDBMySQLPassword() string {
	return receiver.configEnv.DBMySQLPassword
}

func (receiver viperConfigLoaderEnv) GetDBMySQLName() string {
	return receiver.configEnv.DBMySQLName
}

func (receiver viperConfigLoaderEnv) GetDBMySQLCharset() string {
	return receiver.configEnv.DBMySQLCharset
}

func (receiver viperConfigLoaderEnv) GetDBMySQLCollation() string {
	return receiver.configEnv.DBMySQLCollation
}

func (receiver viperConfigLoaderEnv) GetDBMySQLParseTime() bool {
	return receiver.configEnv.DBMySQLParseTime
}

func (receiver viperConfigLoaderEnv) GetDBMySQLTimezone() string {
	return receiver.configEnv.DBMySQLTimezone
}

func (receiver viperConfigLoaderEnv) GetDBMySQLMaxIdleConn() int {
	return receiver.configEnv.DBMySQLMaxIdleConn
}
func (receiver viperConfigLoaderEnv) GetDBMySQLMaxOpenConn() int {
	return receiver.configEnv.DBMySQLMaxOpenConn
}
func (receiver viperConfigLoaderEnv) GetDBMySQLConnMaxLifetime() time.Duration {
	return receiver.configEnv.DBMySQLConnMaxLifetime
}

// postgresql implement

func (receiver viperConfigLoaderEnv) GetDBPostgreSQLHost() string {
	return receiver.configEnv.DBPostgreSQLHost
}

func (receiver viperConfigLoaderEnv) GetDBPostgreSQLPort() int {
	return receiver.configEnv.DBPostgreSQLPort
}

func (receiver viperConfigLoaderEnv) GetDBPostgreSQLUserName() string {
	return receiver.configEnv.DBPostgreSQLUserName
}

func (receiver viperConfigLoaderEnv) GetDBPostgreSQLPassword() string {
	return receiver.configEnv.DBPostgreSQLPassword
}

func (receiver viperConfigLoaderEnv) GetDBPostgreSQLName() string {
	return receiver.configEnv.DBPostgreSQLName
}

func (receiver viperConfigLoaderEnv) GetDBPostgreSQLCharset() string {
	return receiver.configEnv.DBPostgreSQLCharset
}

func (receiver viperConfigLoaderEnv) GetDBPostgreSQLCollation() string {
	return receiver.configEnv.DBPostgreSQLCollation
}

func (receiver viperConfigLoaderEnv) GetDBPostgreSQLParseTime() bool {
	return receiver.configEnv.DBPostgreSQLParseTime
}

func (receiver viperConfigLoaderEnv) GetDBPostgreSQLTimezone() string {
	return receiver.configEnv.DBPostgreSQLTimezone
}

func (receiver viperConfigLoaderEnv) GetDBPostgreSQLMaxIdleConn() int {
	return receiver.configEnv.DBPostgreSQLMaxIdleConn
}
func (receiver viperConfigLoaderEnv) GetDBPostgreSQLMaxOpenConn() int {
	return receiver.configEnv.DBPostgreSQLMaxOpenConn
}
func (receiver viperConfigLoaderEnv) GetDBPostgreSQLConnMaxLifetime() time.Duration {
	return receiver.configEnv.DBPostgreSQLConnMaxLifetime
}
