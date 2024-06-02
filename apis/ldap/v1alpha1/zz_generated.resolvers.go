/*
Copyright 2022 Upbound Inc.
*/
// Code generated by angryjet. DO NOT EDIT.

package v1alpha1

import (
	"context"
	v1alpha13 "github.com/crossplane-contrib/provider-keycloak/apis/client/v1alpha1"
	v1alpha11 "github.com/crossplane-contrib/provider-keycloak/apis/group/v1alpha1"
	v1alpha1 "github.com/crossplane-contrib/provider-keycloak/apis/realm/v1alpha1"
	v1alpha12 "github.com/crossplane-contrib/provider-keycloak/apis/role/v1alpha1"
	common "github.com/crossplane-contrib/provider-keycloak/config/common"
	reference "github.com/crossplane/crossplane-runtime/pkg/reference"
	resource "github.com/crossplane/upjet/pkg/resource"
	errors "github.com/pkg/errors"
	client "sigs.k8s.io/controller-runtime/pkg/client"
)

// ResolveReferences of this CustomMapper.
func (mg *CustomMapper) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.LdapUserFederationID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.LdapUserFederationIDRef,
		Selector:     mg.Spec.ForProvider.LdapUserFederationIDSelector,
		To: reference.To{
			List:    &UserFederationList{},
			Managed: &UserFederation{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.LdapUserFederationID")
	}
	mg.Spec.ForProvider.LdapUserFederationID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.LdapUserFederationIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.RealmID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.RealmIDRef,
		Selector:     mg.Spec.ForProvider.RealmIDSelector,
		To: reference.To{
			List:    &v1alpha1.RealmList{},
			Managed: &v1alpha1.Realm{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.RealmID")
	}
	mg.Spec.ForProvider.RealmID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.RealmIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.LdapUserFederationID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.InitProvider.LdapUserFederationIDRef,
		Selector:     mg.Spec.InitProvider.LdapUserFederationIDSelector,
		To: reference.To{
			List:    &UserFederationList{},
			Managed: &UserFederation{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.LdapUserFederationID")
	}
	mg.Spec.InitProvider.LdapUserFederationID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.LdapUserFederationIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.RealmID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.InitProvider.RealmIDRef,
		Selector:     mg.Spec.InitProvider.RealmIDSelector,
		To: reference.To{
			List:    &v1alpha1.RealmList{},
			Managed: &v1alpha1.Realm{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.RealmID")
	}
	mg.Spec.InitProvider.RealmID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.RealmIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this FullNameMapper.
func (mg *FullNameMapper) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.LdapUserFederationID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.LdapUserFederationIDRef,
		Selector:     mg.Spec.ForProvider.LdapUserFederationIDSelector,
		To: reference.To{
			List:    &UserFederationList{},
			Managed: &UserFederation{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.LdapUserFederationID")
	}
	mg.Spec.ForProvider.LdapUserFederationID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.LdapUserFederationIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.RealmID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.RealmIDRef,
		Selector:     mg.Spec.ForProvider.RealmIDSelector,
		To: reference.To{
			List:    &v1alpha1.RealmList{},
			Managed: &v1alpha1.Realm{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.RealmID")
	}
	mg.Spec.ForProvider.RealmID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.RealmIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.LdapUserFederationID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.InitProvider.LdapUserFederationIDRef,
		Selector:     mg.Spec.InitProvider.LdapUserFederationIDSelector,
		To: reference.To{
			List:    &UserFederationList{},
			Managed: &UserFederation{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.LdapUserFederationID")
	}
	mg.Spec.InitProvider.LdapUserFederationID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.LdapUserFederationIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.RealmID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.InitProvider.RealmIDRef,
		Selector:     mg.Spec.InitProvider.RealmIDSelector,
		To: reference.To{
			List:    &v1alpha1.RealmList{},
			Managed: &v1alpha1.Realm{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.RealmID")
	}
	mg.Spec.InitProvider.RealmID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.RealmIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this GroupMapper.
func (mg *GroupMapper) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.LdapUserFederationID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.LdapUserFederationIDRef,
		Selector:     mg.Spec.ForProvider.LdapUserFederationIDSelector,
		To: reference.To{
			List:    &UserFederationList{},
			Managed: &UserFederation{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.LdapUserFederationID")
	}
	mg.Spec.ForProvider.LdapUserFederationID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.LdapUserFederationIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.RealmID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.RealmIDRef,
		Selector:     mg.Spec.ForProvider.RealmIDSelector,
		To: reference.To{
			List:    &v1alpha1.RealmList{},
			Managed: &v1alpha1.Realm{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.RealmID")
	}
	mg.Spec.ForProvider.RealmID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.RealmIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.LdapUserFederationID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.InitProvider.LdapUserFederationIDRef,
		Selector:     mg.Spec.InitProvider.LdapUserFederationIDSelector,
		To: reference.To{
			List:    &UserFederationList{},
			Managed: &UserFederation{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.LdapUserFederationID")
	}
	mg.Spec.InitProvider.LdapUserFederationID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.LdapUserFederationIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.RealmID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.InitProvider.RealmIDRef,
		Selector:     mg.Spec.InitProvider.RealmIDSelector,
		To: reference.To{
			List:    &v1alpha1.RealmList{},
			Managed: &v1alpha1.Realm{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.RealmID")
	}
	mg.Spec.InitProvider.RealmID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.RealmIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this HardcodedAttributeMapper.
func (mg *HardcodedAttributeMapper) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.LdapUserFederationID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.LdapUserFederationIDRef,
		Selector:     mg.Spec.ForProvider.LdapUserFederationIDSelector,
		To: reference.To{
			List:    &UserFederationList{},
			Managed: &UserFederation{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.LdapUserFederationID")
	}
	mg.Spec.ForProvider.LdapUserFederationID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.LdapUserFederationIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.RealmID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.RealmIDRef,
		Selector:     mg.Spec.ForProvider.RealmIDSelector,
		To: reference.To{
			List:    &v1alpha1.RealmList{},
			Managed: &v1alpha1.Realm{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.RealmID")
	}
	mg.Spec.ForProvider.RealmID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.RealmIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.LdapUserFederationID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.InitProvider.LdapUserFederationIDRef,
		Selector:     mg.Spec.InitProvider.LdapUserFederationIDSelector,
		To: reference.To{
			List:    &UserFederationList{},
			Managed: &UserFederation{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.LdapUserFederationID")
	}
	mg.Spec.InitProvider.LdapUserFederationID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.LdapUserFederationIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.RealmID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.InitProvider.RealmIDRef,
		Selector:     mg.Spec.InitProvider.RealmIDSelector,
		To: reference.To{
			List:    &v1alpha1.RealmList{},
			Managed: &v1alpha1.Realm{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.RealmID")
	}
	mg.Spec.InitProvider.RealmID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.RealmIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this HardcodedGroupMapper.
func (mg *HardcodedGroupMapper) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.Group),
		Extract:      resource.ExtractParamPath("name", false),
		Reference:    mg.Spec.ForProvider.GroupRef,
		Selector:     mg.Spec.ForProvider.GroupSelector,
		To: reference.To{
			List:    &v1alpha11.GroupList{},
			Managed: &v1alpha11.Group{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.Group")
	}
	mg.Spec.ForProvider.Group = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.GroupRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.LdapUserFederationID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.LdapUserFederationIDRef,
		Selector:     mg.Spec.ForProvider.LdapUserFederationIDSelector,
		To: reference.To{
			List:    &UserFederationList{},
			Managed: &UserFederation{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.LdapUserFederationID")
	}
	mg.Spec.ForProvider.LdapUserFederationID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.LdapUserFederationIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.RealmID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.RealmIDRef,
		Selector:     mg.Spec.ForProvider.RealmIDSelector,
		To: reference.To{
			List:    &v1alpha1.RealmList{},
			Managed: &v1alpha1.Realm{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.RealmID")
	}
	mg.Spec.ForProvider.RealmID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.RealmIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.Group),
		Extract:      resource.ExtractParamPath("name", false),
		Reference:    mg.Spec.InitProvider.GroupRef,
		Selector:     mg.Spec.InitProvider.GroupSelector,
		To: reference.To{
			List:    &v1alpha11.GroupList{},
			Managed: &v1alpha11.Group{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.Group")
	}
	mg.Spec.InitProvider.Group = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.GroupRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.LdapUserFederationID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.InitProvider.LdapUserFederationIDRef,
		Selector:     mg.Spec.InitProvider.LdapUserFederationIDSelector,
		To: reference.To{
			List:    &UserFederationList{},
			Managed: &UserFederation{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.LdapUserFederationID")
	}
	mg.Spec.InitProvider.LdapUserFederationID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.LdapUserFederationIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.RealmID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.InitProvider.RealmIDRef,
		Selector:     mg.Spec.InitProvider.RealmIDSelector,
		To: reference.To{
			List:    &v1alpha1.RealmList{},
			Managed: &v1alpha1.Realm{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.RealmID")
	}
	mg.Spec.InitProvider.RealmID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.RealmIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this HardcodedRoleMapper.
func (mg *HardcodedRoleMapper) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.LdapUserFederationID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.LdapUserFederationIDRef,
		Selector:     mg.Spec.ForProvider.LdapUserFederationIDSelector,
		To: reference.To{
			List:    &UserFederationList{},
			Managed: &UserFederation{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.LdapUserFederationID")
	}
	mg.Spec.ForProvider.LdapUserFederationID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.LdapUserFederationIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.RealmID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.RealmIDRef,
		Selector:     mg.Spec.ForProvider.RealmIDSelector,
		To: reference.To{
			List:    &v1alpha1.RealmList{},
			Managed: &v1alpha1.Realm{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.RealmID")
	}
	mg.Spec.ForProvider.RealmID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.RealmIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.Role),
		Extract:      resource.ExtractParamPath("name", false),
		Reference:    mg.Spec.ForProvider.RoleRef,
		Selector:     mg.Spec.ForProvider.RoleSelector,
		To: reference.To{
			List:    &v1alpha12.RoleList{},
			Managed: &v1alpha12.Role{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.Role")
	}
	mg.Spec.ForProvider.Role = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.RoleRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.LdapUserFederationID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.InitProvider.LdapUserFederationIDRef,
		Selector:     mg.Spec.InitProvider.LdapUserFederationIDSelector,
		To: reference.To{
			List:    &UserFederationList{},
			Managed: &UserFederation{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.LdapUserFederationID")
	}
	mg.Spec.InitProvider.LdapUserFederationID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.LdapUserFederationIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.RealmID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.InitProvider.RealmIDRef,
		Selector:     mg.Spec.InitProvider.RealmIDSelector,
		To: reference.To{
			List:    &v1alpha1.RealmList{},
			Managed: &v1alpha1.Realm{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.RealmID")
	}
	mg.Spec.InitProvider.RealmID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.RealmIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.Role),
		Extract:      resource.ExtractParamPath("name", false),
		Reference:    mg.Spec.InitProvider.RoleRef,
		Selector:     mg.Spec.InitProvider.RoleSelector,
		To: reference.To{
			List:    &v1alpha12.RoleList{},
			Managed: &v1alpha12.Role{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.Role")
	}
	mg.Spec.InitProvider.Role = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.RoleRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this MsadLdsUserAccountControlMapper.
func (mg *MsadLdsUserAccountControlMapper) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.LdapUserFederationID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.LdapUserFederationIDRef,
		Selector:     mg.Spec.ForProvider.LdapUserFederationIDSelector,
		To: reference.To{
			List:    &UserFederationList{},
			Managed: &UserFederation{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.LdapUserFederationID")
	}
	mg.Spec.ForProvider.LdapUserFederationID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.LdapUserFederationIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.RealmID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.RealmIDRef,
		Selector:     mg.Spec.ForProvider.RealmIDSelector,
		To: reference.To{
			List:    &v1alpha1.RealmList{},
			Managed: &v1alpha1.Realm{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.RealmID")
	}
	mg.Spec.ForProvider.RealmID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.RealmIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.LdapUserFederationID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.InitProvider.LdapUserFederationIDRef,
		Selector:     mg.Spec.InitProvider.LdapUserFederationIDSelector,
		To: reference.To{
			List:    &UserFederationList{},
			Managed: &UserFederation{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.LdapUserFederationID")
	}
	mg.Spec.InitProvider.LdapUserFederationID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.LdapUserFederationIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.RealmID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.InitProvider.RealmIDRef,
		Selector:     mg.Spec.InitProvider.RealmIDSelector,
		To: reference.To{
			List:    &v1alpha1.RealmList{},
			Managed: &v1alpha1.Realm{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.RealmID")
	}
	mg.Spec.InitProvider.RealmID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.RealmIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this MsadUserAccountControlMapper.
func (mg *MsadUserAccountControlMapper) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.LdapUserFederationID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.LdapUserFederationIDRef,
		Selector:     mg.Spec.ForProvider.LdapUserFederationIDSelector,
		To: reference.To{
			List:    &UserFederationList{},
			Managed: &UserFederation{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.LdapUserFederationID")
	}
	mg.Spec.ForProvider.LdapUserFederationID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.LdapUserFederationIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.RealmID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.RealmIDRef,
		Selector:     mg.Spec.ForProvider.RealmIDSelector,
		To: reference.To{
			List:    &v1alpha1.RealmList{},
			Managed: &v1alpha1.Realm{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.RealmID")
	}
	mg.Spec.ForProvider.RealmID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.RealmIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.LdapUserFederationID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.InitProvider.LdapUserFederationIDRef,
		Selector:     mg.Spec.InitProvider.LdapUserFederationIDSelector,
		To: reference.To{
			List:    &UserFederationList{},
			Managed: &UserFederation{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.LdapUserFederationID")
	}
	mg.Spec.InitProvider.LdapUserFederationID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.LdapUserFederationIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.RealmID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.InitProvider.RealmIDRef,
		Selector:     mg.Spec.InitProvider.RealmIDSelector,
		To: reference.To{
			List:    &v1alpha1.RealmList{},
			Managed: &v1alpha1.Realm{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.RealmID")
	}
	mg.Spec.InitProvider.RealmID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.RealmIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this RoleMapper.
func (mg *RoleMapper) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ClientID),
		Extract:      common.UUIDExtractor(),
		Reference:    mg.Spec.ForProvider.ClientIDRef,
		Selector:     mg.Spec.ForProvider.ClientIDSelector,
		To: reference.To{
			List:    &v1alpha13.OpenIdClientList{},
			Managed: &v1alpha13.OpenIdClient{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ClientID")
	}
	mg.Spec.ForProvider.ClientID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ClientIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.LdapUserFederationID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.LdapUserFederationIDRef,
		Selector:     mg.Spec.ForProvider.LdapUserFederationIDSelector,
		To: reference.To{
			List:    &UserFederationList{},
			Managed: &UserFederation{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.LdapUserFederationID")
	}
	mg.Spec.ForProvider.LdapUserFederationID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.LdapUserFederationIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.RealmID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.RealmIDRef,
		Selector:     mg.Spec.ForProvider.RealmIDSelector,
		To: reference.To{
			List:    &v1alpha1.RealmList{},
			Managed: &v1alpha1.Realm{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.RealmID")
	}
	mg.Spec.ForProvider.RealmID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.RealmIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.ClientID),
		Extract:      common.UUIDExtractor(),
		Reference:    mg.Spec.InitProvider.ClientIDRef,
		Selector:     mg.Spec.InitProvider.ClientIDSelector,
		To: reference.To{
			List:    &v1alpha13.OpenIdClientList{},
			Managed: &v1alpha13.OpenIdClient{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.ClientID")
	}
	mg.Spec.InitProvider.ClientID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.ClientIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.LdapUserFederationID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.InitProvider.LdapUserFederationIDRef,
		Selector:     mg.Spec.InitProvider.LdapUserFederationIDSelector,
		To: reference.To{
			List:    &UserFederationList{},
			Managed: &UserFederation{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.LdapUserFederationID")
	}
	mg.Spec.InitProvider.LdapUserFederationID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.LdapUserFederationIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.RealmID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.InitProvider.RealmIDRef,
		Selector:     mg.Spec.InitProvider.RealmIDSelector,
		To: reference.To{
			List:    &v1alpha1.RealmList{},
			Managed: &v1alpha1.Realm{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.RealmID")
	}
	mg.Spec.InitProvider.RealmID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.RealmIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this UserAttributeMapper.
func (mg *UserAttributeMapper) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.LdapUserFederationID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.LdapUserFederationIDRef,
		Selector:     mg.Spec.ForProvider.LdapUserFederationIDSelector,
		To: reference.To{
			List:    &UserFederationList{},
			Managed: &UserFederation{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.LdapUserFederationID")
	}
	mg.Spec.ForProvider.LdapUserFederationID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.LdapUserFederationIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.RealmID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.RealmIDRef,
		Selector:     mg.Spec.ForProvider.RealmIDSelector,
		To: reference.To{
			List:    &v1alpha1.RealmList{},
			Managed: &v1alpha1.Realm{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.RealmID")
	}
	mg.Spec.ForProvider.RealmID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.RealmIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.LdapUserFederationID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.InitProvider.LdapUserFederationIDRef,
		Selector:     mg.Spec.InitProvider.LdapUserFederationIDSelector,
		To: reference.To{
			List:    &UserFederationList{},
			Managed: &UserFederation{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.LdapUserFederationID")
	}
	mg.Spec.InitProvider.LdapUserFederationID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.LdapUserFederationIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.RealmID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.InitProvider.RealmIDRef,
		Selector:     mg.Spec.InitProvider.RealmIDSelector,
		To: reference.To{
			List:    &v1alpha1.RealmList{},
			Managed: &v1alpha1.Realm{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.RealmID")
	}
	mg.Spec.InitProvider.RealmID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.RealmIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this UserFederation.
func (mg *UserFederation) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.RealmID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.RealmIDRef,
		Selector:     mg.Spec.ForProvider.RealmIDSelector,
		To: reference.To{
			List:    &v1alpha1.RealmList{},
			Managed: &v1alpha1.Realm{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.RealmID")
	}
	mg.Spec.ForProvider.RealmID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.RealmIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.RealmID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.InitProvider.RealmIDRef,
		Selector:     mg.Spec.InitProvider.RealmIDSelector,
		To: reference.To{
			List:    &v1alpha1.RealmList{},
			Managed: &v1alpha1.Realm{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.RealmID")
	}
	mg.Spec.InitProvider.RealmID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.RealmIDRef = rsp.ResolvedReference

	return nil
}
