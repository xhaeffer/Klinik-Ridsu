package routes

import (
	"net/http"
	"strconv"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"projek1/databases"
	"projek1/jwt"
)

func Reservasi (r *gin.Engine, db *gorm.DB) {
	r.GET("/reservasi", func(c *gin.Context) {
		if err := jwt.VerifyToken(c); err != nil {
			if err == jwt.ErrMissingToken {
				c.JSON(http.StatusUnauthorized, gin.H{"error": "Token tidak ditemukan, Silahkan Login terlebih dahulu!"})
				return
			} else {
				c.JSON(http.StatusUnauthorized, gin.H{"error": "Token tidak valid, Silahkan login ulang!"})
				return
			}
		}
		c.HTML(http.StatusOK, "reservasi.html", nil)
	})

	r.POST("/reservasi", func(c *gin.Context) {
		var data databases.Reservasi

		if err := jwt.VerifyToken(c); err != nil {
			if err == jwt.ErrMissingToken {
				c.JSON(http.StatusUnauthorized, gin.H{"error": "Token tidak ditemukan"})
			} else {
				c.JSON(http.StatusUnauthorized, gin.H{"error": "Token tidak valid"})
			}
			return
		}

		// Bind data dari formulir HTML ke struct
		if err := c.Bind(&data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Simpan data ke dalam database
		if err := db.Create(&data).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal menyimpan data ke database"})
			return
		} else {
			c.HTML(http.StatusOK, "reservasi.html", gin.H{"message": "Reservasi berhasil disimpan", "data": data})
		}
	})

	r.GET("/reservasi/api/byID/:start/:end", func(c *gin.Context) {
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

		var data []databases.Reservasi

		startID, _ := strconv.Atoi(start)
		endID, _ := strconv.Atoi(end)

		for i := startID; i <= endID; i++ {
			var item databases.Reservasi
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