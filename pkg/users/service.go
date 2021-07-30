package users

import (
	"context"
	"crypto/sha1"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/nikitanovikovdev/flatsApp-users/pkg/platform/user"
	"github.com/spf13/viper"
	"time"
)

type Service struct {
	repo *Repository
}

func NewService(r *Repository) *Service {
	return &Service{
		repo: r,
	}
}

type TokenClaims struct {
	jwt.StandardClaims
	Username string `json:"username" bson:"username"`
}
//
//func (s *Service) CreateUser(ctx context.Context, u []byte) (interface{}, error) {
//	var user user.User
//	if err := json.Unmarshal(u, &user); err != nil {
//		return user, err
//	}
//	user.Password = generatePasswordHash(user.Password)
//	return s.repo.CreateUser(ctx, user)
//}

func (s *Service) GenerateToken(ctx context.Context, username, password string) (string, error) {
	if err := initConfig(); err != nil {
		fmt.Errorf("error connection to config : %v", err)
	}

	var usr user.User
	usr.Username = username
	usr.Password = password

	user, err := s.repo.GetUser(ctx, usr.Username, generatePasswordHash(usr.Password))
	if err != nil {
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &TokenClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(1 * time.Hour).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		user.Username,
	})

	return token.SignedString([]byte(viper.GetString("keys.signing_key")))
}

func generatePasswordHash(pass string) string {
	hash := sha1.New()
	hash.Write([]byte(pass))

	return fmt.Sprintf("%x", hash.Sum([]byte(viper.GetString("keys.salt"))))
}

func initConfig() error {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")

	return viper.ReadInConfig()
}