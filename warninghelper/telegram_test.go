package warninghelper

import "testing"

func TestTelegramWarning_sendMessage(t *testing.T) {
	type args struct {
		msg interface{}
	}
	tests := []struct {
		name    string
		tr      *TelegramWarning
		args    args
		wantErr bool
	}{
		{
			name: "Test_sendMessage",
			tr: &TelegramWarning{
				botToken: "",
				chatId:   "",
			},
			args: args{
				msg: "早安你好",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.tr.sendMessage(tt.args.msg); (err != nil) != tt.wantErr {
				t.Errorf("TelegramWarning.sendMessage() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
