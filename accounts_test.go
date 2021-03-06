package goAccounts

import (
	"fmt"
	// "strconv"
	"testing"
)

func TestCreateUser(t *testing.T) {
	token := "eyJhbGciOiJSUzI1NiIsInR5cCIgOiAiSldUIiwia2lkIiA6ICJ3Ujh3X1lGOWpKWFRWQ2x2VHF1RkswZkctQXROQUJsb3FBd0N4MmlTTWQ4In0.eyJleHAiOjE2MjA0ODIyODksImlhdCI6MTYyMDQ3NTA4OSwiYXV0aF90aW1lIjoxNjIwNDc1MDU3LCJqdGkiOiJhNmFjODk2My01ZjcxLTQ1NWYtYjYzMS01ODY3ZTcwN2VlNDEiLCJpc3MiOiJodHRwczovL2xvZ2luLmphcXBvdC5vcmcvYXV0aC9yZWFsbXMvamFxcG90IiwiYXVkIjoiYWNjb3VudCIsInN1YiI6IjI0MjVkNzYwLTAxOGQtNDA4YS1hZTBiLWNkZTRjNTYzNTRiOSIsInR5cCI6IkJlYXJlciIsImF6cCI6ImphcXBvdC11aS1jb2RlIiwibm9uY2UiOiI3MjdkNDIxYzllN2JkMGVjNDNkNTk0NjY0MjAyN2RhNTQxU0xiWllVSCIsInNlc3Npb25fc3RhdGUiOiIwYzAwNTQyNi02M2I4LTRhNTMtOWVhMy1mMTM1ZTkzMWIxYjAiLCJhY3IiOiIwIiwiYWxsb3dlZC1vcmlnaW5zIjpbIicqJyIsIioiXSwicmVhbG1fYWNjZXNzIjp7InJvbGVzIjpbIm9mZmxpbmVfYWNjZXNzIiwidW1hX2F1dGhvcml6YXRpb24iXX0sInJlc291cmNlX2FjY2VzcyI6eyJhY2NvdW50Ijp7InJvbGVzIjpbIm1hbmFnZS1hY2NvdW50IiwibWFuYWdlLWFjY291bnQtbGlua3MiLCJ2aWV3LXByb2ZpbGUiXX19LCJzY29wZSI6Im9wZW5pZCBqYXFwb3QtYWNjb3VudHMgZW1haWwgcHJvZmlsZSB3cml0ZSByZWFkIiwiZW1haWxfdmVyaWZpZWQiOmZhbHNlLCJuYW1lIjoiUGFudGVsaXMgS2FyYXR6YXMiLCJwcmVmZXJyZWRfdXNlcm5hbWUiOiJwYW50ZWxpc3BhbmthIiwiZ2l2ZW5fbmFtZSI6IlBhbnRlbGlzIiwiZmFtaWx5X25hbWUiOiJLYXJhdHphcyIsImVtYWlsIjoicGFudGVsaXNwYW5rYUBnbWFpbC5jb20ifQ.Y75tmMu7WNUgW6Ma2SEnx91JDqw4qNhBKMgxQ7OCJvyb29shMdSAd9hFwlXESjnDEjVsvV0-eNid5h3Y0hyA7AfBVTWPqzPrQZQGMt5o-7CIvNVZmYJ2B_yRoO3-v8Z3-8lpjPSkybpWowJgQVfBehOMRXS8AMRp_m57LZ2cRAb4r-N_8JqH1Nbune9Pdm5dCno-yi4oDyO8TN9dJII7lHTVPFLGUD73xjCejDAulEYytgTcQ7cffMYPEuUnv5SA3RsPt0TP6Ptwk-DQaiSuj1kHXAviX7ztxTDuHRmzGn811ZGgS19jAOotyWLxyQuhiOq2fF_5Xmc7i58RDvNfQg"
	q := InitClient("https://accountsapi.jaqpot.org")
	var IAccountsClient = q
	orgUser, err := IAccountsClient.GetOrganization("TqKqs0m8v26mji68", token)
	if err != nil {
		// t.Log("Passing test with error:" + err.Error())
		fmt.Println("Passing test with error:" + err.Error())
	}
	if orgUser.Id != "TqKqs0m8v26mji68" {
		t.Error("Not the same user id")
	}
}
