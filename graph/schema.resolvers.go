package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"Go_Graphql/graph/generated"
	"Go_Graphql/graph/model"
	"context"
	"errors"
	_"fmt"
)

func (r *mutationResolver) CreateComputer(ctx context.Context, input *model.NewComputer) (*model.Computer, error) {
	if input.RAM < 2 {
		return nil, errors.New("Ram must be greater than 1")
	}

	if input.Battery < 20 {
		return nil, errors.New("Battery power must be greater than 19")
	}

	computer := model.Computer{
		RAM: input.RAM,
		Battery: input.Battery,
		Proccessor: input.Proccessor,
		Gpu: input.Gpu,
	}

	_,err := r.DB.Model(&computer).Insert()

	if err != nil {
		return nil, errors.New("Insert new computer failed!")
	}

	return &computer, nil
}

func (r *mutationResolver) UpdateComputer(ctx context.Context, id string, input *model.NewComputer) (*model.Computer, error) {
	var computer model.Computer

	err := r.DB.Model(&computer).Where("id = ?",id).First()

	if err != nil {
		return nil, errors.New("Computer not found!")
	}

	computer.RAM = input.RAM
	computer.Battery = input.Battery
	computer.Proccessor = input.Proccessor
	computer.Gpu = input.Gpu

	_, updateErr := r.DB.Model(&computer).Where("id = ?", id).Update()

	if updateErr != nil {
		return nil, errors.New("Update computer failed!")
	}
	return &computer, nil

}

func (r *mutationResolver) DeleteComputer(ctx context.Context, id string) (bool, error) {
	var computer model.Computer

	err := r.DB.Model(&computer).Where("id = ?",id).First()

	if err != nil {
		return false, errors.New("Computer not found!")
	}

	_, deleteError := r.DB.Model(&computer).Where("id = ?", id).Delete()

	if deleteError != nil {
		return false, errors.New("delete computer failed!")
	}

	return true, nil
}

func (r *queryResolver) Computers(ctx context.Context) ([]*model.Computer, error) {
	var computers []*model.Computer

	err := r.DB.Model(&computers).Order("id").Select()

	if err != nil {
		return nil, errors.New("failed to query computers")
	}

	return computers, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }