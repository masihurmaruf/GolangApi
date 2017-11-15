package main

type HttpResp struct {
	Status     	 int    	`json:"status"`
	Description  string 	`json:description`
	Body       	 string   	`json:body`
}
