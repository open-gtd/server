package factories

import (
	"github.com/open-gtd/server/api"
	tagsApi "github.com/open-gtd/server/tags/api"
)

func GetPushDataToSseFunc() api.PushDataToSseFunc {
	return tagsApi.GetPushDataToSseFunc()
}
