package main

import "testing"

func TestImportJSONHourlyMode(t *testing.T) {
	var invoice Invoice
	if err := importJson([]byte(`{"hourly":true}`), &invoice); err != nil {
		t.Fatalf("import hourly mode: %v", err)
	}
	if !invoice.Hourly {
		t.Error("hourly mode was not enabled")
	}
}
