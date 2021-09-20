package serice_test

import (
	"module/series"
	"testing"
)

func TestSeries(t *testing.T) {
	t.Log(series.GetFibonacciSerie(10))
}