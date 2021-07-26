package myhttp

import "flag"

// ParseCLI parse command line args
func ParseCLI() (int, []string) {
	limit := flag.Int("limit", 10, "int that indicates the limit of concurrent request possible")
	flag.Parse()

	urls := flag.Args()

	return *limit, urls
}
