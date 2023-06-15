package main

const (
	sqlFoo  = `select id, name, count from table1 where id = :id;`
	sqlFoo2 = `select id, name, count from table1 where id in $1;`
)

type Row1 struct {
	Id    string `db:"id"`
	Name  string `db:"name"`
	Count int    `db:"count"`
}

func (r *Row1) LoadInfoFromDb() (err error) {
	return dbFirst.Get(r, sqlFoo)
}

func selectRowsFromDbByIds(ids []string) (info []Row1, err error) {
	info = []Row1{}
	if err = dbFirst.Select(&info, sqlFoo2, ids); err != nil {
		return
	}

	return
}
