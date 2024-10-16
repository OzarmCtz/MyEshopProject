// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0

package mysql

import (
	"context"
	"database/sql"
)

type Querier interface {
	CreateAppSettings(ctx context.Context, arg CreateAppSettingsParams) (sql.Result, error)
	CreateDiscount(ctx context.Context, arg CreateDiscountParams) (sql.Result, error)
	CreateDiscountLink(ctx context.Context, arg CreateDiscountLinkParams) (sql.Result, error)
	CreateGroupPrivilege(ctx context.Context, gpPath NullString) (sql.Result, error)
	CreateGroupPrivilegeLink(ctx context.Context, arg CreateGroupPrivilegeLinkParams) (sql.Result, error)
	CreateGroupUser(ctx context.Context, arg CreateGroupUserParams) (sql.Result, error)
	CreateGroupUserLink(ctx context.Context, arg CreateGroupUserLinkParams) (sql.Result, error)
	CreateGroupUserLinkByGroupName(ctx context.Context, arg CreateGroupUserLinkByGroupNameParams) (sql.Result, error)
	CreateItem(ctx context.Context, arg CreateItemParams) (sql.Result, error)
	CreateItemCategory(ctx context.Context, arg CreateItemCategoryParams) (sql.Result, error)
	CreateItemReview(ctx context.Context, arg CreateItemReviewParams) (sql.Result, error)
	CreateItemSubCategory(ctx context.Context, arg CreateItemSubCategoryParams) (sql.Result, error)
	CreateItemsCategoryLink(ctx context.Context, arg CreateItemsCategoryLinkParams) (sql.Result, error)
	CreateItemsSubCategoryLink(ctx context.Context, arg CreateItemsSubCategoryLinkParams) (sql.Result, error)
	CreateItemsSubCategoryLinkBySubCategoryName(ctx context.Context, arg CreateItemsSubCategoryLinkBySubCategoryNameParams) (sql.Result, error)
	CreateUser(ctx context.Context, arg CreateUserParams) (sql.Result, error)
	CreateUserBasket(ctx context.Context, arg CreateUserBasketParams) (sql.Result, error)
	CreateUserWishList(ctx context.Context, arg CreateUserWishListParams) (sql.Result, error)
	DeleteAppSettings(ctx context.Context, asID int32) (int64, error)
	DeleteDiscount(ctx context.Context, dID int32) (int64, error)
	DeleteDiscountLink(ctx context.Context, dlID int32) (int64, error)
	DeleteGroupPrivilege(ctx context.Context, gpID int32) (int64, error)
	DeleteGroupPrivilegeLink(ctx context.Context, gplID int32) (int64, error)
	DeleteGroupUser(ctx context.Context, guID int32) (int64, error)
	DeleteGroupUserLink(ctx context.Context, gulID int32) (int64, error)
	DeleteItem(ctx context.Context, iID int32) (int64, error)
	DeleteItemCategory(ctx context.Context, icID int32) (int64, error)
	DeleteItemReview(ctx context.Context, irID int32) (int64, error)
	DeleteItemSubCategory(ctx context.Context, iscID int32) (int64, error)
	DeleteItemsCategoryLink(ctx context.Context, iclID int32) (int64, error)
	DeleteItemsSubCategoryLink(ctx context.Context, isclID int32) (int64, error)
	DeleteUser(ctx context.Context, uID int32) (int64, error)
	DeleteUserBasketByUserAndItems(ctx context.Context, arg DeleteUserBasketByUserAndItemsParams) (int64, error)
	DeleteUserWishListByUserAndItems(ctx context.Context, arg DeleteUserWishListByUserAndItemsParams) (int64, error)
	GetAppSettings(ctx context.Context, asID int32) (AppSetting, error)
	GetAppSettingsByKey(ctx context.Context, asKey string) (AppSetting, error)
	GetDiscount(ctx context.Context, dID int32) (Discount, error)
	GetDiscountByCode(ctx context.Context, dCode string) (Discount, error)
	GetDiscountLink(ctx context.Context, dlID int32) (DiscountLink, error)
	GetGroupByName(ctx context.Context, guName string) (GroupsUser, error)
	GetGroupPrivilege(ctx context.Context, gpID int32) (GroupsPrivilege, error)
	GetGroupPrivilegeLink(ctx context.Context, gplID int32) (GroupsPrivilegesLink, error)
	GetGroupPrivilegesByUserId(ctx context.Context, uID int32) ([]GroupsPrivilege, error)
	GetGroupUser(ctx context.Context, guID int32) (GroupsUser, error)
	GetGroupUserLink(ctx context.Context, gulID int32) (GroupsUsersLink, error)
	GetGroupUserLinkByUser(ctx context.Context, gulUserID int32) (GroupsUsersLink, error)
	GetItem(ctx context.Context, iID int32) (Item, error)
	GetItemCategory(ctx context.Context, icID int32) (ItemsCategory, error)
	GetItemCategoryLink(ctx context.Context, iclID int32) (ItemsCategoryLink, error)
	GetItemCategoryLinkBySubCategory(ctx context.Context, iclItemsSubCategoryID int32) (ItemsCategory, error)
	GetItemReview(ctx context.Context, irID int32) (ItemsReview, error)
	GetItemSubCategory(ctx context.Context, iscID int32) (ItemsSubCategory, error)
	GetItemSubCategoryAndCategoryLinked(ctx context.Context, iscID int32) (GetItemSubCategoryAndCategoryLinkedRow, error)
	GetItemsSubCategoryLink(ctx context.Context, isclID int32) (ItemsSubCategoryLink, error)
	GetItemsSubCategoryLinkByItem(ctx context.Context, isclItemsID int32) (ItemsSubCategory, error)
	GetUser(ctx context.Context, uID int32) (User, error)
	GetUserBasket(ctx context.Context, ubID int32) (UsersBasket, error)
	GetUserByEmail(ctx context.Context, uEmail string) (User, error)
	GetUserWishList(ctx context.Context, wlID int32) (UsersWishlist, error)
	ListActiveItems(ctx context.Context, iIsDisabled bool) ([]Item, error)
	ListAppSettings(ctx context.Context) ([]AppSetting, error)
	ListDiscounts(ctx context.Context) ([]Discount, error)
	ListDiscountsLinks(ctx context.Context) ([]DiscountLink, error)
	ListGroupPrivilegesLink(ctx context.Context) ([]GroupsPrivilegesLink, error)
	ListGroupsPrivileges(ctx context.Context) ([]GroupsPrivilege, error)
	ListGroupsUserByUser(ctx context.Context, gulUserID int32) ([]GroupsUser, error)
	ListGroupsUsers(ctx context.Context) ([]GroupsUser, error)
	ListGroupsUsersLink(ctx context.Context) ([]GroupsUsersLink, error)
	ListItemSubCategoryAndCategoryLinked(ctx context.Context) ([]ListItemSubCategoryAndCategoryLinkedRow, error)
	ListItems(ctx context.Context) ([]ListItemsRow, error)
	ListItemsCategory(ctx context.Context) ([]ItemsCategory, error)
	ListItemsCategoryAndOccurence(ctx context.Context) ([]ListItemsCategoryAndOccurenceRow, error)
	ListItemsCategoryLink(ctx context.Context) ([]ItemsCategoryLink, error)
	ListItemsCategoryLinkByCategory(ctx context.Context, iclItemsCategoryID int32) ([]ItemsCategoryLink, error)
	ListItemsReviews(ctx context.Context) ([]ItemsReview, error)
	ListItemsReviewsByItemId(ctx context.Context, irItemsID int32) ([]ItemsReview, error)
	ListItemsReviewsByUserId(ctx context.Context, irUserID int32) ([]ItemsReview, error)
	ListItemsSubCategory(ctx context.Context) ([]ItemsSubCategory, error)
	ListItemsSubCategoryLink(ctx context.Context) ([]ItemsSubCategoryLink, error)
	ListItemsSubCategoryLinkByCategory(ctx context.Context, isclSubCategoryID int32) ([]ItemsSubCategoryLink, error)
	ListUserBasketByUser(ctx context.Context, ubUserID int32) ([]UsersBasket, error)
	ListUserWishListByUser(ctx context.Context, wlUserID int32) ([]UsersWishlist, error)
	ListUsers(ctx context.Context) ([]User, error)
	UpdateAppSettings(ctx context.Context, arg UpdateAppSettingsParams) (int64, error)
	UpdateDiscount(ctx context.Context, arg UpdateDiscountParams) (int64, error)
	UpdateDiscountLink(ctx context.Context, arg UpdateDiscountLinkParams) (int64, error)
	UpdateGroupPrivilege(ctx context.Context, arg UpdateGroupPrivilegeParams) (int64, error)
	UpdateGroupPrivilegeLink(ctx context.Context, arg UpdateGroupPrivilegeLinkParams) (int64, error)
	UpdateGroupUser(ctx context.Context, arg UpdateGroupUserParams) (int64, error)
	UpdateGroupUserLink(ctx context.Context, arg UpdateGroupUserLinkParams) (int64, error)
	UpdateItem(ctx context.Context, arg UpdateItemParams) (int64, error)
	UpdateItemCategory(ctx context.Context, arg UpdateItemCategoryParams) (int64, error)
	UpdateItemSubCategory(ctx context.Context, arg UpdateItemSubCategoryParams) (int64, error)
	UpdateItemsCategoryLink(ctx context.Context, arg UpdateItemsCategoryLinkParams) (int64, error)
	UpdateItemsCategoryLinkBySubCategory(ctx context.Context, arg UpdateItemsCategoryLinkBySubCategoryParams) (int64, error)
	UpdateItemsSubCategoryLink(ctx context.Context, arg UpdateItemsSubCategoryLinkParams) (int64, error)
	UpdateItemsSubCategoryLinkByItem(ctx context.Context, arg UpdateItemsSubCategoryLinkByItemParams) (int64, error)
	UpdateUser(ctx context.Context, arg UpdateUserParams) (int64, error)
}

var _ Querier = (*Queries)(nil)
