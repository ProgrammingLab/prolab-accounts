module github.com/ProgrammingLab/prolab-accounts

go 1.12

require (
	github.com/friendsofgo/errors v0.9.2 // indirect
	github.com/go-redis/redis v6.15.9+incompatible
	github.com/gobuffalo/packr/v2 v2.8.0
	github.com/gobwas/glob v0.2.3
	github.com/gogo/protobuf v1.3.1
	github.com/golang/protobuf v1.4.2
	github.com/google/wire v0.3.0 // indirect
	github.com/grpc-ecosystem/grpc-gateway v1.14.7
	github.com/izumin5210/gex v0.6.0 // indirect
	github.com/izumin5210/grapi v0.5.0
	github.com/joho/godotenv v1.3.0
	github.com/jordan-wright/email v0.0.0-20180115032944-94ae17dedda2
	github.com/kat-co/vala v0.0.0-20140812221447-bfe9b50e828a
	github.com/kelseyhightower/envconfig v1.4.0
	github.com/lib/pq v1.8.0
	github.com/minio/minio-go v0.0.0-20190207005939-34f2b94c43ca
	github.com/mmcdole/gofeed v1.0.0
	github.com/mwitkow/go-proto-validators v0.0.0-20190212092829-1f388280e944
	github.com/ory/hydra v1.7.4
	github.com/ory/hydra-legacy-sdk v0.0.0-20190409103449-1f564942be76
	github.com/pkg/errors v0.9.1
	github.com/rs/cors v1.7.0
	github.com/shurcooL/githubv4 v0.0.0-20190119021625-d9689b595017
	github.com/shurcooL/graphql v0.0.0-20181231061246-d48a9a75455f // indirect
	github.com/spf13/viper v1.7.1
	github.com/srvc/appctx v0.1.0
	github.com/volatiletech/inflect v0.0.0-20170731032912-e7201282ae8d // indirect
	github.com/volatiletech/null v8.0.0+incompatible
	github.com/volatiletech/sqlboiler v3.7.1+incompatible
	golang.org/x/crypto v0.0.0-20200510223506-06a226fb4e37
	golang.org/x/image v0.0.0-20190802002840-cff245a6509b
	golang.org/x/net v0.0.0-20200324143707-d3edc9973b7e
	golang.org/x/oauth2 v0.0.0-20200107190931-bf48bf16ab8d
	golang.org/x/sync v0.0.0-20200317015054-43a5402ce75a
	google.golang.org/genproto v0.0.0-20200513103714-09dca8ec2884
	google.golang.org/grpc v1.31.1
	gopkg.in/DATA-DOG/go-sqlmock.v1 v1.3.0 // indirect
	k8s.io/utils v0.0.0-20191010214722-8d271d903fe4 // indirect
)

replace github.com/gofrs/uuid v0.0.0-20180830191909-370558f003bf => github.com/gofrs/uuid/v3 v3.1.1

replace github.com/google/go-github v0.0.0-20181222022713-a5cb647b1fac => github.com/google/go-github/v21 v21.0.0
