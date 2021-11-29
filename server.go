package main 

import (
	"net/http"
	"encoding/json"
	"fmt"
	"context"
	"os"
	"time"
	"log"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
)

type Request struct {
	Command string `json:"command"` 
	Task string `json:"task"`
}

type Response struct {
	Message string `json:"message"`
}

type Task struct {
    ID        primitive.ObjectID `bson:"_id"`
    CreatedAt time.Time          `bson:"created_at"`
    UpdatedAt time.Time          `bson:"updated_at"`
    Text      string             `bson:"text"`
    Completed bool               `bson:"completed"`
}

var collection *mongo.Collection
var ctx = context.TODO()

func init() {
	err := godotenv.Load()
  	if err != nil {
    	log.Fatal("Error loading .env file")
  	}

	uri := os.Getenv("DATABASE_URI")
    clientOptions := options.Client().ApplyURI(uri)

	client, err := mongo.Connect(ctx, clientOptions)
    if err != nil {
        log.Fatal(err)
    }
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}
	collection = client.Database("tasker").Collection("tasks")
}


func main(){
	port := os.Getenv("PORT")

	http.HandleFunc("/api/index", indexHandler)
	http.HandleFunc("/api/rm", deleteHandler)
	http.HandleFunc("/api/done", doneHandler)
	http.HandleFunc("/api/add", addHandler)

	fs := http.FileServer(http.Dir("./public"))
	http.Handle("/", fs)

	fmt.Println("Server listening on port 5000")
	log.Panic(
		http.ListenAndServe(":" + port, nil),
	)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	filter := bson.D{{}}
	var tasks []*Task

	cur, err := collection.Find(ctx, filter)
    if err != nil {
    	w.WriteHeader(400)
        fmt.Fprintf(w, "Error finding collection.")
	}

    for cur.Next(ctx) {
        var t Task
        err := cur.Decode(&t)
        if err != nil {
        	w.WriteHeader(400)
			fmt.Fprintf(w,"Error parsing data." )
		}

        tasks = append(tasks, &t)
    }

    if err := cur.Err(); err != nil {
        w.WriteHeader(400)
		fmt.Fprintf(w, "Database error.")
    }

    cur.Close(ctx)

    if len(tasks) == 0 {
    	w.WriteHeader(400)
		fmt.Fprintf(w, "No tasks to return.")
	}

	response, err := json.Marshal(tasks)
	if err != nil {
		panic(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(response)
}

func addHandler(w http.ResponseWriter, r *http.Request) {
	request := Request{}

	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		w.WriteHeader(400)
        fmt.Fprintf(w, "Decode error! please check your JSON formating.")
	}

	task := request.Task 	
	fmt.Printf("The task to add is: %s\n", task)
	if task == "" {
	//	w.WriteHeader(400)
	//	fmt.Fprintf(w, "Cannot add an empty task.")
		http.Error(w, "Cannot add an empty task.", http.StatusBadRequest)
		return 
	} //  else {
		addTask := &Task{
			ID:        primitive.NewObjectID(),
        	CreatedAt: time.Now(),
        	UpdatedAt: time.Now(),
        	Text:      task,
        	Completed: false,
		}
    	_, err = collection.InsertOne(ctx, addTask)
	
		response, err := json.Marshal(addTask)
		if err != nil {
			panic(err)
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(response)
//	}	
}

func deleteHandler(w http.ResponseWriter, r *http.Request) {
	request := Request{}

	err := json.NewDecoder(r.Body).Decode(&request)
     if err != nil {
         w.WriteHeader(400)
         fmt.Fprintf(w, "Decode error! please check your JSON formating.")
     }

	task := request.Task
	filter := bson.D{primitive.E{Key: "text", Value: task}}

    res, err := collection.DeleteOne(ctx, filter)
    if err != nil {
    	w.WriteHeader(400)
        fmt.Fprintf(w, "Error deleting task.")
	}

    if res.DeletedCount == 0 {
		w.WriteHeader(400)
        fmt.Fprintf(w, "No tasks were deleted")
    }

	serverResponse := &Response{
		Message: "Task successfully deleted.",
	}

	response, err := json.Marshal(serverResponse)
	if err != nil {
		panic(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(response)
}

func doneHandler(w http.ResponseWriter, r *http.Request) {
	request := Request{}

	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
    	w.WriteHeader(400)
    	fmt.Fprintf(w, "Decode error! please check your JSON formating.")
    }

	task := request.Task
	filter := bson.D{primitive.E{Key: "text", Value: task}}
    update := bson.D{primitive.E{Key: "$set", Value: bson.D{
        primitive.E{Key: "completed", Value: true},
    }}}
    _ = collection.FindOneAndUpdate(ctx, filter, update)

	serverResponse := &Response{
        Message: "Status successfully updated.",
    }
    response, err := json.Marshal(serverResponse)
    if err != nil {
        panic(err)
    }

    w.Header().Set("Content-Type", "application/json")
    w.Write(response)

}
