package web

import (
	"crypto/rand"
	"io"

	"github.com/dgraph-io/badger/v3"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type WxInfo struct {
	Name        string `json:"name"`
	OpenID      string `json:"openid"`
	SessionKey  string `json:"sessionkey"`
	AccessToken string `json:"accesstoken"`
}

type Token struct {
	AccessID    string `json:"accessID"`
	AccessToken string `json:"accesstoken"`
}

func (w *Web) GenerateAccessToken(name, openid, sessionkey string) (string, error) {
	token := &Token{}
	accessToken := make([]byte, 64)
	token.AccessID = uuid.NewString()

	if _, err := io.ReadFull(rand.Reader, accessToken); err != nil {
		return "", err
	}

	txn := w.cache.NewTransaction(true)
	err = txn.SetEntry(badger.NewEntry([]byte(token.AccessID), []byte(token.)))
}
func (w *Web) AdminAuth() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}
