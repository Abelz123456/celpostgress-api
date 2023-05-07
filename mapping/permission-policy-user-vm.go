package mapping

import "github.com/celpostgress-api/entity"

type PermissionPolicyUserVm struct {
	Oid         string  `json:"oid"`
	CompanyName *string `json:"companyName"`
	Address     *string `json:"address"`
	EmailName   string  `json:"emailName"`
	Password    string  `json:"password"`
	Token       *string `json:"token"`
}

type PermissionPolicyUserAuthVm struct {
	Oid          string  `json:"oid"`
	CompanyName  *string `json:"companyName"`
	Address      *string `json:"address"`
	EmailName    string  `json:"emailName"`
	Token        *string `json:"token"`
	RefreshToken *string `json:"refreshToken"`
	Role         *string `json:"role"`
}

func ToPermissionPolicyUserResponse(permissionPolicyUser entity.PermissionPolicyUser) PermissionPolicyUserVm {
	return PermissionPolicyUserVm{
		Oid:       permissionPolicyUser.Oid,
		EmailName: permissionPolicyUser.EmailName,
		Password:  permissionPolicyUser.Password,
	}
}

func ToPermissionPolicyUserResponses(permissionPolicyUsers []entity.PermissionPolicyUser) []PermissionPolicyUserVm {
	var permissionPolicyUserResponses []PermissionPolicyUserVm
	for _, permissionPolicyUser := range permissionPolicyUsers {
		permissionPolicyUserResponses = append(permissionPolicyUserResponses, ToPermissionPolicyUserResponse(permissionPolicyUser))
	}
	return permissionPolicyUserResponses
}

func ToPermissionPolicyUserResponseAuth(permissionPolicyUser entity.PermissionPolicyUser,
	token string, refreshToken string, role string) PermissionPolicyUserAuthVm {
	return PermissionPolicyUserAuthVm{
		Oid:       permissionPolicyUser.Oid,
		EmailName: permissionPolicyUser.EmailName,
		// Password:    permissionPolicyUser.Password,
		Token:        &token,
		RefreshToken: &refreshToken,
		Role:         &role,
	}
}
