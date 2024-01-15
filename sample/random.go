package sample

import (
	"math/rand"
	"time"

	"github.com/google/uuid"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func randomID() string {
	return uuid.New().String()
}

func randomInt(min, max int) int {
	return min + rand.Int()%(max-min+1)
}

func randomFloat32(min, max float32) float32 {
	return min + rand.Float32()*(max-min)
}
