package model

type Contacts struct {
	Id             int64
	Uid            int64
	InteractiveUid int64
	Status         uint
	MessageId      int64
	Dateline       int64
	IsRead         uint
}
