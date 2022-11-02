package models

import (
	"errors"
	"net/http"
	"time"

	"github.com/TendonT52/go-tendon-api/app/services/authService"
	"github.com/TendonT52/go-tendon-api/app/services/mongoService"
	"github.com/TendonT52/go-tendon-api/app/services/validateService"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	Admin   = "admin"
	Teacher = "teacher"
	Student = "student"
)

type User struct {
	UserId          primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	FirstName       string             `json:"firstName" bson:"firstName"`
	LastName        string             `json:"lastName" bson:"lastName"`
	Email           string             `json:"email" bson:"email"`
	HashPassword    string             `json:"password" bson:"password"`
	Role            string             `json:"role" bson:"role"`
	UpdatedAt       primitive.DateTime `json:"updated_at" bson:"updated_at"`
	BriefCurriculum []BriefCurriculum  `json:"briefCurriculum" bson:"briefCurriculum"`
}

type SignInUser struct {
	Email    string `json:"email" bson:"email"`
	Password string `json:"password" bson:"password"`
}

type SignUpUser struct {
	FirstName string `json:"firstName" bson:"firstName" binding:"required,alpha" validate:"required,alpha"`
	LastName  string `json:"lastName" bson:"lastName" binding:"required,alpha" validate:"required,alpha"`
	Email     string `json:"email" bson:"email" binding:"required,email" validate:"required,email"`
	Password  string `json:"password" bson:"password" binding:"required,gt=8" validate:"required,gt=8"`
}

func SignUp(ctx *gin.Context) *User {
	signUpUser := &SignUpUser{}
	user := &User{}
	err := ctx.ShouldBindJSON(signUpUser)
	if err != nil {
		message := validateService.TranslateError(signUpUser)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"ok":      false,
			"message": message,
		})
		ctx.Error(errors.New("error while bind JSON in newSignInUser"))
		ctx.Abort()
		return user
	}
	user = newUserFromSignInUser(signUpUser)
	mongoService.InsertDoc(user)	
	return user
}

func newUserFromSignInUser(signUpUser *SignUpUser) *User{
	return &User{
		UserId: primitive.NewObjectID(),
		FirstName: signUpUser.FirstName,
		LastName: signUpUser.FirstName,
		Email: signUpUser.Email,
		HashPassword: authService.HashPassword(signUpUser.Password),
		Role: Student,
		UpdatedAt: primitive.NewDateTimeFromTime(time.Now()),
		BriefCurriculum: []BriefCurriculum{},
	}
}

func (user *User) InsertMongo() (*mongo.Collection, interface{}) {
	return mongoService.AccountCollection, user
}