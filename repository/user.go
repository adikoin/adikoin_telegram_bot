package repository

import (
	"context"

	model "telegram_bot/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var cntx context.Context = context.TODO()

type UserRepository interface {
	FindByTelegramUserID(id int64) (*model.User, error)
	SaveUser(user *model.User) (*model.User, error)
	// GetAllUser(page int64, limit int64) (*model.PagedUser, error)
	// GetDropdownUsers() (*model.AllDropDownUsers, error)

	// FindByEmail(email string) (*model.User, error)
	// FindByPhone(phone string) (*model.User, error)
	// FindById(id string) (*model.User, error)
	// UpdateUser(id string, user *model.UserUpdate) (*model.UserUpdate, error)
	// DeleteUser(id string) error
	// ChangeUserPassword(id string, user *model.UserChangePassword) (*model.UserChangePassword, error)
}

type userRepositoryImpl struct {
	Connection *mongo.Database
}

func NewUserRepository(Connection *mongo.Database) UserRepository {
	return &userRepositoryImpl{Connection: Connection}
}

func (userRepository *userRepositoryImpl) FindByTelegramUserID(id int64) (*model.User, error) {

	var existingUser model.User

	filter := bson.M{"_telegramUserID": id}

	err := userRepository.Connection.Collection("users").FindOne(cntx, filter).Decode(&existingUser)

	if err != nil {
		return nil, err
	}

	return &existingUser, nil
}

func (userRepository *userRepositoryImpl) SaveUser(user *model.User) (*model.User, error) {

	_, err := userRepository.Connection.Collection("users").InsertOne(cntx, user)
	if err != nil {
		return nil, err
	}

	return user, nil
}

// func (userRepository *userRepositoryImpl) FindByEmail(email string) (*model.User, error) {
// 	var existingUser model.User
// 	filter := bson.M{"email": email}
// 	err := userRepository.Connection.Collection("users").FindOne(cntx, filter).Decode(&existingUser)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return &existingUser, nil
// }

// func (userRepository *userRepositoryImpl) FindByPhone(phone string) (*model.User, error) {
// 	var existingUser model.User
// 	filter := bson.M{"phone": phone}
// 	err := userRepository.Connection.Collection("users").FindOne(cntx, filter).Decode(&existingUser)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return &existingUser, nil
// }

// func (userRepository *userRepositoryImpl) FindById(id string) (*model.User, error) {
// 	var existingUser model.User
// 	objectId, _ := primitive.ObjectIDFromHex(id)
// 	filter := bson.M{"_id": objectId}
// 	err := userRepository.Connection.Collection("users").FindOne(cntx, filter).Decode(&existingUser)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return &existingUser, nil
// }

// func (userRepository *userRepositoryImpl) UpdateUser(id string, user *model.UserUpdate) (*model.UserUpdate, error) {

// 	objectId, _ := primitive.ObjectIDFromHex(id)

// 	filter := bson.M{"_id": objectId}
// 	update := bson.M{
// 		"$set": bson.M{
// 			"firstName": user.FirstName,
// 			"lastName":  user.LastName,
// 		},
// 	}

// 	result, err := userRepository.Connection.Collection("users").UpdateMany(cntx, filter, update)
// 	if err != nil {
// 		return nil, err
// 	}
// 	if result.MatchedCount == 0 {
// 		return nil, exception.ResourceNotFoundException("User", "id", id)
// 	}

// 	user.ID = objectId
// 	user.CurrentPassword = ""
// 	return user, nil
// }

// func (userRepository *userRepositoryImpl) DeleteUser(id string) error {
// 	objectId, _ := primitive.ObjectIDFromHex(id)
// 	filter := bson.M{"_id": objectId}

// 	result, err := userRepository.Connection.Collection("users").DeleteOne(cntx, filter)
// 	if err != nil {
// 		return err
// 	}
// 	if result.DeletedCount == 0 {
// 		return exception.ResourceNotFoundException("User", "id", id)
// 	}

// 	return nil
// }

// func (userRepository *userRepositoryImpl) ChangeUserPassword(id string, user *model.UserChangePassword) (*model.UserChangePassword, error) {

// 	objectId, _ := primitive.ObjectIDFromHex(id)

// 	filter := bson.M{"_id": objectId}
// 	update := bson.M{
// 		"$set": bson.M{
// 			"password": user.Password,
// 		},
// 	}

// 	result, err := userRepository.Connection.Collection("users").UpdateOne(cntx, filter, update)
// 	if err != nil {
// 		return nil, err
// 	}
// 	if result.MatchedCount == 0 {
// 		return nil, exception.ResourceNotFoundException("User", "id", id)
// 	}

// 	user.ID = objectId
// 	user.CurrentPassword = ""
// 	return user, nil
// }

// func (userRepository *userRepositoryImpl) GetAllUser(page int64, limit int64) (*model.PagedUser, error) {
// 	var users []model.User

// 	filter := bson.M{}

// 	collection := userRepository.Connection.Collection("users")

// 	//	projection := bson.D{
// 	//		{"id", 1},
// 	//		{"firstName", 1},
// 	//		{"lastName", 1},
// 	//		{"email", 1},
// 	//	}
// 	//

// 	projection := bson.D{
// 		{
// 			Key:   "id",
// 			Value: 1,
// 		},
// 		{
// 			Key:   "firstName",
// 			Value: 1,
// 		}, {
// 			Key:   "lastName",
// 			Value: 1,
// 		},
// 		{
// 			Key:   "email",
// 			Value: 1,
// 		}}

// 	paginatedData, err := paginate.New(collection).Context(cntx).Limit(limit).Page(page).Select(projection).Filter(filter).Decode(&users).Find()
// 	if err != nil {
// 		return nil, err
// 	}

// 	return &model.PagedUser{
// 		Data:     users,
// 		PageInfo: paginatedData.Pagination,
// 	}, nil
// }

// func (userRepository *userRepositoryImpl) GetDropdownUsers() (*model.AllDropDownUsers, error) {

// 	filter := bson.M{}

// 	collection := userRepository.Connection.Collection("users")

// 	cursor, err := collection.Find(cntx, filter)
// 	if err != nil {
// 		return nil, err
// 	}

// 	var users []model.UserDropDown

// 	if err = cursor.All(cntx, &users); err != nil {
// 		log.Fatal(err)
// 	}

// 	return &model.AllDropDownUsers{
// 		Data: users,
// 	}, nil

// }
