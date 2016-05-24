package imcg

import (
	"testing"
)

func TestGetMobileNumberData(t *testing.T) {
	input1 := "9711889900"    // Valid input.
	input2 := "9818998866"    // Valid input.
	input3 := "8813823345"    // Valid input.
	input4 := "+919711889900" // Valid input.
	input5 := "919711889900"  // Valid input.
	input6 := "8999887788"    // Data does not exist.
	input7 := "99887766"      // Improper length.
	input8 := "819711889900"  // Improper prefix.

	res := GetMobileNumberData(input1)
	if res[0] != "VODAFONE" && res[1] != "DELHI" {
		t.Error("Expected [VODAFONE,DELHI], Got ", res)
	}

	res = GetMobileNumberData(input2)
	if res[0] != "AIRTEL" && res[1] != "DELHI" {
		t.Error("Expected [AIRTEL,DELHI], Got ", res)
	}

	res = GetMobileNumberData(input3)
	if res[0] != "VODAFONE" && res[1] != "HARYANA" {
		t.Error("Expected [VODAFONE,HARYANA], Got ", res)
	}

	res = GetMobileNumberData(input4)
	if res[0] != "VODAFONE" && res[1] != "DELHI" {
		t.Error("Expected [VODAFONE,DELHI], Got ", res)
	}

	res = GetMobileNumberData(input5)
	if res[0] != "VODAFONE" && res[1] != "DELHI" {
		t.Error("Expected [VODAFONE,DELHI], Got ", res)
	}

	res = GetMobileNumberData(input6)
	if res != nil {
		t.Error("Expected nil, Got ", res)
	}

	res = GetMobileNumberData(input7)
	if res != nil {
		t.Error("Expected nil, Got ", res)
	}

	res = GetMobileNumberData(input8)
	if res != nil {
		t.Error("Expected nil, Got ", res)
	}
}
