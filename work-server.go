package sanzhizhouComponent

import (
	"errors"
	"fmt"
	//sanzhizhouComponentLib "sanzhizhouComponent/lib"
	sanzhizhouComponentLib "github.com/tangwenru/sanzhizhou-component-go/lib"
)

type WorkServer struct {
}

type WorkServerRandDetail struct {
	Id             int64  `json:"id"`
	Title          string `json:"title"`
	Domain         string `json:"domain"`
	DomainProtocol string `json:"domainProtocol"`
	Port           int64  `json:"port"`
	WorkType       string `json:"workType"`
	Enabled        bool   `json:"enabled"`
}

type WorkServerRandDetailQuery struct {
	WorkType string `json:"workType"`
}

type WorkServerRandDetailResult struct {
	Success bool                 `json:"success"`
	Message string               `json:"message"`
	Data    WorkServerRandDetail `json:"data"`
}

// ----------------------
type WorkServerDetailResult struct {
	Success bool             `json:"success"`
	Message string           `json:"message"`
	Data    WorkServerDetail `json:"data"`
}
type WorkServerDetailQuery struct {
	Id int64 `json:"id"`
}

type WorkServerDetail struct {
	Id                       int64  `json:"id"`
	Title                    string `json:"title"`
	Domain                   string `json:"domain"`
	DomainProtocol           string `json:"domainProtocol"`
	Port                     int64  `json:"port"`
	SshPort                  int64  `json:"sshPort"`
	WorkType                 string `json:"workType"` // '','auto-video','data-agent') COLLATE utf8mb3_unicode_ci NOT NULL,
	Enabled                  bool   `json:"enabled"`
	SelfStart                bool   `json:"selfStart"`
	Introduction             string `json:"introduction"`
	LastGetAutoVideoTaskTime int64  `json:"lastGetAutoVideoTaskTime"`
	ClientServeVersion       string `json:"clientServeVersion"`
	LastHeart                int64  `json:"lastHeart"`
	WorkStatus               string `json:"workStatus"`
	Created                  int64  `json:"created"`
	Updated                  int64  `json:"updated"`
}

type WorkServerUpLastWorkTaskTimeQuery struct {
	WorkServerId int64  `json:"workServerId"`
	UpColName    string `json:"upColName"`
	UpTime       int64  `json:"upTime"`
}

type WorkServerUpLastWorkTaskTimeResult struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

// /////////
type WorkServerListQuery struct {
}

type WorkServerListResult struct {
	Success bool               `json:"success"`
	Message string             `json:"message"`
	Data    []WorkServerDetail `json:"data"`
}

func init() {

}

func (this *WorkServer) RandDetail(userToken string, workType string) (*WorkServerRandDetail, error) {
	query := WorkServerRandDetailQuery{
		WorkType: workType,
	}
	vipListResult := WorkServerRandDetailResult{}

	bytesResult, err := sanzhizhouComponentLib.MainSystem(userToken, "workServer/randDetail", &query, &vipListResult)

	if err != nil {
		fmt.Println("Work Server RandDetail err:", string(bytesResult), err)
	}

	if !vipListResult.Success {
		return nil, errors.New(vipListResult.Message)
	}

	return &vipListResult.Data, nil
}

func (this *WorkServer) Detail(userToken string, id int64) (*WorkServerDetail, error) {
	query := WorkServerDetailQuery{
		Id: id,
	}
	vipListResult := WorkServerDetailResult{}

	if id <= 0 {
		return &WorkServerDetail{}, nil
	}

	bytesResult, err := sanzhizhouComponentLib.MainSystem(userToken, "workServer/detail", &query, &vipListResult)

	if err != nil {
		fmt.Println("Work Server Detail err:", string(bytesResult), err)
	}

	if !vipListResult.Success {
		return &WorkServerDetail{}, errors.New(vipListResult.Message)
	}

	return &vipListResult.Data, nil
}

func (this *WorkServer) List(userToken string) (*[]WorkServerDetail, error) {
	query := WorkServerListQuery{}
	vipListResult := WorkServerListResult{}

	bytesResult, err := sanzhizhouComponentLib.MainSystem(userToken, "workServer/list", &query, &vipListResult)

	if err != nil {
		fmt.Println("work Server List err:", string(bytesResult), err)
	}

	if !vipListResult.Success {
		return &[]WorkServerDetail{}, errors.New(vipListResult.Message)
	}

	return &vipListResult.Data, nil
}

func (this *WorkServer) UpLastWorkTaskTime(
	workServerId int64,
	upColName string,
	upTime int64,
) error {
	query := WorkServerUpLastWorkTaskTimeQuery{
		WorkServerId: workServerId,
		UpColName:    upColName,
		UpTime:       upTime,
	}
	vipListResult := WorkServerUpLastWorkTaskTimeResult{}

	bytesResult, err := sanzhizhouComponentLib.MainSystem("", "workServer/upLastWorkTaskTime", &query, &vipListResult)

	if err != nil {
		fmt.Println("Work Server Detail err:", string(bytesResult), err)
		return err
	}

	if !vipListResult.Success {
		return errors.New("WorkServer UpLastWorkTaskTime:" + vipListResult.Message)
	}

	return nil
}
