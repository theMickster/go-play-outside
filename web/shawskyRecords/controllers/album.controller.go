package controllers

import (
	"net/http"
	"shawskyRecords/models"
	"shawskyRecords/services"

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
//
// @Param X-ApplicationId header string true "Application  ID"
func (c *Controller) GetAlbums(ctx *gin.Context) {
	service := services.NewAlbumService()
	ctx.IndentedJSON(http.StatusOK, service.GetAlbums())
}

// GetAlbumById godoc
//
//		@Summary		Retrieve a single album
//		@ID Retrieve an album by id
//		@Tags			Albums
//		@Produce		json
//	 @Param id path string true "album id"
//		@Success		200	{object}	models.Album
//		@Failure		400	{object}	httputil.HTTPError
//		@Failure		404	{object}	httputil.HTTPError
//		@Failure		500	{object}	httputil.HTTPError
//		@Router			/albums/{id} [get]
//
// @Param X-ApplicationId header string true "Application  ID"
func (c *Controller) GetAlbumById(ctx *gin.Context) {
	id := ctx.Param("id")
	service := services.NewAlbumService()
	result, err := service.GetAlbumById(id)
	if err != nil {
		ctx.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
		return
	}
	ctx.IndentedJSON(http.StatusOK, result)
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
//
// @Param X-ApplicationId header string true "Application  ID"
func (c *Controller) CreateAlbum(ctx *gin.Context) {
	service := services.NewAlbumService()
	var newAlbum models.Album

	if err := ctx.BindJSON(&newAlbum); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Invalid album body"})
		return
	}
	result := service.CreateAlbum(newAlbum)
	ctx.IndentedJSON(http.StatusCreated, result)
}
