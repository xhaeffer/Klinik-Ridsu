package jadwal

import (
	"encoding/base64"
	"net/http"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	
	"KlinikRidsu/databases"
	"KlinikRidsu/session" 
)

func Jadwal (r *gin.Engine, db *gorm.DB) {
	r.GET("/api/jadwal", func(c *gin.Context) {
		var data []databases.ProfilDokter
		if err := db.Preload("JadwalDokter").Find(&data).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal mengambil data!"})
			return
		}
		for i := range data {
			encodedImage := base64.StdEncoding.EncodeToString(data[i].Gambar)
			data[i].EncodedGambar = encodedImage
		}
		c.JSON(http.StatusOK, data)
	})

	r.GET("/api/jadwal/byID/:id", session.VerifyToken(), func(c *gin.Context) {
		param := c.Param("id")

		var data []databases.ProfilDokter

		if err := db.Preload("JadwalDokter").Where("id_dokter = ?", param).Find(&data).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal mengambil data!"})
			return
		}
	
		if len(data) == 0 {
			c.JSON(http.StatusNotFound, gin.H{"error": "Data tidak ditemukan!"})
			return
		}
	
		c.JSON(http.StatusOK, data)
	})

	r.GET("/api/jadwal/byPoli", session.VerifyToken(), func(c *gin.Context) {
		var poliList []string
	
		if err := db.Model(&databases.ProfilDokter{}).Distinct("poli").Pluck("poli", &poliList).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal mengambil data poli!"})
			return
		}
	
		c.JSON(http.StatusOK, poliList)
	})

	r.GET("/api/jadwal/byPoli/:poli", session.VerifyToken(), func(c *gin.Context) {
		param := c.Param("poli")
		
		var data []databases.ProfilDokter
		if err := db.Preload("JadwalDokter").Where("poli = ?", param).Find(&data).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal mengambil data!"})
			return
		}

		if len(data) == 0 {
			c.JSON(http.StatusNotFound, gin.H{"error": "Data tidak ditemukan!"})
			return
		}
		
		c.JSON(http.StatusOK, data)
	})


}