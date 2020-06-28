package todo

import (
	"github.com/Masato4556/kurayami-go/db"
	"github.com/Masato4556/kurayami-go/entity"
)

// Service procides user's behavior
type Service struct{}

// Todo is alias of entity.Todo struct
type Todo entity.Todo

//DB追加
func (s Service) Insert(text string, status string) {
	db := db.GetDB()
	if err != nil {
		panic("データベース開けず！（dbInsert)")
	}
	db.Create(&entity.Todo{Text: text, Status: status})
	defer db.Close()
}

//DB更新
func (s Service) Update(id int, text string, status string) {
	db := db.GetDB()
	if err != nil {
		panic("データベース開けず！（dbUpdate)")
	}
	var todo entity.Todo
	db.First(&todo, id)
	todo.Text = text
	todo.Status = status
	db.Save(&todo)
	db.Close()
}

//DB削除
func (s Service) Delete(id int) {
	db := db.GetDB()
	if err != nil {
		panic("データベース開けず！（dbDelete)")
	}
	var todo entity.Todo
	db.First(&todo, id)
	db.Delete(&todo)
	db.Close()
}

//DB全取得
func (s Service) GetAll() []entity.Todo {
	db := db.GetDB()
	if err != nil {
		panic("データベース開けず！(dbGetAll())")
	}
	var todos []entity.Todo
	db.Order("created_at desc").Find(&todos)
	db.Close()
	return todos
}

//DB一つ取得
func (s Service) GetOne(id int) entity.Todo {
	db := db.GetDB()
	if err != nil {
		panic("データベース開けず！(dbGetOne())")
	}
	var todo entity.Todo
	db.First(&todo, id)
	db.Close()
	return todo
}
