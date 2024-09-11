package services

import (
	"fmt"
	"shawskyRecords/models"
	"strings"
	"sync"

	"github.com/google/uuid"
)

var albums = []models.Album{
	{Id: "41cbc53b-9434-4321-8542-6a1fdac1fc8f", Title: "Up All Night", Artist: "Kip Moore", Price: 9.99},
	{Id: "dfea4654-d5da-425a-9866-ab6c9dd10d7e", Title: "Wild Ones", Artist: "Kip Moore", Price: 11.99},
	{Id: "c8b3a3e1-bec4-4a8a-9879-2ae50b187cbc", Title: "Slowheart", Artist: "Kip Moore", Price: 7.99},
	{Id: "be9fa363-e9c8-4a75-b2c8-68c19a457710", Title: "Wild World", Artist: "Kip Moore", Price: 10.39},
	{Id: "69e12637-1b77-465e-b0ba-d2b70831f92e", Title: "Damn Love", Artist: "Kip Moore", Price: 11.19},
	{Id: "a1161ffe-c006-416d-9a3b-f53813431bbb", Title: "The Honky Tonk Kid", Artist: "Aaron Watson", Price: 10.00},
	{Id: "83f3b3a5-9275-4e9f-bc95-c93aedafa5af", Title: "San Angelo", Artist: "Aaron Watson", Price: 10.50},
	{Id: "6be2d50f-3787-477f-b3c1-69fe31dc3dba", Title: "Angels & Outlaws", Artist: "Aaron Watson", Price: 8.79},
	{Id: "8d1ae8ac-ec55-46ad-a496-7e0438a15c04", Title: "Unwanted Man", Artist: "Aaron Watson", Price: 9.99},
	{Id: "403fab66-2a4b-45fe-91ca-50422dba1bd2", Title: "Vaquero", Artist: "Aaron Watson", Price: 9.99},
	{Id: "c9288909-8da9-4c45-b591-02d44b6fcd91", Title: "Cover Girl", Artist: "Aaron Watson", Price: 14.89},
	{Id: "a35f9a71-1ef0-419f-ab7e-fefd07ef628a", Title: "American Soul", Artist: "Aaron Watson", Price: 9.49},
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
		if strings.EqualFold(a.Id, id) {
			return a, nil
		}
	}
	return models.Album{}, fmt.Errorf("album not found with id %v", id)
}

func (s AlbumService) CreateAlbum(input models.Album) (models.Album, error) {
	m.Lock()
	defer m.Unlock()
	for _, a := range albums {
		if strings.EqualFold(a.Artist, input.Artist) && strings.EqualFold(a.Title, input.Title) {
			return models.Album{}, fmt.Errorf("cannot create a duplicate album for artist :: %v  & title :: %v", input.Artist, input.Title)
		}
	}
	input.Id = uuid.NewString()
	albums = append(albums, input)
	return input, nil
}

func (s AlbumService) UpdateAlbum(input models.Album) (models.Album, error) {
	m.RLock()
	defer m.RUnlock()
	for a := range albums {
		if albums[a].Id == input.Id {
			albums[a] = input
			return input, nil
		}
	}
	return models.Album{}, fmt.Errorf("album not found with id %v", input.Id)
}

func (s AlbumService) DeleteAlbum(id string) (bool, error) {
	m.Lock()
	defer m.Unlock()
	for a := range albums {
		if strings.EqualFold(albums[a].Id, id) {
			albums = append(albums[:a], albums[a+1:]...)
			return true, nil
		}
	}
	return false, fmt.Errorf("album not found with id %v", id)
}
