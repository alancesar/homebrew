package measure

import "testing"

func TestExtractSymbolAndValue(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name       string
		args       args
		wantSymbol string
		wantValue  float64
		wantMatch  bool
	}{
		{
			name: "Should extract elements from decimal wantValue",
			args: args{
				input: "12.34kg",
			},
			wantSymbol: "kg",
			wantValue:  12.34,
			wantMatch:  true,
		},
		{
			name: "Should extract elements from integer wantValue",
			args: args{
				input: "12oz",
			},
			wantSymbol: "oz",
			wantValue:  12,
			wantMatch:  true,
		},
		{
			name: "Should extract elements with blank space between wantValue and wantSymbol",
			args: args{
				input: "34 g",
			},
			wantSymbol: "g",
			wantValue:  34,
			wantMatch:  true,
		},
		{
			name: "Should extract elements with uppercase and return lowercase",
			args: args{
				input: "56 L",
			},
			wantSymbol: "l",
			wantValue:  56,
			wantMatch:  true,
		},
		{
			name: "Should not wantMatch if not contains symbol",
			args: args{
				input: "12",
			},
			wantMatch: false,
		},
		{
			name: "Should not wantMatch if not contains value",
			args: args{
				input: "g",
			},
			wantMatch: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			symbol, value, match := ExtractSymbolAndValue(tt.args.input)
			if match != tt.wantMatch {
				t.Errorf("ExtractSymbolAndValue() wantMatch = %v, wantMatch %v", match, tt.wantMatch)
				return
			}
			if symbol != tt.wantSymbol {
				t.Errorf("ExtractSymbolAndValue() symbol = %v, wantSymbol %v", symbol, tt.wantSymbol)
			}
			if value != tt.wantValue {
				t.Errorf("ExtractSymbolAndValue() value = %v, wantSymbol %v", value, tt.wantValue)
			}
		})
	}
}
