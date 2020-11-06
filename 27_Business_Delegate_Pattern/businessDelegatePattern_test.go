package _7_Business_Delegate_Pattern

import "testing"

func TestBusinessDelegatePattern(t *testing.T) {

	businessDelegate := &BusinessDelegate{}
	businessDelegate.serviceType = "EJB"
	client := &Client{businessDelegate}

	tests := []struct {
		name        string
		serviceType string
		want        string
	}{
		{"EJB ", "EJB", "Processing task by invoking EJB Service"},
		{"JMS ", "JMS", "Processing task by invoking JMS Service"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			businessDelegate.serviceType = tt.serviceType
			if got := client.doTask(); got != tt.want {
				t.Errorf("doTask() = %v, want %v", got, tt.want)
			}
		})
	}
}
