package osmdata

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"
	"time"

	ovp "github.com/captchanjack/osmdata/overpass"
)

func TestPresetQuery(t *testing.T) {
	var testCases = []struct {
		queryMethod QueryMethod
		args        []interface{}
	}{
		{
			Radius,
			[]interface{}{
				Drive,
				true,
				ovp.JSON,
				5.0,
				-37.740347,
				144.930127,
			},
		},
		{
			Polygon,
			[]interface{}{
				Walk,
				false,
				ovp.XML,
				[][][]float64{
					{
						{
							145.10180044651912,
							-37.8444294644357,
						},
						{
							145.10180044651912,
							-37.845678389372324,
						},
						{
							145.10359037007055,
							-37.845678389372324,
						},
						{
							145.10359037007055,
							-37.8444294644357,
						},
						{
							145.10180044651912,
							-37.8444294644357,
						},
					},
				},
			},
		},
		{
			BoundingBox,
			[]interface{}{
				Bike,
				true,
				ovp.CSV,
				-37.845678389372324,
				145.10180044651912,
				-37.8444294644357,
				145.10359037007055,
				"(::id,::type,::lat,::lon,true,\",\")", // CSV header selection, required
			},
		},
		{
			PlaceName,
			[]interface{}{
				Rail,
				true,
				ovp.JSON,
				"henley, new south wales, australia",
			},
		},
	}

	for _, test := range testCases {
		query := GetPresetQuery(test.queryMethod, test.args...)
		resp, err := query.Execute(10)

		if len(resp) == 0 {
			t.Errorf("%v: string response should be non-empty\n", test.queryMethod)
		}

		if err != nil {
			t.Errorf("%v: %v\n", test.queryMethod, err)
		}

		if len(resp) != 0 && err == nil {
			fmt.Printf("%v: test successful\n", test.queryMethod)
		}

		time.Sleep(10 * time.Second) // prevent getting rated limited
	}
}

func TestExport(t *testing.T) {
	time.Sleep(10 * time.Second) // prevent getting rated limited
	
	cwd, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	filename := filepath.Join(filepath.Dir(cwd), "osmdata/test.json")

	query := GetPresetQuery(
		Radius,
		Drive,
		true,
		ovp.JSON,
		5.0,
		-37.740347,
		144.930127,
	)
	resp, err := query.ExecuteAndExport(filename, 10)

	defer os.Remove(filename)

	if len(resp) == 0 {
		t.Errorf("%v: string response should be non-empty\n", Radius)
	}

	if err != nil {
		t.Errorf("%v: %v\n", Radius, err)
	}

	if len(resp) != 0 && err == nil {
		fmt.Println("Export: test successful")
	}

	time.Sleep(10 * time.Second) // prevent getting rated limited
}
