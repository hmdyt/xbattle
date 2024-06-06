package misc

import (
	"fmt"
	"github.com/gofrs/uuid"
)

type UuidGenerator interface {
	Generate() string
}

type UuidGeneratorV7 struct{}

func NewUuidGeneratorV7() UuidGenerator {
	return &UuidGeneratorV7{}
}

func (u *UuidGeneratorV7) Generate() string {
	id, err := uuid.NewV7()
	if err != nil {
		panic(fmt.Sprintf("UUIDv7の生成に失敗しました: %v", err))
	}
	return id.String()
}
