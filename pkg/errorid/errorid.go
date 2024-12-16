package errorid

import "math/rand"

func GenErrorID() string {
	c := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890"
	var res string
	for range 8 {
		res += string(c[rand.Intn(len(c))])
	}
	return res
}
