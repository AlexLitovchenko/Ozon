package Handler

import (
	"math/rand"
	"time"
)

func CreateShortString() string {
	rand.Seed(time.Now().UnixNano())
	digits := "0123456789"
	specials := "_"
	low := "abcdefghijklmnopqrstuvwxyz"
	all := "ABCDEFGHIJKLMNOPQRSTUVWXYZ" + low + digits + specials
	length := 10
	buf := make([]byte, length)
	buf[0] = digits[rand.Intn(len(digits))]
	buf[1] = specials[rand.Intn(len(specials))]
	buf[2] = low[rand.Intn(len(low))]
	for i := 3; i < length; i++ {
		buf[i] = all[rand.Intn(len(all))]
	}
	rand.Shuffle(len(buf), func(i, j int) {
		buf[i], buf[j] = buf[j], buf[i]
	})
	str := string(buf)
	return str
}
