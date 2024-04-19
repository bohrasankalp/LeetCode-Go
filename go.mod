module github.com/halfrost/LeetCode-Go

go 1.19

replace github.com/halfrost/LeetCode-Go/structures => ./structures

replace github.com/halfrost/LeetCode-Go/template => ./template

replace github.com/halfrost/LeetCode-Go/ctl/util => ./ctl/util

replace github.com/halfrost/LeetCode-Go/ctl/models => ./ctl/models

require (
	github.com/BurntSushi/toml v1.2.1
	github.com/halfrost/LeetCode-Go/ctl/models v0.0.0-20221125083011-ed716f02a412
	github.com/halfrost/LeetCode-Go/ctl/util v0.0.0-20221125083011-ed716f02a412
	github.com/halfrost/LeetCode-Go/structures v0.0.0-20221125083011-ed716f02a412
	github.com/halfrost/LeetCode-Go/template v0.0.0-20221125083011-ed716f02a412
	github.com/mozillazg/request v0.8.0
	github.com/spf13/cobra v1.6.1
)

require (
	github.com/bitly/go-simplejson v0.5.0 // indirect
	github.com/bmizerany/assert v0.0.0-20160611221934-b7ed37b82869 // indirect
	github.com/inconshreveable/mousetrap v1.1.0 // indirect
	github.com/kr/pretty v0.3.0 // indirect
	github.com/spf13/pflag v1.0.5 // indirect
	golang.org/x/net v0.23.0 // indirect
)
