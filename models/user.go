package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID                    primitive.ObjectID `json:"id" form:"id" bson:"_id,omitempty"`
	Email                 string             `json:"email" bson:"email" validate:"email"`
	Balance               int                `json:"balance" bson:"balance"`
	TelegramUserID        int64              `json:"telegramUserID" bson:"_telegramUserID,omitempty"`
	IsBot                 bool               `json:"isBot" bson:"isBot" validate:"required"`
	FirstName             string             `json:"firstName" bson:"firstName" validate:"required"`
	LastName              string             `json:"lastName" bson:"lastName" validate:"required"`
	Username              string             `json:"username" bson:"username" validate:"required"`
	LanguageCode          string             `json:"languageCode" bson:"languageCode" validate:"required"`
	IsPremium             bool               `json:"isPremium" bson:"isPremium"`
	AddedToAttachmentMenu bool               `json:"addedToAttachmentMenu" bson:"addedToAttachmentMenu"`
	AllowsWriteToPm       bool               `json:"allowsWriteToPm" bson:"allowsWriteToPm"`
	PhotoUrl              string             `json:"photoUrl" bson:"photoUrl"`
}

// type User struct {
// 	ID primitive.ObjectID `json:"id" form:"id" bson:"_id,omitempty"`
// 	// CompanyID  primitive.ObjectID `json:"companyID" form:"companyID" bson:"_companyID,omitempty"`
// 	*UserInput `bson:",inline"`
// }

// type UserInput struct {
// 	FirstName string `json:"firstName" bson:"firstName" validate:"required"`
// 	LastName  string `json:"lastName" bson:"lastName" validate:"required"`
// 	Email     string `json:"email" bson:"email" validate:"required,email"`
// 	Password  string `json:"password,omitempty"  bson:"password" validate:"required"`
// }

type UserUpdate struct {
	*UserInputUpdate `bson:",inline"`
	ID               primitive.ObjectID `json:"id" form:"id" bson:"_id,omitempty"`
}

type UserChangePassword struct {
	*ChangeUserPassword `bson:",inline"`
	ID                  primitive.ObjectID `json:"id" form:"id" bson:"_id,omitempty"`
}

type UserInputUpdate struct {
	FirstName       string `json:"firstName" xml:"firstName" form:"firstName" bson:"firstName" validate:"required"`
	LastName        string `json:"lastName" xml:"lastName" form:"lastName" bson:"lastName" validate:"required"`
	CurrentPassword string `json:"currentPassword,omitempty" form:"currentPassword" xml:"currentPassword,omitempty" bson:"currentPassword" validate:"required"`
	CompanyCategory string `json:"companyCategory" bson:"companyCategory"`
}

type ChangeUserPassword struct {
	CurrentPassword string `json:"currentPassword,omitempty" form:"currentPassword" xml:"currentPassword,omitempty" bson:"currentPassword" validate:"required"`
	Password        string `json:"password" form:"password" xml:"password" bson:"password" validate:"required"`
}

type LoginInput struct {
	Email string `json:"email" bson:"email" validate:"required,email"`
	// Phone    string `json:"phone" form:"phone" bson:"phone" validate:"required"`
	Password string `json:"password" form:"password" bson:"password" validate:"required"`
}

type UserDropDown struct {
	ID        primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	FirstName string             `json:"firstName" bson:"firstName"`
}

type AllDropDownUsers struct {
	Data []UserDropDown `json:"data" xml:"data"`
}
