/*
Copyright 2022 Upbound Inc.
*/

// Code generated by upjet. DO NOT EDIT.

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type GroupMembershipProtocolMapperInitParameters struct {

	// Indicates if the property should be added as a claim to the access token. Defaults to true.
	AddToAccessToken *bool `json:"addToAccessToken,omitempty" tf:"add_to_access_token,omitempty"`

	// Indicates if the property should be added as a claim to the id token. Defaults to true.
	AddToIDToken *bool `json:"addToIdToken,omitempty" tf:"add_to_id_token,omitempty"`

	// Indicates if the property should be added as a claim to the UserInfo response body. Defaults to true.
	AddToUserinfo *bool `json:"addToUserinfo,omitempty" tf:"add_to_userinfo,omitempty"`

	// The name of the claim to insert into a token.
	ClaimName *string `json:"claimName,omitempty" tf:"claim_name,omitempty"`

	// The client this protocol mapper should be attached to. Conflicts with client_scope_id. One of client_id or client_scope_id must be specified.
	// The mapper's associated client. Cannot be used at the same time as client_scope_id.
	// +crossplane:generate:reference:type=github.com/crossplane-contrib/provider-keycloak/apis/openidclient/v1alpha1.Client
	// +crossplane:generate:reference:extractor=github.com/crossplane-contrib/provider-keycloak/config/common.UUIDExtractor()
	ClientID *string `json:"clientId,omitempty" tf:"client_id,omitempty"`

	// Reference to a Client in openidclient to populate clientId.
	// +kubebuilder:validation:Optional
	ClientIDRef *v1.Reference `json:"clientIdRef,omitempty" tf:"-"`

	// Selector for a Client in openidclient to populate clientId.
	// +kubebuilder:validation:Optional
	ClientIDSelector *v1.Selector `json:"clientIdSelector,omitempty" tf:"-"`

	// The client scope this protocol mapper should be attached to. Conflicts with client_id. One of client_id or client_scope_id must be specified.
	// The mapper's associated client scope. Cannot be used at the same time as client_id.
	// +crossplane:generate:reference:type=github.com/crossplane-contrib/provider-keycloak/apis/openidclient/v1alpha1.ClientScope
	ClientScopeID *string `json:"clientScopeId,omitempty" tf:"client_scope_id,omitempty"`

	// Reference to a ClientScope in openidclient to populate clientScopeId.
	// +kubebuilder:validation:Optional
	ClientScopeIDRef *v1.Reference `json:"clientScopeIdRef,omitempty" tf:"-"`

	// Selector for a ClientScope in openidclient to populate clientScopeId.
	// +kubebuilder:validation:Optional
	ClientScopeIDSelector *v1.Selector `json:"clientScopeIdSelector,omitempty" tf:"-"`

	// Indicates whether the full path of the group including its parents will be used. Defaults to true.
	FullPath *bool `json:"fullPath,omitempty" tf:"full_path,omitempty"`

	// The display name of this protocol mapper in the GUI.
	// A human-friendly name that will appear in the Keycloak console.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The realm this protocol mapper exists within.
	// The realm id where the associated client or client scope exists.
	// +crossplane:generate:reference:type=github.com/crossplane-contrib/provider-keycloak/apis/realm/v1alpha1.Realm
	RealmID *string `json:"realmId,omitempty" tf:"realm_id,omitempty"`

	// Reference to a Realm in realm to populate realmId.
	// +kubebuilder:validation:Optional
	RealmIDRef *v1.Reference `json:"realmIdRef,omitempty" tf:"-"`

	// Selector for a Realm in realm to populate realmId.
	// +kubebuilder:validation:Optional
	RealmIDSelector *v1.Selector `json:"realmIdSelector,omitempty" tf:"-"`
}

type GroupMembershipProtocolMapperObservation struct {

	// Indicates if the property should be added as a claim to the access token. Defaults to true.
	AddToAccessToken *bool `json:"addToAccessToken,omitempty" tf:"add_to_access_token,omitempty"`

	// Indicates if the property should be added as a claim to the id token. Defaults to true.
	AddToIDToken *bool `json:"addToIdToken,omitempty" tf:"add_to_id_token,omitempty"`

	// Indicates if the property should be added as a claim to the UserInfo response body. Defaults to true.
	AddToUserinfo *bool `json:"addToUserinfo,omitempty" tf:"add_to_userinfo,omitempty"`

	// The name of the claim to insert into a token.
	ClaimName *string `json:"claimName,omitempty" tf:"claim_name,omitempty"`

	// The client this protocol mapper should be attached to. Conflicts with client_scope_id. One of client_id or client_scope_id must be specified.
	// The mapper's associated client. Cannot be used at the same time as client_scope_id.
	ClientID *string `json:"clientId,omitempty" tf:"client_id,omitempty"`

	// The client scope this protocol mapper should be attached to. Conflicts with client_id. One of client_id or client_scope_id must be specified.
	// The mapper's associated client scope. Cannot be used at the same time as client_id.
	ClientScopeID *string `json:"clientScopeId,omitempty" tf:"client_scope_id,omitempty"`

	// Indicates whether the full path of the group including its parents will be used. Defaults to true.
	FullPath *bool `json:"fullPath,omitempty" tf:"full_path,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The display name of this protocol mapper in the GUI.
	// A human-friendly name that will appear in the Keycloak console.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The realm this protocol mapper exists within.
	// The realm id where the associated client or client scope exists.
	RealmID *string `json:"realmId,omitempty" tf:"realm_id,omitempty"`
}

type GroupMembershipProtocolMapperParameters struct {

	// Indicates if the property should be added as a claim to the access token. Defaults to true.
	// +kubebuilder:validation:Optional
	AddToAccessToken *bool `json:"addToAccessToken,omitempty" tf:"add_to_access_token,omitempty"`

	// Indicates if the property should be added as a claim to the id token. Defaults to true.
	// +kubebuilder:validation:Optional
	AddToIDToken *bool `json:"addToIdToken,omitempty" tf:"add_to_id_token,omitempty"`

	// Indicates if the property should be added as a claim to the UserInfo response body. Defaults to true.
	// +kubebuilder:validation:Optional
	AddToUserinfo *bool `json:"addToUserinfo,omitempty" tf:"add_to_userinfo,omitempty"`

	// The name of the claim to insert into a token.
	// +kubebuilder:validation:Optional
	ClaimName *string `json:"claimName,omitempty" tf:"claim_name,omitempty"`

	// The client this protocol mapper should be attached to. Conflicts with client_scope_id. One of client_id or client_scope_id must be specified.
	// The mapper's associated client. Cannot be used at the same time as client_scope_id.
	// +crossplane:generate:reference:type=github.com/crossplane-contrib/provider-keycloak/apis/openidclient/v1alpha1.Client
	// +crossplane:generate:reference:extractor=github.com/crossplane-contrib/provider-keycloak/config/common.UUIDExtractor()
	// +kubebuilder:validation:Optional
	ClientID *string `json:"clientId,omitempty" tf:"client_id,omitempty"`

	// Reference to a Client in openidclient to populate clientId.
	// +kubebuilder:validation:Optional
	ClientIDRef *v1.Reference `json:"clientIdRef,omitempty" tf:"-"`

	// Selector for a Client in openidclient to populate clientId.
	// +kubebuilder:validation:Optional
	ClientIDSelector *v1.Selector `json:"clientIdSelector,omitempty" tf:"-"`

	// The client scope this protocol mapper should be attached to. Conflicts with client_id. One of client_id or client_scope_id must be specified.
	// The mapper's associated client scope. Cannot be used at the same time as client_id.
	// +crossplane:generate:reference:type=github.com/crossplane-contrib/provider-keycloak/apis/openidclient/v1alpha1.ClientScope
	// +kubebuilder:validation:Optional
	ClientScopeID *string `json:"clientScopeId,omitempty" tf:"client_scope_id,omitempty"`

	// Reference to a ClientScope in openidclient to populate clientScopeId.
	// +kubebuilder:validation:Optional
	ClientScopeIDRef *v1.Reference `json:"clientScopeIdRef,omitempty" tf:"-"`

	// Selector for a ClientScope in openidclient to populate clientScopeId.
	// +kubebuilder:validation:Optional
	ClientScopeIDSelector *v1.Selector `json:"clientScopeIdSelector,omitempty" tf:"-"`

	// Indicates whether the full path of the group including its parents will be used. Defaults to true.
	// +kubebuilder:validation:Optional
	FullPath *bool `json:"fullPath,omitempty" tf:"full_path,omitempty"`

	// The display name of this protocol mapper in the GUI.
	// A human-friendly name that will appear in the Keycloak console.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The realm this protocol mapper exists within.
	// The realm id where the associated client or client scope exists.
	// +crossplane:generate:reference:type=github.com/crossplane-contrib/provider-keycloak/apis/realm/v1alpha1.Realm
	// +kubebuilder:validation:Optional
	RealmID *string `json:"realmId,omitempty" tf:"realm_id,omitempty"`

	// Reference to a Realm in realm to populate realmId.
	// +kubebuilder:validation:Optional
	RealmIDRef *v1.Reference `json:"realmIdRef,omitempty" tf:"-"`

	// Selector for a Realm in realm to populate realmId.
	// +kubebuilder:validation:Optional
	RealmIDSelector *v1.Selector `json:"realmIdSelector,omitempty" tf:"-"`
}

// GroupMembershipProtocolMapperSpec defines the desired state of GroupMembershipProtocolMapper
type GroupMembershipProtocolMapperSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     GroupMembershipProtocolMapperParameters `json:"forProvider"`
	// THIS IS A BETA FIELD. It will be honored
	// unless the Management Policies feature flag is disabled.
	// InitProvider holds the same fields as ForProvider, with the exception
	// of Identifier and other resource reference fields. The fields that are
	// in InitProvider are merged into ForProvider when the resource is created.
	// The same fields are also added to the terraform ignore_changes hook, to
	// avoid updating them after creation. This is useful for fields that are
	// required on creation, but we do not desire to update them after creation,
	// for example because of an external controller is managing them, like an
	// autoscaler.
	InitProvider GroupMembershipProtocolMapperInitParameters `json:"initProvider,omitempty"`
}

// GroupMembershipProtocolMapperStatus defines the observed state of GroupMembershipProtocolMapper.
type GroupMembershipProtocolMapperStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        GroupMembershipProtocolMapperObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// GroupMembershipProtocolMapper is the Schema for the GroupMembershipProtocolMappers API.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,keycloak}
type GroupMembershipProtocolMapper struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.claimName) || (has(self.initProvider) && has(self.initProvider.claimName))",message="spec.forProvider.claimName is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name) || (has(self.initProvider) && has(self.initProvider.name))",message="spec.forProvider.name is a required parameter"
	Spec   GroupMembershipProtocolMapperSpec   `json:"spec"`
	Status GroupMembershipProtocolMapperStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// GroupMembershipProtocolMapperList contains a list of GroupMembershipProtocolMappers
type GroupMembershipProtocolMapperList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []GroupMembershipProtocolMapper `json:"items"`
}

// Repository type metadata.
var (
	GroupMembershipProtocolMapper_Kind             = "GroupMembershipProtocolMapper"
	GroupMembershipProtocolMapper_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: GroupMembershipProtocolMapper_Kind}.String()
	GroupMembershipProtocolMapper_KindAPIVersion   = GroupMembershipProtocolMapper_Kind + "." + CRDGroupVersion.String()
	GroupMembershipProtocolMapper_GroupVersionKind = CRDGroupVersion.WithKind(GroupMembershipProtocolMapper_Kind)
)

func init() {
	SchemeBuilder.Register(&GroupMembershipProtocolMapper{}, &GroupMembershipProtocolMapperList{})
}
