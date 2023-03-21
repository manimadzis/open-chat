package entities

type File struct {
	Id        FileId
	Name      string
	Path      string
	CreatorId UserId
}

type FileView struct {
	Id   FileId `json:"id"`
	Path string `json:"path"`
}
