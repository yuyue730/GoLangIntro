// Unit test for ParseCarModel
package parser

import (
	"fmt"
	"io/ioutil"
	"testing"
)

func TestParseCarModel(t *testing.T) {
	contents, err := ioutil.ReadFile("./carmodel_test_data.html")

	if err != nil {
		panic(err)
	}

	results := ParseCarModel(contents, "")

	const resultURLSize = 5
	if len(results.Requests) != resultURLSize {
		t.Errorf("result size should have %d, but had %d", resultURLSize,
			len(results.Requests))
	}

	expectedDetailUrls := []string{
		"http://newcar.xcar.com.cn/m49497/",
		"http://newcar.xcar.com.cn/m49498/",
		"http://newcar.xcar.com.cn/m49499/",
		"http://newcar.xcar.com.cn/m49500/",
		"http://newcar.xcar.com.cn/m43936/",
	}

	fmt.Println(results.Requests)

	for i, url := range expectedDetailUrls {
		if results.Requests[i].Url != url {
			t.Errorf("model result index %d url should have %s, but had %s", i,
				results.Requests[i].Url, url)
		}
	}
}
