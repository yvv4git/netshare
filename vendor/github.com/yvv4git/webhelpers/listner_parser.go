package webhelpers

import "fmt"

// GetListenServerString function for generate server host:port string.
func GetListenServerString(host string, port int) string {
	return fmt.Sprintf("%s:%d", host, port)
}
