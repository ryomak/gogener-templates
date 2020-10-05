package handler

import (
	"net/http"

	"[[.ModName]]/application/usecase"
	"[[.ModName]]/application/usecase/request"
	"[[.ModName]]/pkg/log"
	"[[.ModName]]/presentation/v1/resource/factory"
)

type Me interface {
	Get(w http.ResponseWriter, r *http.Request)
}

type me struct {
	meUsecase usecase.Me
}

func NewMe(
	meUsecase usecase.Me,
) Me {
	return &me{
		meUsecase: meUsecase,
	}
}

// MeGet
// @Summary 自分の情報を取得
// @Accept  json
// @Produce  json
// @Param data body request.MeGet true "Request body"
// @Success 200 {object} response.MeGet
// @Failure 500 {object} resource.Error "something went wrong"
// @Router /api/v1/me [get]
func (m me) Get(w http.ResponseWriter, r *http.Request) {
	req, err := request.NewMeGet(r)
	if err != nil {
		log.Warningf(r.Context(), "Me.Get.Request %v", err)
		factory.ErrorJSON(w, err, http.StatusBadRequest)
		return
	}
	res, err := m.meUsecase.Get(r.Context(), req)
	if err != nil {
		log.Errorf(r.Context(), "Me.Get: %v", err)
		factory.ErrorJSON(w, err, http.StatusInternalServerError)
		return
	}
	factory.JSON(w, res)
}
