package random

import (
	"math/rand"
	"strconv"
	"strings"
	"time"
)

var r *rand.Rand

const keyword = "ABCDEFGHIJKLNMOPQRSTUVWXYZabcdefghijklnmopqrstuvwxyz0123456789"

func RandomString(n int) string {
	var sb strings.Builder
	k := len(keyword)

	for i := 0; i < n; i++ {
		c := keyword[rand.Intn(k)]
		sb.WriteByte(c)
	}

	return sb.String()
}

func RandomInt(maxInt int) int {
	r = rand.New(rand.NewSource(time.Now().UnixNano()))
	return r.Intn(maxInt)
}

func RandomPinCode() string {
	var ints strings.Builder
	for i := 0; i < 6; i++ {
		s := RandomInt(9)
		ints.WriteString(strconv.Itoa(s))
	}
	return ints.String()
}
