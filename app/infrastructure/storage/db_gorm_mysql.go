package storage

import (
	"log"
	"strconv"
	"sync"
	"time"

	"github.com/go-sql-driver/mysql"
	mysqlGorm "gorm.io/driver/mysql"
	"gorm.io/gorm"
	gormLogger "gorm.io/gorm/logger"

	"github.com/phanphuctho7760/go-clean-architecture/app/infrastructure/pkgs/colorconst"
	"github.com/phanphuctho7760/go-clean-architecture/app/infrastructure/pkgs/configs"
	"github.com/phanphuctho7760/go-clean-architecture/app/utils/helpers"
)

type DBGormMySQL struct {
	DB *gorm.DB
}

func NewDBGormMySql() *DBGormMySQL {
	return &DBGormMySQL{}
}

func (receiver *DBGormMySQL) Connect() (err error) {
	receiver.DB, err = receiver.connectMySQL()
	if err != nil {
		log.Printf("%sFail to disconnect gorm because receiver.DB is nil at %s%s\n", colorconst.ANSIColorRed, helpers.GetCallerLocationSkip(1), colorconst.ANSIColorWhite)
		return
	}
	return
}

func (receiver *DBGormMySQL) Disconnect() (err error) {
	if receiver.DB == nil {
		log.Printf("%sFail to disconnect gorm because receiver.DB is nil at %s%s\n", colorconst.ANSIColorRed, helpers.GetCallerLocationSkip(1), colorconst.ANSIColorWhite)
		return
	}
	sqlDB, err := receiver.DB.DB()
	if err != nil {
		log.Printf("%sFail to get connection to prepare disconnect gorm error: %s at %s%s\n", colorconst.ANSIColorRed, err.Error(), helpers.GetCallerLocationSkip(1), colorconst.ANSIColorWhite)
		return
	}
	if sqlDB == nil {
		log.Printf("%sFail to disconnect gorm because sqlDB is nil at %s%s\n", colorconst.ANSIColorRed, helpers.GetCallerLocationSkip(1), colorconst.ANSIColorWhite)
		return
	}
	err = sqlDB.Close()
	if err != nil {
		log.Printf("%sFail to close connection gorm error: %s at %s%s\n", colorconst.ANSIColorRed, err.Error(), helpers.GetCallerLocationSkip(1), colorconst.ANSIColorWhite)
		return
	}
	receiver.DB = nil

	return
}

func (receiver *DBGormMySQL) connectMySQL() (db *gorm.DB, err error) {
	var once sync.Once
	host := configs.ConfigEnvGlobalInstance.GetDBMySQLHost()
	port := configs.ConfigEnvGlobalInstance.GetDBMySQLPort()
	username := configs.ConfigEnvGlobalInstance.GetDBMySQLUserName()
	password := configs.ConfigEnvGlobalInstance.GetDBMySQLPassword()
	name := configs.ConfigEnvGlobalInstance.GetDBMySQLName()
	// charset :=  configs.ConfigEnvGlobalInstance.GetDBMySQLCharset
	collation := configs.ConfigEnvGlobalInstance.GetDBMySQLCollation()
	parseTime := configs.ConfigEnvGlobalInstance.GetDBMySQLParseTime()
	timezone := configs.ConfigEnvGlobalInstance.GetDBMySQLTimezone()
	timeLocation, err := time.LoadLocation(timezone)
	maxIdleConn := configs.ConfigEnvGlobalInstance.GetDBMySQLMaxIdleConn()
	maxOpenConn := configs.ConfigEnvGlobalInstance.GetDBMySQLMaxOpenConn()
	connMaxLifetime := configs.ConfigEnvGlobalInstance.GetDBMySQLConnMaxLifetime()

	gormLogConfig := gorm.Config{
		Logger: func() gormLogger.Interface {
			return gormLogger.Default.LogMode(gormLogger.Silent)
		}(),
	}

	mysqlCfg := mysql.Config{
		User:                     username,
		Passwd:                   password,
		Net:                      "tcp",
		Addr:                     host + ":" + strconv.Itoa(port),
		DBName:                   name,
		Params:                   nil,
		Collation:                collation,
		Loc:                      timeLocation,
		MaxAllowedPacket:         0,
		ServerPubKey:             "",
		TLSConfig:                "",
		TLS:                      nil,
		Timeout:                  0,
		ReadTimeout:              0,
		WriteTimeout:             0,
		AllowAllFiles:            false,
		AllowCleartextPasswords:  false,
		AllowFallbackToPlaintext: false,
		AllowNativePasswords:     false,
		AllowOldPasswords:        false,
		CheckConnLiveness:        false,
		ClientFoundRows:          false,
		ColumnsWithAlias:         false,
		InterpolateParams:        false,
		MultiStatements:          false,
		ParseTime:                parseTime,
		RejectReadOnly:           false,
	}

	once.Do(func() {
		db, err = gorm.Open(mysqlGorm.Open(mysqlCfg.FormatDSN()), &gormLogConfig)
	})

	if err != nil {
		log.Fatalf("%sFail to open connect mysql gorm. Error: %s at %s%s\n", colorconst.ANSIColorRed, err.Error(), helpers.GetCallerLocationSkip(1), colorconst.ANSIColorWhite)
	}

	if err == nil {
		sqlDB, _ := db.DB()
		sqlDB.SetMaxIdleConns(maxIdleConn)
		sqlDB.SetMaxOpenConns(maxOpenConn)
		sqlDB.SetConnMaxLifetime(connMaxLifetime)
	}

	return
}
