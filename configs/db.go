package configs

import (
	"DavidJChavez/rw-demo-api/graph/model"
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type DB struct {
	client *mongo.Client
}

var Db *DB

func ConnectDB() {
	client, err := mongo.NewClient(options.Client().ApplyURI(os.Getenv("MDB_URI")))
	if err != nil {
		log.Fatal(err)
	}

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB!")
	Db = &DB{client: client}
}

func colHelper(db *DB, collectionName string) *mongo.Collection {
	return db.client.Database("rw-demo").Collection(collectionName)
}

func (db *DB) CreateUser(input *model.NewUser) (*model.User, error) {
	collection := colHelper(db, "users")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	input.CreatedAt = time.Now().Format(time.RFC3339Nano)

	res, err := collection.InsertOne(ctx, input)

	if err != nil {
		log.Fatal(err)
	}

	user := &model.User{
		ID:        res.InsertedID.(primitive.ObjectID).Hex(),
		FirstName: input.FirstName,
		LastName:  input.LastName,
		Email:     input.Email,
		Phone:     input.Phone,
		CreatedAt: input.CreatedAt,
	}

	return user, nil
}

func (db *DB) ListUsers() ([]*model.User, error) {
	collection := colHelper(db, "users")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		log.Fatal(err)
	}

	defer cursor.Close(ctx)

	var users []*model.User
	for cursor.Next(ctx) {
		var singleUser *model.User
		err = cursor.Decode(&singleUser)
		if err != nil {
			log.Fatal(err)
		}
		users = append(users, singleUser)
	}

	return users, nil
}
