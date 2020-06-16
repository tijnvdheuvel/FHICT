package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	uri := "https://api.fhict.nl/location/current"
	token := "eyJ0eXAiOiJKV1QiLCJhbGciOiJSUzI1NiIsIng1dCI6ImhPWFBYUmtteU5kc1ViMFMtN2Vlc2xOUEI0OCIsImtpZCI6ImhPWFBYUmtteU5kc1ViMFMtN2Vlc2xOUEI0OCJ9.eyJpc3MiOiJodHRwczovL2lkZW50aXR5LmZoaWN0Lm5sIiwiYXVkIjoiaHR0cHM6Ly9pZGVudGl0eS5maGljdC5ubC9yZXNvdXJjZXMiLCJleHAiOjE1ODE1OTQzNDEsIm5iZiI6MTU4MTU4NzE0MSwiY2xpZW50X2lkIjoiYXBpLWNsaWVudCIsInVybjpubC5maGljdDp0cnVzdGVkX2NsaWVudCI6InRydWUiLCJzY29wZSI6WyJvcGVuaWQiLCJwcm9maWxlIiwiZW1haWwiLCJmaGljdCIsImZoaWN0X3BlcnNvbmFsIiwiZmhpY3RfbG9jYXRpb24iXSwic3ViIjoiYmI1YmFkOWEtYWQ4YS00OGFhLTk4OWUtNWM5MTRjZTU3Yzc0IiwiYXV0aF90aW1lIjoxNTgxNTg3MTQxLCJpZHAiOiJmaGljdC1zc28iLCJyb2xlIjpbInVzZXIiLCJtZWRld2Vya2VyIl0sInVwbiI6Ikk4ODI1OTNAZmhpY3QubmwiLCJuYW1lIjoiTWljaGllbHNlbixCYXMgQi5TLkguVC4iLCJlbWFpbCI6ImIubWljaGllbHNlbkBmb250eXMubmwiLCJ1cm46bmwuZmhpY3Q6c2NoZWR1bGUiOiJ0ZWFjaGVyfE1pY2giLCJmb250eXNfdXBuIjoiODgyNTkzQGZvbnR5cy5ubCIsImFtciI6WyJleHRlcm5hbCJdfQ.BHkqdaMD0X7hPTc0HtJMpyDWm_6enmZzI5zLtgxUqnvLMNZSG0JQo-MH3vdna3gPLxkkis13G8YYeCWxhLtWOP6CYIKp5WddSMTyJVnWDkIsl5kp553NYgpPLCFVciXR6dHLB49COcmen8GIsJ3mMgWRQBbEkb5gy1QwomcpPxXYfl6yMS3Gz85jhQBGPGwsczcoG3qvV6TBeL8O_Hl4NF0bvb5rxZlkj7cCp5eF1JxI2ZFboTs4QH7I52wSzb8k2t2CDbeMpLkpvf8TGnE9tmvdxN75Yo2bbjMwCt_YMk1zKHddkbZNXG0cBD5SBCtLF7Rx5-e5srHi4ejxTQT58A"

	request, err := http.NewRequest("GET", uri, nil)
	if err != nil {
		panic(err)
	}
	request.Header.Set("Authorization", "Bearer "+token)

	client := http.Client{}
	response, err := client.Do(request)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(data))
}
