package testmod

import (
	"errors"
	"fmt"
)

func Hi(name, lang string) (string, error) {
	switch lang {
	case "en":
		return fmt.Sprintf("Hi, %s!", name), nil
	case "chinese":
		return fmt.Sprintf("你好，%s！", name), nil
	default:
		return "", errors.New("unknown lang")
	}
}
