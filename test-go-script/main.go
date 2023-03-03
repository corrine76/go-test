package main

import (
	"crypto/tls"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	gogit "github.com/go-git/go-git/v5"
	gogithttp "github.com/go-git/go-git/v5/plumbing/transport/http"
	gitlab "github.com/xanzy/go-gitlab"
)

func exists(path string) bool {
	_, err := os.Stat(path) //os.Stat获取文件信息
	if err != nil {
		if os.IsExist(err) {
			return true
		}
		return false
	}
	return true
}

func main() {
	transCfg := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true}, // ignore expired SSL certificates
	}
	httpClient := &http.Client{
		Transport: transCfg,
	}
	git, err := gitlab.NewClient("LmiZY_m-bADsBecyMtKs", gitlab.WithBaseURL("https://git.dustess.com/api/v4"), gitlab.WithHTTPClient(httpClient))
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

	var page, pageSize int = 1, 1000
	match := "/biz/"
	dir := "/tmp/dustess/biz"
	exist := exists(dir)
	if !exist {
		os.MkdirAll(dir, os.ModePerm)
	}

	for {
		projects, resp, err := git.Projects.ListProjects(&gitlab.ListProjectsOptions{
			Simple:      gitlab.Bool(true),
			ListOptions: gitlab.ListOptions{Page: page, PerPage: pageSize},
		})
		if err != nil {
			panic(err)
		}
		fmt.Printf("totalItem=%d, TotalPages=%d, ItemsPerPage=%d CurrentPage=%d\n", resp.TotalItems, resp.TotalPages, resp.ItemsPerPage, resp.CurrentPage)
		if len(projects) == 0 {
			break
		}
		for _, project := range projects {
			if !strings.Contains(project.HTTPURLToRepo, match) {
				continue
			}
			fmt.Printf("正在处理 %s\n", project.HTTPURLToRepo)
			projectpath := fmt.Sprintf("%s/%s", dir, project.Name)
			exist := exists(projectpath)
			if !exist {
				os.MkdirAll(projectpath, os.ModePerm)
				fmt.Printf("开始克隆 %s\n", project.HTTPURLToRepo)
				_, err := gogit.PlainClone(projectpath, false, &gogit.CloneOptions{
					Auth: &gogithttp.BasicAuth{
						Username: "liuyao",
						Password: "LmiZY_m-bADsBecyMtKs",
					},
					URL: project.HTTPURLToRepo,
				})
				if err != nil {
					fmt.Printf("git clone failed project=%s addr=%s err=%v\n", project.Name, project.HTTPURLToRepo, err)
				} else {
					fmt.Printf("完成克隆 %s\n", project.HTTPURLToRepo)
				}
			} else {
				fmt.Printf("开始更新 %s\n", project.HTTPURLToRepo)
				r, err := gogit.PlainOpen(projectpath)
				if err != nil {
					fmt.Printf("git pull failed project=%s addr=%s err=%v\n", project.Name, project.HTTPURLToRepo, err)
					continue
				}
				w, err := r.Worktree()
				if err != nil {
					fmt.Printf("git pull failed project=%s addr=%s err=%v\n", project.Name, project.HTTPURLToRepo, err)
					continue
				}
				err = w.Pull(&gogit.PullOptions{
					Auth: &gogithttp.BasicAuth{
						Username: "liuyao",
						Password: "LmiZY_m-bADsBecyMtKs",
					},
				})
				if err != nil {
					fmt.Printf("git pull failed project=%s addr=%s err=%v\n", project.Name, project.HTTPURLToRepo, err)
					continue
				}
				fmt.Printf("完成更新 %s\n", project.HTTPURLToRepo)
			}
		}
		page++
	}
}
