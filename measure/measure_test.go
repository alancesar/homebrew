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
		wantErr    bool
	}{
		{
			name: "Should extract elements from decimal wantValue",
			args: args{
				input: "12.34kg",
			},
			wantSymbol: "kg",
			wantValue:  12.34,
			wantErr:    false,
		},
		{
			name: "Should extract elements from integer wantValue",
			args: args{
				input: "12oz",
			},
			wantSymbol: "oz",
			wantValue:  12,
			wantErr:    false,
		},
		{
			name: "Should extract elements with blank space between wantValue and wantSymbol",
			args: args{
				input: "34 g",
			},
			wantSymbol: "g",
			wantValue:  34,
			wantErr:    false,
		},
		{
			name: "Should extract elements with uppercase and return lowercase",
			args: args{
				input: "56 L",
			},
			wantSymbol: "l",
			wantValue:  56,
			wantErr:    false,
		},
		{
			name: "Should return error if not contains symbol",
			args: args{
				input: "12",
			},
			wantErr: true,
		},
		{
			name: "Should return error if not contains value",
			args: args{
				input: "g",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, err := ExtractSymbolAndValue(tt.args.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("ExtractSymbolAndValue() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.wantSymbol {
				t.Errorf("ExtractSymbolAndValue() got = %v, wantSymbol %v", got, tt.wantSymbol)
			}
			if got1 != tt.wantValue {
				t.Errorf("ExtractSymbolAndValue() got1 = %v, wantSymbol %v", got1, tt.wantValue)
			}
		})
	}
}
