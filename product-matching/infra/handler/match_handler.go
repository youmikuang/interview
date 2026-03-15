package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"product-matching/application"
	"product-matching/domain/model"
)

// MatchRequest 匹配请求体
type MatchRequest struct {
	Phone     string `json:"phone"`
	Name      string `json:"name"`
	Age       int    `json:"age"`
	Gender    string `json:"gender"`
	Region    string `json:"region"`
	HasHouse  bool   `json:"hasHouse"`
	HasCar    bool   `json:"hasCar"`
	HasSocial bool   `json:"hasSocial"`
}

// ProductItem 产品响应项
type ProductItem struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// MatchResponse 匹配成功响应
type MatchResponse struct {
	Products []ProductItem `json:"products"`
}

// ErrorResponse 错误响应
type ErrorResponse struct {
	Error string `json:"error"`
}

// WriteJSON 写入 JSON 响应
func WriteJSON(w http.ResponseWriter, status int, v interface{}) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(v)
}

// NewMatchHandler 返回匹配处理函数
func NewMatchHandler(matchApp *application.ProductMatchingApp) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// 方法校验
		if r.Method != http.MethodPost {
			WriteJSON(w, http.StatusMethodNotAllowed, ErrorResponse{Error: "only POST is allowed"})
			return
		}

		// 提取 channel_id
		channelID := r.URL.Query().Get("channel_id")
		if channelID == "" {
			WriteJSON(w, http.StatusBadRequest, ErrorResponse{Error: "missing channel_id"})
			return
		}

		// 解析请求体
		var req MatchRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			WriteJSON(w, http.StatusBadRequest, ErrorResponse{Error: fmt.Sprintf("invalid request body: %v", err)})
			return
		}

		// 转为 domain model
		user := &model.User{
			Phone:     req.Phone,
			Name:      req.Name,
			Age:       req.Age,
			Gender:    req.Gender,
			Region:    req.Region,
			HasHouse:  req.HasHouse,
			HasCar:    req.HasCar,
			HasSocial: req.HasSocial,
		}

		// 调用匹配
		products, err := matchApp.MatchProducts(user, channelID)
		if err != nil {
			WriteJSON(w, http.StatusInternalServerError, ErrorResponse{Error: err.Error()})
			return
		}

		// 构造响应
		items := make([]ProductItem, 0, len(products))
		for _, p := range products {
			items = append(items, ProductItem{ID: p.ID, Name: p.Name})
		}
		WriteJSON(w, http.StatusOK, MatchResponse{Products: items})
	}
}
