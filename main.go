package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Task struct {
	Id        primitive.ObjectID `bson:"_id,omitempty"`
	Title     string             `bson:"title"`
	Completed bool               `bson:"completed"`
}

func (t Task) String() string {
	return fmt.Sprintf(
		"Task{_id='%s', Title='%s', Completed='%t'}",
		t.Id, t.Title, t.Completed)
}

var (
	databaseName = "todo"
	collTask     = "tasks"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}
}

func main() {
	uri := os.Getenv("MONGODB_URI")
	ctx := context.Background()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		log.Fatal(err)
	}

	if err := client.Ping(ctx, nil); err != nil {
		log.Fatal(err)
	}
	fmt.Println("Mongo Connected.")

	coll := client.Database(databaseName).Collection(collTask)

	DeleteMany(ctx, coll, bson.D{})
	InsertOne(ctx, coll)
	InsertMany(ctx, coll)
	// Update(ctx, coll, "66377fde0a64de919e76e735")
	// Delete(ctx, coll, "6637f5e2c8849ab3cd545a21")
	FindOne(ctx, coll, bson.D{{Key: "title", Value: "hug cat"}})
	Find(ctx, coll, bson.D{})

	// jsonData, err := json.MarshalIndent(cursors, "", "    ")
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Printf("%s\n", jsonData)

	defer func() {
		if err := client.Disconnect(ctx); err != nil {
			log.Fatal(err)
		}
		fmt.Println("Mongo Disconnected.")
	}()

}

func InsertOne(ctx context.Context, coll *mongo.Collection) {
	task := Task{Title: "play with cat"}
	result, err := coll.InsertOne(ctx, task)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(result.InsertedID)
}

func InsertMany(ctx context.Context, coll *mongo.Collection) {
	tasks := []interface{}{
		Task{Title: "hug cat", Completed: false},
		Task{Title: "feed cat", Completed: false},
	}
	results, err := coll.InsertMany(ctx, tasks)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(results.InsertedIDs)
}

func FindOne(ctx context.Context, coll *mongo.Collection, filter interface{}) {
	var task Task
	if err := coll.FindOne(ctx, filter).Decode(&task); err != nil {
		if err == mongo.ErrNoDocuments {
			// fmt.Println("no documents in result")
			log.Fatal("no documents in result")
		} else {
			log.Fatal(err)
		}
	}
	fmt.Println(task)
}

func Find(ctx context.Context, coll *mongo.Collection, filter interface{}) {
	cursor, err := coll.Find(ctx, filter)
	if err != nil {
		log.Fatal(err)
	}
	for cursor.Next(ctx) {
		var task Task
		if err := cursor.Decode(&task); err != nil {
			log.Fatal(err)
		}
		fmt.Println(task)
	}
}

func Update(ctx context.Context, coll *mongo.Collection, id string) {
	objectID, _ := primitive.ObjectIDFromHex(id)
	filter := bson.D{{Key: "_id", Value: objectID}}
	update := bson.D{{Key: "$set", Value: bson.D{{Key: "completed", Value: true}}}}
	result, err := coll.UpdateOne(ctx, filter, update)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(result)
}

func Delete(ctx context.Context, coll *mongo.Collection, id string) {
	objectID, _ := primitive.ObjectIDFromHex(id)
	filter := bson.D{{Key: "_id", Value: objectID}}
	result, err := coll.DeleteOne(ctx, filter)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(result)
}
func DeleteMany(ctx context.Context, coll *mongo.Collection, filter interface{}) {
	result, err := coll.DeleteMany(ctx, filter)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(result)
}
