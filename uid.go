package uid

import (
	"crypto/rand"
	"strings"
	"time"
)

func HumanUid() string {
	r, _ := rand.Prime(rand.Reader, 64)
	id := time.Now().UTC().Format("20060102150405.0000000")
	id = strings.ReplaceAll(id, ".", "")
	id += r.String()
	return id[0:32]
}

func NanoUid() string {
	r, _ := rand.Prime(rand.Reader, 64)
	id := time.Now().UTC().Format("20060102150405.0000000")
	id = strings.ReplaceAll(id, ".", "")
	id += r.String()
	return id[0:23]
}

func MicroUid() string {
	r, _ := rand.Prime(rand.Reader, 64)
	id := time.Now().UTC().Format("20060102150405.0000000")
	id = strings.ReplaceAll(id, ".", "")
	id += r.String()
	return id[0:20]
}

func SecUid() string {
	r, _ := rand.Prime(rand.Reader, 64)
	id := time.Now().UTC().Format("20060102150405.0000000")
	id = strings.ReplaceAll(id, ".", "")
	id += r.String()
	return id[0:23]
}
