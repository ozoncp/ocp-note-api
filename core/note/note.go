package note

import "fmt"

type Note struct {
	Id          uint
	UserId      uint
	ClassroomId uint
	DocumentId  uint
}

func (n *Note) Init(id uint, userId uint, classroomId uint, documentId uint) {
	n.Id = id
	n.UserId = userId
	n.ClassroomId = classroomId
	n.DocumentId = documentId
}

func (n *Note) Create() {
	fmt.Println("Note created")
}

func (n *Note) Delete() {
	fmt.Println("Note deleted")
}

func (n *Note) Change() {
	fmt.Println("Note changed")
}

func (n *Note) String() {
	fmt.Printf("Id: %d, UserId: %d, ClassroomId: %d, DocumentId: %d", n.Id, n.UserId, n.ClassroomId, n.DocumentId)
}
