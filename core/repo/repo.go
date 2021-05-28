package repo

type Repo interface {
	AddTasks(task []Note) error
}
