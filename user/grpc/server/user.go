package server

import (
	"errors"

	"golang.org/x/net/context"

	internal "github.com/Marlos-Rodriguez/go-postgres-wallet-back/internal/storage"
	"github.com/Marlos-Rodriguez/go-postgres-wallet-back/user/models"
	"github.com/Marlos-Rodriguez/go-postgres-wallet-back/user/storage"
	"github.com/jinzhu/gorm"
)

//Server User Server struct
type Server struct {
}

func getStorageService() *storage.UserStorageService {
	var DB *gorm.DB = internal.ConnectDB()

	return storage.NewUserStorageService(DB)
}

//CheckUser check if the user exits
func (s *Server) CheckUser(ctx context.Context, request *UserRequest) (*UserResponse, error) {
	storageService := getStorageService()

	if len(request.ID) < 0 || request.ID == "" {
		return &UserResponse{Exits: false, Active: false}, errors.New("Must send a ID")
	}

	exits, isActive, err := storageService.CheckExistingUser(request.ID)

	if err != nil {
		return &UserResponse{Exits: false, Active: false}, err
	}

	storageService.CloseDB()

	return &UserResponse{Exits: exits, Active: isActive}, nil
}

//CheckRelation Check if exits a Relation
func (s *Server) CheckRelation(ctx context.Context, request *RelationRequest) (*RelationResponse, error) {
	storageService := getStorageService()

	if len(request.FromUsername) < 0 || request.FromUsername == "" && len(request.ToUsername) < 0 || request.ToUsername == "" {
		return &RelationResponse{Exits: false}, errors.New("Must send ID")
	}

	exits, err := storageService.CheckExistingRelation(request.FromUsername, request.ToUsername, false)

	if err != nil {
		return &RelationResponse{Exits: false}, err
	}

	storageService.CloseDB()

	return &RelationResponse{Exits: exits}, nil
}

//ChangeAvatar Change the avatar in DB
func (s *Server) ChangeAvatar(ctx context.Context, request *AvatarName) (*AvatarResponse, error) {
	storageService := getStorageService()

	if len(request.Name) < 0 || request.Name == "" {
		return &AvatarResponse{Sucess: false}, errors.New("Must send the avatar name")
	}

	var userDB *models.User = new(models.User)

	userDB.Profile.Avatar = request.Name

	if sucess, err := storageService.ModifyUser(userDB, "", ""); sucess == false || err != nil {
		return &AvatarResponse{Sucess: false}, err
	}

	return &AvatarResponse{Sucess: true}, nil
}