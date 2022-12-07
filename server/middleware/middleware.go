package middleware

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var collection *mongo.collection

func init(){
	loadTheEnv()
	createDBInstance()
}


func loadTheEnv(){
	 err := godotenv.Load(".env")
		if err!= nil{
			log.Fatal("Error loading the .env file")
		}
}

func createDBInstance(){
	connectionString := os.Getenv("DB_URI")
	dbName := os.Getenv("DB_NAME")
	collName := os.Getenv("DB_COLLECTION_NAME")

	clientOptions := options.Client().ApplyURL(connectionString)

	client, err := mongo.Connect(context.TODO(), clientOptions)
    
	if err !=nil{
		log.Fatal(err)
	}

	err = client.Ping(context.TODO(), nil)
	if err!=nil{
		log.Fatal(err)
	}

	fmt.Println("connect to mongoDb!")

	collection = client.Database(dbName).collection(collName)
	fmt.Println("collection instance created")
}

func GetAllTasks(w http.ResponseWriter, r *http.Request){
      w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	  w.Header().Set("Access-Control-Allow-Origin", "")
	  payload := getAllTasks()
	json.NewEncoder(w).Encode(payload)  
}

func CreateTask(w http.ResponseWriter, r *http.Request){
    w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST")
    var task models.ToDoList
	json.NewDecoder(r.Body).Decode(&task)
	insertOneTask(task)
	json.NewEncoder(w).Encode(task)
}

func TaskComplete(w http.ResponseWriter, r *http.Request){
    w.Header().Set("Content-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "PUT")
	w.Header().Set("Access-Control-Allow_Headers", "Content-Type" )

	params:= mux.Vars(r)
	taskComplete(params["id"])
	json.NewEncoder(w).Encode(params["id"])
}

func undoTask(w http.ResponseWriter, r *http.Request){
     w.Header().Set("Content-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "PUT")
	w.Header().Set("Access-Control-Allow_Headers", "Content-Type" )

	params:= mux.Vars(r)
	UndoTask(params["id"])
	json.NewEncoder(w).Encode(params["id"])

}

func DeleteTask(w http.ResponseWriter, r *http.Request){
 w.Header().Set("Content-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "DELETE")
	w.Header().Set("Access-Control-Allow_Headers", "Content-Type" )

	params := mux.Vars(r)
	deleteOneTask(params["id"])
}

func DeleteAllTasks(w http.ResponseWriter, r *http.Request){
     w.Header().Set("Content-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	count := deleteAllTasks()
	json.NewEncoder(w).Encode(count)
}

func getAllTasks() []primitive.M{
  cur, err := collection.Find(context.Background(), bson.D{{}})
	if err!=nil{
		log.Fatal(err)
	}

	var results []primitive.M
	for cur.Next(context.Background()){
		var result bson.M
		e := cur.Decode(&result)
		if e !=nil{
			log.Fatal(e)
		}
		results = append(results, result)
	}
	if err := cur.Err(); err != nil{

		log.Fatal(err)
	}
	cur.Close(context.Background())
	return results
}

func taskComplete(task string){
     id, _ := primitive.ObjectIDFromHex(task)
	 filter :=bson.M{"id": id}
	 update := bson.M{"$set":bson.M{"status": true}}
	 result, err := collection.updateOne(context.Background(), filter, update)
	 if err!=nil{
		log.Fatal(err)
	 }
	 fmt.Println("modified count:", result.ModifiedCount)
}

func insertOneTask(){

}

func UndoTask(){

}

func deleteOneTask(){

}

func deleteAllTasks(){

}