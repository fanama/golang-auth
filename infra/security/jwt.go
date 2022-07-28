package security

import (
	"fanama/auth/domain"
	"log"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	jwt "github.com/golang-jwt/jwt/v4"
)

func VerifyToken(tokenString string) error {
	claims, err := ExtractClaims(tokenString)

	if err != nil {
		return err
	}

	_, err = jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_SECRET")), nil
	})

	return err
}

func ExtractClaims(tokenStr string) (jwt.MapClaims, error) {
	hmacSecret := []byte(os.Getenv("JWT_SECRET"))
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		// check token signing method etc
		return hmacSecret, nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	} else {
		log.Printf("Invalid JWT Token")
		return nil, err
	}
}

func GenerateToken(user domain.User) (string, string, error) {
	claims := jwt.MapClaims{
		"iss": user.Name,
		"exp": time.Now().Add(time.Hour * 15).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Generate encoded token
	t, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))

	if err != nil {
		log.Fatalf("An Error Occured %v", err)
	}

	rtClaims := jwt.MapClaims{
		"iss": user.Name,
		"exp": time.Now().Add(time.Hour * time.Duration(24)).Unix(),
	}

	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, rtClaims)

	rt, err := refreshToken.SignedString([]byte(os.Getenv("JWT_SECRET")))

	if err != nil {
		log.Fatalf("An Error Occured %v", err)
	}

	return t, rt, err
}

func RenewToken(refreshToken string) (string, string, error) {
	err := VerifyToken(refreshToken)
	if err != nil {
		return "", "", err
	}

	rtClaims, err := ExtractClaims(refreshToken)

	if err != nil {
		return "", "", err
	}
	// Create token

	claims := jwt.MapClaims{
		"iss": rtClaims["iss"],
		"exp": time.Now().Add(time.Hour * 15).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Generate encoded token
	t, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))

	if err != nil {
		log.Fatalf("An Error Occured %v", err)
	}

	// Create refresh token
	newRtClaims := jwt.MapClaims{
		"iss": rtClaims["iss"],
		"exp": time.Now().Add(time.Hour * time.Duration(24)).Unix(),
	}

	refreshTkn := jwt.NewWithClaims(jwt.SigningMethodHS256, newRtClaims)

	rt, _ := refreshTkn.SignedString([]byte(os.Getenv("JWT_SECRET")))

	return t, rt, nil
}

func GetAuthCookies(accessToken, refreshToken string) (*fiber.Cookie, *fiber.Cookie) {
	accessCookie := &fiber.Cookie{
		Name:     "access_token",
		Value:    accessToken,
		Expires:  time.Now().Add(24 * time.Hour),
		HTTPOnly: true,
		Secure:   true,
	}

	refreshCookie := &fiber.Cookie{
		Name:     "refresh_token",
		Value:    refreshToken,
		Expires:  time.Now().Add(2 * 24 * time.Hour),
		HTTPOnly: true,
		Secure:   true,
	}

	return accessCookie, refreshCookie
}

func InvalidToken(accessToken, refreshToken string) (*fiber.Cookie, *fiber.Cookie) {
	accessCookie := &fiber.Cookie{
		Name:     "access_token",
		Value:    accessToken,
		Expires:  time.Now().Add(0),
		HTTPOnly: true,
		Secure:   true,
	}

	refreshCookie := &fiber.Cookie{
		Name:     "refresh_token",
		Value:    refreshToken,
		Expires:  time.Now().Add(0),
		HTTPOnly: true,
		Secure:   true,
	}

	return accessCookie, refreshCookie
}
