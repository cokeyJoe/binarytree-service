package server

import (
	"binarytree/pkg/logging"
	"os"

	"github.com/pkg/errors"
)

var (
	ErrBstServiceIsNil = errors.New("bst service is nil")
	ErrServerStopped   = errors.New("server is stopped")
	ErrListenerIsNil   = errors.New("listener is nil")
)

type Server struct {
	treeService TreeServiceSource

	listener Listener

	logger Logger
}

func New(listener Listener, source TreeServiceSource, logger Logger) *Server {
	return &Server{
		treeService: source,
		listener:    listener,
		logger:      logger,
	}
}

type Listener interface {
	ListenAndServe() error
	Shutdown() error
}

type Logger interface {
	ErrorWithFields(string, logging.Fields)
}

type TreeServiceSource interface {
	Open() error
	Close() error
}

func (s *Server) Start() error {

	var err error

	err = s.initBST()
	if err != nil {
		s.logger.ErrorWithFields(err.Error(), logging.Fields{
			"event": "initbstService",
		})
		return err
	}

	s.registerOSSignal()

	err = s.startListener()
	if err != nil {
		s.logger.ErrorWithFields(err.Error(), logging.Fields{
			"event": "startListener",
		})
	}

	return ErrServerStopped
}

func (s *Server) Stop() {
	if s.treeService != nil {
		s.treeService.Close()
	}

	if s.listener != nil {
		s.listener.Shutdown()
	}
}

func (s *Server) initBST() error {

	if s.treeService == nil {
		return errors.Wrap(ErrBstServiceIsNil, "failed to init bst")
	}

	return s.treeService.Open()

}

func (s *Server) registerOSSignal() {
	sig := make(chan os.Signal, 1)

	go func() {
		<-sig
		s.Stop()
	}()
}

func (s *Server) startListener() error {
	if s.listener == nil {
		return errors.Wrap(ErrListenerIsNil, "failed to start listener")
	}

	return s.listener.ListenAndServe()
}
