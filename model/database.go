package model

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
)

type User struct {
	tableName struct{} `pg:"public.user"`
	Id        string
	FirstName string
	LastName  string
	Username  string
	Password  string
	Secret    string
}

func (u *User) IsValidPassword(password string) bool {
	h := hmac.New(sha256.New, []byte(u.Secret))
	h.Write([]byte(password))
	sha := hex.EncodeToString(h.Sum(nil))
	return sha == u.Password
}

func (u *User) EncrpytPasswordWithHashHMAC_SHA256() {
	h := hmac.New(sha256.New, []byte(u.Secret))
	h.Write([]byte(u.Password))
	u.Password = hex.EncodeToString(h.Sum(nil))
}
