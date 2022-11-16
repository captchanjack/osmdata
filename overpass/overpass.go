package overpass

import (
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"
)

const (
	overpassInterpreterEndpoint = "https://overpass-api.de/api/interpreter"
	overpassStatusEndpoint      = "https://overpass-api.de/api/status"
)

// Specify maxAttempts tp retry when rate limited, defaults to 1
// Returns string
func QueryOverpass(query string, maxAttempts ...int) (response string, err error) {
	body, err := QueryOverpassBytes(query, maxAttempts...)
	sb := string(body)
	return sb, err
}

// Specify maxAttempts tp retry when rate limited, defaults to 1
// Returns byte array
func QueryOverpassBytes(query string, maxAttempts ...int) (response []byte, err error) {
	var isAvailable bool
	_maxAttempts := 1
	if len(maxAttempts) > 0 {
		_maxAttempts = maxAttempts[0]
	}
	i := 0

	for i < _maxAttempts {
		isAvailable, err = overpassIsAvailable()

		if isAvailable {
			break
		}

		if err != nil {
			fmt.Println(err)
		}

		i++

		if i < _maxAttempts {
			fmt.Println("Sleeping 5 seconds before trying again")
			time.Sleep(5 * time.Second)
		}
	}

	if !isAvailable {
		return make([]byte, 0), fmt.Errorf("overpass server is not available for download, please wait and try again")
	}

	if err != nil {
		return make([]byte, 0), fmt.Errorf("encountered error during POST request to Overpass API: %w", err)
	}

	resp, err := http.Post(overpassInterpreterEndpoint, "application/json", strings.NewReader(query))

	if err != nil {
		return make([]byte, 0), fmt.Errorf("encountered error during POST request to Overpass API: %w", err)
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)

	if resp.StatusCode != 200 {
		return make([]byte, 0), fmt.Errorf("overpass engine error (status code %v):\n%s", resp.StatusCode, body)
	}

	if err != nil {
		return make([]byte, 0), fmt.Errorf("encountered error during POST request to Overpass API: %w", err)
	}

	return body, err
}

func overpassIsAvailable() (isAvailable bool, err error) {
	resp, err := http.Get(overpassStatusEndpoint)

	if err != nil {
		return false, fmt.Errorf("encountered error during GET request to Overpass Status API: %w", err)
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)

	if err != nil {
		return false, fmt.Errorf("encountered error during POST request to Overpass Status API: %w", err)
	}

	sb := string(body)

	if resp.StatusCode != 200 && strings.Contains(sb, "slots available after") {
		return false, fmt.Errorf("overpass server is not available for download, please wait and try again:\n%s", sb)
	}

	return true, nil
}
