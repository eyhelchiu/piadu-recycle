package main

import (
	"os"
	"net/http"
	"encoding/json"
	"github.com/gorilla/mux"
	"labix.org/v2/mgo"
)

//{"name": "拳头垃圾回收点", "telephone": "18902265187", "province": "广东省", "city": "梅州市", "hometown": "梅县", "location": "江北黄金花园", "scope": "黄金花园|胡权新家|渡江津|江北花园", "geolocation": "******"}

var (
	config map[string]string
	db *mgo.Database
)

func init() {
	loadConfigure()
	connectDatabase()
}

func loadConfigure() {
	file, err := os.Open("config.json.default")
	if err != nil {
		println("load configure failed.")
		panic(err)
		os.Exit(1)
	}
	defer file.Close()

	dec := json.NewDecoder(file)
	err = dec.Decode(&config)
	if err != nil {
		println("read configure failed.")
		panic(err)
		os.Exit(1)
	}
}

func connectDatabase() {
	session, err := mgo.Dial(config["db"])
	if err != nil {
		panic(err)
	}
	//defer session.Close()
	session.SetMode(mgo.Monotonic, true)
	db = session.DB(config["db_name"])
}

func main() {
	http.Handle("/assets/", http.FileServer(http.Dir(".")))

	r := mux.NewRouter()
	for url, handle := range handlers {
		r.HandleFunc(url, handle)
	}
	http.Handle("/", r)

	port := config["port"]
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		println("server listen failed.")
		return
	}
}
