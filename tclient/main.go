package main

import (
	"flag"
	"fmt"
	"log"
	"net/rpc"
	"time"

	"github.com/raven4752/gtimer"
)

type outOfTimeError int

func (f outOfTimeError) Error() string {
	return fmt.Sprintf("out Of Time due to long delay %d ms", f)
}
func getUTCTime(address string, token string) (time.Time, error) {
	client, err := rpc.Dial("tcp", address)
	if err != nil {
		log.Fatal("dialing", err)
	}
	args := token
	var reply time.Time
	start := time.Now()
	if err = client.Call(gtimer.GetServiceName(), args, &reply); err != nil {
		return time.Now(), err
	}
	end := time.Now()
	if elipsed := end.Sub(start).Nanoseconds() / int64(time.Millisecond); elipsed < 6000 {

		return reply.Add(time.Duration(elipsed / 2)), nil
	} else {
		fmt.Printf("aborted due to delay")
		return reply, outOfTimeError(elipsed)
	}
}

func queryTime(addr string, token string, format string, duration time.Duration) {
	for {
		if timestamp, err := getUTCTime(addr, token); err != nil {
			fmt.Println(err.Error())
		} else {
			fmt.Println(timestamp.Format(format))
		}
		time.Sleep(duration)
	}
}
func main() {
	log.Print(" client start")
	addr := flag.String("server", "127.0.0.1:12345", "address:port of the server")
	token := flag.String("token", "gouliguojiashengsiyi", "token used to authentication")
	mode := flag.String("mode", "digital", "mode to show time. digital: show time per second.analogue:show time per minute")
	flag.Parse()
	switch *mode {
	case "digital":
		queryTime(*addr, *token, "15:04:05", time.Second)
	case "analogue":
		queryTime(*addr, *token, "15:04", time.Minute)
	default:
		queryTime(*addr, *token, "15:04:05", time.Second)

	}

}
