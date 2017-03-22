package main

import (
	"log"

	"github.com/emicklei/go-restful"

	device "github.com/gbaeke/go-device/proto"
	"github.com/micro/go-micro/client"
	_ "github.com/micro/go-plugins/registry/kubernetes"
	"github.com/micro/go-web"

	"golang.org/x/net/context"
)

// DevSvc represents the API
type DevSvc struct{}

var (
	cl device.DevSvcClient
)

// Root gets called at /
func (d *DevSvc) Root(req *restful.Request, rsp *restful.Response) {
	log.Print("Received API request at /")
	rsp.WriteEntity(map[string]string{
		"message": "Device API at /",
	})
}

// Get gets called at /device
func (d *DevSvc) Get(req *restful.Request, rsp *restful.Response) {
	log.Print("Received API request at /device")

	// get device name from the request
	name := req.PathParameter("name")

	// call Device API and get the device
	response, err := cl.Get(context.TODO(), &device.DeviceName{Name: name})

	if err != nil {
		rsp.WriteEntity(err)
	} else {
		rsp.WriteEntity(response)
	}

}

func main() {
	service := web.NewService(
		web.Name("go.micro.api.device"),
	)

	service.Init()

	// DevSvc client
	cl = device.NewDevSvcClient("go.micro.srv.device", client.DefaultClient)

	// handle restful
	devsvc := new(DevSvc)
	ws := new(restful.WebService)
	wc := restful.NewContainer()
	ws.Consumes(restful.MIME_XML, restful.MIME_JSON)
	ws.Produces(restful.MIME_JSON, restful.MIME_XML)
	ws.Path("/device")
	ws.Route(ws.GET("/").To(devsvc.Root))
	ws.Route(ws.GET("/{name}").To(devsvc.Get))
	wc.Add(ws)

	// register handler
	service.Handle("/", wc)

	// run server
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}

}
