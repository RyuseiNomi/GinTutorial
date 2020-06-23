package db

import (
	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
)

type Todo struct {
	gorm.Model
	Text   string
	Status string
}

/**
 * DBにレコードを追加する
 */
func Init() {
	db, err := gorm.Open("sqlite3", "test.sqlite3")
	if err != nil {
		panic("DBの初期化に失敗しました。%v", err)
	}
	db.AutoMigrate(&Todo{})
	defer db.Close()
}

/**
 * DBにレコードを追加する
 * @param text   string TODOの内容
 * @param status string TODOのステータス
 */
func Insert(text string, status string) {
	db, err := gorm.Open("sqlite3", "test.sqlite3")
	if err != nil {
		panic("DBの初期化に失敗しました。%v", err)
	}
	db.Create(&Todo{Text: text, Status: status})
	defer db.Close()
}

/**
 * 対象のDBのレコードの内容を更新する
 * @param id     int    更新対象のTODOレコードのID
 * @param text   string TODOの内容
 * @param status string TODOのステータス
 */
func Update(id int, text string, status string) {
	db, err := gorm.Open("sqlite3", "test.sqlite3")
	if err != nil {
		panic("DBの初期化に失敗しました。%v", err)
	}
	var todo Todo
	db.First(&todo, id)
	todo.Text = text
	todo.Status = status
	db.Save(&todo)
	db.Close()
}

/**
 * 対象のDBのレコードを削除する
 * @param id int 更新対象のTODOレコードのID
 */
func Delete(id int) {
	db, err := gorm.Open("sqlite3", "test.sqlite3")
	if err != nil {
		panic("DBの初期化に失敗しました。%v", err)
	}
	var todo Todo
	db.First(&todo, id)
	db.Delete(&todo)
	db.Close()
}

/**
 * TODOレコードを作成日順に全て取得する
 * @return array Todo DBに保存されている全てのTodo
 */
func FindAll() []Todo {
	db, err := gorm.Open("sqlite3", "test.sqlite3")
	if err != nil {
		panic("DBの初期化に失敗しました。%v", err)
	}
	var todos []Todo
	db.Order("created_at desc").Find(&todos)
	db.Close()
	return todos
}

/**
 * 指定のTODOレコードを1件取得する
 * @param  int  id 取得したいレコードのID
 * @return Todo DBに保存されているTODO
 */
func FetchOne(id int) Todo {
	db, err := gorm.Open("sqlite3", "test.sqlite3")
	if err != nil {
		panic("DBの初期化に失敗しました。%v", err)
	}
	var todo Todo
	db.First(&todo, id)
	db.Close()
	return todo
}
