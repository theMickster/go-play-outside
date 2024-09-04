package services

import (
	"fmt"
	"shawskyRecords/models"
	"sync"

	"github.com/google/uuid"
)

var albums = []models.Album{
	{Id: "41cbc53b-9434-4321-8542-6a1fdac1fc8f", Title: "Up All Night", Artist: "Kip Moore", Price: 9.99},
	{Id: "dfea4654-d5da-425a-9866-ab6c9dd10d7e", Title: "Wild Ones", Artist: "Kip Moore", Price: 11.99},
	{Id: "c8b3a3e1-bec4-4a8a-9879-2ae50b187cbc", Title: "Slowheart", Artist: "Kip Moore", Price: 7.99},
	{Id: "be9fa363-e9c8-4a75-b2c8-68c19a457710", Title: "Wild World", Artist: "Kip Moore", Price: 10.39},
	{Id: "69e12637-1b77-465e-b0ba-d2b70831f92e", Title: "Damn Love", Artist: "Kip Moore", Price: 11.19},
}

type AlbumService struct{}

var m sync.RWMutex

func NewAlbumService() *AlbumService {
	return &AlbumService{}
}

func (s AlbumService) GetAlbums() []models.Album {
	m.RLock()
	defer m.RUnlock()
	return albums
}

func (s AlbumService) GetAlbumById(id string) (models.Album, error) {
	m.RLock()
	defer m.RUnlock()
	for _, a := range albums {
		if a.Id == id {
			return a, nil
		}
	}
	return models.Album{}, fmt.Errorf("album not found with id %v", id)
}

func (s AlbumService) CreateAlbum(input models.Album) models.Album {
	m.Lock()
	defer m.Unlock()
	input.Id = uuid.NewString()
	albums = append(albums, input)
	return input
}
