// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: users.proto

package api_pb

import (
	fmt "fmt"
	math "math"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/golang/protobuf/ptypes/empty"
	_ "github.com/ProgrammingLab/prolab-accounts/api/type"
	_ "."
	_ "google.golang.org/genproto/googleapis/api/annotations"
	regexp "regexp"
	github_com_mwitkow_go_proto_validators "github.com/mwitkow/go-proto-validators"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

func (this *User) Validate() error {
	if this.Role != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Role); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Role", err)
		}
	}
	if this.Department != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Department); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Department", err)
		}
	}
	if this.Icon != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Icon); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Icon", err)
		}
	}
	return nil
}
func (this *ListUsersRequest) Validate() error {
	if !(len(this.Query) < 128) {
		return github_com_mwitkow_go_proto_validators.FieldError("Query", fmt.Errorf(`value '%v' must length be less than '128'`, this.Query))
	}
	return nil
}
func (this *ListUsersResponse) Validate() error {
	for _, item := range this.Users {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Users", err)
			}
		}
	}
	return nil
}
func (this *GetUserRequest) Validate() error {
	return nil
}
func (this *UpdateRoleRequest) Validate() error {
	return nil
}

var _regex_CreateUserRequest_Name = regexp.MustCompile(`^[A-Za-z0-9_]{1,20}$`)

func (this *CreateUserRequest) Validate() error {
	if !_regex_CreateUserRequest_Name.MatchString(this.Name) {
		return github_com_mwitkow_go_proto_validators.FieldError("Name", fmt.Errorf(`value '%v' must be a string conforming to regex "^[A-Za-z0-9_]{1,20}$"`, this.Name))
	}
	if !(len(this.FullName) < 128) {
		return github_com_mwitkow_go_proto_validators.FieldError("FullName", fmt.Errorf(`value '%v' must length be less than '128'`, this.FullName))
	}
	if !(len(this.RegisterationToken) < 128) {
		return github_com_mwitkow_go_proto_validators.FieldError("RegisterationToken", fmt.Errorf(`value '%v' must length be less than '128'`, this.RegisterationToken))
	}
	return nil
}
func (this *GetCurrentUserRequest) Validate() error {
	return nil
}
func (this *UpdateUserProfileRequest) Validate() error {
	if !(len(this.FullName) < 128) {
		return github_com_mwitkow_go_proto_validators.FieldError("FullName", fmt.Errorf(`value '%v' must length be less than '128'`, this.FullName))
	}
	if !(len(this.Description) < 1024) {
		return github_com_mwitkow_go_proto_validators.FieldError("Description", fmt.Errorf(`value '%v' must length be less than '1024'`, this.Description))
	}
	if !(this.Grade > 0) {
		return github_com_mwitkow_go_proto_validators.FieldError("Grade", fmt.Errorf(`value '%v' must be greater than '0'`, this.Grade))
	}
	if !(this.Grade < 6) {
		return github_com_mwitkow_go_proto_validators.FieldError("Grade", fmt.Errorf(`value '%v' must be less than '6'`, this.Grade))
	}
	if !(len(this.TwitterScreenName) < 128) {
		return github_com_mwitkow_go_proto_validators.FieldError("TwitterScreenName", fmt.Errorf(`value '%v' must length be less than '128'`, this.TwitterScreenName))
	}
	if !(len(this.GithubUserName) < 128) {
		return github_com_mwitkow_go_proto_validators.FieldError("GithubUserName", fmt.Errorf(`value '%v' must length be less than '128'`, this.GithubUserName))
	}
	if _, ok := ProfileScope_name[int32(this.ProfileScope)]; !ok {
		return github_com_mwitkow_go_proto_validators.FieldError("ProfileScope", fmt.Errorf(`value '%v' must be a valid ProfileScope field`, this.ProfileScope))
	}
	if !(len(this.AtcoderUserName) < 128) {
		return github_com_mwitkow_go_proto_validators.FieldError("AtcoderUserName", fmt.Errorf(`value '%v' must length be less than '128'`, this.AtcoderUserName))
	}
	if !(len(this.DisplayName) < 51) {
		return github_com_mwitkow_go_proto_validators.FieldError("DisplayName", fmt.Errorf(`value '%v' must length be less than '51'`, this.DisplayName))
	}
	return nil
}
func (this *UpdateUserIconRequest) Validate() error {
	return nil
}
