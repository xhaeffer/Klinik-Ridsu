package routes

import (
	"encoding/base64"
	"net/http"
	"strconv"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"projek1/databases"
	"projek1/jwt" 
)

func Jadwal (r *gin.Engine, db *gorm.DB) {
	r.GET("/jadwal", func(c *gin.Context) {
		var data []databases.ProfilDokter
		if err := db.Preload("JadwalDokter").Find(&data).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal mengambil data dari database"})
			return
		}
		for i := range data {
			encodedImage := base64.StdEncoding.EncodeToString(data[i].Gambar)
			data[i].EncodedGambar = encodedImage
		}
		c.HTML(http.StatusOK, "jadwal.html", gin.H{"data": data})
	})

	r.GET("/jadwal/api/byID/:id", func(c *gin.Context) {
		start := c.Param("id")
	
		if err := jwt.VerifyToken(c); err != nil {
			if err == jwt.ErrMissingToken {
				c.JSON(http.StatusUnauthorized, gin.H{"error": "Token tidak ditemukan"})
			} else {
				c.JSON(http.StatusUnauthorized, gin.H{"error": "Token tidak valid"})
			}
			return
		}
	
		var data []databases.ProfilDokter
		startID, _ := strconv.Atoi(start)

		if err := db.Preload("JadwalDokter").Where("id_dokter = ?", startID).Find(&data).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal mengambil data dari database"})
			return
		}
	
		if len(data) == 0 {
			c.JSON(http.StatusNotFound, gin.H{"error": "Data tidak ditemukan"})
			return
		}
	
		c.JSON(http.StatusOK, data)
	})

	// r.GET("/jadwal/api/byIDRange/:start/:end", func(c *gin.Context) {
	// 	start := c.Param("start")
	// 	end := c.Param("end")
	
	// 	if err := jwt.VerifyToken(c); err != nil {
	// 		if err == jwt.ErrMissingToken {
	// 			c.JSON(http.StatusUnauthorized, gin.H{"error": "Token tidak ditemukan"})
	// 		} else {
	// 			c.JSON(http.StatusUnauthorized, gin.H{"error": "Token tidak valid"})
	// 		}
	// 		return
	// 	}
	
	// 	var data []databases.ProfilDokter
	
	// 	// Periksa apakah parameter "end" diisi atau tidak
	// 	startID, _ := strconv.Atoi(start)
	// 	endID, _ := strconv.Atoi(end)
	
	// 	if err := db.Preload("JadwalDokter").Where("id_dokter BETWEEN ? AND ?", startID, endID).Find(&data).Error; err != nil {
	// 		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal mengambil data dari database"})
	// 		return
	// 	}
	
	// 	if len(data) == 0 {
	// 		c.JSON(http.StatusNotFound, gin.H{"error": "Data tidak ditemukan"})
	// 		return
	// 	}
	
	// 	c.JSON(http.StatusOK, data)
	// })
	
	

	r.GET("/jadwal/api/byPoli/:poli", func(c *gin.Context) {
		poli := c.Param("poli")

		if err := jwt.VerifyToken(c); err != nil {
			if err == jwt.ErrMissingToken {
				c.JSON(http.StatusUnauthorized, gin.H{"error": "Token tidak ditemukan"})
			} else {
				c.JSON(http.StatusUnauthorized, gin.H{"error": "Token tidak valid"})
			}
			return
		}
		
		var data []databases.ProfilDokter
		if err := db.Preload("JadwalDokter").Where("poli = ?", poli).Find(&data).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal mengambil data dari database"})
			return
		}
		
		c.JSON(http.StatusOK, data)
		
	})

	r.GET("/jadwal/api/getPoli", func(c *gin.Context) {

		if err := jwt.VerifyToken(c); err != nil {
			if err == jwt.ErrMissingToken {
				c.JSON(http.StatusUnauthorized, gin.H{"error": "Token tidak ditemukan"})
			} else {
				c.JSON(http.StatusUnauthorized, gin.H{"error": "Token tidak valid"})
			}
			return
		}

		var poliList []string
	
		if err := db.Model(&databases.ProfilDokter{}).Distinct("poli").Pluck("poli", &poliList).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal mengambil data poli dari database"})
			return
		}
	
		c.JSON(http.StatusOK, poliList)
	})
}