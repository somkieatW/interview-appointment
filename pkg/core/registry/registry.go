package registry

import "github.com/somkieatW/interview-appointment/pkg/repository"

type RepositoryRegistry struct {
	AppointmentRepository repository.AppointmentRepository
	UserRepository        repository.UserRepository
	CommentRepository     repository.CommentRepository
}
