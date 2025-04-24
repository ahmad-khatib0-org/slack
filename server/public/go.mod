module github.com/ahmad-khatib0-org/slack/server/public

go 1.24.2

require (
	github.com/blang/semver/v4 v4.0.0
	github.com/dyatlov/go-opengraph/opengraph v0.0.0-20220524092352-606d7b1e5f8a
	github.com/francoispqt/gojay v1.2.13
	github.com/go-sql-driver/mysql v1.9.2
	github.com/gorilla/websocket v1.5.3
	github.com/hashicorp/go-hclog v1.6.3
	github.com/hashicorp/go-multierror v1.1.1
	github.com/hashicorp/go-plugin v1.6.3
	github.com/lib/pq v1.10.9
	github.com/mattermost/go-i18n v1.11.0
	github.com/mattermost/gosaml2 v0.8.0
	github.com/mattermost/ldap v3.0.4+incompatible
	github.com/mattermost/logr/v2 v2.0.22
	github.com/pborman/uuid v1.2.1
	github.com/pkg/errors v0.9.1
	github.com/stretchr/testify v1.10.0
	github.com/tinylib/msgp v1.2.5
	github.com/vmihailenco/msgpack/v5 v5.4.1
	golang.org/x/crypto v0.37.0
	golang.org/x/mod v0.24.0
	golang.org/x/net v0.39.0
	golang.org/x/text v0.24.0
	golang.org/x/tools v0.32.0
	gopkg.in/yaml.v3 v3.0.1
)

require (
	filippo.io/edwards25519 v1.1.0 // indirect
	github.com/beevik/etree v1.1.0 // indirect
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/fatih/color v1.13.0 // indirect
	github.com/golang/protobuf v1.5.3 // indirect
	github.com/google/uuid v1.3.0 // indirect
	github.com/hashicorp/errwrap v1.0.0 // indirect
	github.com/hashicorp/yamux v0.1.1 // indirect
	github.com/jonboulle/clockwork v0.2.2 // indirect
	github.com/mattermost/xml-roundtrip-validator v0.1.0 // indirect
	github.com/mattn/go-colorable v0.1.12 // indirect
	github.com/mattn/go-isatty v0.0.17 // indirect
	github.com/oklog/run v1.0.0 // indirect
	github.com/pelletier/go-toml v1.2.0 // indirect
	github.com/philhofer/fwd v1.1.3-0.20240916144458-20a13a1f6b7c // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/russellhaering/goxmldsig v1.2.0 // indirect
	github.com/stretchr/objx v0.5.2 // indirect
	github.com/vmihailenco/tagparser/v2 v2.0.0 // indirect
	github.com/wiggin77/merror v1.0.2 // indirect
	github.com/wiggin77/srslog v1.0.1 // indirect
	golang.org/x/sync v0.13.0 // indirect
	golang.org/x/sys v0.32.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20230711160842-782d3b101e98 // indirect
	google.golang.org/grpc v1.58.3 // indirect
	google.golang.org/protobuf v1.36.1 // indirect
	gopkg.in/asn1-ber.v1 v1.0.0-20181015200546-f715ec2f112d // indirect
	gopkg.in/natefinch/lumberjack.v2 v2.0.0 // indirect
	gopkg.in/yaml.v2 v2.2.2 // indirect
)

// Hack to prevent the willf/bitset module from being upgraded to 1.2.0.
// They changed the module path from github.com/willf/bitset to
// github.com/bits-and-blooms/bitset and a couple of dependent repos are yet
// to update their module paths.
exclude (
	github.com/RoaringBitmap/roaring v0.7.0
	github.com/RoaringBitmap/roaring v0.7.1
	github.com/dyatlov/go-opengraph v0.0.0-20210112100619-dae8665a5b09
	github.com/willf/bitset v1.2.0
)
