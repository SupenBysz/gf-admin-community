package performance

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/glog"
)

// MetricType 指标类型
type MetricType string

const (
	MetricTypeAPI      MetricType = "api"
	MetricTypeDB       MetricType = "database"
	MetricTypeCache    MetricType = "cache"
	MetricTypeExternal MetricType = "external"
)

// PerformanceMetric 性能指标
type PerformanceMetric struct {
	Type        MetricType            `json:"type"`
	Name        string                `json:"name"`
	StartTime   time.Time             `json:"start_time"`
	EndTime     time.Time             `json:"end_time"`
	Duration    time.Duration         `json:"duration"`
	Success     bool                  `json:"success"`
	Error       string                `json:"error,omitempty"`
	RequestID   string                `json:"request_id,omitempty"`
	UserID      int64                 `json:"user_id,omitempty"`
	Extra       map[string]interface{} `json:"extra,omitempty"`
}

// PerformanceMonitor 性能监控器
type PerformanceMonitor struct {
	logger    *glog.Logger
	metrics   []PerformanceMetric
	mutex     sync.RWMutex
	enabled   bool
	threshold time.Duration // 慢查询阈值
}

// NewPerformanceMonitor 创建性能监控器
func NewPerformanceMonitor() *PerformanceMonitor {
	return &PerformanceMonitor{
		logger:    g.Log(),
		metrics:   make([]PerformanceMetric, 0),
		enabled:   true,
		threshold: 1 * time.Second, // 默认1秒阈值
	}
}

// SetEnabled 设置是否启用监控
func (pm *PerformanceMonitor) SetEnabled(enabled bool) {
	pm.enabled = enabled
}

// SetThreshold 设置慢查询阈值
func (pm *PerformanceMonitor) SetThreshold(threshold time.Duration) {
	pm.threshold = threshold
}

// StartTimer 开始计时
func (pm *PerformanceMonitor) StartTimer(ctx context.Context, metricType MetricType, name string, extra ...map[string]interface{}) *Timer {
	if !pm.enabled {
		return &Timer{enabled: false}
	}

	timer := &Timer{
		monitor:   pm,
		metric: PerformanceMetric{
			Type:      metricType,
			Name:      name,
			StartTime: time.Now(),
			Success:   true,
			Extra:     make(map[string]interface{}),
		},
		enabled: true,
	}

	// 从上下文中获取请求信息
	if requestID := ctx.Value("RequestId"); requestID != nil {
		timer.metric.RequestID = fmt.Sprintf("%v", requestID)
	}
	if userID := ctx.Value("UserId"); userID != nil {
		if uid, ok := userID.(int64); ok {
			timer.metric.UserID = uid
		}
	}

	// 合并额外信息
	for _, e := range extra {
		for k, v := range e {
			timer.metric.Extra[k] = v
		}
	}

	return timer
}

// RecordMetric 记录性能指标
func (pm *PerformanceMonitor) RecordMetric(metric PerformanceMetric) {
	if !pm.enabled {
		return
	}

	pm.mutex.Lock()
	pm.metrics = append(pm.metrics, metric)
	pm.mutex.Unlock()

	// 记录日志
	pm.logMetric(metric)
}

// GetMetrics 获取性能指标
func (pm *PerformanceMonitor) GetMetrics() []PerformanceMetric {
	pm.mutex.RLock()
	defer pm.mutex.RUnlock()
	
	// 返回副本
	metrics := make([]PerformanceMetric, len(pm.metrics))
	copy(metrics, pm.metrics)
	return metrics
}

// ClearMetrics 清空性能指标
func (pm *PerformanceMonitor) ClearMetrics() {
	pm.mutex.Lock()
	pm.metrics = pm.metrics[:0]
	pm.mutex.Unlock()
}

// GetSlowQueries 获取慢查询
func (pm *PerformanceMonitor) GetSlowQueries() []PerformanceMetric {
	pm.mutex.RLock()
	defer pm.mutex.RUnlock()
	
	var slowQueries []PerformanceMetric
	for _, metric := range pm.metrics {
		if metric.Duration > pm.threshold {
			slowQueries = append(slowQueries, metric)
		}
	}
	return slowQueries
}

// logMetric 记录性能指标日志
func (pm *PerformanceMonitor) logMetric(metric PerformanceMetric) {
	logData := g.Map{
		"type":       metric.Type,
		"name":       metric.Name,
		"duration":   metric.Duration.String(),
		"success":    metric.Success,
		"request_id": metric.RequestID,
	}

	if metric.UserID != 0 {
		logData["user_id"] = metric.UserID
	}

	if metric.Error != "" {
		logData["error"] = metric.Error
	}

	if len(metric.Extra) > 0 {
		logData["extra"] = metric.Extra
	}

	// 根据性能情况选择日志级别
	if metric.Duration > pm.threshold {
		pm.logger.Warning(context.Background(), "Slow operation detected", logData)
	} else if !metric.Success {
		pm.logger.Error(context.Background(), "Operation failed", logData)
	} else {
		pm.logger.Info(context.Background(), "Operation completed", logData)
	}
}

// Timer 计时器
type Timer struct {
	monitor *PerformanceMonitor
	metric  PerformanceMetric
	enabled bool
}

// Stop 停止计时
func (t *Timer) Stop() {
	if !t.enabled {
		return
	}

	t.metric.EndTime = time.Now()
	t.metric.Duration = t.metric.EndTime.Sub(t.metric.StartTime)
	t.monitor.RecordMetric(t.metric)
}

// StopWithError 停止计时并记录错误
func (t *Timer) StopWithError(err error) {
	if !t.enabled {
		return
	}

	t.metric.Success = false
	if err != nil {
		t.metric.Error = err.Error()
	}
	t.Stop()
}

// AddExtra 添加额外信息
func (t *Timer) AddExtra(key string, value interface{}) {
	if !t.enabled {
		return
	}

	if t.metric.Extra == nil {
		t.metric.Extra = make(map[string]interface{})
	}
	t.metric.Extra[key] = value
}

// 全局性能监控器实例
var globalMonitor = NewPerformanceMonitor()

// StartTimer 全局开始计时函数
func StartTimer(ctx context.Context, metricType MetricType, name string, extra ...map[string]interface{}) *Timer {
	return globalMonitor.StartTimer(ctx, metricType, name, extra...)
}

// SetEnabled 全局设置是否启用监控
func SetEnabled(enabled bool) {
	globalMonitor.SetEnabled(enabled)
}

// SetThreshold 全局设置慢查询阈值
func SetThreshold(threshold time.Duration) {
	globalMonitor.SetThreshold(threshold)
}

// GetMetrics 全局获取性能指标
func GetMetrics() []PerformanceMetric {
	return globalMonitor.GetMetrics()
}

// GetSlowQueries 全局获取慢查询
func GetSlowQueries() []PerformanceMetric {
	return globalMonitor.GetSlowQueries()
}

// ClearMetrics 全局清空性能指标
func ClearMetrics() {
	globalMonitor.ClearMetrics()
}

// APIMiddleware API性能监控中间件
func APIMiddleware(r *ghttp.Request) {
	timer := StartTimer(r.Context(), MetricTypeAPI, fmt.Sprintf("%s %s", r.Method, r.URL.Path), map[string]interface{}{
		"method": r.Method,
		"path":   r.URL.Path,
		"ip":     r.GetClientIp(),
	})

	defer func() {
		if err := r.GetError(); err != nil {
			timer.StopWithError(err)
		} else {
			timer.Stop()
		}
	}()

	r.Middleware.Next()
}

// DBQueryTimer 数据库查询计时器
func DBQueryTimer(ctx context.Context, sql string, args ...interface{}) *Timer {
	return StartTimer(ctx, MetricTypeDB, "db_query", map[string]interface{}{
		"sql":  sql,
		"args": args,
	})
}

// CacheTimer 缓存操作计时器
func CacheTimer(ctx context.Context, operation, key string) *Timer {
	return StartTimer(ctx, MetricTypeCache, operation, map[string]interface{}{
		"key": key,
	})
}

// ExternalAPITimer 外部API调用计时器
func ExternalAPITimer(ctx context.Context, url, method string) *Timer {
	return StartTimer(ctx, MetricTypeExternal, "external_api", map[string]interface{}{
		"url":    url,
		"method": method,
	})
}