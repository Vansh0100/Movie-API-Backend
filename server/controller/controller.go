package controller

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/Vansh0100/movieapi/model"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var db *mongo.Database
var movieCollection *mongo.Collection



func Connection(connectionString string) {
	clientOptions:=options.Client().ApplyURI(connectionString)

	client,err:=mongo.Connect(context.TODO(),clientOptions)
	if err!=nil{
		log.Fatal(err)
	}
	fmt.Println("MongoDB Connection Successfull")

	Controller(client)
}




func Controller(client *mongo.Client)  {
	db=client.Database("MovieAPI")
	movieCollection=db.Collection("movies")

}
func getOneMovie(m map[string] string) model.Movie{
	var result model.Movie
	// filter:=model.Movie.Title["Brahmastra"]}
	filter:=bson.D{{"title",m["title"]}}

	data:=movieCollection.FindOne(context.Background(),filter).Decode(&result)
	if data!=nil{
		log.Fatal(data)
	}

	// finalJson,_:=json.MarshalIndent(result,"","\t")
	// fmt.Println("The data is:\n",string(finalJson))
	return result
}

func getAllMovies() []*model.Movie {
	var result []*model.Movie

	data,err:=movieCollection.Find(context.Background(),bson.D{{}})
	if err!=nil{
		log.Fatal(err)
	}

	for data.Next(context.Background()){
		var item model.Movie
		err=data.Decode(&item)
		if err!=nil{
			log.Fatal(err)
		}
		result = append(result, &item)
	}
	data.Close(context.Background())
	// finalJson,_:=json.MarshalIndent(result,"","\t")
	// fmt.Println(string(finalJson))
	return result
}

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	}


func HomePage(w http.ResponseWriter,r *http.Request){
	w.Write([]byte("<h1>Welcome to Golang API with MongoDB Connection</h1>"))
}
func GetAllMovies(w http.ResponseWriter,r *http.Request)  {
	enableCors(&w)
	var result []*model.Movie=getAllMovies()

	// finalJson,_:=json.MarshalIndent(result,"","\t")
	json.NewEncoder(w).Encode(result)
}

func GetOneMovie(w http.ResponseWriter, r *http.Request){
	params:=mux.Vars(r)
	var result model.Movie=getOneMovie(params)
	// finalJson,_:=json.MarshalIndent(result,"","\t");
	json.NewEncoder(w).Encode(result)
}