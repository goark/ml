module github.com/spiegel-im-spiegel/mklink

go 1.13

require (
	github.com/PuerkitoBio/goquery v1.5.1
	github.com/atotto/clipboard v0.1.2
	github.com/mattn/go-encoding v0.0.2
	github.com/spf13/cobra v1.0.1-0.20201001152800-40d34bca1bff
	github.com/spiegel-im-spiegel/errs v1.0.2
	github.com/spiegel-im-spiegel/gocli v0.10.3
	github.com/spiegel-im-spiegel/gprompt v0.9.7
	golang.org/x/net v0.0.0-20200930145003-4acb6c075d10
)

replace github.com/coreos/etcd v3.3.13+incompatible => github.com/coreos/etcd v3.3.25+incompatible
