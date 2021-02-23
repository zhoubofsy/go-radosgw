package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"

	"liyongcool.nat300.top/bozhou/go-radosgw/pkg/api"
)

func printRawMode(out io.Writer, data interface{}) error {
	js, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return err
	}
	fmt.Fprintf(out, "%s\n", js)
	return nil
}

func main() {
	//api, err := radosAPI.New(os.Getenv("RADOSGW_API"), os.Getenv("RADOSGW_ACCESS"), os.Getenv("RADOSGW_SECRET"))
	api, err := radosAPI.New("http://120.48.27.190:80", "C3ZBITE3VS5AD4Y3YEZB", "3cZZ8D7mP0hNCiqUIYnxKmhEPmbzcCFBkr7Bz4ey")
	if err != nil {
		log.Fatal(err)
	}

	// create a new user named JohnDoe
	user, err, _ := api.CreateUser(radosAPI.UserConfig{
		Tenant:      "ccb",
		UID:         "JohnDoe",
		DisplayName: "John Doe",
	})
	/*
		if err != nil {
			log.Fatal(err)
		}
	*/
	printRawMode(os.Stdout, user)

	// get user
	user, err = api.GetUser("U456123")

	// get user quota
	userQuotas, err := api.GetQuotas(radosAPI.QuotaConfig{UID: "U456123", QuotaType: "user"})
	bucketQuotas, err := api.GetQuotas(radosAPI.QuotaConfig{UID: "U456123", QuotaType: "bucket"})

	printRawMode(os.Stdout, userQuotas)
	printRawMode(os.Stdout, bucketQuotas)

	userInfo, err := api.GetUserInfo(radosAPI.UserInfoConfig{UID: "U456123", Stats: "true", Sync: "true"})
	printRawMode(os.Stdout, userInfo)
	// remove JohnDoe
	//err = api.RemoveUser(radosAPI.UserConfig{
	//	UID: "JohnDoe",
	//})
	//if err != nil {
	//	log.Fatal(err)
	//}
}
