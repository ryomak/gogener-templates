package middleware

import (
	"net/http"

	"[[.ModName]]/domain/service"
	"[[.ModName]]/presentation/api/v1/resource/factory"
)

type MaintenanceMiddleware struct {
	maintenanceService service.Maintenance
}

func NewMaintenance(
	maintenanceService service.Maintenance,
) *MaintenanceMiddleware {
	return &MaintenanceMiddleware{
		maintenanceService: maintenanceService,
	}
}

func (m MaintenanceMiddleware) All(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		if err := m.maintenanceService.CheckByName(r.Context(), "all"); err != nil {
			factory.ErrorJSON(
				w,
				err,
				http.StatusServiceUnavailable,
			)
			return
		}
		next.ServeHTTP(w, r)
	}
	return http.HandlerFunc(fn)
}
