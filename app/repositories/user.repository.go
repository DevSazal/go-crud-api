package repositories

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/DevSazal/go-crud-api/app/models"
)

type UserRepository struct {
	collection *mongo.Collection
}

func NewUserRepository(client *mongo.Client) *UserRepository {
	db := client.Database("test")
	collection := db.Collection("users")
	return &UserRepository{collection: collection}
}

func (r *UserRepository) GetUsers() ([]models.User, error) {
	var users []models.User


	// // Debugging the query
	// query := bson.D{{}} // Empty filter for getting all documents
	// log.Println(query)
	// queryBytes, _ := bson.Marshal(query)
	// log.Println("Query:", string(queryBytes))
	// // Debug End


	cur, err := r.collection.Find(context.Background(), nil)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	err = cur.All(context.Background(), &users)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return users, nil
}

func (r *UserRepository) GetUser(id primitive.ObjectID) (*models.User, error) {
	user := new(models.User)

	err := r.collection.FindOne(context.Background(), models.User{ID: id}).Decode(user)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return user, nil
}

func (r *UserRepository) CreateUser(user *models.User) error {
	_, err := r.collection.InsertOne(context.Background(), user)
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}

func (r *UserRepository) UpdateUser(id primitive.ObjectID, user *models.User) error {
	update := bson.M{
		"$set": bson.M{
			"name":     user.Name,
			"email":    user.Email,
			"password": user.Password,
		},
	}

	_, err := r.collection.UpdateOne(context.Background(), models.User{ID: id}, update)
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}

func (r *UserRepository) DeleteUser(id primitive.ObjectID) error {
	_, err := r.collection.DeleteOne(context.Background(), models.User{ID: id})
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}
