package util

import(
	"time"
	"math/rand"
)

func RandomString(n int) string{
	letters:=[]byte("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	result:=make([]byte,n)
	rand.Seed(time.Now().Unix())
	for i:=range result{
		result[i]=letters[rand.Intn(len(letters))]
	}
	return string(result)
}
