package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		name, _ := os.Hostname()
		// 打印主机名
		log.Printf("%s ping", name)
		fmt.Fprintln(w, "Pong")
	})
	http.HandleFunc("/service", func(w http.ResponseWriter, r *http.Request) {
		// 发出一个 HTTP 请求给服务名为 k8s-combat-service 的 Service
		// 通过 Kubernetes 的 DNS 机制解析服务名，找到相应的 Service（ClusterIP）
		// 然后请求被转发给某个后端的 pod（不是当前 pod）
		// Service 负载均衡策略	kube-proxy 实现，基于 iptables/ipvs，随机、轮询、本地优先
		resp, err := http.Get("http://k8s-combat-service:8081/ping")
		if err != nil {
			log.Println(err)
			fmt.Fprint(w, err)
			return
		}
		fmt.Fprint(w, resp.Status)
	})
	http.ListenAndServe(":8081", nil)
}
