package options

import (
	"github.com/spf13/pflag"
)

var edco = &ElasticsearchDocumentCleanerOptions{}

func init() {
	edco.addFlags(pflag.CommandLine)
	pflag.Parse()
}

// GetElasticsearchDocumentCleanerOptions returns the pointer of VarnishCacheInvalidatorOptions
func GetElasticsearchDocumentCleanerOptions() *ElasticsearchDocumentCleanerOptions {
	return edco
}

// ElasticsearchDocumentCleanerOptions contains frequent command line and application options.
type ElasticsearchDocumentCleanerOptions struct {
	// ElasticsearchUrl is the url of the target Elasticsearch instance. It may be ip:port or DNS. ex: http://localhost:9200
	ElasticsearchUrl string
	// IndexName is the target index to clean documents
	IndexName string
	// TimestampField is the timestamp field specifier for our documents in index to pick correct documents. Defaults to @timestamp
	TimestampField string
	// Operator is the specifier to pick correct documents according to TimestampField, possible values are lt, lte, gt, gte. Defaults to lt
	Operator string
	// Mode is the specifier that our older format, possible values are minute, hour, day. Defaults to day
	Mode string
	// Value is the specifier to delete documents with specified Mode and Operator, defaults to 30
	Value int32
}

func (edco *ElasticsearchDocumentCleanerOptions) addFlags(fs *pflag.FlagSet) {
	fs.StringVar(&edco.ElasticsearchUrl, "elasticsearchUrl", "http://localhost:9200", "url of the "+
		"target Elasticsearch instance. It may be ip:port or DNS. defaults to http://localhost:9200")
	fs.StringVar(&edco.IndexName, "indexName", "test-index", "target index to clean documents, defaults "+
		"to test-index")
	fs.StringVar(&edco.TimestampField, "timestampField", "@timestamp", "timestamp field specifier for "+
		"our documents in index to pick correct documents. Defaults to @timestamp")
	fs.StringVar(&edco.Operator, "operator", "lt", "specifier to pick correct documents according "+
		"to timestampField, possible values are lt, lte, gt, gte. Defaults to lt")
	fs.StringVar(&edco.Mode, "mode", "day", "specifier that our older format, possible values are "+
		"minute, hour, day. Defaults to day")
	fs.Int32Var(&edco.Value, "value", 30, "specifier to delete documents with specified mode and "+
		"operator, defaults to 30")
}
