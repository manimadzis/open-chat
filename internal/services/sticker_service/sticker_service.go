package sticker_service

import (
	"context"
	"open-chat/internal/entities"
	"open-chat/internal/services"
)

type stickerService struct {
	stickerRepo   services.StickerRepository
	stickerLoader services.StickerLoader
}

func (s stickerService) CreateStickerPack(
	ctx context.Context,
	stickerPack entities.StickerPack,
) (
	entities.StickerPackId,
	error,
) {
	var err error
	if stickerPack.Id, err = s.stickerRepo.CreateStickerPack(
		ctx,
		stickerPack,
	); err != nil {
		return 0, err
	}

	if err := s.stickerLoader.Upload(stickerPack.Stickers); err != nil {
		return 0, err
	}
	return stickerPack.Id, nil
}

func (s stickerService) DeleteStickerPack(
	ctx context.Context,
	stickerPackId entities.StickerPackId,
) error {
	var err error
	if err = s.stickerRepo.DeleteStickerPack(
		ctx,
		stickerPackId,
	); err != nil {
		return err
	}

	return nil
}

func (s stickerService) FindStickersByStickerPackId(
	ctx context.Context,
	stickerPackId entities.StickerPackId,
) (
	[]entities.Sticker,
	error,
) {
	stickers, err := s.stickerRepo.FindStickersByStickerPackId(
		ctx,
		stickerPackId,
	)
	if err != nil {
		return nil, err
	}

	if err = s.stickerLoader.Download(stickers); err != nil {
		return nil, err
	}

	return stickers, nil
}

func NewStickerService(
	stickerRepo services.StickerRepository,
	stickerLoader services.StickerLoader,
) services.StickerService {
	return &stickerService{
		stickerRepo:   stickerRepo,
		stickerLoader: stickerLoader,
	}
}
