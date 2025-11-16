package main

import (
	"context"
	"encoding/json"
	"log"
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
var ctx = context.Background()
var collection *mongo.Collection

type ErrorMessage struct {
	Error string `json:"error"`
}

type RequestBody struct {
	Task string `json:"task"`
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
		log.Printf("error connecting to Mongodb: %v\n", err)
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
		log.Printf("error decoding data: %v\n", err)
	}

	path := request.Path[5:]
	switch path {
	case "getTasks":
		return getRouteHandler()
	case "addTask":
		return addRouteHandler(jsonBody.Task)
	case "upateStatus":
		return updateStatusHandler(jsonBody.Task)
	case "deleteTask":
		return deleteTaskHandler(jsonBody.Task)
	default:
		return errorResponse(http.StatusBadRequest, &ErrorMessage{Error: "route not found."})
	}
}

func getRouteHandler() (events.APIGatewayProxyResponse, error) {
	var tasks []Task
	cursor, err := collection.Find(ctx, bson.D{{}})
	if err != nil {
		return errorResponse(http.StatusInternalServerError, &ErrorMessage{Error: "error fetching tasks."})
	}
	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		var currentTask Task
		err := cursor.Decode(&currentTask)
		if err != nil {
			log.Printf("error decoding data: %v\n", err)
		}
		tasks = append(tasks, currentTask)
	}
	jsonTasks, err := json.Marshal(tasks)
	if err != nil {
		log.Printf("error marshaling tasks %v\n", err)
	}
	return successfulResponse(string(jsonTasks))

}

func addRouteHandler(task string) (events.APIGatewayProxyResponse, error) {
	if task == "" {
		errorResponse(http.StatusBadRequest, &ErrorMessage{Error: "invalid input: task must not be empty."})
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
		return errorResponse(http.StatusInternalServerError, &ErrorMessage{Error: "error inserting new task."})
	}
	jsonTaskToAdd, err := json.Marshal(taskToAdd)
	if err != nil {
		log.Printf("error encoding new task: %v\n", err)
	}
	return successfulResponse(string(jsonTaskToAdd))
}

func updateStatusHandler(id string) (events.APIGatewayProxyResponse, error) {
	filter := getIdFilter(id)
	updateTaskOperation := bson.D{{Key: "$set", Value: bson.D{primitive.E{Key: "completed", Value: true}, primitive.E{Key: "updated_at", Value: time.Now()}}}}
	_, err := collection.UpdateOne(ctx, filter, updateTaskOperation)
	if err != nil {
		return errorResponse(http.StatusInternalServerError, &ErrorMessage{Error: "error updating task."})
	}
	return successfulResponse(id)
}

func deleteTaskHandler(id string) (events.APIGatewayProxyResponse, error) {
	filter := getIdFilter(id)
	_, err := collection.DeleteOne(ctx, filter)
	if err != nil {
		return errorResponse(http.StatusInternalServerError, &ErrorMessage{Error: "error deleting task."})
	}
	return successfulResponse(id)
}

func getIdFilter(id string) primitive.D {
	idPrimitive, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Printf("error creating mongo object id: %v\n", err)
	}
	idFilter := bson.D{{Key: "_id", Value: idPrimitive}}
	return idFilter
}

func successfulResponse(responseBody string) (events.APIGatewayProxyResponse, error) {
	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusOK,
		Headers:    headers,
		Body:       responseBody,
	}, nil
}

func errorResponse(statusCode int, message *ErrorMessage) (events.APIGatewayProxyResponse, error) {
	jsonMessage, err := json.Marshal(message)
	if err != nil {
		log.Printf("error marshaling error message %v\n", err)
	}
	return events.APIGatewayProxyResponse{
		StatusCode: statusCode,
		Headers:    headers,
		Body:       string(jsonMessage),
	}, nil
}
