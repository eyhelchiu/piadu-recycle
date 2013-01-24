package main

import (
	"labix.org/v2/mgo/bson"
)

type RecyclePoint struct {
	Id_ bson.ObjectId `bson:"_id"`
	Name string
	Province string
	City string
	Hometown string
	Location string
	Scope string
	Telephone string
	GeoLatitude float64
	GeoLongitude float64
}
