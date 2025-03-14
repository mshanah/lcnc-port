package port

import "github.com/mshanah/lcnc-domain/domain"

type WorkspaceRepository interface {
	Save(workspace *domain.Workspace) error
	FindByID(id string) (*domain.Workspace, error)
	FindAll() ([]*domain.Workspace, error)
	Delete(id string) error
}
