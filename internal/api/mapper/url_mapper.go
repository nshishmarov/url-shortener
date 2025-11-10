package mapper

import (
	"url-shortener/internal/api/dto"
	"url-shortener/internal/storage/entity"
)

func MapUrlEntityToUrlDto(urlEntity *entity.URL) *dto.URL {
	return &dto.URL{
		URL: urlEntity.URL,
	}
}