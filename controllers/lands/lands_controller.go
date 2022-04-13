package lands

import (
	"eden/domain/lands"
	"eden/services"
	"eden/utils/errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func Create(c *gin.Context) {
	var land lands.Land
	if err := c.ShouldBindJSON(&land); err != nil {
		restErr := errors.NewBadRequestError("invalid json body")
		c.JSON(restErr.Status, restErr)
		return
	}
	result, saveErr := services.CreateLand(land)
	if saveErr != nil {
		c.JSON(saveErr.Status, saveErr)
		return
	}
	c.JSON(http.StatusCreated, result)
}

func Get(c *gin.Context) {
	landId, landErr := strconv.ParseInt(c.Param("land_id"), 10, 64)
	if landErr != nil {
		err := errors.NewBadRequestError("invalid land id")
		c.JSON(err.Status, err)
		return
	}
	land, getErr := services.GetLand(landId)
	if getErr != nil {
		c.JSON(getErr.Status, getErr)
		return
	}
	c.JSON(http.StatusOK, land)
}

func Search(c *gin.Context) {
	c.String(http.StatusNotImplemented, "implement me!")
}
