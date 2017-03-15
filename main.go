package main

import (
	"fmt"
	"os"

	device "github.com/gbaeke/go-device/proto"
	"github.com/micro/cli"
	"github.com/micro/go-micro"
	"golang.org/x/net/context"
)

// DevSvc defines the service
type DevSvc struct {
	devs map[string]*device.Device
}

// Get function of service
func (d *DevSvc) Get(ctx context.Context, req *device.DeviceName, rsp *device.Device) error {
	fmt.Println("Will respond with ", d.devs[req.Name])

	// this also works
	rsp.Name = d.devs[req.Name].Name
	rsp.Active = d.devs[req.Name].Active

	return nil
}

// LoadDevices creates and returns a map of devices
func LoadDevices() map[string]*device.Device {
	// initialise the map
	devices := make(map[string]*device.Device)

	// add some dummy devices; should come from some db
	devices["device1"] = &device.Device{Name: "device1", Active: true}
	devices["device2"] = &device.Device{Name: "device2", Active: true}
	devices["device3"] = &device.Device{Name: "device3", Active: true}

	return devices
}

// Setup and the client
func runClient(service micro.Service) {
	// Create new Dev client
	DevClient := device.NewDevSvcClient("DevSvc", service.Client())

	// Call the greeter
	rsp, err := DevClient.Get(context.TODO(), &device.DeviceName{Name: "device2"})
	if err != nil {
		fmt.Println(err)
		return
	}

	// Print response
	fmt.Println("Response: ", rsp)
}

func main() {
	// keep it extremely simple for now
	service := micro.NewService(
		micro.Name("DevSvc"),
		micro.Version("latest"),

		micro.Flags(cli.BoolFlag{
			Name:  "run_client",
			Usage: "Launch the client",
		}),
	)

	service.Init(
		// Add runtime action
		// We could actually do this above
		micro.Action(func(c *cli.Context) {
			if c.Bool("run_client") {
				runClient(service)
				os.Exit(0)
			}
		}),
	)

	// register handler and initialise devs map with a list of devices
	device.RegisterDevSvcHandler(service.Server(), &DevSvc{devs: LoadDevices()})

	//run Server
	if err := service.Run(); err != nil {
		fmt.Println(err)
	}

}
