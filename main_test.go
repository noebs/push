package push

import (
	_ "embed"
	"testing"
)

func TestNewMessage(t *testing.T) {
	type args struct {
		title string
		body  string
		topic string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{"testing successful messagee7", args{title: "Service Notice", body: "Dear Cashq users, we are running through EBS issues now. Will feedback promptly!", topic: "general"}, false},
		{"testing successful message", args{title: "Hear hearj", body: "yeah that would do it!", topic: "general"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := NewMessage(tt.args.title, tt.args.body, tt.args.topic); (err != nil) != tt.wantErr {
				t.Errorf("NewMessage() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
