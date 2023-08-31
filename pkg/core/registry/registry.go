package registry

import "github.com/somkieatW/candidate/pkg/repository"

type RepositoryRegistry struct {
	AppointmentRepository repository.AppointmentRepository
	UserRepository        repository.UserRepository
	CommentRepository     repository.CommentRepository
}
