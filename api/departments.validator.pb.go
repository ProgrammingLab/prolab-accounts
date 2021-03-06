// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: departments.proto

package api_pb

import (
	fmt "fmt"
	math "math"
	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	_ "github.com/golang/protobuf/ptypes/empty"
	github_com_mwitkow_go_proto_validators "github.com/mwitkow/go-proto-validators"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

func (this *Department) Validate() error {
	return nil
}
func (this *ListDepartmentsRequest) Validate() error {
	return nil
}
func (this *ListDepartmentsResponse) Validate() error {
	for _, item := range this.Departments {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Departments", err)
			}
		}
	}
	return nil
}
func (this *GetDepartmentRequest) Validate() error {
	return nil
}
