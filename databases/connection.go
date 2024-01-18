package databases

import (
	"fmt"
	"log"
	"KlinikRidsu/configs"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func InitDatabase(dbName string) *gorm.DB {
	configs.DbConfig()

	dbConfigs := configs.ReadDatabaseConfigs()
	config, err := configs.FindDatabaseConfig(dbConfigs, dbName)
	if err != nil {
		log.Fatalf("Database configuration not found: %v", err)
	}

	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s",
		config.Username,
		config.Password,
		config.Host,
		config.Port,
		config.DbName,
	)

	var errConnect error
	db, errConnect = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if errConnect != nil {
		log.Fatalf("Failed to connect to the database: %v", errConnect)
	}

	db.Table(new(Reservasi).TableName()).AutoMigrate(&Reservasi{})
	db.Table(new(ProfilDokter).TableName()).AutoMigrate(&ProfilDokter{})
	db.Table(new(JadwalDokter).TableName()).AutoMigrate(&JadwalDokter{})
	db.Table(new(User).TableName()).AutoMigrate(&User{})

	return db
}

