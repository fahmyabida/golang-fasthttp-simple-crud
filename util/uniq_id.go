package util

import (
	"math/rand"
	"strings"
	"time"
)

func CreateTraceID() string {
	timeNow := time.Now().Format("060102150405.000")
	traceID := strings.Replace(timeNow, ".", "", 1) + CreateRandomHex(5)
	return traceID
}

func CreateRandomHex(n int) string {

	const letterBytes = "ABCDEF0123456789"
	const (
		letterIdxBits = 4                    // 4 bits to represent a letter index
		letterIdxMask = 1<<letterIdxBits - 1 // All 1-bits, as many as letterIdxBits
		letterIdxMax  = 63 / letterIdxBits   // # of letter indices fitting in 63 bits
	)

	var src = rand.NewSource(time.Now().UnixNano())

	// RandStringBytesMaskImprSrc ...
	// Src: https://stackoverflow.com/a/31832326/710955
	b := make([]byte, n)
	// A src.Int63() generates 63 random bits, enough for letterIdxMax characters!
	for i, cache, remain := n-1, src.Int63(), letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = src.Int63(), letterIdxMax
		}
		if idx := int(cache & letterIdxMask); idx < len(letterBytes) {
			b[i] = letterBytes[idx]
			i--
		}
		cache >>= letterIdxBits
		remain--
	}

	return string(b)
}
