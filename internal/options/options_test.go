package options

import "testing"

func TestGetElasticsearchDocumentCleanerOptions(t *testing.T) {
	t.Log("fetching default options.ElasticsearchDocumentCleanerOptions")
	opts := GetElasticsearchDocumentCleanerOptions()
	t.Logf("fetched default options.ElasticsearchDocumentCleanerOptions, %v\n", opts)
}
