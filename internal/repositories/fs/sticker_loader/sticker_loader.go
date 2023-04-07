package sticker_loader

import (
	"fmt"
	"open-chat/internal/entities"
	"open-chat/internal/services"
	"os"
	"path"
)

type stickerLoader struct {
	dir string
}

func (s stickerLoader) Download(stickers []entities.Sticker) error {
	for i, sticker := range stickers {
		data, err := os.ReadFile(path.Join(s.dir, sticker.Path))
		if err != nil {
			if err == os.ErrNotExist {
				return services.ErrNoSuchSticker
			}
			return services.NewUnknownError(err)
		}
		stickers[i].Data = data
	}
	return nil
}

func (s stickerLoader) Upload(stickers []entities.Sticker) error {
	for _, sticker := range stickers {
		filename := fmt.Sprintf("%d.png", sticker.Id)
		err := os.WriteFile(path.Join(s.dir, filename), sticker.Data, 0744)
		if err != nil {
			if err == os.ErrExist {
				return services.ErrStickerAlreadyExists
			}
			return services.NewUnknownError(err)
		}
	}
	return nil
}

func NewStickerLoader(dir string) services.StickerLoader {
	return &stickerLoader{dir: dir}
}
