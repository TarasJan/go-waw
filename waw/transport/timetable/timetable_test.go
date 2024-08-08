package timetable

import "testing"

func TestZTMTimeNormalization24(t *testing.T) {
	timestamp := "24:15:00"
	expected := "00:15:00"
	got := normalizeZTMTime(timestamp)
	if got != expected {
		t.Errorf("Normalization failed. Expected: %s, Got: %s", expected, got)
	}
}

func TestZTMTimeNormalizationOk(t *testing.T) {
	timestamp := "17:15:00"
	expected := "17:15:00"
	got := normalizeZTMTime(timestamp)
	if got != expected {
		t.Errorf("Normalization failed. Expected: %s, Got: %s", expected, got)
	}
}
