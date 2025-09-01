package class

import "fmt"

// Helper function for claims extraction
func GetClaims(claims map[string]interface{}, field string) string {
	if v, ok := claims[field]; ok && v != nil {
		if str, ok := v.(string); ok {
			return str
		}
		return fmt.Sprintf("%v", v)
	}
	return ""
}
