package authentication

import (
	"ecommerce-backend/pkg/common/database"
	"ecommerce-backend/pkg/models"
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"os"
	"time"
)

type controller interface {
	signIn(user SignInDTO) (string, error)
	signUp(user models.User) error
}

type controllerImpl struct {
	repository repository
	jwtSecret  string
}

func newController() controller {
	return controllerImpl{
		repository: repositoryImpl{},
		jwtSecret:  os.Getenv("JWT_SECRET"),
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
		claims["username"] = dbUser.Username
		tokenString, err = token.SignedString([]byte(c.jwtSecret))

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
