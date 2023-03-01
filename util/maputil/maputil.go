// maputil.go
package maputil

import "log"

var MyMap map[string]interface{} // declare the map with a capital letter

func init() {

	log.Println("maputil init() called...")
	MyMap = make(map[string]interface{}) // initialize the map
}

func Add(key string, value interface{}) {

	MyMap[key] = value
}
