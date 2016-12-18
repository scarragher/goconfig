package config

import (
	"fmt"
	"log"
	"testing"

	"github.com/scarragher/goconfig"
)

func TestConfigurations(t *testing.T) {

	boolConfig := config.NewConfiguration("IsTestEnabled", true)
	stringConfig := config.NewConfiguration("MyName", "Stephen")
	intConfig := config.NewConfiguration("AttemptsAtGoal", 9)
	float64Config := config.NewConfiguration("Heat", 14.)

	fmt.Println(boolConfig)
	fmt.Println(stringConfig)
	fmt.Println(intConfig)
	fmt.Println(float64Config)

	config.AddConfiguration(*boolConfig)
	config.AddConfiguration(*stringConfig)
	config.AddConfiguration(*intConfig)
	config.AddConfiguration(*float64Config)

	boolGConfig, _ := config.GetConfiguration("IsTestEnabled")
	stringGConfig, _ := config.GetConfiguration("MyName")
	intGConfig, _ := config.GetConfiguration("AttemptsAtGoal")
	float64GConfig, _ := config.GetConfiguration("Heat")

	fmt.Println(boolGConfig)
	fmt.Println(stringGConfig)
	fmt.Println(intGConfig)
	fmt.Println(float64GConfig)

	if config.GetConfigurationValue("IsTestEnabled").(bool) == true {
		fmt.Println("OK")
	}

	if config.GetConfigurationValue("AttemptsAtGoal").(int) > 3 {
		fmt.Println("OK")
	}

	err := config.Save("config.json", "c:\\test")

	if err != nil {
		fmt.Println(err)
	}

	if !boolGConfig.ValueBool() {
		log.Fatal("fail")
	}

}
