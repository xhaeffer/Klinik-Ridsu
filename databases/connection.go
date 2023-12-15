package databases

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func InitDatabase() *gorm.DB {
	// dsn := "xhaeffer:hahalol123@tcp(xhaeffer.me:11095)/klinik"
	dsn := "xhaeffer:hahalol123@tcp(localhost:3306)/klinik"
	var err error
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Gagal terhubung ke database")
	}

	db.Table(new(Reservasi).TableName()).AutoMigrate(&Reservasi{})
	db.Table(new(ProfilDokter).TableName()).AutoMigrate(&ProfilDokter{})
	db.Table(new(JadwalDokter).TableName()).AutoMigrate(&JadwalDokter{})
	db.Table(new(User).TableName()).AutoMigrate(&User{})

	return db
}
