package gtimer

import (
	"fmt"
	"time"
)

type RemoteTimer struct {
	token string
}

func (t *RemoteTimer) NewTimer(string string) {
	t.token = string
}

type NoAuthError int

func (f NoAuthError) Error() string {
	return fmt.Sprintf("wrong authentication")
}
func GetServiceName() string {
	return "RemoteTimer.GetUTCTime"
	//return runtime.FuncForPC(reflect.ValueOf((*RemoteTimer).GetUTCTime).Pointer()).Name()
}
func (t *RemoteTimer) GetUTCTime(arg string, reply *time.Time) error {
	if arg != (string(t.token)) {
		return NoAuthError(0)
	}
	*reply = time.Now()
	return nil
}
