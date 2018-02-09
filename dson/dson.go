package dson

import (
	"encoding/json"
	"log"
)

// JSONObject is a json object
type JSONObject map[string]interface{}

// JSONArray is a json array
type JSONArray []interface{}

func toBoolean(v interface{}) bool {
	if result, ok := v.(bool); ok {
		return result
	}
	return false
}

func toInt(v interface{}) int {
	return int(toFloat64(v))
}

func toInt64(v interface{}) int64 {
	return int64(toFloat64(v))
}

func toFloat32(v interface{}) float32 {
	return float32(toFloat64(v))
}

func toFloat64(v interface{}) float64 {
	if result, ok := v.(float64); ok {
		return result
	}
	return 0.0
}

func toString(v interface{}) string {
	if result, ok := v.(string); ok {
		return result
	}
	return ""
}

// ParseObject parse json string to a JSONObject instance
func ParseObject(text string) *JSONObject {
	var mapObj map[string]interface{}
	err := json.Unmarshal([]byte(text), &mapObj)
	if err != nil {
		log.Panic(err)
	}
	resJs := JSONObject(mapObj)
	return &resJs
}

// ParseArray parse json string to a JSONArray instance
func ParseArray(text string) *JSONArray {
	var array []interface{}
	err := json.Unmarshal([]byte(text), &array)
	if err != nil {
		log.Panic(err)
	}
	resArray := JSONArray(array)
	return &resArray
}

// String return json string of this JSONObject
func (jo *JSONObject) String() string {
	data, err := json.Marshal(*jo)
	if err != nil {
		log.Panic(err)
	}
	return string(data)
}

// PrettyString return json string with indent, and is easy to see
func (jo *JSONObject) PrettyString(indent string) string {
	data, err := json.MarshalIndent(*jo, "", indent)
	if err != nil {
		log.Panic(err)
	}
	return string(data)
}

func (jo *JSONObject) getVal(key string) interface{} {
	mapObj := map[string]interface{}(*jo)
	return mapObj[key]
}

// String return json string of this JSONArray
func (ja *JSONArray) String() string {
	data, err := json.Marshal(*ja)
	if err != nil {
		log.Panic(err)
	}
	return string(data)
}

// PrettyString return json string with indent, and is easy to see
func (ja *JSONArray) PrettyString(indent string) string {
	data, err := json.MarshalIndent(*ja, "", indent)
	if err != nil {
		log.Panic(err)
	}
	return string(data)
}

// Size the size of this JSONArray
func (ja *JSONArray) Size() int {
	var array = []interface{}(*ja)
	return len(array)
}

func (ja *JSONArray) getVal(i int) interface{} {
	array := []interface{}(*ja)
	return array[i]
}

// GetObject get JSONObject by key
func (jo *JSONObject) GetObject(key string) (jsonObj *JSONObject) {
	val := jo.getVal(key)
	if result, ok := val.(map[string]interface{}); ok {
		resJs := JSONObject(result)
		jsonObj = &resJs
	}
	return
}

// GetArray get JSONArray by key
func (jo *JSONObject) GetArray(key string) (jsonArray *JSONArray) {
	val := jo.getVal(key)
	if result, ok := val.([]interface{}); ok {
		resArray := JSONArray(result)
		jsonArray = &resArray
	}
	return
}

func (ja *JSONArray) GetObject(i int) (jsonObj *JSONObject) {
	val := ja.getVal(i)
	if result, ok := val.(map[string]interface{}); ok {
		resJs := JSONObject(result)
		jsonObj = &resJs
	}
	return
}

func (ja *JSONArray) GetArray(i int) (jsonArray *JSONArray) {
	val := ja.getVal(i)
	if result, ok := val.([]interface{}); ok {
		resArray := JSONArray(result)
		jsonArray = &resArray
	}
	return
}

func (jo *JSONObject) GetBoolean(key string) bool {
	return toBoolean(jo.getVal(key))
}

// GetInt
func (jo *JSONObject) GetInt(key string) int {
	return toInt(jo.getVal(key))
}

func (jo *JSONObject) GetInt64(key string) int64 {
	return toInt64(jo.getVal(key))
}

func (jo *JSONObject) GetFloat32(key string) float32 {
	return toFloat32(jo.getVal(key))
}

func (jo *JSONObject) GetFloat64(key string) float64 {
	return toFloat64(jo.getVal(key))
}

func (jo *JSONObject) GetString(key string) string {
	return toString(jo.getVal(key))
}

func (ja *JSONArray) GetBoolean(i int) bool {
	return toBoolean(ja.getVal(i))
}

func (ja *JSONArray) GetInt(i int) int {
	return toInt(ja.getVal(i))
}

func (ja *JSONArray) GetInt64(i int) int64 {
	return toInt64(ja.getVal(i))
}

func (ja *JSONArray) GetFloat32(i int) float32 {
	return toFloat32(ja.getVal(i))
}

func (ja *JSONArray) GetFloat64(i int) float64 {
	return toFloat64(ja.getVal(i))
}

func (ja *JSONArray) GetString(i int) string {
	return toString(ja.getVal(i))
}
