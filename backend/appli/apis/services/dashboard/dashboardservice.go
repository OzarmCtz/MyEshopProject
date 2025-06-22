package userservice

import (
	"errors"

	"github.com/OzarmCtz/e_shop_backend_v1/app/data/global"
	aadg "github.com/OzarmCtz/e_shop_backend_v1/appli/apis/domain/groups/users"
	aadgu "github.com/OzarmCtz/e_shop_backend_v1/appli/apis/domain/groups/users/link"
	adu "github.com/OzarmCtz/e_shop_backend_v1/appli/apis/domain/users"
)

var (
	DashboardService dashboardServiceInterface = &dashboardService{}
)

type dashboardService struct{}

type dashboardServiceInterface interface {
	CheckLoginDashboardAccess(currentUser adu.AppliUserLogin) error
}

func (us *dashboardService) CheckLoginDashboardAccess(currentUser adu.AppliUserLogin) error {
	// Récupérer le groupe de l'utilisateur
	userGroup, err := aadgu.GetGroupUserLinkByUser(currentUser.User.UID)
	if err != nil {
		return err
	}

	// Récupérer les détails du groupe
	userGroupName, err := aadg.GetGroupUser(userGroup.GulGroupID)
	if err != nil {
		return err
	}

	// Vérifier les autorisations
	if userGroupName.GuName != global.ADMIN_STATUS && userGroupName.GuName != global.SUPER_ADMIN_STATUS {
		return errors.New("you are not authorized to login")
	}

	// Si on arrive ici, l'utilisateur est autorisé
	return nil
}
