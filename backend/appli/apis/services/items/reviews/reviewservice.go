package reviews

import (
	"errors"

	aadir "github.com/OzarmCtz/e_shop_backend_v1/appli/apis/domain/items/reviews"
	adu "github.com/OzarmCtz/e_shop_backend_v1/appli/apis/domain/users"
	adm "github.com/OzarmCtz/e_shop_backend_v1/appli/datasources/mysql"
	aua "github.com/OzarmCtz/e_shop_backend_v1/appli/utils/auth"
)

var ReviewsService reviewsServiceInterface = &reviewsService{}

type Reviews []adm.ItemsReview
type Review adm.ItemsReview
type reviewsService struct{}

type reviewsServiceInterface interface {
	ListItemReviewByItemIdPublic(itemId int32) ([]adm.ItemsReview, error)
	ListItemReviewByUserIdPrivate(currentUser adu.AppliUserLogin, userId int32) ([]adm.ItemsReview, error)
	ListItemReviewPrivate() ([]adm.ItemsReview, error)
	GetItemReviewPrivate(currentUser adu.AppliUserLogin, reviewId int32) (adm.ItemsReview, error)
	InsertItemReviewPrivate(currentUser adu.AppliUserLogin, reviewParams adm.CreateItemReviewParams) (adm.ItemsReview, error)
	DeleteItemReviewByUserAndItemPrivate(currentUser adu.AppliUserLogin, userReviewId int32) (int64, error)
}

func (rs *reviewsService) ListItemReviewByItemIdPublic(itemId int32) ([]adm.ItemsReview, error) {
	reviews, err := aadir.ListItemsReviewByItemId(itemId)
	if err != nil {
		return reviews, err
	}
	return reviews, err
}

func (rs *reviewsService) ListItemReviewByUserIdPrivate(currentUser adu.AppliUserLogin, userId int32) ([]adm.ItemsReview, error) {

	reviews, err := aadir.ListItemsReviewsByUser(userId)
	if err != nil {
		return reviews, err
	}
	return reviews, err

}

func (rs *reviewsService) ListItemReviewPrivate() ([]adm.ItemsReview, error) {
	reviews, err := aadir.ListItemsReviews()
	if err != nil {
		return reviews, err
	}
	return reviews, err
}

func (rs *reviewsService) GetItemReviewPrivate(currentUser adu.AppliUserLogin, reviewId int32) (adm.ItemsReview, error) {
	review, err := aadir.GetUserReview(reviewId)
	if err != nil {
		return review, err
	}
	return review, err
}

func (rs *reviewsService) InsertItemReviewPrivate(currentUser adu.AppliUserLogin, reviewParams adm.CreateItemReviewParams) (adm.ItemsReview, error) {
	var review adm.ItemsReview

	if reviewParams.IrUserID == currentUser.User.UID {
		res, err := aadir.InsertReview(reviewParams)
		if err != nil {
			return review, err
		}

		id, err := res.LastInsertId()
		if err != nil {
			return review, err
		}

		review, err := aadir.GetUserReview(int32(id))
		if err != nil {
			return review, err
		}

		return review, nil
	}
	return review, errors.New("user does not have permission to insert this review")
}

func (rs *reviewsService) DeleteItemReviewByUserAndItemPrivate(currentUser adu.AppliUserLogin, userReviewId int32) (int64, error) {

	reviewInfo, err := aadir.GetUserReview(userReviewId)
	if err != nil {
		return 0, err
	}

	isSuperAdmin, err := aua.IsRealySuperAdmin(currentUser.User.UID)
	if err != nil {
		return 0, err
	}

	isAdmin, err := aua.IsRealyAdmin(currentUser.User.UID)
	if err != nil {
		return 0, err
	}

	// Vérifie si l'utilisateur n'est ni un administrateur ni le créateur de la revue
	if !isSuperAdmin && !isAdmin && reviewInfo.IrUserID != currentUser.User.UID {
		return 0, errors.New("user does not have permission to delete this review")
	}

	rows, err := aadir.DeleteItemReviewByUserAndItem(userReviewId)
	if err != nil {
		return rows, err
	}
	return rows, nil
}
