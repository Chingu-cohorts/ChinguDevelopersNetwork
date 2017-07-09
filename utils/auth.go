package utils

import (
	"crypto/rsa"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/Chingu-cohorts/ChinguCentral/models"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/dgrijalva/jwt-go/request"
	"github.com/julienschmidt/httprouter"
)

var (
	privateKey *rsa.PrivateKey
	publicKey  *rsa.PublicKey
)

func init() {
	privateKeyFile, err := ioutil.ReadFile("private.rsa")
	if err != nil {
		log.Fatalf("Couldn't read private key file: %v", err)
	}

	publicKeyFile, err := ioutil.ReadFile("public.rsa.pub")
	if err != nil {
		log.Fatalf("Couldn't read public key file: %v", err)
	}

	privateKey, err = jwt.ParseRSAPrivateKeyFromPEM(privateKeyFile)
	if err != nil {
		log.Fatalf("Error parsing private key file: %v", err)
	}

	publicKey, err = jwt.ParseRSAPublicKeyFromPEM(publicKeyFile)
	if err != nil {
		log.Fatalf("Error parsing public key file: %v", err)
	}
}

// GenerateJWT generates a JWT token for a given user
func GenerateJWT(user models.User) string {
	claims := models.Claim{
		User: user,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 744).Unix(),
			Issuer:    "ChinguCentral",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)
	result, err := token.SignedString(privateKey)
	if err != nil {
		log.Fatalf("Couldn't sign the token: %v", err)
	}

	return result
}

// ReadJWT reads the data from the JWT token
func ReadJWT(token string) (*models.Claim, error) {
	authToken, _ := jwt.ParseWithClaims(token, &models.Claim{}, func(token *jwt.Token) (interface{}, error) {
		return nil, errors.New("Error fetching claims")
	})

	authTokenClaims, ok := authToken.Claims.(*models.Claim)
	if !ok {
		return nil, errors.New("Error fetching claims")
	}

	return authTokenClaims, nil
}

// AuthRequest verifies if the JWT token is valid
func AuthRequest(handleFunc httprouter.Handle) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		token, err := request.ParseFromRequestWithClaims(r, request.OAuth2Extractor, &models.Claim{}, func(token *jwt.Token) (interface{}, error) {
			return publicKey, nil
		})

		if err != nil {
			switch err.(type) {
			case *jwt.ValidationError:
				vErr := err.(*jwt.ValidationError)
				switch vErr.Errors {
				case jwt.ValidationErrorExpired:
					JSONMessage(w, "Expired token", http.StatusBadRequest)
					return
				case jwt.ValidationErrorSignatureInvalid:
					JSONMessage(w, "Invalid signature", http.StatusBadRequest)
					return
				default:
					JSONMessage(w, "Invalid token", http.StatusBadRequest)
					return
				}
			default:
				JSONMessage(w, "Invalid token", http.StatusBadRequest)
				return
			}
		}

		if token.Valid {
			handleFunc(w, r, ps)
		} else {
			http.Error(w, "Error: "+http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}
	}
}

// ValidateToken verifies if a given token is valid
func ValidateToken(token string) bool {
	authToken, err := jwt.ParseWithClaims(token, &models.Claim{}, func(token *jwt.Token) (interface{}, error) {
		return publicKey, nil
	})

	if err != nil {
		return false
	}

	if authToken.Valid {
		return true
	}

	return false
}
