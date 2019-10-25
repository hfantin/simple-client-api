package eureka

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"time"

	u "github.com/hfantin/clientgo/utils"
)

type Client struct {
	instanceId string
	appName    string
}

func (c *Client) Register() {
	c.appName = "clientgo"
	c.instanceId = c.appName + ":" + u.GetUUID()

	dir, _ := os.Getwd()
	data, _ := ioutil.ReadFile(dir + "/templates/regtpl.json")
	tpl := string(data)
	tpl = strings.Replace(tpl, "${ipAddress}", u.GetLocalIP(), -1)
	tpl = strings.Replace(tpl, "${port}", "3000", -1)
	tpl = strings.Replace(tpl, "${instanceId}", c.instanceId, -1)
	tpl = strings.Replace(tpl, "${appName}", c.appName, -1)
	// Register.
	registerAction := HttpAction{
		Url:         "http://localhost:8761/eureka/apps/clientgo",
		Method:      "POST",
		ContentType: "application/json",
		Body:        tpl,
	}
	var result bool
	for {
		result = DoHttpRequest(registerAction)
		if result {
			break
		} else {
			time.Sleep(time.Second * 5)
		}
	}
}

func (c *Client) StartHeartbeat() {
	for {
		time.Sleep(time.Second * 30)
		c.heartbeat()
	}
}

func (c *Client) heartbeat() {
	log.Println("eurekaclient - send heartBeat")
	heartbeatAction := HttpAction{
		Url:    "http://localhost:8761/eureka/apps/clientgo/" + c.instanceId,
		Method: "PUT",
	}
	DoHttpRequest(heartbeatAction)
}

func (c *Client) Deregister() {
	fmt.Println("Trying to deregister application...")
	// Deregister
	deregisterAction := HttpAction{
		Url:    "http://localhost:8761/eureka/apps/clientgo/" + c.instanceId,
		Method: "DELETE",
	}
	DoHttpRequest(deregisterAction)
	fmt.Println("Deregistered application, exiting. Check Eureka...")
}
