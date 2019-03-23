// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: achievements.proto

package api_pb

import (
	fmt "fmt"
	math "math"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/mwitkow/go-proto-validators"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	_ "github.com/golang/protobuf/ptypes/empty"
	_ "github.com/golang/protobuf/ptypes/timestamp"
	github_com_mwitkow_go_proto_validators "github.com/mwitkow/go-proto-validators"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

func (this *Achievement) Validate() error {
	if this.HappenedAt != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.HappenedAt); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("HappenedAt", err)
		}
	}
	for _, item := range this.Members {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Members", err)
			}
		}
	}
	return nil
}
func (this *ListAchievementsRequest) Validate() error {
	return nil
}
func (this *ListAchievementsResponse) Validate() error {
	for _, item := range this.Achievements {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Achievements", err)
			}
		}
	}
	return nil
}
func (this *GetAchievementRequest) Validate() error {
	return nil
}
func (this *CreateAchievementRequest) Validate() error {
	if this.Achievement != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Achievement); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Achievement", err)
		}
	}
	return nil
}
func (this *UpdateAchievementRequest) Validate() error {
	if this.Achievement != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Achievement); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Achievement", err)
		}
	}
	return nil
}
func (this *DeleteAchievementRequest) Validate() error {
	return nil
}
