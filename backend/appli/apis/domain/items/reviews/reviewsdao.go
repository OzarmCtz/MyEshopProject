package reviews

import (
	"context"
	"database/sql"

	"github.com/OzarmCtz/e_shop_backend_v1/app/data/global"
	adm "github.com/OzarmCtz/e_shop_backend_v1/appli/datasources/mysql"
)

func GetUserReview(reviewId int32) (adm.ItemsReview, error) {
	ctx, cancel := context.WithTimeout(context.Background(), global.Timeout)
	defer cancel()
	review, err := adm.QueriesDb.GetItemReview(ctx, reviewId)
	return review, err
}

func ListItemsReviewByItemId(itemId int32) ([]adm.ItemsReview, error) {
	ctx, cancel := context.WithTimeout(context.Background(), global.Timeout)
	defer cancel()
	review, err := adm.QueriesDb.ListItemsReviewsByItemId(ctx, itemId)
	return review, err
}

func ListItemsReviewsByUser(userId int32) ([]adm.ItemsReview, error) {
	ctx, cancel := context.WithTimeout(context.Background(), global.Timeout)
	defer cancel()
	review, err := adm.QueriesDb.ListItemsReviewsByUserId(ctx, userId)
	return review, err
}

func ListItemsReviews() ([]adm.ItemsReview, error) {
	ctx, cancel := context.WithTimeout(context.Background(), global.Timeout)
	defer cancel()
	review, err := adm.QueriesDb.ListItemsReviews(ctx)
	return review, err
}

func InsertReview(reviewParams adm.CreateItemReviewParams) (sql.Result, error) {
	ctx, cancel := context.WithTimeout(context.Background(), global.Timeout)
	defer cancel()
	res, err := adm.QueriesDb.CreateItemReview(ctx, reviewParams)
	return res, err
}

func DeleteItemReviewByUserAndItem(reviewItemId int32) (int64, error) {
	ctx, cancel := context.WithTimeout(context.Background(), global.Timeout)
	defer cancel()
	rows, err := adm.QueriesDb.DeleteItemReview(ctx, reviewItemId)
	return rows, err
}
