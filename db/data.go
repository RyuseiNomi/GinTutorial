package db

import (
	"fmt"
	"log"
	"reflect"

	"golang.org/x/net/context"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"

	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
)

type FirebaseInstance struct {
	Records []map[string]interface{}
}

type Todo struct {
	gorm.Model
	Text   string
	Status string
}

/**
 * Google Cloud Firestoreの初期化を実行する
 */
func Init() (*firestore.Client, context.Context) {
	ctx := context.Background()
	opt := option.WithCredentialsFile("firebase/gintutorial-89639-firebase-adminsdk-ifdrh-d9d477af86.json")
	app, err := firebase.NewApp(ctx, nil, opt)
	if err != nil {
		log.Fatal("Database Initialization Error: %v", err)
	}
	client, err := app.Firestore(ctx)
	if err != nil {
		log.Fatal("Database InitializationError: %v", err)
	}
	return client, ctx
}

/**
 * DBにレコードを追加する
 * @param text   string TODOの内容
 * @param status string TODOのステータス
 */
func Insert(client *firestore.Client, ctx context.Context, text string, status string) {
	_, _, err := client.Collection("todos").Add(ctx, map[string]interface{}{
		"text":   text,
		"status": status,
	})
	if err != nil {
		log.Fatalf("Failed insert record: &v", err)
	}
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
		log.Printf("DBの初期化に失敗しました。%v", err)
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
		log.Printf("DBの初期化に失敗しました。%v", err)
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
func FetchAll(ctx context.Context, client *firestore.Client) []Todo {
	var todos []Todo
	iter := client.Collection("todos").Documents(ctx)
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Fatalf("Failed to load documents: %v", err)
		}
		todo := doc.Data()
		fmt.Println(reflect.TypeOf(todo["text"]))
		todos = append(
			todos,
			Todo{
				Text:   todo["text"].(string),
				Status: todo["status"].(string),
			})
	}
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
		log.Printf("DBの初期化に失敗しました。%v", err)
	}
	var todo Todo
	db.First(&todo, id)
	db.Close()
	return todo
}
