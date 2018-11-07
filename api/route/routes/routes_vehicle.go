package routes

import "gbmchallenge/api/service"

func VehicleRoutes() []Route {
	ws := service.NewVehicleWS()
	routes := []Route{
		{
			Method:  "GET",
			Path:    "/api/vehicle",
			Handler: ws.GetVehiclesHandler,
		},
		{
			Method:  "GET",
			Path:    "/api/vehicle/{vehicleId}/position",
			Handler: ws.GetVehiclePositionHandler,
		},
		{
			Method:  "PATCH",
			Path:    "/api/vehicle/{vehicleId}/position",
			Handler: ws.UpdateVehiclePositionHandler,
		},
	}

	return routes
}