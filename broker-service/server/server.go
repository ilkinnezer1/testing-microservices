package server

import "broker/routes"

const BrokerPort = ":7000"

func Start() {
	e := routes.BaseRoutes()

	if err := e.Start(BrokerPort); err != nil {
		panic(err)
	}
}
