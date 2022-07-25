package uniqueId

import(
	"time"
	// "strconv"
	"github.com/catinello/base62"

	// log "github.com/sirupsen/logrus"
)

func generateShortUniqueIdService() string {
	now := time.Now().UnixMilli()
	epochTime := int(now)
	sEnc := base62.Encode(epochTime)
	return sEnc
}