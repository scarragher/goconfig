package config

import (
	"fmt"
	"testing"
)

func TestConfigurations(t *testing.T) {

	boolConfig := NewConfiguration("IsTestEnabled", true)
	stringConfig := NewConfiguration("MyName", "Stephen")
	intConfig := NewConfiguration("AttemptsAtGoal", 9)
	float64Config := NewConfiguration("Heat", 14.)

	fmt.Println(boolConfig)
	fmt.Println(stringConfig)
	fmt.Println(intConfig)
	fmt.Println(float64Config)

	AddConfiguration(*boolConfig)
	AddConfiguration(*stringConfig)
	AddConfiguration(*intConfig)
	AddConfiguration(*float64Config)

	boolGConfig, _ := GetConfiguration("IsTestEnabled")
	stringGConfig, _ := GetConfiguration("MyName")
	intGConfig, _ := GetConfiguration("AttemptsAtGoal")
	float64GConfig, _ := GetConfiguration("Heat")

	fmt.Println(boolGConfig)
	fmt.Println(stringGConfig)
	fmt.Println(intGConfig)
	fmt.Println(float64GConfig)

	if GetConfigurationValue("IsTestEnabled").(bool) == true {
		fmt.Println("does the compiler know this is a bool?")
	}

	if GetConfigurationValue("AttemptsAtGoal").(int) > 3 {
		fmt.Println("is this an int?")
	}

	err := SaveSpecified("config.json", "c:\\test")

	if err != nil {
		fmt.Println(err)
	}

	enabled := ValueBool("IsTestEnabled")

	fmt.Println(enabled)

}
