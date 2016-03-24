# geocoder
A geocoding server to interact with Google Geocoding API

Building
----

 - Go 1.6
 - [Godeps](https://github.com/tools/godep)

Run `godep get` on the project directory to fetch dependencies or get it from `vendors/`

Running
---

We need a Google Maps Geocoding API key, if you don't have one, go to your console and create.
Then set it on environment as `GOOGLE_GEO_KEY`

To use Redis you will need a environment variable called `REDIS_HOST` for Redis' host and `REDIS_PORT`
for Redis' port, and if you desire, a `REDIS_PASSWORD` too. The database defaults to `0`

Then just run `go run server.go` and open at `:4000`

URLs
-----

`/address?latlng=40.7127840,-74.0059410`  translate the latitude and longitude into human readable address

**Response:**

```
{
  "address": "New York City Hall, New York, NY 10007, USA"
}

```


`/coordinates?address=New_York` translate the location name into a latitude and longitude pair

**Response:**

```
{
    "location": {
        "lat": 40.7127837,
        "lng": -74.0059413
    }
}
```