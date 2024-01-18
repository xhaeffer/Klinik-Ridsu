package reservasi

import (
	"KlinikRidsu/databases"
	"KlinikRidsu/session"
	"net/http"

	"github.com/dpapathanasiou/go-recaptcha"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Reservasi (r *gin.Engine, db *gorm.DB) {
	r.POST("/api/reservasi", session.VerifyToken(), func(c *gin.Context) {
		// var requestData map[string]interface{}
		// if err := c.ShouldBindJSON(&requestData); err != nil {
		// 	c.JSON(http.StatusBadRequest, gin.H{"error": "Captcha tidak valid!"})
		// 	return
		// }
		
		var data databases.Reservasi
		if err := c.Bind(&data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Format data tidak valid!"})
			return
		}

		if result, err := recaptcha.Confirm(c.ClientIP(), data.RecaptchaResponse); err != nil || !result {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Verifikasi reCAPTCHA gagal!"})
			return
		}

		if err := db.Create(&data).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal mengirim data!"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Reservasi berhasil dikirim!"})
	})

	r.PUT("/api/reservasi/byID/:id", session.VerifyToken(), func(c *gin.Context) {
		id := c.Param("id")

		type RequestData struct {
			TglKunjungan 		string `json:"tgl_kunjungan"`
			Pembayaran   		string `json:"pembayaran"`
			NoAsuransi   		int    `json:"no_asuransi"`
			Email        		string `json:"email"`
			NoTelp       		string `json:"no_telp"`
			RecaptchaResponse 	string `json:"recaptchaResponse"`
		}

		var requestData RequestData
		if err := c.Bind(&requestData); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Format data tidak valid!"})
			return
		}

		if result, err := recaptcha.Confirm(c.ClientIP(), requestData.RecaptchaResponse); err != nil || !result {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Verifikasi reCAPTCHA gagal!"})
			return
		}

		if err := db.Model(&databases.Reservasi{}).Where("id_reservasi = ?", id).Updates(requestData).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal melakukan update data!"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Data berhasil diupdate!"})
	})

	r.DELETE("/api/reservasi/byID/:id", session.VerifyToken(), func(c *gin.Context) {
		id := c.Param("id")

		var requestData map[string]interface{}
		if err := c.ShouldBindJSON(&requestData); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Ada kesalahan pada server, mohon coba lagi nanti!"})
			return
		}
		
		recaptchaResponse := requestData["recaptchaResponse"].(string)
		if result, err := recaptcha.Confirm(c.ClientIP(), recaptchaResponse); err != nil || !result {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Verifikasi reCAPTCHA gagal!"})
			return
		}

		var reservasi databases.Reservasi
		if err := db.Where("id_reservasi = ?", id).Delete(&reservasi).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal menghapus data!"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Data berhasil dihapus!"})
	})

	r.GET("/api/reservasi/byID/:id", session.VerifyToken(), func(c *gin.Context) {
		id := c.Param("id")

		var data []databases.Reservasi
		if err := db.Model(&databases.Reservasi{}).Where("id_reservasi = ?", id).Find(&data).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal mengambil data!"})
			return
		}
		
		if len(data) == 0 {
			c.JSON(http.StatusNotFound, gin.H{"error": "Data tidak ditemukan!"})
			return
		}

		c.JSON(http.StatusOK, data)
	})

	r.GET("/api/reservasi/byNoRS/:param", session.VerifyToken(), func(c *gin.Context) {
		no_rs := c.Param("param")

		var data []databases.Reservasi
		if err := db.Model(&databases.Reservasi{}).Where("no_rs = ?", no_rs).Find(&data).Error; err != nil {
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
