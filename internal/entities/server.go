package entities

type Server struct {
	Id      ServerId
	Name    string
	OwnerId UserId
}

type ServerView struct {
	Id    ServerId `json:"id,omitempty"`
	Name  string   `json:"name,omitempty"`
	Users []User   `json:"users,omitempty"`
	Owner *User    `json:"owner,omitempty"`
}
