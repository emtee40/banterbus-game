package entities

type NewPlayer struct {
	ID        string
	Nickname  string
	Avatar    []byte
	SessionID int64
}

type NewRoom struct {
	GameName string
	RoomCode string
}

type Player struct {
	ID       string
	Nickname string
	Avatar   string
}

type Room struct {
	Code    string
	Players []Player
}
