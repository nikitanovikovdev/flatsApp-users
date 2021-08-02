package users

import (
	"context"
	"crypto/sha1"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/nikitanovikovdev/flatsApp-users/pkg/platform/user"
	"github.com/spf13/viper"
	"log"
	"net/http"
	"time"
)

var m = make(map[string]string)

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

func (s *Service) CreateUser(ctx context.Context, user user.User) (interface{}, error) {

	user.Password = generatePasswordHash(user.Password)

	return s.repo.CreateUser(ctx, user)
}

func (s *Service) GenerateToken(ctx context.Context, user user.User) (string, error) {
	if err := initConfig(); err != nil {
		fmt.Errorf("error connection to config : %v", err)
	}

	user, err := s.repo.GetUser(ctx, user.Username, generatePasswordHash(user.Password))
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

	m["token"], err = token.SignedString([]byte(viper.GetString("keys.signing_key")))
	if err != nil {
		log.Fatal(err)
	}

	return token.SignedString([]byte(viper.GetString("keys.signing_key")))
}

func(s *Service) GetCookie() (http.Cookie, error) {
	expirationTime := time.Now().Add(1 *time.Hour)

	return http.Cookie{
		Name: "token",
		Value: m["token"],
		Expires: expirationTime,
	}, nil
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