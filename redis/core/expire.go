package core

import (
	"log"
	"time"
)

func expireSample() float32 {
	var limit int = 20
	var expiredCount int = 0

	for key, obj := range store {
		if obj.ExpiresAt != -1 {
			limit--

			// If the key is expired
			if obj.ExpiresAt <= time.Now().UnixMilli() {
				delete(store, key)
				expiredCount++
			}
		}

		// Once we iterated to 20 keys that have some expiration set
		// we break the loop
		if limit == 0 {
			break
		}
	}

	return float32(expiredCount) / float32(limit)
}

// Deletes all the expired keys
func DeleteExpiredKeys() {
	for {
		frac := expireSample()

		// If the sample had less than 25% keys expired
		// we break the loop

		if frac < 0.25 {
			break
		}
	}
	log.Println("Deleted the expired but undeleted keys. total keys ", len(store))
}
