package alert

import (
	"context"
	"fmt"
	"github.com/kooroshh/ss-health-check/config"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"
)

func SendAlert(callbacks []config.Callback,message string) {
	for _ , callback := range callbacks {
		if callback.Type == "webhook" {
			data := url.Values{}
			params := strings.Split(callback.Params,"&")
			for index , param := range params {
				key := strings.Split(param,"=")[0]
				value := strings.Split(param,"=")[1]
				if index == (len(params) -1) {
					value = fmt.Sprintf(value,message)
					fmt.Println(value)
				}
				data.Add(key,value)
				client := http.Client{}

				ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
				defer cancel()
				req, err := http.NewRequestWithContext(ctx,callback.Method,callback.Url,strings.NewReader(data.Encode()))
				req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

				if err !=nil {
					fmt.Fprintln(os.Stderr,err)
					return
				}
				_, err = client.Do(req)
				if err != nil {
					return
				}
			}
		}
	}
}
