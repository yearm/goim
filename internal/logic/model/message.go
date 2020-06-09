package model

type Message struct {
	Id          int64
	Poster      int64
	Receiver    int64
	MessageType uint
	Message     string
	Dateline    int64
	Seq         int64
}
