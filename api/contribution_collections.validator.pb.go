// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: contribution_collections.proto

package api_pb

import (
	fmt "fmt"
	math "math"
	proto "github.com/golang/protobuf/proto"
	_ "."
	_ "google.golang.org/genproto/googleapis/api/annotations"
	_ "github.com/golang/protobuf/ptypes/empty"
	_ "github.com/golang/protobuf/ptypes/timestamp"
	github_com_mwitkow_go_proto_validators "github.com/mwitkow/go-proto-validators"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

func (this *ContributionCollection) Validate() error {
	if this.User != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.User); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("User", err)
		}
	}
	for _, item := range this.Days {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Days", err)
			}
		}
	}
	return nil
}
func (this *ContributionDay) Validate() error {
	if this.Date != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Date); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Date", err)
		}
	}
	return nil
}
func (this *ListContributionCollectionsRequest) Validate() error {
	if !(this.UsersCount > -1) {
		return github_com_mwitkow_go_proto_validators.FieldError("UsersCount", fmt.Errorf(`value '%v' must be greater than '-1'`, this.UsersCount))
	}
	if !(this.UsersCount < 101) {
		return github_com_mwitkow_go_proto_validators.FieldError("UsersCount", fmt.Errorf(`value '%v' must be less than '101'`, this.UsersCount))
	}
	return nil
}
func (this *ListContributionCollectionsResponse) Validate() error {
	for _, item := range this.ContributionCollections {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("ContributionCollections", err)
			}
		}
	}
	return nil
}
