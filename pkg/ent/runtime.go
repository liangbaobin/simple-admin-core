// Code generated by ent, DO NOT EDIT.

package ent

import (
	"time"

	"github.com/suyuan32/simple-admin-core/pkg/ent/api"
	"github.com/suyuan32/simple-admin-core/pkg/ent/dictionary"
	"github.com/suyuan32/simple-admin-core/pkg/ent/dictionarydetail"
	"github.com/suyuan32/simple-admin-core/pkg/ent/menu"
	"github.com/suyuan32/simple-admin-core/pkg/ent/menuparam"
	"github.com/suyuan32/simple-admin-core/pkg/ent/oauthprovider"
	"github.com/suyuan32/simple-admin-core/pkg/ent/role"
	"github.com/suyuan32/simple-admin-core/pkg/ent/schema"
	"github.com/suyuan32/simple-admin-core/pkg/ent/token"
	"github.com/suyuan32/simple-admin-core/pkg/ent/user"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	apiMixin := schema.API{}.Mixin()
	apiMixinFields0 := apiMixin[0].Fields()
	_ = apiMixinFields0
	apiFields := schema.API{}.Fields()
	_ = apiFields
	// apiDescCreatedAt is the schema descriptor for created_at field.
	apiDescCreatedAt := apiMixinFields0[1].Descriptor()
	// api.DefaultCreatedAt holds the default value on creation for the created_at field.
	api.DefaultCreatedAt = apiDescCreatedAt.Default.(func() time.Time)
	// apiDescUpdatedAt is the schema descriptor for updated_at field.
	apiDescUpdatedAt := apiMixinFields0[2].Descriptor()
	// api.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	api.DefaultUpdatedAt = apiDescUpdatedAt.Default.(func() time.Time)
	// api.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	api.UpdateDefaultUpdatedAt = apiDescUpdatedAt.UpdateDefault.(func() time.Time)
	// apiDescMethod is the schema descriptor for method field.
	apiDescMethod := apiFields[3].Descriptor()
	// api.DefaultMethod holds the default value on creation for the method field.
	api.DefaultMethod = apiDescMethod.Default.(string)
	dictionaryMixin := schema.Dictionary{}.Mixin()
	dictionaryMixinFields0 := dictionaryMixin[0].Fields()
	_ = dictionaryMixinFields0
	dictionaryMixinFields1 := dictionaryMixin[1].Fields()
	_ = dictionaryMixinFields1
	dictionaryFields := schema.Dictionary{}.Fields()
	_ = dictionaryFields
	// dictionaryDescCreatedAt is the schema descriptor for created_at field.
	dictionaryDescCreatedAt := dictionaryMixinFields0[1].Descriptor()
	// dictionary.DefaultCreatedAt holds the default value on creation for the created_at field.
	dictionary.DefaultCreatedAt = dictionaryDescCreatedAt.Default.(func() time.Time)
	// dictionaryDescUpdatedAt is the schema descriptor for updated_at field.
	dictionaryDescUpdatedAt := dictionaryMixinFields0[2].Descriptor()
	// dictionary.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	dictionary.DefaultUpdatedAt = dictionaryDescUpdatedAt.Default.(func() time.Time)
	// dictionary.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	dictionary.UpdateDefaultUpdatedAt = dictionaryDescUpdatedAt.UpdateDefault.(func() time.Time)
	// dictionaryDescStatus is the schema descriptor for status field.
	dictionaryDescStatus := dictionaryMixinFields1[0].Descriptor()
	// dictionary.DefaultStatus holds the default value on creation for the status field.
	dictionary.DefaultStatus = dictionaryDescStatus.Default.(uint8)
	dictionarydetailMixin := schema.DictionaryDetail{}.Mixin()
	dictionarydetailMixinFields0 := dictionarydetailMixin[0].Fields()
	_ = dictionarydetailMixinFields0
	dictionarydetailMixinFields1 := dictionarydetailMixin[1].Fields()
	_ = dictionarydetailMixinFields1
	dictionarydetailFields := schema.DictionaryDetail{}.Fields()
	_ = dictionarydetailFields
	// dictionarydetailDescCreatedAt is the schema descriptor for created_at field.
	dictionarydetailDescCreatedAt := dictionarydetailMixinFields0[1].Descriptor()
	// dictionarydetail.DefaultCreatedAt holds the default value on creation for the created_at field.
	dictionarydetail.DefaultCreatedAt = dictionarydetailDescCreatedAt.Default.(func() time.Time)
	// dictionarydetailDescUpdatedAt is the schema descriptor for updated_at field.
	dictionarydetailDescUpdatedAt := dictionarydetailMixinFields0[2].Descriptor()
	// dictionarydetail.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	dictionarydetail.DefaultUpdatedAt = dictionarydetailDescUpdatedAt.Default.(func() time.Time)
	// dictionarydetail.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	dictionarydetail.UpdateDefaultUpdatedAt = dictionarydetailDescUpdatedAt.UpdateDefault.(func() time.Time)
	// dictionarydetailDescStatus is the schema descriptor for status field.
	dictionarydetailDescStatus := dictionarydetailMixinFields1[0].Descriptor()
	// dictionarydetail.DefaultStatus holds the default value on creation for the status field.
	dictionarydetail.DefaultStatus = dictionarydetailDescStatus.Default.(uint8)
	menuMixin := schema.Menu{}.Mixin()
	menuMixinFields0 := menuMixin[0].Fields()
	_ = menuMixinFields0
	menuFields := schema.Menu{}.Fields()
	_ = menuFields
	// menuDescCreatedAt is the schema descriptor for created_at field.
	menuDescCreatedAt := menuMixinFields0[1].Descriptor()
	// menu.DefaultCreatedAt holds the default value on creation for the created_at field.
	menu.DefaultCreatedAt = menuDescCreatedAt.Default.(func() time.Time)
	// menuDescUpdatedAt is the schema descriptor for updated_at field.
	menuDescUpdatedAt := menuMixinFields0[2].Descriptor()
	// menu.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	menu.DefaultUpdatedAt = menuDescUpdatedAt.Default.(func() time.Time)
	// menu.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	menu.UpdateDefaultUpdatedAt = menuDescUpdatedAt.UpdateDefault.(func() time.Time)
	// menuDescPath is the schema descriptor for path field.
	menuDescPath := menuFields[3].Descriptor()
	// menu.DefaultPath holds the default value on creation for the path field.
	menu.DefaultPath = menuDescPath.Default.(string)
	// menuDescRedirect is the schema descriptor for redirect field.
	menuDescRedirect := menuFields[5].Descriptor()
	// menu.DefaultRedirect holds the default value on creation for the redirect field.
	menu.DefaultRedirect = menuDescRedirect.Default.(string)
	// menuDescComponent is the schema descriptor for component field.
	menuDescComponent := menuFields[6].Descriptor()
	// menu.DefaultComponent holds the default value on creation for the component field.
	menu.DefaultComponent = menuDescComponent.Default.(string)
	// menuDescOrderNo is the schema descriptor for order_no field.
	menuDescOrderNo := menuFields[7].Descriptor()
	// menu.DefaultOrderNo holds the default value on creation for the order_no field.
	menu.DefaultOrderNo = menuDescOrderNo.Default.(uint32)
	// menuDescDisabled is the schema descriptor for disabled field.
	menuDescDisabled := menuFields[8].Descriptor()
	// menu.DefaultDisabled holds the default value on creation for the disabled field.
	menu.DefaultDisabled = menuDescDisabled.Default.(bool)
	// menuDescHideMenu is the schema descriptor for hide_menu field.
	menuDescHideMenu := menuFields[11].Descriptor()
	// menu.DefaultHideMenu holds the default value on creation for the hide_menu field.
	menu.DefaultHideMenu = menuDescHideMenu.Default.(bool)
	// menuDescHideBreadcrumb is the schema descriptor for hide_breadcrumb field.
	menuDescHideBreadcrumb := menuFields[12].Descriptor()
	// menu.DefaultHideBreadcrumb holds the default value on creation for the hide_breadcrumb field.
	menu.DefaultHideBreadcrumb = menuDescHideBreadcrumb.Default.(bool)
	// menuDescCurrentActiveMenu is the schema descriptor for current_active_menu field.
	menuDescCurrentActiveMenu := menuFields[13].Descriptor()
	// menu.DefaultCurrentActiveMenu holds the default value on creation for the current_active_menu field.
	menu.DefaultCurrentActiveMenu = menuDescCurrentActiveMenu.Default.(string)
	// menuDescIgnoreKeepAlive is the schema descriptor for ignore_keep_alive field.
	menuDescIgnoreKeepAlive := menuFields[14].Descriptor()
	// menu.DefaultIgnoreKeepAlive holds the default value on creation for the ignore_keep_alive field.
	menu.DefaultIgnoreKeepAlive = menuDescIgnoreKeepAlive.Default.(bool)
	// menuDescHideTab is the schema descriptor for hide_tab field.
	menuDescHideTab := menuFields[15].Descriptor()
	// menu.DefaultHideTab holds the default value on creation for the hide_tab field.
	menu.DefaultHideTab = menuDescHideTab.Default.(bool)
	// menuDescFrameSrc is the schema descriptor for frame_src field.
	menuDescFrameSrc := menuFields[16].Descriptor()
	// menu.DefaultFrameSrc holds the default value on creation for the frame_src field.
	menu.DefaultFrameSrc = menuDescFrameSrc.Default.(string)
	// menuDescCarryParam is the schema descriptor for carry_param field.
	menuDescCarryParam := menuFields[17].Descriptor()
	// menu.DefaultCarryParam holds the default value on creation for the carry_param field.
	menu.DefaultCarryParam = menuDescCarryParam.Default.(bool)
	// menuDescHideChildrenInMenu is the schema descriptor for hide_children_in_menu field.
	menuDescHideChildrenInMenu := menuFields[18].Descriptor()
	// menu.DefaultHideChildrenInMenu holds the default value on creation for the hide_children_in_menu field.
	menu.DefaultHideChildrenInMenu = menuDescHideChildrenInMenu.Default.(bool)
	// menuDescAffix is the schema descriptor for affix field.
	menuDescAffix := menuFields[19].Descriptor()
	// menu.DefaultAffix holds the default value on creation for the affix field.
	menu.DefaultAffix = menuDescAffix.Default.(bool)
	// menuDescDynamicLevel is the schema descriptor for dynamic_level field.
	menuDescDynamicLevel := menuFields[20].Descriptor()
	// menu.DefaultDynamicLevel holds the default value on creation for the dynamic_level field.
	menu.DefaultDynamicLevel = menuDescDynamicLevel.Default.(uint32)
	// menuDescRealPath is the schema descriptor for real_path field.
	menuDescRealPath := menuFields[21].Descriptor()
	// menu.DefaultRealPath holds the default value on creation for the real_path field.
	menu.DefaultRealPath = menuDescRealPath.Default.(string)
	menuparamMixin := schema.MenuParam{}.Mixin()
	menuparamMixinFields0 := menuparamMixin[0].Fields()
	_ = menuparamMixinFields0
	menuparamFields := schema.MenuParam{}.Fields()
	_ = menuparamFields
	// menuparamDescCreatedAt is the schema descriptor for created_at field.
	menuparamDescCreatedAt := menuparamMixinFields0[1].Descriptor()
	// menuparam.DefaultCreatedAt holds the default value on creation for the created_at field.
	menuparam.DefaultCreatedAt = menuparamDescCreatedAt.Default.(func() time.Time)
	// menuparamDescUpdatedAt is the schema descriptor for updated_at field.
	menuparamDescUpdatedAt := menuparamMixinFields0[2].Descriptor()
	// menuparam.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	menuparam.DefaultUpdatedAt = menuparamDescUpdatedAt.Default.(func() time.Time)
	// menuparam.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	menuparam.UpdateDefaultUpdatedAt = menuparamDescUpdatedAt.UpdateDefault.(func() time.Time)
	oauthproviderMixin := schema.OauthProvider{}.Mixin()
	oauthproviderMixinFields0 := oauthproviderMixin[0].Fields()
	_ = oauthproviderMixinFields0
	oauthproviderFields := schema.OauthProvider{}.Fields()
	_ = oauthproviderFields
	// oauthproviderDescCreatedAt is the schema descriptor for created_at field.
	oauthproviderDescCreatedAt := oauthproviderMixinFields0[1].Descriptor()
	// oauthprovider.DefaultCreatedAt holds the default value on creation for the created_at field.
	oauthprovider.DefaultCreatedAt = oauthproviderDescCreatedAt.Default.(func() time.Time)
	// oauthproviderDescUpdatedAt is the schema descriptor for updated_at field.
	oauthproviderDescUpdatedAt := oauthproviderMixinFields0[2].Descriptor()
	// oauthprovider.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	oauthprovider.DefaultUpdatedAt = oauthproviderDescUpdatedAt.Default.(func() time.Time)
	// oauthprovider.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	oauthprovider.UpdateDefaultUpdatedAt = oauthproviderDescUpdatedAt.UpdateDefault.(func() time.Time)
	roleMixin := schema.Role{}.Mixin()
	roleMixinFields0 := roleMixin[0].Fields()
	_ = roleMixinFields0
	roleMixinFields1 := roleMixin[1].Fields()
	_ = roleMixinFields1
	roleFields := schema.Role{}.Fields()
	_ = roleFields
	// roleDescCreatedAt is the schema descriptor for created_at field.
	roleDescCreatedAt := roleMixinFields0[1].Descriptor()
	// role.DefaultCreatedAt holds the default value on creation for the created_at field.
	role.DefaultCreatedAt = roleDescCreatedAt.Default.(func() time.Time)
	// roleDescUpdatedAt is the schema descriptor for updated_at field.
	roleDescUpdatedAt := roleMixinFields0[2].Descriptor()
	// role.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	role.DefaultUpdatedAt = roleDescUpdatedAt.Default.(func() time.Time)
	// role.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	role.UpdateDefaultUpdatedAt = roleDescUpdatedAt.UpdateDefault.(func() time.Time)
	// roleDescStatus is the schema descriptor for status field.
	roleDescStatus := roleMixinFields1[0].Descriptor()
	// role.DefaultStatus holds the default value on creation for the status field.
	role.DefaultStatus = roleDescStatus.Default.(uint8)
	// roleDescDefaultRouter is the schema descriptor for default_router field.
	roleDescDefaultRouter := roleFields[2].Descriptor()
	// role.DefaultDefaultRouter holds the default value on creation for the default_router field.
	role.DefaultDefaultRouter = roleDescDefaultRouter.Default.(string)
	// roleDescRemark is the schema descriptor for remark field.
	roleDescRemark := roleFields[3].Descriptor()
	// role.DefaultRemark holds the default value on creation for the remark field.
	role.DefaultRemark = roleDescRemark.Default.(string)
	// roleDescOrderNo is the schema descriptor for order_no field.
	roleDescOrderNo := roleFields[4].Descriptor()
	// role.DefaultOrderNo holds the default value on creation for the order_no field.
	role.DefaultOrderNo = roleDescOrderNo.Default.(uint32)
	tokenMixin := schema.Token{}.Mixin()
	tokenMixinFields0 := tokenMixin[0].Fields()
	_ = tokenMixinFields0
	tokenMixinFields1 := tokenMixin[1].Fields()
	_ = tokenMixinFields1
	tokenFields := schema.Token{}.Fields()
	_ = tokenFields
	// tokenDescCreatedAt is the schema descriptor for created_at field.
	tokenDescCreatedAt := tokenMixinFields0[1].Descriptor()
	// token.DefaultCreatedAt holds the default value on creation for the created_at field.
	token.DefaultCreatedAt = tokenDescCreatedAt.Default.(func() time.Time)
	// tokenDescUpdatedAt is the schema descriptor for updated_at field.
	tokenDescUpdatedAt := tokenMixinFields0[2].Descriptor()
	// token.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	token.DefaultUpdatedAt = tokenDescUpdatedAt.Default.(func() time.Time)
	// token.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	token.UpdateDefaultUpdatedAt = tokenDescUpdatedAt.UpdateDefault.(func() time.Time)
	// tokenDescStatus is the schema descriptor for status field.
	tokenDescStatus := tokenMixinFields1[0].Descriptor()
	// token.DefaultStatus holds the default value on creation for the status field.
	token.DefaultStatus = tokenDescStatus.Default.(uint8)
	userMixin := schema.User{}.Mixin()
	userMixinFields0 := userMixin[0].Fields()
	_ = userMixinFields0
	userMixinFields1 := userMixin[1].Fields()
	_ = userMixinFields1
	userFields := schema.User{}.Fields()
	_ = userFields
	// userDescCreatedAt is the schema descriptor for created_at field.
	userDescCreatedAt := userMixinFields0[1].Descriptor()
	// user.DefaultCreatedAt holds the default value on creation for the created_at field.
	user.DefaultCreatedAt = userDescCreatedAt.Default.(func() time.Time)
	// userDescUpdatedAt is the schema descriptor for updated_at field.
	userDescUpdatedAt := userMixinFields0[2].Descriptor()
	// user.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	user.DefaultUpdatedAt = userDescUpdatedAt.Default.(func() time.Time)
	// user.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	user.UpdateDefaultUpdatedAt = userDescUpdatedAt.UpdateDefault.(func() time.Time)
	// userDescStatus is the schema descriptor for status field.
	userDescStatus := userMixinFields1[0].Descriptor()
	// user.DefaultStatus holds the default value on creation for the status field.
	user.DefaultStatus = userDescStatus.Default.(uint8)
	// userDescSideMode is the schema descriptor for side_mode field.
	userDescSideMode := userFields[4].Descriptor()
	// user.DefaultSideMode holds the default value on creation for the side_mode field.
	user.DefaultSideMode = userDescSideMode.Default.(string)
	// userDescBaseColor is the schema descriptor for base_color field.
	userDescBaseColor := userFields[5].Descriptor()
	// user.DefaultBaseColor holds the default value on creation for the base_color field.
	user.DefaultBaseColor = userDescBaseColor.Default.(string)
	// userDescActiveColor is the schema descriptor for active_color field.
	userDescActiveColor := userFields[6].Descriptor()
	// user.DefaultActiveColor holds the default value on creation for the active_color field.
	user.DefaultActiveColor = userDescActiveColor.Default.(string)
	// userDescRoleID is the schema descriptor for role_id field.
	userDescRoleID := userFields[7].Descriptor()
	// user.DefaultRoleID holds the default value on creation for the role_id field.
	user.DefaultRoleID = userDescRoleID.Default.(uint64)
	// userDescAvatar is the schema descriptor for avatar field.
	userDescAvatar := userFields[10].Descriptor()
	// user.DefaultAvatar holds the default value on creation for the avatar field.
	user.DefaultAvatar = userDescAvatar.Default.(string)
}
