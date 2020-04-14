package helpers

import (
	"strings"

	"github.com/torkashvand/goshortener/config"
)

//ConvertBase convert an int to a base len()
func ConvertBase(num uint) string {
	cfg := config.Config()
	shortenerBase := cfg.GetString("SHORTENER_BASE")
	var shortenedSlice []string
	base := uint(len(shortenerBase))

	for num > 0 {
		mod := num % base
		num = num / base

		shortenedSlice = append(shortenedSlice, string(shortenerBase[mod]))

		for i, j := 0, len(shortenedSlice)-1; i < j; i, j = i+1, j-1 {
			shortenedSlice[i], shortenedSlice[j] = shortenedSlice[j], shortenedSlice[i]
		}

	}

	shortenedString := strings.Join(shortenedSlice, "")

	return shortenedString
}
