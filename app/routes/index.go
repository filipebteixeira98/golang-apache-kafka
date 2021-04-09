package routes

import (
	"bufio"
	"encoding/json"
	"errors"
	"os"
	"strconv"
	"strings"
)

type Route struct {
	ID        string     `json:"routeId"`
	ClientID  string     `json:"clientId"`
	Positions []Position `json:"positions"`
}

type Position struct {
	Lat  float64 `json:"lat"`
	Long float64 `json:"long"`
}

type PartialRoutePosition struct {
	ID       string    `json:"routeId"`
	ClientID string    `json:"clientId"`
	Position []float64 `json:"position"`
	Finished bool      `json:"finished"`
}

func NewRoute() *Route {
	return &Route{}
}

func (r *Route) LoadPositions() error {
	if r.ID == "" {
		return errors.New("route ID not provided")
	}
	f, err := os.Open("docs/" + r.ID + ".txt")
	if err != nil {
		return err
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		data := strings.Split(scanner.Text(), ",")
		latitude, err := strconv.ParseFloat(data[0], 64)
		if err != nil {
			return err
		}
		longitude, err := strconv.ParseFloat(data[1], 64)
		if err != nil {
			return err
		}
		r.Positions = append(r.Positions, Position{
			Lat:  latitude,
			Long: longitude,
		})
	}
	return nil
}

func (r *Route) ExportJSONPositions() ([]string, error) {
	var partialRoute PartialRoutePosition
	var result []string
	total := len(r.Positions)
	for k, v := range r.Positions {
		partialRoute.ID = r.ID
		partialRoute.ClientID = r.ClientID
		partialRoute.Position = []float64{v.Lat, v.Long}
		partialRoute.Finished = false
		if total-1 == k {
			partialRoute.Finished = true
		}
		partialRouteJSON, err := json.Marshal(partialRoute)
		if err != nil {
			return nil, err
		}
		result = append(result, string(partialRouteJSON))
	}
	return result, nil
}
