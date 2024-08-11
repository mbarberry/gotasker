// GPL-3.0
// **Copyright (c) Mike Barberry**

package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var headers = map[string]string{"Access-Control-Allow-Origin": "*", "Content-Type": "application/json"}
var ctx = context.TODO()
var collection *mongo.Collection

type ErrorMessage struct {
	Error string `json:"string"`
}

type RequestBody struct {
	Data string `json:"task"`
}

type Task struct {
	ID        primitive.ObjectID `bson:"_id"`
	CreatedAt time.Time          `bson:"created_at"`
	UpdatedAt time.Time          `bson:"updated_at"`
	Text      string             `bson:"text"`
	Completed bool               `bson:"completed"`
}

func init() {
	uri := os.Getenv("DATABASE_URI")
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	collection = client.Database("tasker").Collection("tasks")
}

func main() {
	lambda.Start(router)
}

func router(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	var jsonBody RequestBody
	err := json.Unmarshal([]byte(request.Body), &jsonBody)
	if err != nil {
		fmt.Println("Error decoding data.")
		panic(err)
	}

	path := request.Path[5:]
	switch path {
	case "getTasks":
		return getRouteHandler()
	case "addTask":
		return addRouteHandler(jsonBody.Data)
	case "upateStatus":
		return updateStatusHandler(jsonBody.Data)
	case "deleteTask":
		return deleteTaskHandler(jsonBody.Data)
	default:
		jsonErrorMessage := getJSONErrorMesage("Route not found.")
		return errorResponse(400, string(jsonErrorMessage))
	}
}

func getRouteHandler() (events.APIGatewayProxyResponse, error) {
	var tasks []Task
	queryFilter := bson.D{{}}
	queryResults, err := collection.Find(ctx, queryFilter)
	if err != nil {
		jsonErrorMessage := getJSONErrorMesage("Error fetching tasks.")
		return errorResponse(500, string(jsonErrorMessage))
	}
	for queryResults.Next(ctx) {
		var currentTask Task
		err := queryResults.Decode(&currentTask)
		if err != nil {
			fmt.Println("Error decoding data.")
		}
		tasks = append(tasks, currentTask)
	}
	jsonTasks, _ := json.Marshal(tasks)
	return successfulResponse(string(jsonTasks))

}

func addRouteHandler(task string) (events.APIGatewayProxyResponse, error) {
	if task == "" {
		fmt.Println("Invalid input: task must not be empty.")
	}
	taskToAdd := Task{
		ID:        primitive.NewObjectID(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Text:      task,
		Completed: false,
	}
	_, err := collection.InsertOne(ctx, taskToAdd)
	if err != nil {
		jsonErrorMessage := getJSONErrorMesage("Error inserting new task.")
		return errorResponse(500, string(jsonErrorMessage))
	}
	jsonTaskToAdd, err := json.Marshal(taskToAdd)
	if err != nil {
		fmt.Println("Error encoding new task.")
	}
	return successfulResponse(string(jsonTaskToAdd))
}

func updateStatusHandler(id string) (events.APIGatewayProxyResponse, error) {
	filter := getIdFilter(id)
	updateTaskOperation := bson.D{{Key: "$set", Value: bson.D{primitive.E{Key: "completed", Value: true}, primitive.E{Key: "updated_at", Value: time.Now()}}}}
	_, err := collection.UpdateOne(ctx, filter, updateTaskOperation)
	if err != nil {
		jsonErrorMessage := getJSONErrorMesage("Error updating task.")
		return errorResponse(500, string(jsonErrorMessage))
	}
	return successfulResponse(id)
}

func deleteTaskHandler(id string) (events.APIGatewayProxyResponse, error) {
	filter := getIdFilter(id)
	_, err := collection.DeleteOne(ctx, filter)
	if err != nil {
		jsonErrorMessage := getJSONErrorMesage("Error deleting task.")
		return errorResponse(500, string(jsonErrorMessage))
	}
	return successfulResponse(id)
}

func getIdFilter(id string) primitive.D {
	idPrimitive, _ := primitive.ObjectIDFromHex(id)
	idFilter := bson.D{{Key: "_id", Value: idPrimitive}}
	return idFilter
}

func getJSONErrorMesage(message string) []byte {
	errorMessage := ErrorMessage{Error: message}
	jsonErrorMessage, _ := json.Marshal(errorMessage)
	return jsonErrorMessage
}

func successfulResponse(data string) (events.APIGatewayProxyResponse, error) {
	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusOK,
		Headers:    headers,
		Body:       data,
	}, nil
}

func errorResponse(statusCode int, data string) (events.APIGatewayProxyResponse, error) {
	return events.APIGatewayProxyResponse{
		StatusCode: statusCode,
		Headers:    headers,
		Body:       data,
	}, nil
}
