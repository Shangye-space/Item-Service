package helpers_test

import (
	"testing"

	"github.com/Shangye-space/Item-Service/src/api/helpers"
)

func TestCheckIDRight(t *testing.T) {
	id := 5
	err := helpers.CheckID(&id)
	if err != nil {
		t.Error("ID is not valid")
	}

}

func TestCheckIDWrong(t *testing.T) {
	var id int
	err := helpers.CheckID(&id)
	if err == nil {
		t.Error("ID is not valid")
	}

}

func TestCheckBoolTrue(t *testing.T) {
	bool := true
	num, err := helpers.CheckBool(&bool)
	if num != 1 {
		t.Errorf("Bool is not valid, got error %v", err)
	}
}

func TestCheckBoolFalse(t *testing.T) {
	bool := false
	num, err := helpers.CheckBool(&bool)
	if num != 0 {
		t.Errorf("Bool is not valid, got error %v", err)
	}
}

func TestCheckBoolInvalid(t *testing.T) {
	var bool bool
	num, err := helpers.CheckBool(&bool)
	if err != nil {
		t.Errorf("%v is not valid, got error %v", num, err)
	}
}

func TestCheckString(t *testing.T){
	string := "text" 
	err := helpers.CheckString(&string)
	if err != nil{
		t.Errorf("Got error %v", err)
	}
}
func TestCheckStringInvalid(t *testing.T){
	var string string
	err := helpers.CheckString(&string)
	if err == nil{
		t.Errorf("String was accepted, but should not be, error: %v", err)
	}
}

func TestCheckStringEmpty(t *testing.T){
	string := "" 
	err := helpers.CheckString(&string)
	if err == nil{
		t.Errorf("String was accepted, but should not be, error: %v", err)
	}
}


func TestCheckNumber(t *testing.T){
	var num float32
	num = 155.00
	err := helpers.CheckNumber(&num)
	if err != nil{
		t.Errorf("Number should be accepted, error: %v", err)
	}
}

func TestCheckNumberNil(t *testing.T){
	var num float32
	err := helpers.CheckNumber(&num)
	if err == nil{
		t.Errorf("Number should not be accepted, error: %v", err)
	}
}

func TestCheckNumberNegative(t *testing.T){
	var num float32
	num = -55.00
	err := helpers.CheckNumber(&num)
	if err == nil{
		t.Errorf("Number should not be accepted, error: %v", err)
	}
}

func TestCheckNumberInt(t *testing.T){
	num := 55
	err := helpers.CheckNumberInt(&num)
	if err != nil{
		t.Errorf("Number should be accepted, error: %v", err)
	}
}

func TestCheckNumberIntNil(t *testing.T){
	var num int
	err := helpers.CheckNumberInt(&num)
	if err == nil{
		t.Errorf("Number should not be accepted, error: %v", err)
	}
}

func TestCheckNumberIntNegative(t *testing.T){
	num := -55
	err := helpers.CheckNumberInt(&num)
	if err == nil{
		t.Errorf("Number should not be accepted, error: %v", err)
	}
}