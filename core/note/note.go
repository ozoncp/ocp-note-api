package note

import "fmt"

type Note struct {
	Id          uint64
	UserId      uint32
	ClassroomId uint32
	DocumentId  uint32
}

func (n *Note) Create() {
	fmt.Printf("Note '%d' created", n.Id)
}

func (n *Note) Delete() {
	fmt.Printf("Note '%d' deleted", n.Id)
}

func (n *Note) Change() {
	fmt.Printf("Note '%d' changed", n.Id)
}

func (n *Note) String() {
	fmt.Printf("Id: %d, UserId: %d, ClassroomId: %d, DocumentId: %d", n.Id, n.UserId, n.ClassroomId, n.DocumentId)
}
