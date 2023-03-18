package account

import (
	"ecommerce-backend/pkg/common/database"
	"ecommerce-backend/pkg/common/env"
	"ecommerce-backend/pkg/models"
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"time"
)

type controller interface {
	signIn(user SignInDTO) (string, error)
	signUp(user models.User) error
	getUserInformation(userId string) (UserResponseDTO, error)
}

type controllerImpl struct {
	repository repository
}

func newController() controller {
	return controllerImpl{
		repository: repositoryImpl{},
	}
}

func (c controllerImpl) signIn(user SignInDTO) (string, error) {
	var tokenString string
	err := database.GetInstance().Transaction(func(tx *gorm.DB) error {
		dbUser, err := c.repository.getUserByUsername(tx, user.Username)
		if err != nil {
			return errors.New("wrong username or password")
		}

		err = bcrypt.CompareHashAndPassword([]byte(dbUser.Password), []byte(user.Password))
		if err != nil {
			return errors.New("wrong username or password")
		}

		token := jwt.New(jwt.SigningMethodHS256)
		claims := token.Claims.(jwt.MapClaims)
		claims["exp"] = time.Now().Add(24 * time.Hour)
		claims["userId"] = dbUser.ID
		tokenString, err = token.SignedString([]byte(env.JWT_SECRET))

		return nil
	})
	return tokenString, err
}

func (c controllerImpl) signUp(user models.User) error {
	hashedBytes, err := bcrypt.GenerateFromPassword([]byte(user.Password), 12)
	if err != nil {
		return errors.New("cannot hash the password")
	}
	user.Password = string(hashedBytes)

	return c.repository.save(database.GetInstance(), user)
}

func (c controllerImpl) getUserInformation(userId string) (UserResponseDTO, error) {
	dbUser, err := c.repository.getUserByUserId(database.GetInstance(), userId)
	if err != nil {
		return UserResponseDTO{}, err
	}
	return UserResponseDTO{
		FirstName: dbUser.FirstName,
		LastName:  dbUser.LastName,
	}, nil
}
