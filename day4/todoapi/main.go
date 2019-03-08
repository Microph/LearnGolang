package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Task struct {
	Desc string `bson:"desc" json:"desc"`
	Done bool   `bson:"done" json:"done"`
}

type TodoList struct {
	ID    primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Title string             `bson:"title" json:"title"`
	Tasks []Task             `bson:"tasks" json:"tasks"`
}

func check(c *gin.Context, err error) {
	if err != nil {
		fmt.Println(err)
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
}

func main() {
	mongoURI := "mongodb://u30cbkfrc67odtpvtbyq:Zyl4Q2lK6zQKr4fRga4L@bnbfqre3af2qu3i-mongodb.services.clever-cloud.com:27017/bnbfqre3af2qu3i"
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(mongoURI))
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	if err := client.Ping(context.Background(), nil); err != nil {
		log.Println(err)
		os.Exit(1)
	}
	collection := client.Database("bnbfqre3af2qu3i").Collection("todos")
	r := gin.Default()

	r.GET("/todos", func(c *gin.Context) {
		ctx := context.TODO()
		cur, err := collection.Find(ctx, bson.D{})
		check(c, err)
		defer cur.Close(ctx)

		var todos []TodoList
		for cur.Next(ctx) {
			todo := TodoList{
				Tasks: []Task{},
			}
			err := cur.Decode(&todo)
			check(c, err)

			todos = append(todos, todo)
		}

		c.JSON(http.StatusOK, todos)
	})

	r.POST("/todos", func(c *gin.Context) {
		ctx := context.TODO()
		var todo TodoList
		err := c.Bind(&todo)
		check(c, err)

		//insert
		res, err := collection.InsertOne(ctx, todo)
		check(c, err)

		todo.ID = res.InsertedID.(primitive.ObjectID)
		c.JSON(http.StatusOK, todo)
	})

	r.GET("/todos/:id", func(c *gin.Context) {
		ctx := context.TODO()
		var todo TodoList
		id, _ := primitive.ObjectIDFromHex(c.Param("id"))
		err := collection.FindOne(ctx, bson.D{{"_id", id}}).Decode(&todo)
		check(c, err)

		c.JSON(http.StatusOK, todo)
	})

	r.PUT("/todos/:id", func(c *gin.Context) {
		ctx := context.TODO()
		id, _ := primitive.ObjectIDFromHex(c.Param("id"))
		var todo TodoList
		err := c.Bind(&todo)
		check(c, err)

		var param struct {
			Title string `json:"title" bson:"title"`
		}

		param.Title = todo.Title
		_, err = collection.UpdateMany(ctx,
			bson.D{{"_id", id}},
			bson.D{{"$set", bson.M{"title": param.Title}}})
		check(c, err)

		c.JSON(http.StatusOK, nil)
	})

	r.DELETE("/todos/:id", func(c *gin.Context) {
		ctx := context.TODO()
		id, _ := primitive.ObjectIDFromHex(c.Param("id"))
		_, err := collection.DeleteMany(ctx, bson.D{{"_id", id}})
		check(c, err)

		c.JSON(http.StatusOK, nil)
	})
	r.Run(":8000")
}
