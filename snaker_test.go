package snaker

import (
	"testing"
)

func TestCamelToSnake(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name      string
		args      args
		want, msg string
	}{
		{"empty", args{""}, "", "should return an empty string on an empty input"},
		{"One", args{"One"}, "one", "should work with one word"},
		{"ONE", args{"ONE"}, "o_n_e", "should return an uppercase string as separate words"},
		{"ID", args{"ID"}, "id", "should return ID as lowercase"},
		{"i", args{"i"}, "i", "should work with a single lowercase character"},
		{"I", args{"I"}, "i", "should work with a single uppercase character"},
		{"ThisHasToBeConvertedCorrectlyID", args{"ThisHasToBeConvertedCorrectlyID"}, "this_has_to_be_converted_correctly_id", "should return a long text as expected"},
		{"ThisIDIsFine", args{"ThisIDIsFine"}, "this_id_is_fine", "should return the text as expected if the initialism is in the middle"},
		{"ThisHTTPSConnection", args{"ThisHTTPSConnection"}, "this_https_connection", "should work with long initialism"},
		{"HTTPSID", args{"HTTPSID"}, "https_id", "should work with multi initializes"},
		{"OAuthClient", args{"OAuthClient"}, "oauth_client", "should work with initialism where only certain characters are uppercase"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CamelToSnake(tt.args.s); got != tt.want {
				t.Errorf("CamelToSnake() = %v, want %v, %s!", got, tt.want, tt.msg)
			}
		})
	}
}

func TestSnakeToCamel(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name      string
		args      args
		want, msg string
	}{
		{"empty", args{""}, "", "should return an empty string on an empty input"},
		{"potato_", args{"potato_"}, "Potato", "should not blow up on trailing _"},
		{"this_has_to_be_uppercased", args{"this_has_to_be_uppercased"}, "ThisHasToBeUppercased", "should return a snaked text as camel case"},
		{"this_is_an_id", args{"this_is_an_id"}, "ThisIsAnID", "should return a snaked text as camel case, except the word ID"},
		{"this_is_an_identifier", args{"this_is_an_identifier"}, "ThisIsAnIdentifier", "should return 'id' not as uppercase"},
		{"id", args{"id"}, "ID", "should simply work with id"},
		{"oauth_client", args{"oauth_client"}, "OAuthClient", "should work with initialism where only certain characters are uppercase"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SnakeToCamel(tt.args.s); got != tt.want {
				t.Errorf("SnakeToCamel() = %v, want %v, %s!", got, tt.want, tt.msg)
			}
		})
	}
}

func TestSnakeToCamelLower(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name      string
		args      args
		want, msg string
	}{
		{"empty", args{""}, "", "should return an empty string on an empty input"},
		{"potato_", args{"potato_"}, "potato", "should not blow up on trailing _"},
		{"this_has_to_be_uppercased", args{"this_has_to_be_uppercased"}, "thisHasToBeUppercased", "should return a snaked text as camel case"},
		{"this_is_an_id", args{"this_is_an_id"}, "thisIsAnID", "should return a snaked text as camel case, except the word ID"},
		{"this_is_an_identifier", args{"this_is_an_identifier"}, "thisIsAnIdentifier", "should return 'id' not as uppercase"},
		{"id", args{"id"}, "id", "should simply work with id"},
		{"id_me_please", args{"id_me_please"}, "idMePlease", "should simply work with leading id"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SnakeToCamelLower(tt.args.s); got != tt.want {
				t.Errorf("SnakeToCamelLower() = %v, want %v", got, tt.want)
			}
		})
	}
}
