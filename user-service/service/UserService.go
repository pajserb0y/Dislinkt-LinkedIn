package service

import (
	"math/rand"
	"strconv"
	"time"
	"user-service/model"
	"user-service/repo"

	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

type Credentials struct {
	Password string `json:"password"`
	Username string `json:"username"`
}

type Claims struct {
	Username string `json:"username"`
	Id       string `json:"id"`
	Role     string `json:"role"`
	jwt.StandardClaims
}

var jwtKey = []byte("tajni_kljuc_za_jwt_hash")

type UserService struct {
	userRepo *repo.UserRepository
}

func New() (*UserService, error) {

	userRepo, err := repo.New()
	if err != nil {
		return nil, err
	}

	return &UserService{
		userRepo: userRepo,
	}, nil
}

func (s *UserService) SearchUsers(username string) []model.User {
	return s.userRepo.SearchUsers(username)
}

func (s *UserService) GetByID(id int) model.User {
	return s.userRepo.GetByID(id)
}

func (s *UserService) GetByUsername(username string) model.User {
	return s.userRepo.GetByUsername(username)
}

func (s *UserService) GetMe(id int) model.User {
	return s.userRepo.GetMe(id)
}

func (s *UserService) RemoveExperience(id int) int {
	return s.userRepo.RemoveExperience(id)
}

func (s *UserService) RemoveInterest(id int) int {
	return s.userRepo.RemoveInterest(id)
}

func (s *UserService) CloseDB() error {
	return s.userRepo.Close()
}

func (s *UserService) CreateUser(name string, email string, password string, username string, gender model.Gender, phonenumber string, dateofbirth time.Time, biography string) int {
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), 8)
	apiKey, _ := bcrypt.GenerateFromPassword([]byte(GenerateRandomString(10)), 8)
	return s.userRepo.CreateUser(name, email, string(hashedPassword), username, gender, phonenumber, dateofbirth, biography, string(apiKey))
}

func (s *UserService) AddInterest(interest string, userId uint) int {

	return s.userRepo.AddInterest(interest, userId)
}

func (s *UserService) AddExperience(company string, position string, from time.Time, until time.Time, userId uint) int {

	return s.userRepo.AddExperience(company, position, from, until, userId)
}

func (s *UserService) UpdateUser(id uint, name string, email string, password string, username string, gender model.Gender, phonenumber string, dateofbirth time.Time, biography string, isPublic bool) int {

	return s.userRepo.UpdateUser(id, name, email, password, username, gender, phonenumber, dateofbirth, biography, isPublic)
}

func (s *UserService) Login(username string, password string) string {
	expectedPassword := s.GetByUsername(username).Password
	id := strconv.FormatUint(uint64(s.GetByUsername(username).ID), 10)
	err := bcrypt.CompareHashAndPassword([]byte(expectedPassword), []byte(password))
	if err != nil {
		return ""
	}
	expirationTime := time.Now().Add(15 * time.Minute)

	claims := &Claims{
		Username: username,
		Id:       id,
		Role:     "user",
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return ""
	}
	return tokenString
}

func (s *UserService) BlockUser(userId int, blockedUserId int) {
	user := s.GetByID(userId)
	block := s.GetByID(blockedUserId)
	if userId == blockedUserId || user.ID == 0 || block.ID == 0 {
		return
	}
	s.userRepo.BlockUser(userId, blockedUserId)
	return
}

func (s *UserService) GetApiKeyForUserCredentials(username string, password string) string {
	user := s.GetByUsername(username)
	expectedPassword := user.Password
	err := bcrypt.CompareHashAndPassword([]byte(expectedPassword), []byte(password))
	if err != nil {
		return ""
	}
	return user.ApiKey
}

func (s *UserService) CreateJobOffer(id int, info string, position string, companyName string, qualifications string, key string) {
	s.userRepo.CreateJobOffer(id, info, position, companyName, qualifications, key)
	return
}

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func GenerateRandomString(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}
