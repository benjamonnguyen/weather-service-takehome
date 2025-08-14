package nws

import (
	"testing"
)

var sut = NewApiClient()

func TestGetPoint(t *testing.T) {
	resp, err := sut.GetPointForecasts(39.7456, -97.0892)
	if err != nil {
		t.Fail()
	}
	t.Logf("%+v\n", resp)
}
