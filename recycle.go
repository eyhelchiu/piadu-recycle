package main

import (
	"strconv"
	"net/http"
	"labix.org/v2/mgo/bson"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, r, "index.html", nil)
	return
}

func recycleAppendHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		renderTemplate(w, r, "recycle_append.html", nil);
		return
	}
	if r.Method == "POST" {
		name := r.FormValue("name")
		province := r.FormValue("province")
		city := r.FormValue("city")
		hometown := r.FormValue("hometown")
		location := r.FormValue("location")
		scope := r.FormValue("scope")
		telephone := r.FormValue("telephone")
		geolatitude, _ := strconv.ParseFloat(r.FormValue("geolatitude"), 64)
		geolongitude, _ := strconv.ParseFloat(r.FormValue("geolongitude"), 64)

		c := db.C("recycle_point")
		err := c.Insert(&RecyclePoint{
			Id_:  bson.NewObjectId(),
			Name: name,
			Province: province,
			City: city,
			Hometown: hometown,
			Location: location,
			Scope: scope,
			Telephone: telephone,
			GeoLatitude: geolatitude,
			GeoLongitude: geolongitude,
		})

		if err != nil {
		}

		url := "/search?province=" + province + "&city=" + city + "&hometown=" + hometown
		http.Redirect(w, r, url, http.StatusFound)
	}
}

func recycleSearchHandler(w http.ResponseWriter, r *http.Request) {
	province := r.FormValue("province")
	city := r.FormValue("city")
	hometown := r.FormValue("hometown")

	if province=="" || city=="" || hometown=="" {
		renderTemplate(w, r, "index.html", map[string]interface{}{"province":province,"city":city,"hometown":hometown});
		return
	}

	var recyclePoints []RecyclePoint
	c := db.C("recycle_point")
	c.Find(bson.M{"province":province,"city":city,"hometown":hometown}).All(&recyclePoints)
	renderTemplate(w, r, "recycle_points.html", map[string]interface{}{"province":province,"city":city,"hometown":hometown,"recyclePoints":recyclePoints})
}

func mobileHandler(w http.ResponseWriter, r *http.Request) {
	province := r.FormValue("province")
	city := r.FormValue("city")
	hometown := r.FormValue("hometown")

	if province=="" || city=="" || hometown=="" {
		renderTemplate2(w, r, "recycle_wap.html", map[string]interface{}{"province":province,"city":city,"hometown":hometown})
		return
	}

	var recyclePoints []RecyclePoint
	c := db.C("recycle_point")
	c.Find(bson.M{"province":province,"city":city,"hometown":hometown}).All(&recyclePoints)
	renderTemplate2(w, r, "recycle_wap_points.html", map[string]interface{}{"recyclePoints":recyclePoints})
}

func mobileGeoHandler(w http.ResponseWriter, r *http.Request) {
	latitude := r.FormValue("latitude")
	longitude := r.FormValue("longitude")

	if latitude=="" || longitude=="" {
		http.Redirect(w, r, "/wap", http.StatusFound)
		return
	}

	latitude_, _ := strconv.ParseFloat(latitude, 0)
	longitude_, _ := strconv.ParseFloat(longitude, 0)

	var recyclePoints []RecyclePoint
	c := db.C("recycle_point")
	c.Find(bson.M{"geolatitude":bson.M{"$gt":latitude_-1,"$lt":latitude_+1}, "geolongitude":bson.M{"$gt":longitude_-1,"$lt":longitude_+1}}).All(&recyclePoints)
	renderTemplate2(w, r, "recycle_wap_points.html", map[string]interface{}{"recyclePoints":recyclePoints})
}
