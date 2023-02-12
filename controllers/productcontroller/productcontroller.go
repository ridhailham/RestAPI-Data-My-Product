package productcontroller

import (
	"encoding/json"
	"net/http"

	"go-restapi-gin/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Index(c *gin.Context) {
	var products []models.Product

	models.DB.Find(&products)
	c.JSON(http.StatusOK, gin.H{
		"products": products,
	})

}

func Show(c *gin.Context) {
	var product models.Product
	id := c.Param("id")

	err := models.DB.First(&product, id).Error

	if err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
				"massage": "Data tidak ditemukan",
			})
			return
		default:
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"massage": err.Error(),
			})
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"product": product,
	})
}

func Create(c *gin.Context) {
	var product models.Product

	err := c.ShouldBindJSON(&product)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"massage": err.Error(),
		})
		return
	}

	models.DB.Create(&product)
	c.JSON(http.StatusOK, gin.H{
		"product": product,
	})
}

func Update(c *gin.Context) {
	var product models.Product
	id := c.Param("id")

	err := c.ShouldBindJSON(&product)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"massage": err.Error(),
		})
		return
	}

	if models.DB.Model(&product).Where("id = ?", id).Updates(&product).RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"massage": "tidak dapat mengupdate product",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"massage": "Data berhasil diperbaruhi",
	})
}

func Delete(c *gin.Context) {
	var product models.Product

	var input struct {
		Id json.Number
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"massage": product,
		})
		return
	}

	id, _ := input.Id.Int64()
	if models.DB.Delete(&product, id).RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"massage": "tidak dapat mengupdate data",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"massage": "Data berhasil diupdate",
	})
}
