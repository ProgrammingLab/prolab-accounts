module github.com/ProgrammingLab/prolab-accounts

go 1.12

require (
	contrib.go.opencensus.io/exporter/stackdriver v0.7.0 // indirect
	git.apache.org/thrift.git v0.0.0-20180902110319-2566ecd5d999 // indirect
	github.com/PuerkitoBio/goquery v1.5.0 // indirect
	github.com/go-redis/redis v6.14.1+incompatible
	github.com/gobuffalo/envy v1.7.0 // indirect
	github.com/gobuffalo/packr/v2 v2.1.0
	github.com/gobwas/glob v0.2.3
	github.com/gogo/protobuf v1.2.1
	github.com/golang/lint v0.0.0-20180702182130-06c8688daad7 // indirect
	github.com/golang/protobuf v1.3.2
	github.com/grpc-ecosystem/grpc-gateway v1.9.6
	github.com/izumin5210/grapi v0.4.0
	github.com/joho/godotenv v1.3.0
	github.com/jordan-wright/email v0.0.0-20180115032944-94ae17dedda2
	github.com/kat-co/vala v0.0.0-20140812221447-bfe9b50e828a
	github.com/kelseyhightower/envconfig v1.3.0
	github.com/lib/pq v1.0.0
	github.com/minio/minio-go v0.0.0-20190207005939-34f2b94c43ca
	github.com/mmcdole/gofeed v1.0.0-beta2
	github.com/mmcdole/goxpp v0.0.0-20181012175147-0068e33feabf // indirect
	github.com/mwitkow/go-proto-validators v0.0.0-20190212092829-1f388280e944
	github.com/openzipkin/zipkin-go v0.1.1 // indirect
	github.com/ory/hydra v1.0.3
	github.com/ory/hydra-legacy-sdk v0.0.0-20190409103449-1f564942be76
	github.com/ory/sqlcon v0.0.7 // indirect
	github.com/pkg/errors v0.8.1
	github.com/rs/cors v1.6.0
	github.com/shurcooL/githubv4 v0.0.0-20190119021625-d9689b595017
	github.com/shurcooL/graphql v0.0.0-20181231061246-d48a9a75455f // indirect
	github.com/spf13/viper v1.4.0
	github.com/volatiletech/inflect v0.0.0-20170731032912-e7201282ae8d // indirect
	github.com/volatiletech/null v8.0.0+incompatible
	github.com/volatiletech/sqlboiler v3.2.0+incompatible
	github.com/xanzy/go-gitlab v0.13.0 // indirect
	golang.org/x/crypto v0.0.0-20190701094942-4def268fd1a4
	golang.org/x/image v0.0.0-20190802002840-cff245a6509b
	golang.org/x/net v0.0.0-20190813141303-74dc4d7220e7
	golang.org/x/oauth2 v0.0.0-20190604053449-0f29369cfe45
	golang.org/x/sync v0.0.0-20190423024810-112230192c58
	google.golang.org/genproto v0.0.0-20190817000702-55e96fffbd48
	google.golang.org/grpc v1.23.0
	gopkg.in/ini.v1 v1.42.0 // indirect
	gopkg.in/yaml.v1 v1.0.0-20140924161607-9f9df34309c0 // indirect
)

replace github.com/gofrs/uuid v0.0.0-20180830191909-370558f003bf => github.com/gofrs/uuid/v3 v3.1.1

replace github.com/google/go-github v0.0.0-20181222022713-a5cb647b1fac => github.com/google/go-github/v21 v21.0.0
