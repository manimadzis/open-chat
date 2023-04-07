package services

import (
	"open-chat/internal/entities"
)

type StickerLoader interface {
	Download(stickers []entities.Sticker) error
	Upload(stickers []entities.Sticker) error
}
