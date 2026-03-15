package handler

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"runtime"
	"testing"

	"product-matching/application"
	domain_service "product-matching/domain/service"
	infra_repository "product-matching/infra/repository"
	infra_service "product-matching/infra/service"
)

func TestMain(m *testing.M) {
	// 切换到项目根目录，使 config/config.json 的相对路径可用
	_, filename, _, _ := runtime.Caller(0)
	projectRoot := filepath.Join(filepath.Dir(filename), "..", "..")
	os.Chdir(projectRoot)
	os.Exit(m.Run())
}

func setupHandler() http.HandlerFunc {
	remoteChecker := infra_service.NewRemoteAPI("https://mock-api.com/check-md5")
	productRepo := &infra_repository.MockProductRepo{}
	channelRepo := &infra_repository.MockChannelRepo{}
	filterService := domain_service.NewProductFilter(remoteChecker, productRepo, channelRepo)
	matchApp := application.NewProductMatchingApp(filterService)
	return NewMatchHandler(matchApp)
}

func TestMatchHandler(t *testing.T) {
	handler := setupHandler()

	tests := []struct {
		name         string
		method       string
		url          string
		body         interface{}
		wantStatus   int
		wantProducts int // -1 means check error instead
		wantError    bool
	}{
		{
			name:   "用户1（张三）匹配成功",
			method: http.MethodPost,
			url:    "/match?channel_id=C001",
			body: MatchRequest{
				Phone: "123456", Name: "张三", Age: 30, Gender: "男",
				Region: "北京", HasHouse: true, HasCar: true, HasSocial: true,
			},
			wantStatus:   http.StatusOK,
			wantProducts: 2,
		},
		{
			name:   "用户2（李四）不符合条件",
			method: http.MethodPost,
			url:    "/match?channel_id=C001",
			body: MatchRequest{
				Phone: "654321", Name: "李四", Age: 18, Gender: "女",
				Region: "上海", HasHouse: false, HasCar: false, HasSocial: true,
			},
			wantStatus:   http.StatusOK,
			wantProducts: 0,
		},
		{
			name:       "缺少 channel_id",
			method:     http.MethodPost,
			url:        "/match",
			body:       MatchRequest{Phone: "123456"},
			wantStatus: http.StatusBadRequest,
			wantError:  true,
		},
		{
			name:       "错误的 HTTP 方法",
			method:     http.MethodGet,
			url:        "/match?channel_id=C001",
			body:       nil,
			wantStatus: http.StatusMethodNotAllowed,
			wantError:  true,
		},
		{
			name:   "手机号为空",
			method: http.MethodPost,
			url:    "/match?channel_id=C001",
			body: MatchRequest{
				Phone: "", Name: "无名", Age: 25, Gender: "男",
				Region: "北京", HasHouse: true, HasCar: true, HasSocial: true,
			},
			wantStatus: http.StatusInternalServerError,
			wantError:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var reqBody *bytes.Buffer
			if tt.body != nil {
				b, err := json.Marshal(tt.body)
				if err != nil {
					t.Fatalf("failed to marshal request body: %v", err)
				}
				reqBody = bytes.NewBuffer(b)
			} else {
				reqBody = &bytes.Buffer{}
			}

			req := httptest.NewRequest(tt.method, tt.url, reqBody)
			req.Header.Set("Content-Type", "application/json")
			rec := httptest.NewRecorder()

			handler.ServeHTTP(rec, req)

			if rec.Code != tt.wantStatus {
				t.Errorf("status = %d, want %d", rec.Code, tt.wantStatus)
			}

			if tt.wantError {
				var errResp ErrorResponse
				if err := json.NewDecoder(rec.Body).Decode(&errResp); err != nil {
					t.Fatalf("failed to decode error response: %v", err)
				}
				if errResp.Error == "" {
					t.Error("expected non-empty error message")
				}
			} else {
				var resp MatchResponse
				if err := json.NewDecoder(rec.Body).Decode(&resp); err != nil {
					t.Fatalf("failed to decode response: %v", err)
				}
				if len(resp.Products) != tt.wantProducts {
					t.Errorf("products count = %d, want %d", len(resp.Products), tt.wantProducts)
				}
			}
		})
	}
}
