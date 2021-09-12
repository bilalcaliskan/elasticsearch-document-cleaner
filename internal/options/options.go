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
	// TODO: implement
}

func (edco *ElasticsearchDocumentCleanerOptions) addFlags(fs *pflag.FlagSet) {
	// TODO: implement
}
