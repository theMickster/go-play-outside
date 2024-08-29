package controllers

import (
	"net/http"
	"shawskyRecords/models"

	"github.com/gin-gonic/gin"
)

// GetAlbums godoc
//
//	@Summary		Retrieve all albums
//	@Tags			Albums
//	@Produce		json
//	@Success		200	{array}		models.Album
//	@Failure		400	{object}	httputil.HTTPError
//	@Failure		404	{object}	httputil.HTTPError
//	@Failure		500	{object}	httputil.HTTPError
//	@Router			/albums [get]
func (c *Controller) GetAlbums(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, models.Albums)
}

// CreateAlbum godoc
//
//	@Summary		Create a single album
//	@Tags			Albums
//	@Produce		json
//	@Accept			json
//	@Param			album	body		models.Album	true	"Add album"
//	@Success		200	{object}	models.Album
//	@Failure		400	{object}	httputil.HTTPError
//	@Failure		404	{object}	httputil.HTTPError
//	@Failure		500	{object}	httputil.HTTPError
//	@Router			/albums [post]
func (c *Controller) CreateAlbum(content *gin.Context) {
	var newAlbum models.Album

	if err := content.BindJSON(&newAlbum); err != nil {
		return
	}

	models.Albums = append(models.Albums, newAlbum)
	content.IndentedJSON(http.StatusCreated, newAlbum)
}
