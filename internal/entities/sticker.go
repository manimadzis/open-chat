package entities

type StickerPack struct {
	Id       StickerPackId
	Name     string
	Stickers []Sticker
}

type Sticker struct {
	Id  StickerId
	URL string
	StickerPackId
}
