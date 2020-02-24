// Unit test for ParseCarList
package parser

import (
	"io/ioutil"
	"testing"
)

func TestParseCarList(t *testing.T) {
	contents, err := ioutil.ReadFile("./carlist_test_data.html")

	if err != nil {
		panic(err)
	}

	results := ParseCarList(contents)

	const resultModelSize = 20
	if len(results.Requests) != resultModelSize {
		t.Errorf("result size should have %d, but had %d", resultModelSize, 
			len(results.Requests))
	}

	expectedModelUrls := []string {
		"http://newcar.xcar.com.cn/3428/",
		"http://newcar.xcar.com.cn/3796/",
		"http://newcar.xcar.com.cn/4063/",
	}

	for i, url := range expectedModelUrls {
		if results.Requests[i].Url != url {
			t.Errorf("result index %d url should have %s, but had %s", i, 
				expectedModelUrls[i], url)
		}
	}
}