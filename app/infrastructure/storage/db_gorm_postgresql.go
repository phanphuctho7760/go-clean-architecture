package storage

import (
	"fmt"
	"log"
	"sync"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	gormLogger "gorm.io/gorm/logger"

	"github.com/phanphuctho7760/go-clean-architecture/app/infrastructure/pkgs/colorconst"
	"github.com/phanphuctho7760/go-clean-architecture/app/infrastructure/pkgs/configs"
	"github.com/phanphuctho7760/go-clean-architecture/app/utils/helpers"
)

type DBGormPostgreSQL struct {
	DB *gorm.DB
}

func NewDBGormPostgreSQL() *DBGormPostgreSQL {
	return &DBGormPostgreSQL{}
}

func (receiver *DBGormPostgreSQL) Connect() (err error) {
	receiver.DB, err = receiver.connectPostgreSQL()
	if err != nil {
		log.Printf("%sFail to disconnect gorm because receiver.DB is nil at %s%s\n", colorconst.ANSIColorRed, helpers.GetCallerLocationSkip(1), colorconst.ANSIColorWhite)
		return
	}
	return
}

func (receiver *DBGormPostgreSQL) Disconnect() (err error) {
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

func (receiver *DBGormPostgreSQL) connectPostgreSQL() (db *gorm.DB, err error) {
	var once sync.Once
	host := configs.ConfigEnvGlobalInstance.GetDBPostgreSQLHost()
	port := configs.ConfigEnvGlobalInstance.GetDBPostgreSQLPort()
	username := configs.ConfigEnvGlobalInstance.GetDBPostgreSQLUserName()
	password := configs.ConfigEnvGlobalInstance.GetDBPostgreSQLPassword()
	name := configs.ConfigEnvGlobalInstance.GetDBPostgreSQLName()
	//charset := configs.ConfigEnvGlobalInstance.GetDBPostgreSQLCharset()
	//collation := configs.ConfigEnvGlobalInstance.GetDBPostgreSQLCollation()
	//parseTime := configs.ConfigEnvGlobalInstance.GetDBPostgreSQLParseTime()
	timezone := configs.ConfigEnvGlobalInstance.GetDBPostgreSQLTimezone()
	//timeLocation, err := time.LoadLocation(timezone)
	maxIdleConn := configs.ConfigEnvGlobalInstance.GetDBPostgreSQLMaxIdleConn()
	maxOpenConn := configs.ConfigEnvGlobalInstance.GetDBPostgreSQLMaxOpenConn()
	connMaxLifetime := configs.ConfigEnvGlobalInstance.GetDBPostgreSQLConnMaxLifetime()

	gormLogConfig := gorm.Config{
		Logger: func() gormLogger.Interface {
			return gormLogger.Default.LogMode(gormLogger.Silent)
		}(),
	}

	dsn := fmt.Sprintf("host=%s user=%s password=%s port=%d dbname=%s  sslmode=disable TimeZone=%s", host, username, password, port, name, timezone)
	//	dsn = "host=db_postgre user=postgres password=123456789 port=5432 dbname=go-clean-architecture  sslmode=disable TimeZone=Asia/Ho_Chi_Minh"
	fmt.Println(dsn)
	once.Do(func() {
		db, err = gorm.Open(postgres.Open(dsn), &gormLogConfig)
	})

	if err != nil {
		log.Printf("%sFail to open connect PostgreSQL gorm. Error: %s at %s%s\n", colorconst.ANSIColorRed, err.Error(), helpers.GetCallerLocationSkip(1), colorconst.ANSIColorWhite)
	}

	if err == nil {
		sqlDB, _ := db.DB()
		sqlDB.SetMaxIdleConns(maxIdleConn)
		sqlDB.SetMaxOpenConns(maxOpenConn)
		sqlDB.SetConnMaxLifetime(connMaxLifetime)
	}

	return
}
