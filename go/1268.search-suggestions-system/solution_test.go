package searchsuggestionssystem

import (
	"reflect"
	"testing"
)

func Test_suggestedProducts(t *testing.T) {
	type args struct {
		products   []string
		searchWord string
	}
	tests := []struct {
		name    string
		args    args
		wantAns [][]string
	}{
		{
			name: "tc1",
			args: args{
				products:   []string{"mobile", "mouse", "moneypot", "monitor", "mousepad"},
				searchWord: "mouse",
			},
			wantAns: [][]string{
				{"mobile", "moneypot", "monitor"},
				{"mobile", "moneypot", "monitor"},
				{"mouse", "mousepad"},
				{"mouse", "mousepad"},
				{"mouse", "mousepad"},
			},
		},
		{
			name: "tc2",
			args: args{
				products:   []string{"havana"},
				searchWord: "havana",
			},
			wantAns: [][]string{
				{"havana"},
				{"havana"},
				{"havana"},
				{"havana"},
				{"havana"},
				{"havana"},
			},
		},
		{
			name: "tc3",
			args: args{
				products:   []string{"bags", "baggage", "banner", "box", "cloths"},
				searchWord: "bags",
			},
			wantAns: [][]string{
				{"baggage", "bags", "banner"},
				{"baggage", "bags", "banner"},
				{"baggage", "bags"},
				{"bags"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := suggestedProducts(tt.args.products, tt.args.searchWord); !reflect.DeepEqual(gotAns, tt.wantAns) {
				t.Errorf("suggestedProducts() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
