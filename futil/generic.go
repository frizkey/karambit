package futil

import (
	"encoding/json"
	"github.com/frizkey/karambit/flogger"
	"hash/fnv"
	"math/rand"
	"regexp"
	"strings"
)

// CoallesceString is a helper function to coalesce string
//
// Deprecated: Use CoalesceString instead. CoallesceString will be removed in v1.0.0.
func CoallesceString(s ...string) string {
	return CoalesceString(s...)
}

func CoalesceString(s ...string) string {
	for _, v := range s {
		if v != "" {
			return v
		}
	}
	return ""
}

// Contains is
func Contains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}
	return false
}

// IsNumber is
func IsNumber(s string) bool {
	match, _ := regexp.MatchString("([a-z]+)", s)
	if !match {
		return true
	} else {
		return false
	}
}

func Hash(s string) uint32 {
	h := fnv.New32a()
	if _, err := h.Write([]byte(s)); err != nil {
		flogger.Errorf("Error hashing string: %s", err.Error())
		return 0
	}
	return h.Sum32()
}

func JsonEscape(i string) string {
	b, err := json.Marshal(i)
	if err != nil {
		return ""
	}
	s := string(b)
	return s[1 : len(s)-1]
}

func StrEscape(str string) string {

	str = strings.ReplaceAll(str, "\\", "\\\\")
	str = strings.ReplaceAll(str, "\"", "\\\"")
	str = strings.ReplaceAll(str, "\b", "\\b")
	str = strings.ReplaceAll(str, "\f", "\\f")
	str = strings.ReplaceAll(str, "\r", "\\r")
	str = strings.ReplaceAll(str, "\n", "\\n")
	str = strings.ReplaceAll(str, "\t", "\\t")

	return str
}

func ToJSONString(data interface{}) string {

	js, _ := json.Marshal(data)

	return string(js)
}

// Coalesce is
func Coalesce(values ...interface{}) interface{} {
	for _, value := range values {
		if value != nil {
			return value
		}
	}
	return nil
}

// GenerateRandomString is a helper function to generate random string by i length
func GenerateRandomString(i int) string {
	const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	b := make([]byte, i)
	for i := range b {
		b[i] = letterBytes[GenerateRandomInt(0, len(letterBytes))]
	}
	return string(b)
}

// GenerateRandomInt is a helper function to generate random int by min and max
func GenerateRandomInt(min, max int) int {
	return min + rand.Intn(max-min)
}

// ToUpper is a helper function to convert string to uppercase
func ToUpper(s string) string {
	return strings.ToUpper(s)
}
