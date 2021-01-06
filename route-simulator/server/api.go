package main

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

// writeResponse is helper to marshal return data into http response body
func writeResponse(w http.ResponseWriter, data interface{}, statusCode ...int) {
	status := 200
	if len(statusCode) > 0 && statusCode[0] != 0 {
		status = statusCode[0]
	}

	b, err := json.Marshal(data)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(status)
	w.Write(b)
}

func serveAPI(router *mux.Router) {
	subRouter := router.PathPrefix("/api").Subrouter()

	subRouter.HandleFunc("/routes", getRoutes)
	subRouter.HandleFunc("/routes/{route}/coordinates", getRouteCoordinates)
	subRouter.HandleFunc("/routes/{route}/start", getRouteStart)
	subRouter.HandleFunc("/routes/{route}/next", getNextCoordinate)

}

func getRoutes(w http.ResponseWriter, r *http.Request) {
	var routeNames []string
	for name := range routeMap {
		routeNames = append(routeNames, name)
	}

	writeResponse(w, routeNames)
}

func getRouteCoordinates(w http.ResponseWriter, r *http.Request) {
	routeName := mux.Vars(r)["route"]

	route, ok := routeMap[routeName]
	if !ok {
		writeResponse(w, "Route not found", http.StatusNotFound)
		return
	}

	writeResponse(w, route.Points)
}

func getRouteStart(w http.ResponseWriter, r *http.Request) {
	routeName := mux.Vars(r)["route"]
	reverse := r.URL.Query().Get("direction") == "reverse"

	// TODO: get route starting point, direction forward or reverse
	route, ok := routeMap[routeName]
	if !ok {
		writeResponse(w, "Route not found", http.StatusNotFound)
		return
	}
	if !reverse{
		writeResponse(w,route.Points[0])
	}else {
		writeResponse(w, route.Points[len(route.Points)-1])
	}
}

func getNextCoordinate(w http.ResponseWriter, r *http.Request) {
	routeName := mux.Vars(r)["route"]
	reverse := r.URL.Query().Get("direction") == "reverse"
	currentPoint := r.URL.Query().Get("current")

	// TODO: get next coordinate from current coordinate. need to aware reverse route.
	route, ok := routeMap[routeName]
	if !ok {
		writeResponse(w, "Route not found", http.StatusNotFound)
		return
	}
	var curPoint Point
	err := json.Unmarshal([]byte(currentPoint), &curPoint)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	if !reverse{
		for i, e := range route.Points {
			if e.Longitude == curPoint.Longitude && e.Latitude == curPoint.Latitude {
				writeResponse(w,route.Points[i+1])
			}
		}
	}else {
		for i := len(route.Points) - 1;i >= 0; i--{
			if route.Points[i].Longitude == curPoint.Longitude && route.Points[i].Latitude == curPoint.Latitude {
				writeResponse(w,route.Points[i-1])
			}
		}
	}
}
