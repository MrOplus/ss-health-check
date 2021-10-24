package main

import (
	"encoding/json"
	"fmt"
	"github.com/DavidGamba/go-getoptions"
	"github.com/kooroshh/bridge/adapter/outbound"
	"github.com/kooroshh/ss-health-check/alert"
	"github.com/kooroshh/ss-health-check/config"
	"github.com/kooroshh/ss-health-check/healthcheck"
	"github.com/kooroshh/ss-health-check/parser"
	"io/ioutil"
	"os"
	"sync"
)
var wg sync.WaitGroup
func main() {
	var configPath string

	opt := getoptions.New()
	opt.Bool("help", false, opt.Alias("h", "?"))
	opt.StringVar(&configPath, "config", "conf.json",opt.Required(),opt.Alias("c"),opt.Description("configuration file"))
	_ , err := opt.Parse(os.Args[1:])
	if opt.Called("help") {
		fmt.Fprintf(os.Stderr, opt.Help())
		os.Exit(1)
	}
	if err != nil {
		fmt.Fprintf(os.Stderr, "ERROR: %s\n\n", err)
		fmt.Fprintf(os.Stderr, opt.Help(getoptions.HelpSynopsis))
		os.Exit(1)
	}
	content,err := ioutil.ReadFile(configPath)
	if err != nil {
		fmt.Fprintf(os.Stderr,"Unable to read the file")
		os.Exit(1)
	}
	var conf config.Config
	err = json.Unmarshal(content,&conf)
	if err != nil {
		fmt.Fprintln(os.Stderr,err)
		os.Exit(1)
	}
	fmt.Fprintf(os.Stdout,"Total Servers : %d\n",len(conf.Servers))
	fmt.Fprintln(os.Stdout,"Parsing Servers")
	var proxies []map[string]interface{}
	for i , el := range conf.Servers {
		proxy , err := parser.Parse(el)
		if err != nil {
			fmt.Fprintf(os.Stderr,"Unable to parse server at index : %d",i)
			continue
		}
		proxies = append(proxies,proxy.ToMap())
	}
	var hcs []*healthcheck.HealthCheck

	for _,proxy := range proxies {
		parsedProxy, err :=  outbound.ParseProxy(proxy)
		if err != nil {
			fmt.Println("Unable to parse proxy",err)
			continue
		}
		healthCheck := healthcheck.HealthCheck{
			Proxy:  parsedProxy,
			Config: proxy,
		}
		hcs = append(hcs,&healthCheck)
		wg.Add(1)
		go healthCheck.StartMonitoring(&wg,&conf.Callbacks)
	}
	wg.Wait()
	alert.SendAlert(conf.Callbacks,"There is no server left for monitoring.")
}
