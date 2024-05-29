package tfcli

import (
	"context"

	"github.com/hashicorp/go-tfe"
)

var _ Project = (*project)(nil)

type Project interface {
	List(organization string, name string, count int) (*tfe.ProjectList, error)
	Create(organization string, name string, description string) (*tfe.Project, error)
	Show(organization string, name string) (*tfe.Project, error)
	Update(id string, newName string, NewDescription string) (*tfe.Project, error)
	Delete(organization string, name string) error
}

type project struct {
	client *tfe.Client
}

func NewProject(client *tfe.Client) Project {
	return &project{client: client}
}

func (p *project) List(organization string, name string, count int) (*tfe.ProjectList, error) {
	ctx := context.Background()
	projects, err := p.client.Projects.List(ctx, organization, &tfe.ProjectListOptions{
		Name: name,
		ListOptions: tfe.ListOptions{
			PageSize: count,
		},
	})
	if err != nil {
		return &tfe.ProjectList{}, err
	}
	return projects, nil
}

func (p *project) Create(organization string, name string, description string) (*tfe.Project, error) {
	ctx := context.Background()
	project, err := p.client.Projects.Create(ctx, organization, tfe.ProjectCreateOptions{
		Name:        name,
		Description: &description,
	})
	if err != nil {
		return nil, err
	}
	return project, nil
}

func (p *project) Show(organization string, name string) (*tfe.Project, error) {
	projects, err := p.List(organization, name, 1)
	if err != nil {
		return &tfe.Project{}, err
	}
	return projects.Items[0], nil
}

func (p *project) Update(id string, newName string, NewDescription string) (*tfe.Project, error) {
	ctx := context.Background()
	project, err := p.client.Projects.Update(ctx, id, tfe.ProjectUpdateOptions{
		Name:        &newName,
		Description: &NewDescription,
	})
	if err != nil {
		return &tfe.Project{}, err
	}
	return project, nil
}

func (p *project) Delete(organization string, name string) error {
	ctx := context.Background()
	project, err := p.Show(organization, name)
	if err != nil {
		return err
	}
	err = p.client.Projects.Delete(ctx, project.ID)
	return err
}
