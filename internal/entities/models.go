package entities

type NewPlayer struct {
	Nickname  string
	Avatar    []byte
	SessionID int64
}

type NewRoom struct {
	GameName string
	RoomCode string
}

type Player struct {
	ID       int64
	Nickname string
	Avatar   string
}

type Room struct {
	Code    string
	Players []Player
}
