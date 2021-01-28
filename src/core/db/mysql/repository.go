package mysql

type Reader interface {
	Find(id entity.ID)
}
