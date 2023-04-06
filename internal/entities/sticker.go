package entities

type StickerPack struct {
	Id        StickerPackId
	Name      string
	Stickers  []Sticker
	CreatorId UserId
}

type Sticker struct {
	Id   StickerId
	Path string
	StickerPackId
	Data []byte
}
