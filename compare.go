package main

// https://blog.kowalczyk.info/article/JyRZ/generating-good-unique-ids-in-go.html
// To run:
// go run main.go

import (
	"fmt"
	gouuid1 "github.com/nu7hatch/gouuid"
	"log"
	"math/rand"
	"time"

	"github.com/chilts/sid"
	guuid "github.com/google/uuid"
	"github.com/kjk/betterguid"
	"github.com/lithammer/shortuuid"
	"github.com/oklog/ulid"
	"github.com/rs/xid"
	"github.com/satori/go.uuid"
	"github.com/segmentio/ksuid"
	"github.com/sony/sonyflake"
)

func main() {
	fmt.Printf("github.com/rs/xid:              %s\n", genXid())
	fmt.Printf("github.com/segmentio/ksuid:     %s\n", genKsuid())
	fmt.Printf("github.com/kjk/betterguid:      %s\n", genBetterGUID())
	fmt.Printf("github.com/oklog/ulid:          %s\n", genUlid())
	fmt.Printf("github.com/sony/sonyflake:      %x\n", genSonyflake())
	fmt.Printf("github.com/chilts/sid:          %s\n", genSid())
	fmt.Printf("github.com/lithammer/shortuuid: %s\n", genShortUUID())
	fmt.Printf("github.com/satori/go.uuid:      %s\n", genUUIDv4())
	fmt.Printf("github.com/nu7hatch/gouuid:     %s\n", genUUID1v4())
	fmt.Printf("github.com/google/uuid:         %s\n", genUUID())
}

func genXid() string {
	id := xid.New()
	return id.String()
}

func genKsuid() string {
	id := ksuid.New()
	return id.String()
}

func genBetterGUID() string {
	id := betterguid.New()
	return id
}

func genUlid() string {
	t := time.Now().UTC()
	entropy := rand.New(rand.NewSource(t.UnixNano()))
	id := ulid.MustNew(ulid.Timestamp(t), entropy)
	return id.String()
}

func genSonyflake() uint64 {
	flake := sonyflake.NewSonyflake(sonyflake.Settings{})
	id, err := flake.NextID()
	if err != nil {
		log.Fatalf("flake.NextID() failed with %s\n", err)
	}
	// Note: this is base16, could shorten by encoding as base62 string
	return id
}

func genSid() string {
	id := sid.Id()
	return id
}

func genShortUUID() string {
	id := shortuuid.New()
	return id
}

func genUUIDv4() string {
	id, err := uuid.NewV4()
	if err != nil {
		log.Fatalf("uuid.NewV4() failed with %s\n", err)
	}
	return id.String()
}

func genUUID1v4() string {
	id, err := gouuid1.NewV4()
	if err != nil {
		log.Fatalf("uuid.NewV4() failed with %s\n", err)
	}
	return id.String()
}

func genUUID() string {
	id := guuid.New()
	return id.String()
}
