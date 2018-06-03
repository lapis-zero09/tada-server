package v1

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/labstack/echo"
	"github.com/lapis-zero09/tada-server/models"
)

func GetTags(c echo.Context) error {
	db := models.InitTagTable()
	defer db.Close()

	var tags []models.Tag
	db.Find(&tags)

	return c.JSON(http.StatusOK, tags)
}

func GetTag(c echo.Context) error {
	db := models.InitTagTable()
	defer db.Close()
	id := c.Param("tagId")
	var tag models.Tag
	db.First(&tag, id)

	if tag.ID > 0 {
		return c.JSON(http.StatusOK, tag)
	}
	err := fmt.Errorf("tagId=%s is not found", id)
	return echo.NewHTTPError(http.StatusNotFound, err.Error())
}

func PostTag(c echo.Context) error {
	db := models.InitTagTable()
	defer db.Close()

	var tag models.Tag
	c.Bind(&tag)
	if tag.Name != "" {
		db.Create(&tag)
		return c.JSON(http.StatusCreated, tag)
	}
	err := errors.New("Values must be int")
	return echo.NewHTTPError(http.StatusBadRequest, err.Error())
}

func UpdateTag(c echo.Context) error {
	db := models.InitTagTable()
	defer db.Close()
	id := c.Param("tagId")
	var tag models.Tag
	db.First(&tag, id)

	if tag.ID > 0 {
		var newTag models.Tag
		c.Bind(&newTag)
		if newTag.Name != "" {
			result := models.Tag{
				ID:   tag.ID,
				Name: newTag.Name,
			}
			db.Save(&result)
			return c.JSON(http.StatusOK, result)
		}
		err := errors.New("Values must be int")
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	err := fmt.Errorf("tagId=%s is not found", id)
	return echo.NewHTTPError(http.StatusBadRequest, err.Error())
}

func DeleteTag(c echo.Context) error {
	db := models.InitTagTable()
	defer db.Close()
	id := c.Param("tagId")
	var tag models.Tag
	db.First(&tag, id)

	if tag.ID > 0 {
		db.Delete(&tag)
		return c.NoContent(http.StatusNoContent)
	}
	err := fmt.Errorf("tagId=%s is not found", id)
	return echo.NewHTTPError(http.StatusNotFound, err.Error())
}
