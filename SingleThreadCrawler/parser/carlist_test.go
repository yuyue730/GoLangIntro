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

	results := ParseCarList(contents, "")

	const resultUrlSize = 89
	if len(results.Requests) != resultUrlSize {
		t.Errorf("result size should have %d, but had %d", resultUrlSize,
			len(results.Requests))
	}

	expectedModelUrls := []string{
		"http://newcar.xcar.com.cn/3428/",
		"http://newcar.xcar.com.cn/3796/",
		"http://newcar.xcar.com.cn/4063/",
	}

	for i, url := range expectedModelUrls {
		if results.Requests[i].Url != url {
			t.Errorf("model result index %d url should have %s, but had %s", i,
				results.Requests[i].Url, url)
		}
	}

	const resultModelSize = 20
	expectedOtherListUrls := []string{
		"http://newcar.xcar.com.cn/car/0-0-0-0-0-0-0-0-0-0-0-1",
		"http://newcar.xcar.com.cn/car/0-0-0-0-0-0-6-0-0-0-0-1",
		"http://newcar.xcar.com.cn/car/0-0-0-0-0-0-8-0-0-0-0-1",
	}

	for i, url := range expectedOtherListUrls {
		if results.Requests[resultModelSize+i].Url != url {
			t.Errorf("other list result index %d url should have %s, but had %s",
				i, expectedOtherListUrls[i],
				results.Requests[resultModelSize+i].Url)
		}
	}
}
