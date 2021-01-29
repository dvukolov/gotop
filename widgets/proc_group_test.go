package widgets

import "testing"

func Test_getGroup(t *testing.T) {
	tests := []struct {
		name      string
		fullCmd   string
		wantGroup string
	}{
		{name: "path", fullCmd: "/usr/lib/firefox/firefox", wantGroup: "firefox"},
		{name: "args", fullCmd: "/usr/lib/firefox/firefox -new-window", wantGroup: "firefox"},
		{name: "brackets", fullCmd: "[watchdogd]", wantGroup: "watchdogd"},
		{name: "slash", fullCmd: "[kworker/2:1H-events_highpri]", wantGroup: "kworker"},
		{name: "colon", fullCmd: "avahi-daemon: chroot helper", wantGroup: "avahi-daemon"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotGroup := getGroup(tt.fullCmd); gotGroup != tt.wantGroup {
				t.Errorf("getGroup() = %v, want %v", gotGroup, tt.wantGroup)
			}
		})
	}
}
