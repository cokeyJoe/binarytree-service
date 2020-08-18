package server

import "testing"

func TestServer_Start(t *testing.T) {
	type fields struct {
		treeService TreeServiceSource
		listener    Listener
		logger      Logger
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Server{
				treeService: tt.fields.treeService,
				listener:    tt.fields.listener,
				logger:      tt.fields.logger,
			}
			if err := s.Start(); (err != nil) != tt.wantErr {
				t.Errorf("Server.Start() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
