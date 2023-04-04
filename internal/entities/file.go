package entities

type File struct {
	Id        FileId
	Name      string
	Path      string
	CreatorId UserId
}
