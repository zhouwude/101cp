package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
	authentication "k8s.io/api/authentication/v1beta1"
)

func main() {
	http.HandleFunc("/authenticate", func(w http.ResponseWriter, r *http.Request) {
		decoder := json.NewDecoder(r.Body)
		var tr authentication.TokenReview
		err := decoder.Decode(&tr)
		if err != nil {
			log.Println("[Error]", err.Error())
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(map[string]interface{}{
				"apiVersion": "authentication.k8s.io/v1beta1",
				"kind":       "TokenReview",
				"status": authentication.TokenReviewStatus{
					Authenticated: false,
				},
			})
			return
		}
		log.Print("receving request")
		// Check User这个 token 就是github 的 token 设置成~/.kube/config里面 k8s 会把这个token 带过来
		ts := oauth2.StaticTokenSource(
			&oauth2.Token{AccessToken: tr.Spec.Token},
		)
		tc := oauth2.NewClient(context.Background(), ts)
		//那这个 token 去 github 验证
		client := github.NewClient(tc)
		user, _, err := client.Users.Get(context.Background(), "")
		if err != nil {
			log.Println("[Error]", err.Error())
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(map[string]interface{}{
				"apiVersion": "authentication.k8s.io/v1beta1",
				"kind":       "TokenReview",
				"status": authentication.TokenReviewStatus{
					Authenticated: false,
				},
			})
			return
		}
		log.Printf("[Success] login as %s", *user.Login)
		w.WriteHeader(http.StatusOK)
		// 拿到 USer 的信息 继续去 k8s 验证
		trs := authentication.TokenReviewStatus{
			Authenticated: true,
			User: authentication.UserInfo{
				Username: *user.Login,
				UID:      *user.Login,
			},
		}
		//  把信息写进 response
		json.NewEncoder(w).Encode(map[string]interface{}{
			"apiVersion": "authentication.k8s.io/v1beta1",
			"kind":       "TokenReview",
			"status":     trs,
		})
	})
	log.Println(http.ListenAndServe(":3000", nil))
}
