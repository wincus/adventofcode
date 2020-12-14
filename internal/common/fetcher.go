package common

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strings"
)

// GetData retrieves data used to solve day n
func GetData(n int) ([]string, error) {

	h, ok := os.LookupEnv("SESSION")

	if !ok {
		return nil, fmt.Errorf("SESSION env not found")
	}

	u := &url.URL{
		Scheme: "https",
		Host:   "adventofcode.com",
		Path:   fmt.Sprintf("%v/day/%v/input", 2020, n),
	}

	req, err := http.NewRequest(http.MethodGet, u.String(), nil)

	if err != nil {
		return nil, fmt.Errorf("could not create request: %v", err)
	}

	req.Header.Set("cookie", fmt.Sprintf("session=%v", h))

	c := &http.Client{}

	res, err := c.Do(req)

	b, err := ioutil.ReadAll(res.Body)

	if err != nil {
		return nil, fmt.Errorf("could not read response body: %v", err)
	}

	defer res.Body.Close()

	return strings.Split(string(b), "\n"), nil

}

func ShowData(d []string) {

	for _, s := range d {
		fmt.Printf("%v\n", string(s))
	}

}
