/*
 * Inventory API
 *
 * This is a API wromo
 *
 * API version: 1.0.0
 * Contact: info.wromo@gmail.com
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

import (
	"net/http"
)

func SearchInventory(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}
