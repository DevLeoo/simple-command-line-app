package app

import (
	"fmt"
	"log"
	"net"

	"github.com/urfave/cli"
)

// Generate will return a command line application ready to run
func Generate() *cli.App {
	app := cli.NewApp()

	flags := []cli.Flag{
		cli.StringFlag{
			Name:  "host",
			Value: "amazon.com.br",
		},
	}

	app.Name = "Command line application"
	app.Usage = "Search IPs and server names across the internet"
	app.Commands = []cli.Command{
		{
			Name:   "ip",
			Usage:  "Search IPs addresses",
			Flags:  flags,
			Action: searchIPs,
		}, {
			Name:   "server",
			Usage:  "Search server names across the internet",
			Flags:  flags,
			Action: searchServers,
		},
	}

	return app
}

func searchServers(c *cli.Context) {
	host := c.String("host")

	servers, err := net.LookupNS(host)

	if err != nil {
		log.Fatal(err)
	}

	for _, server := range servers {
		fmt.Println(server.Host)
	}
}

func searchIPs(c *cli.Context) {
	host := c.String("host")

	ips, erro := net.LookupIP(host)

	if erro != nil {
		log.Fatal(erro)
	}

	for _, ip := range ips {
		fmt.Println(ip)
	}
}
