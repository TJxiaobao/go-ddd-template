package vo

type ID struct {
	id int64
}

func NewID(id int64) ID {
	return ID{id: id}
}

func (i ID) GetId() int64 {
	return i.id
}
