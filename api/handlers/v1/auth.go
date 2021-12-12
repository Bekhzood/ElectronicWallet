package v1

import (
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/Bekhzood/ElectronicWallet/config"
	"github.com/Bekhzood/ElectronicWallet/storage"
	"github.com/gin-gonic/gin"
)

const (
	authorizationHeaderKey = "authorization"
)

func (h *Handler) DigestAuth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authorizationHeader := ctx.GetHeader(authorizationHeaderKey)

		if len(authorizationHeader) == 0 {
			err := errors.New("authorization header is not provided")
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, err)
			return
		}

		headers := digestParts(authorizationHeader)
		pwd, err := storage.NewAccountRepo(h.storagePostgres).GetPassword(headers["username"])
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, err)
			return
		}
		headers["password"] = pwd
		serverResponse := createServerResponse(headers)

		if !checkResponses(headers["response"], serverResponse) {
			err := errors.New("invalid username or password is provided")
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, err)
			return
		}

		fields := strings.Fields(authorizationHeader)
		if len(fields) < 6 {
			err := errors.New("invalid authorization header format")
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, err)
			return
		}

		ctx.Next()
	}
}

func checkResponses(userResponse, serverResponse string) bool {
	return userResponse == serverResponse
}

func createServerResponse(headers map[string]string) string {
	cfg := config.Load()
	serverHeaders := make(map[string]string)
	serverHeaders["username"] = headers["username"]
	serverHeaders["password"] = headers["password"]
	serverHeaders["uri"] = headers["uri"]
	serverHeaders["realm"] = cfg.Realm
	serverHeaders["nonce"] = cfg.Nonce
	serverHeaders["method"] = cfg.Method
	response := createDigestResponse(serverHeaders)
	return response
}

func digestParts(header string) map[string]string {
	result := map[string]string{}
	if len(header) > 0 {
		wantedHeaders := []string{"username", "nonce", "realm", "response", "uri"}
		responseHeaders := strings.Split(header, ",")
		for _, r := range responseHeaders {
			for _, w := range wantedHeaders {
				if strings.Contains(r, w) {
					result[w] = strings.Split(r, `"`)[1]
				}
			}
		}
	}
	return result
}

func getSha256(text string) string {
	hasher := sha256.New()
	hasher.Write([]byte(text))
	return hex.EncodeToString(hasher.Sum(nil))
}

func createDigestResponse(digestHeaders map[string]string) string {
	ha1 := getSha256(digestHeaders["username"] + ":" + digestHeaders["realm"] + ":" + digestHeaders["password"])
	ha2 := getSha256(digestHeaders["method"] + ":" + digestHeaders["uri"])
	response := getSha256(fmt.Sprintf("%s:%s:%s", ha1, digestHeaders["nonce"], ha2))
	return response
}
