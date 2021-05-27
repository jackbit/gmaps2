# gmaps2

### Main Example

```
go run main.go s2trip --start '-33.43535986456468, -70.70988860468003' --end '-33.46268514612078, -70.68886433781532' --pickup '-33.447130720203354, -70.70814176197621' --drop '-33.43672804765354, -70.6861326682225' --radius 13 --tripmaxlevel 16 --tripminlevel 16 --tripmaxcell 50
```

Video Explanation:
https://youtube.com/watch?v=vmJxjaTARRY
https://www.youtube.com/watch?v=_-ewGKC4ZRM

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
##### search direction from google map api
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
##### geojson convert google direction result to geojson format

```
go run main.go geojson -h

Usage:
  gmaps2 geojson -t 'multiboxes' -f 'input/multiboxes.json' [flags]

Flags:
  -f, --file string   json source for geo convertion
  -h, --help          help for geojson
  -t, --type string   type=line|multilines|multiboxes (default "line")
```

#### Convert LineString to MultiPolygon
```
 go run main.go geojson -t multipolygon -f input/linestring.json -r 0.2 -u km
```

#### Convert LineString to 1D Polygon
```
go run main.go geojson -t linetogon -f input/linestring.json
```

#### Convert Origin and Destination to Polygon
```
go run main.go geojson -t multiboxes -f input/multiboxes.json
```
 
#### Convert Google Direction to Linestring and Destination to Polygon
```
go run main.go geojson -t line -f input/direction.json
```

  
### S2Polyline
##### s2polyline convert geojson linestring to s2polyline

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

### GEOS
##### geos convert linestring to polygon by GEOS

#### Installation

The project depends on [geos](https://github.com/libgeos/geos) (GEOS is a C++ port of the ​JTS Topology Suite), you need to complete the installation of `geos` first. The installation of `geos`:

1. Mac OS X(via brew)
```sh
$ brew install geos
```
2. Ubuntu or Debian
```sh
$ apt-get install libgeos-dev
```
3. Build from source code
```sh
$ wget http://download.osgeo.org/geos/geos-3.9.0.tar.bz2
$ tar xvfj geos-3.9.0.tar.bz2
$ cd geos-3.9.0
$ ./configure
$ make
$ sudo make install
```

#### Usage
```
go run main.go geos -h


Usage:
  gmaps2 geos -f 'input/linestring.json' -s 2 [flags]

Flags:
  -f, --file string     json source for geo convertion
  -h, --help            help for geos
  -s, --sample string   example operation in integer (s=1,2,3)
```

##### Result polyline from google direction
![polyline](https://github.com/jackbit/gmaps2/raw/main/assets/polyline_direction-min.png)

##### Result conversion to polygon
![Result conversion to polygon](https://github.com/jackbit/gmaps2/raw/main/assets/polygon_direction-min.png)

##### Comparison polyline and polygon
![Comparison polyline and polygon](https://github.com/jackbit/gmaps2/raw/main/assets/merge_polygon_direction-min.png)

##### Radius polygon to polyline
![Radius polygon to polyline](https://github.com/jackbit/gmaps2/raw/main/assets/distance_middle_radius-min.png)

##### Diameter polygon
![Diameter polygon](https://github.com/jackbit/gmaps2/raw/main/assets/full_radius-min.png)


##### S2 Polygon
![Diameter polygon](https://github.com/jackbit/gmaps2/raw/main/assets/googles2polyline.png)