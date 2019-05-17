package model

import (
	"time"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/guregu/dynamo"
	"github.com/google/uuid"
)

// Todo structure of DynamoDB
type Todo struct {
	UUID        string `dynamo:"uuid"`
	Title     string    `dynamo:"title"`
	CreatedAt time.Time `dynamo:"created_at"`
}

// Get todo item by id
func Get(id int) (Todo, error) {
	table := getTable()

	var todo Todo
	err := table.Get("id", id).One(&todo)
	
	return todo, err
}

// Put todo item by title
func Put(title string) error {
	uuid, err := uuid.NewUUID()
	if err != nil {
		return err
	}

	item := Todo{
		UUID:      uuid.String(),
		Title:     title,
		CreatedAt: time.Now().UTC(),
	}

	table := getTable()

	return table.Put(item).Run()
}

// getTable get todo table connection
func getTable() dynamo.Table {
	db := dynamo.New(session.New())

	return db.Table("Todo")
}
