package routes

import (
	"encoding/base64"
	"net/http"
	"strconv"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"projek1/databases"
	"projek1/api" 
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

	// r.POST("/jadwal", func(c *gin.Context) {
	// 	c.HTML(http.StatusOK, "jadwal2.html", nil)
	// })

	r.GET("/jadwal/:start/:end", func(c *gin.Context) {
		start := c.Param("start")
		end := c.Param("end")

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
		endID, _ := strconv.Atoi(end)

		for i := startID; i <= endID; i++ {
			var item databases.ProfilDokter
			if err := db.First(&item, i).Error; err == nil {
				data = append(data, item)
			}
		}

		if len(data) == 0 {
			c.JSON(http.StatusNotFound, gin.H{"error": "Data tidak ditemukan"})
			return
		}

		c.JSON(http.StatusOK, data)
	})
}