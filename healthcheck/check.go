package healthcheck

import (
	"context"
	"fmt"
	"github.com/kooroshh/bridge/constant"
	"github.com/kooroshh/ss-health-check/alert"
	"github.com/kooroshh/ss-health-check/config"
	"math/rand"
	"os"
	"sync"
	"time"
)

type HealthCheck struct {
	Proxy constant.Proxy
	Config map[string]interface{}
}

func (h HealthCheck) StartMonitoring(wg *sync.WaitGroup,callbacks *[]config.Callback) {
	failed := 0
	for {
		ctx , done := context.WithTimeout(context.Background(),10 * time.Second)
		latency , err := h.Proxy.URLTest(ctx,"https://google.com")
		done()
		if err != nil{
			fmt.Fprintln(os.Stderr,err)
			failed++
		}else{
			fmt.Fprintf(os.Stdout,"got response with %dms delay using %s:%d\n" ,latency,h.Config["server"],h.Config["port"])
			failed = 0
		}
		if failed == 3 {
			h.callHooks(*callbacks)
			break
		}
		rand.Seed(time.Now().UnixNano())
		sleepDuration := time.Duration(randInt(30,60)) * time.Second
		fmt.Fprintf(os.Stdout,"Going to Sleep %s:%d with %.1f seconds\n",h.Config["server"],h.Config["port"],sleepDuration.Seconds())
		time.Sleep(sleepDuration)
	}
	wg.Done()

}
func (h HealthCheck) callHooks(callbacks []config.Callback)  {
	alert.SendAlert(callbacks,h.Config["name"].(string))
}
func randInt(min int, max int) int {
	return min + rand.Intn(max-min)
}