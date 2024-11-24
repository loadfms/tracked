package customers

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"time"
	"tracked/internal/constants"

	"github.com/golang-jwt/jwt"
)

type Customer struct {
	PK       string `dynamodbav:"pk"`
	SK       string `dynamodbav:"sk"`
	Name     string `json:"name" dynamodbav:"name"`
	Email    string `json:"email" dynamodbav:"email"`
	Password string `json:"password" dynamodbav:"password"`
	Salt     string `dynamodbav:"salt"`
}

type Claims struct {
	Email        string `json:"email"`
	CustomerUUID string `json:"customer_uuid"`
	jwt.StandardClaims
}

func NewCustomer(name string, email string, password string) (*Customer, error) {
	salt, err := generateSalt()
	if err != nil {
		return nil, err
	}

	hashedPassword := hashPassword(password, salt)

	return &Customer{
		Name:     name,
		Email:    email,
		Password: hashedPassword,
		Salt:     salt,
		PK:       GeneratePKByEmail(email),
		SK:       GeneratePKByEmail(email),
	}, nil
}

func CheckPassword(customer *Customer, password string) bool {
	hashedPassword := hashPassword(password, customer.Salt)
	return hashedPassword == customer.Password
}

func generateSalt() (string, error) {
	salt := make([]byte, 16)
	_, err := rand.Read(salt)
	if err != nil {
		return "", err
	}

	return hex.EncodeToString(salt), nil
}

func GeneratePKByEmail(email string) string {
	hash := sha256.New()
	hash.Write([]byte(email + constants.EmailSalt))
	hashed := hash.Sum(nil)
	return fmt.Sprintf("CUSTOMER##%s", hex.EncodeToString(hashed)[:45])
}

func hashPassword(password, salt string) string {
	hash := sha256.New()
	hash.Write([]byte(password + salt))
	hashed := hash.Sum(nil)
	return hex.EncodeToString(hashed)
}

func GenerateJWTToken(customer *Customer) (string, error) {
	claims := &Claims{
		Email:        customer.Email,
		CustomerUUID: customer.PK,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(constants.JWTTokenSalt))

	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func GetCustomerUUIDFromToken(tokenString string) (string, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(constants.JWTTokenSalt), nil
	})

	if err != nil {
		return "", err
	}

	claims, ok := token.Claims.(*Claims)
	if !ok {
		return "", fmt.Errorf("failed to parse claims")
	}

	return claims.CustomerUUID, nil
}
