package main

import (
	"fmt"
	"viper/util"
	// "github.com/fsnotify/fsnotify"
	// "github.com/spf13/viper"
)

// func main() {
// 	vp := viper.New()
// 	vp.SetConfigName("test")
// 	vp.SetConfigType("json")
// 	vp.AddConfigPath(".")
// 	err := vp.ReadInConfig()
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	fmt.Println(vp.GetString("hello"))

// 	vp.Set("name", "bikam")
// 	fmt.Println(vp.GetString("name"))

// 	vp.WriteConfig()
// 	vp.OnConfigChange(func(in fsnotify.Event) {
// 		fmt.Printf("file changed: %s\n", in.Name)
// 	})

// 	vp.WatchConfig()
// 	for {
// 	}
// }

func main() {
	config, err := util.LoadConfig()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(config)
	}
}
