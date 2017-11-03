package ddgquery

import "testing"

func TestQuery(t *testing.T) {
	_, query := Query("test", 3)

	if query != "https://duckduckgo.com/html/?q=test" {
		t.Errorf("Query was incorrect, got: %s, want: %s.", query, "https://duckduckgo.com/html/?q=test")
	}
}