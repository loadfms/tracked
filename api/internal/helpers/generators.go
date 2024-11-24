package helpers

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"strconv"
	"strings"
	"time"
)

func GenerateSK(prefix string) string {
	timestamp := time.Now().Unix()
	hash := sha256.New()
	hash.Write([]byte(strconv.FormatInt(timestamp, 10)))
	hashed := hash.Sum(nil)

	fmt.Println(hex.EncodeToString(hashed)[:45])

	return fmt.Sprintf("%s##%s", strings.ToUpper(prefix), hex.EncodeToString(hashed)[:45])
}
