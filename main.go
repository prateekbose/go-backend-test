package main

import (
	"context"
	"crypto/sha1"
	"errors"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"crypto/sha1"
	"encoding/hex"

	"encoding/json"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type users struct {
	Id       string
	Name     string
	Email    string
	Password string
}

type posts struct {
	UserId   string
	Id       string
	Caption  string
	ImageURL string
	Time     string
}

func CreateUserAndPost(w http.ResponseWriter, r *http.Request) {

	if r.Method != "POST" {
		err := errors.New("only POST Method is supported")
		w.Write([]byte(err.Error()))
		return
	}

	var Path string = strings.Split(r.URL.Path, "/")[1]

	client, err := mongo.NewClient(options.Client().ApplyURI("YOUR_MONGODB_ATLAS_CONNECT_STRING_HERE"))
	if err != nil {
		log.Fatal(err)
		w.WriteHeader(http.StatusServiceUnavailable)
		w.Write([]byte("503 - Service Unavailable"))
		return
	}

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
		w.WriteHeader(http.StatusServiceUnavailable)
		w.Write([]byte("503 - Service Unavailable"))
		return
	}

	if Path == "users" {

		collection := client.Database("task").Collection("users")

		tempPass := r.FormValue("Password")
		hasher := sha1.New()
		hasher.Write([]byte(tempPass))
		sha1_hash := hex.EncodeToString(hasher.Sum(nil))

		temp := users{}
		temp.Id = r.FormValue("Id")
		temp.Name = r.FormValue("Name")
		temp.Email = r.FormValue("Email")
		temp.Password = sha1_hash
		fmt.Println("POST: user")
		fmt.Println(temp)

		res, err := collection.InsertOne(ctx, temp)
		if err != nil {
			log.Fatal(err)
			w.WriteHeader(http.StatusServiceUnavailable)
			w.Write([]byte("503 - Service Unavailable"))
		}
		fmt.Println(res)
		w.WriteHeader(http.StatusCreated)
		w.Write([]byte("201 - Created"))
	} else if Path == "posts" {

		temp := posts{}
		temp.UserId = r.FormValue("UserId")
		temp.Id = r.FormValue("Id")
		temp.Caption = r.FormValue("Caption")
		temp.ImageURL = r.FormValue("ImageURL")
		temp.Time = time.Now().String()
		fmt.Println("POST: posts")
		fmt.Println(temp)

		collection := client.Database("task").Collection("users")

		if err = collection.FindOne(ctx,
			bson.M{"id": temp.UserId},
			options.FindOne().SetProjection(bson.M{"_id": 0}),
		).Decode(&temp); err != nil {
			fmt.Println(err)
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("401 - Unauthorized"))
			client.Disconnect(ctx)
			return
		}

		collection = client.Database("task").Collection("posts")

		res, err := collection.InsertOne(ctx, temp)
		if err != nil {
			log.Fatal(err)
			w.WriteHeader(http.StatusServiceUnavailable)
			w.Write([]byte("503 - Service Unavailable"))
		}
		fmt.Println(res)
		w.WriteHeader(http.StatusCreated)
		w.Write([]byte("201 - Created"))
	}

	defer client.Disconnect(ctx)
}

func GetUserAndPost(w http.ResponseWriter, r *http.Request) {

	if r.Method != "GET" {
		err := errors.New("only GET Method is supported")
		w.Write([]byte(err.Error()))
		return
	}

	var Id string = strings.Split(r.URL.Path, "/")[2]
	var Path string = strings.Split(r.URL.Path, "/")[1]
	fmt.Printf("GET: %s\nID: %s\n", Path, Id)

	client, err := mongo.NewClient(options.Client().ApplyURI("YOUR_MONGODB_ATLAS_CONNECT_STRING_HERE"))
	if err != nil {
		log.Fatal(err)
		w.WriteHeader(http.StatusServiceUnavailable)
		w.Write([]byte("503 - Service Unavailable"))
		return
	}

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
		w.WriteHeader(http.StatusServiceUnavailable)
		w.Write([]byte("503 - Service Unavailable"))
		return
	}

	if Path == "users" {
		collection := client.Database("task").Collection("users")

		var temp users
		if err = collection.FindOne(ctx,
			bson.M{"id": Id},
			options.FindOne().SetProjection(bson.M{"_id": 0}),
		).Decode(&temp); err != nil {
			fmt.Println(err)
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte("404 - Not Found"))
			client.Disconnect(ctx)
			return
		}

		bytes, err := json.Marshal(temp)

		if err != nil {
			log.Fatal(err)
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("500 - Internal Server Error"))
		}

		fmt.Println("Fetching: SUCCESS")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(bytes))
	} else if Path == "posts" {
		collection := client.Database("task").Collection("posts")

		var temp posts
		if err = collection.FindOne(ctx,
			bson.M{"id": Id},
			options.FindOne().SetProjection(bson.M{"_id": 0}),
		).Decode(&temp); err != nil {
			fmt.Println(err)
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte("404 - Not Found"))
			client.Disconnect(ctx)
			return
		}

		bytes, err := json.Marshal(temp)

		if err != nil {
			log.Fatal(err)
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("500 - Internal Server Error"))
		}

		fmt.Println("Fetching: SUCCESS")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(bytes))
	}

	defer client.Disconnect(ctx)
}

func GetUserPosts(w http.ResponseWriter, r *http.Request) {

	if r.Method != "GET" {
		err := errors.New("only GET Method is supported")
		w.Write([]byte(err.Error()))
		return
	}
	var Id string = strings.Split(r.URL.Path, "/")[3]
	fmt.Printf("GET: User Posts\nID: %s\n", Id)

	client, err := mongo.NewClient(options.Client().ApplyURI("YOUR_MONGODB_ATLAS_CONNECT_STRING_HERE"))
	if err != nil {
		log.Fatal(err)
		w.WriteHeader(http.StatusServiceUnavailable)
		w.Write([]byte("503 - Service Unavailable"))
		return
	}

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
		w.WriteHeader(http.StatusServiceUnavailable)
		w.Write([]byte("503 - Service Unavailable"))
		return
	}

	collection := client.Database("task").Collection("posts")

	var result []posts

	cursor, err := collection.Find(context.TODO(), bson.M{"userid": Id})
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("404 - Not Found"))
		client.Disconnect(ctx)
		return
	}

	if err = cursor.All(context.TODO(), &result); err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("404 - Not Found"))
		client.Disconnect(ctx)
		return
	}

	bytes, err := json.Marshal(result)

	fmt.Println(result)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(bytes))
}

func main() {
	fmt.Println("Serving on PORT: 3000")

	http.HandleFunc("/users", CreateUserAndPost)
	http.HandleFunc("/posts", CreateUserAndPost)

	http.HandleFunc("/users/", GetUserAndPost)
	http.HandleFunc("/posts/", GetUserAndPost)

	http.HandleFunc("/posts/users/", GetUserPosts)

	log.Fatal(http.ListenAndServe(":3000", nil))
}
