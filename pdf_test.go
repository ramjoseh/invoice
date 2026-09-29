package main

import (
	"math"
	"testing"

	"github.com/signintech/gopdf"
)

func TestWriteItemRowsAddsContinuationPage(t *testing.T) {
	pdf := newTestPDF(t)
	invoice := Invoice{
		Id:         "88",
		From:       "Jose H. Ramirez",
		To:         "ML Orchestrator Project",
		Title:      "INVOICE",
		Date:       "Sep 28, 2026",
		Items:      make([]string, 13),
		Quantities: make([]float64, 13),
		Rates:      make([]float64, 13),
	}
	for i := range invoice.Items {
		invoice.Items[i] = "Software and Machine Learning development"
		invoice.Quantities[i] = 1.5
		invoice.Rates[i] = 25
	}

	writeLogo(pdf, invoice.Logo, invoice.From)
	writeTitle(pdf, invoice.Title, invoice.Id, invoice.Date)
	writeBillTo(pdf, invoice.To)
	writeHeaderRow(pdf, invoice.Hourly)
	subtotal := writeItemRows(pdf, invoice)

	if got, want := pdf.GetNumberOfPages(), 2; got != want {
		t.Fatalf("page count = %d, want %d", got, want)
	}
	if want := 13 * 1.5 * 25; math.Abs(subtotal-want) > 0.001 {
		t.Errorf("subtotal = %v, want %v", subtotal, want)
	}
	if pdf.GetY() > summaryY {
		t.Errorf("item rows ended at y=%v, below summary boundary %v", pdf.GetY(), summaryY)
	}
}

func TestQuantityHeader(t *testing.T) {
	tests := []struct {
		name   string
		hourly bool
		want   string
	}{
		{name: "quantity mode", want: "QTY"},
		{name: "hourly mode", hourly: true, want: "HOURS"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := quantityHeader(tt.hourly); got != tt.want {
				t.Errorf("quantityHeader(%v) = %q, want %q", tt.hourly, got, tt.want)
			}
		})
	}
}

func newTestPDF(t *testing.T) *gopdf.GoPdf {
	t.Helper()

	pdf := &gopdf.GoPdf{}
	pdf.Start(gopdf.Config{PageSize: *gopdf.PageSizeA4})
	pdf.SetMargins(40, 40, 40, 40)
	pdf.AddPage()
	if err := pdf.AddTTFFontData("Inter", interFont); err != nil {
		t.Fatalf("add Inter font: %v", err)
	}
	if err := pdf.AddTTFFontData("Inter-Bold", interBoldFont); err != nil {
		t.Fatalf("add Inter Bold font: %v", err)
	}
	return pdf
}
