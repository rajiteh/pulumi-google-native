// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v1

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Gets the definition of a Role.
func LookupOrganizationRole(ctx *pulumi.Context, args *LookupOrganizationRoleArgs, opts ...pulumi.InvokeOption) (*LookupOrganizationRoleResult, error) {
	var rv LookupOrganizationRoleResult
	err := ctx.Invoke("google-native:iam/v1:getOrganizationRole", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupOrganizationRoleArgs struct {
	OrganizationId string `pulumi:"organizationId"`
	RoleId         string `pulumi:"roleId"`
}

type LookupOrganizationRoleResult struct {
	// The current deleted state of the role. This field is read only. It will be ignored in calls to CreateRole and UpdateRole.
	Deleted bool `pulumi:"deleted"`
	// Optional. A human-readable description for the role.
	Description string `pulumi:"description"`
	// Used to perform a consistent read-modify-write.
	Etag string `pulumi:"etag"`
	// The names of the permissions this role grants when bound in an IAM policy.
	IncludedPermissions []string `pulumi:"includedPermissions"`
	// The name of the role. When Role is used in CreateRole, the role name must not be set. When Role is used in output and other input such as UpdateRole, the role name is the complete path, e.g., roles/logging.viewer for predefined roles and organizations/{ORGANIZATION_ID}/roles/logging.viewer for custom roles.
	Name string `pulumi:"name"`
	// The current launch stage of the role. If the `ALPHA` launch stage has been selected for a role, the `stage` field will not be included in the returned definition for the role.
	Stage string `pulumi:"stage"`
	// Optional. A human-readable title for the role. Typically this is limited to 100 UTF-8 bytes.
	Title string `pulumi:"title"`
}