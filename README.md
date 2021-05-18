# gmaps2

### Setup
```
go mod tidy
go run main.go -h

Usage:
  gmaps2 [flags]
  gmaps2 [command]

Available Commands:
  direction   search direction from google map api
  geojson     geojson convert google direction result to geojson format
  help        Help about any command

Flags:
  -h, --help   help for gmaps2

Use "gmaps2 [command] --help" for more information about a command.
```

### Google Direction

```
go run main.go direction -h

Usage:
  gmaps2 direction -o '7.123,144.567' -d '7.456,144.890' -w '7.234,144.678|7.345,144.789' [flags]

Flags:
  -a, --avoid string           avoid=tolls|highways|ferries
  -c, --configfile string      config file
  -t, --departuretime string   departure time in unixtimestamp
  -d, --destination string     destination latitude,longitude
  -h, --help                   help for direction
  -o, --origin string          origin latitude,longitude
  -w, --waypoints string       waypoint addrlatitude,longitude with '|' as separator
```
  
### GeoJSON

```
go run main.go geojson -h

Usage:
  gmaps2 geojson -t 'multiboxes' -f 'input/multiboxes.json' [flags]

Flags:
  -f, --file string   json source for geo convertion
  -h, --help          help for geojson
  -t, --type string   type=line|multilines|multiboxes (default "line")
```

  
### S2Polyline

```
go run main.go s2polyline -h

Usage:
  gmaps2 s2polyline -f 'input/linestring.json' [flags]

Flags:
  -c, --contain string     contain s2.cell from lat,lng
  -f, --file string        json source for geo convertion
  -h, --help               help for s2polyline
  -i, --intersect string   intersect s2.cell from lat,lng
```