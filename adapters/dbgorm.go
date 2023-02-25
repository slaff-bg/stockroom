package adapters

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type dbGorm struct {
	INST *gorm.DB
}

var db dbGorm

type credentials struct {
	DB_HOST      string
	DB_DRIVER    string
	DB_USER      string
	DB_PASSWORD  string
	DB_NAME      string
	DB_PORT      string
	DB_SSL_MODE  string
	DB_TIME_ZONE string
}

var (
	cr    credentials
	dbmei string // db mod env info
)

func ConnectDB() {
	err := godotenv.Load(".env")
	if err != nil {
		panic(err)
	}

	// Depending on the setting of POSTGRES_LOCAL_INSTANCE in the
	// configuration file, it loads the appropriate credentials.
	if os.Getenv("POSTGRES_LOCAL_INSTANCE") == "true" {
		setEnvCredentials(true)
	} else {
		setEnvCredentials(false)
	}

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s",
		cr.DB_HOST,
		cr.DB_USER,
		cr.DB_PASSWORD,
		cr.DB_NAME,
		cr.DB_PORT,
		cr.DB_SSL_MODE,
		cr.DB_TIME_ZONE,
	)
	dbo, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatal("failed to connect to database. \n", err)
		os.Exit(2)
	}

	log.Printf("\nDB connected successfully and ran in a %s environment ...\n", dbmei)
	// dbo.Logger = logger.Default.LogMode(logger.Info)

	// log.Println("running migrations")
	// dbo.AutoMigrate(&models.Fact{})

	db = dbGorm{
		INST: dbo,
	}
}

// setEnvCredentials sets the credentials depending on a configuration where the
// argument (l) means Local.
func setEnvCredentials(l bool) {
	if l {
		cr.DB_HOST = os.Getenv("LOCAL_DB_HOST")
		cr.DB_DRIVER = os.Getenv("LOCAL_DB_DRIVER")
		cr.DB_USER = os.Getenv("LOCAL_DB_USER")
		cr.DB_PASSWORD = os.Getenv("LOCAL_DB_PASSWORD")
		cr.DB_NAME = os.Getenv("LOCAL_DB_NAME")
		cr.DB_PORT = os.Getenv("LOCAL_DB_PORT")
		cr.DB_SSL_MODE = os.Getenv("LOCAL_DB_SSL_MODE")
		cr.DB_TIME_ZONE = os.Getenv("LOCAL_DB_TIME_ZONE")
		dbmei = "local"
		return
	}
	cr.DB_HOST = os.Getenv("DB_HOST")
	cr.DB_DRIVER = os.Getenv("DB_DRIVER")
	cr.DB_USER = os.Getenv("DB_USER")
	cr.DB_PASSWORD = os.Getenv("DB_PASSWORD")
	cr.DB_NAME = os.Getenv("DB_NAME")
	cr.DB_PORT = os.Getenv("DB_PORT")
	cr.DB_SSL_MODE = os.Getenv("DB_SSL_MODE")
	cr.DB_TIME_ZONE = os.Getenv("DB_TIME_ZONE")
	dbmei = "docker"
}
