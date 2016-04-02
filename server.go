package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"

	"gopkg.in/go-redis/cache.v1"

	"time"

	"github.com/hashicorp/consul/api"
	"github.com/joaoevangelista/aetherius/conn"
	"github.com/joaoevangelista/aetherius/rest"
)

// Defining the Google Geocoding API
const (
	apiBase    string = "https://maps.googleapis.com/maps/api/geocode/json"
	expiration        = time.Hour
)

var apiKey = os.Getenv("GOOGLE_GEO_KEY")
var codec *cache.Codec
var client http.Client
var consul *api.Client

func init() {
	client = http.Client{
		Timeout: 2 * time.Second,
	}

	codec = conn.InitRedis()
	consul = conn.InitConsul()
}

func main() {
	if apiKey == "" {
		log.Fatal("Apikey not present on environment, requests will fail that way!")
	}

	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		if sts, err := json.Marshal(rest.Health{Status: "Ok"}); err == nil {
			w.Header().Set(rest.ContentType, rest.ApplicationJSON)
			fmt.Fprint(w, string(sts))
		}
	})

	http.HandleFunc("/coordinates", addrToCoord)
	http.HandleFunc("/address", coordToAddr)
	http.ListenAndServe(":4000", nil)
}

func addrToCoord(w http.ResponseWriter, r *http.Request) {
	address := r.URL.Query().Get("address")
	response := &rest.GeoResponse{}
	err := toGeo(response, address)
	rest.ThrowAPIErrorIfPresent(w, err)
	js, err := json.Marshal(rest.ExtractLocation(response))
	rest.ThrowJSONErrorIfPresent(w, err)
	w.Header().Set(rest.ContentType, rest.ApplicationJSON)
	fmt.Fprint(w, string(js))

}

func coordToAddr(w http.ResponseWriter, r *http.Request) {
	latlng := r.URL.Query().Get("latlng")
	if len(latlng) > 0 && strings.Contains(latlng, ",") {
		location := convertParam(latlng, w)
		if location.Latitude != 0 && location.Longitude != 0 {
			response := &rest.GeoResponse{}
			err := toAddr(response, location)
			rest.ThrowAPIErrorIfPresent(w, err)
			js, err := json.Marshal(rest.ExtractAddress(response))
			rest.ThrowJSONErrorIfPresent(w, err)
			w.Header().Set(rest.ContentType, rest.ApplicationJSON)
			fmt.Fprint(w, string(js))
		}
	} else {
		w.WriteHeader(http.StatusBadRequest)
		rest.PrintError(w, "PARAMETER_MISMATCH",
			"latlng parameter is empty or malformated, must be latlng=12.2121,-12.121")
	}
}

func toGeo(gr *rest.GeoResponse, addr string) error {
	if e := codec.Get(addr, gr); e == nil && gr.Status != "" {
		log.Printf("Got value from cache:  %v", gr)
		return nil
	} else {
		resp, err := client.Get(fmt.Sprintf("%s?key=%s&address=%s", apiBase, apiKey, addr))
		if e := GeoDecoder(gr, resp, err); e == nil {
			codec.Set(&cache.Item{
				Key:        addr,
				Object:     gr,
				Expiration: expiration,
			})
			return nil
		}
		return e
	}

}

func toAddr(gr *rest.GeoResponse, coord rest.Location) (err error) {
	key := asKey(coord)
	if e := codec.Get(key, gr); e == nil && gr.Status != "" {
		log.Println("Got value from cache:  %v", gr)
		return nil
	} else {
		resp, err := client.Get(fmt.Sprintf("%s?key=%s&latlng=%f,%f", apiBase, apiKey, coord.Latitude, coord.Longitude))
		if e = GeoDecoder(gr, resp, err); e == nil {
			codec.Set(&cache.Item{
				Key:        key,
				Object:     gr,
				Expiration: expiration,
			})
		}
		return nil
	}

}

func convertParam(latlng string, w http.ResponseWriter) rest.Location {
	parts := strings.Split(latlng, ",")
	lat, err := strconv.ParseFloat(parts[0], 64)
	lng, err := strconv.ParseFloat(parts[1], 64)
	if err != nil {
		log.Printf("Error while decoding parameter into float64 %v", err)
		rest.ThrowEncodeErrorIfPresent(w, err)
		return rest.Location{}
	}
	return rest.Location{Latitude: lat, Longitude: lng}
}

// GeoDecoder does the heavy work of parsing the response into a GeoResponse
func GeoDecoder(gr *rest.GeoResponse, resp *http.Response, err error) error {
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	c, err := ioutil.ReadAll(resp.Body)
	json.Unmarshal(c, &gr)
	if err != nil {
		log.Printf("Error %v", err)
		return fmt.Errorf("Error while decoding response %v", err)
	}
	return nil
}

func asKey(v interface{}) string {
	if js, e := json.Marshal(v); e == nil {
		return string(js)
	} else {
		panic(e)
	}
}
