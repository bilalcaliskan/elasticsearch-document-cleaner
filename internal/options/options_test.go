package options

import "testing"

func TestGetElasticsearchDocumentCleanerOptions(t *testing.T) {
	t.Log("getting options")
	cases := []struct {
		caseName string
		edco     ElasticsearchDocumentCleanerOptions
	}{
		{"case1", ElasticsearchDocumentCleanerOptions{
			ElasticsearchUrl: "http://localhost:9200",
			Mode:             "day",
			Operator:         "lt",
			TimestampField:   "@timestamp",
			IndexName:        "test-index",
			Value:            30,
		}},
		{"case1", ElasticsearchDocumentCleanerOptions{
			ElasticsearchUrl: "http://localhost:9200",
			Mode:             "month",
			Operator:         "gt",
			TimestampField:   "time",
			IndexName:        "test-index-2",
			Value:            30,
		}},
	}

	for _, tc := range cases {
		t.Logf("generated options on case %s: %v\n", tc.caseName, tc.edco)
	}
}
