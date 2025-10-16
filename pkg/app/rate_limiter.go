package app

import (
	"sync"
	"time"
)

// RateLimiter 限流器
type RateLimiter struct {
	limit    int           // 限制次数
	duration time.Duration // 时间窗口
	mu       sync.Mutex
	records  map[string][]time.Time
}

// NewRateLimiter 创建新的限流器
func NewRateLimiter(limit int, duration time.Duration) *RateLimiter {
	return &RateLimiter{
		limit:    limit,
		duration: duration,
		records:  make(map[string][]time.Time),
	}
}

// Allow 检查是否允许请求
func (rl *RateLimiter) Allow(key string) bool {
	rl.mu.Lock()
	defer rl.mu.Unlock()

	now := time.Now()
	times := rl.records[key]

	// 过滤掉时间窗口外的记录
	validTimes := make([]time.Time, 0)
	for _, t := range times {
		if now.Sub(t) < rl.duration {
			validTimes = append(validTimes, t)
		}
	}

	// 检查是否超过限制
	if len(validTimes) >= rl.limit {
		return false
	}

	// 记录当前请求
	rl.records[key] = append(validTimes, now)
	return true
}

// Clean 清理过期记录
func (rl *RateLimiter) Clean() {
	rl.mu.Lock()
	defer rl.mu.Unlock()

	now := time.Now()
	for key, times := range rl.records {
		validTimes := make([]time.Time, 0)
		for _, t := range times {
			if now.Sub(t) < rl.duration {
				validTimes = append(validTimes, t)
			}
		}
		if len(validTimes) == 0 {
			delete(rl.records, key)
		} else {
			rl.records[key] = validTimes
		}
	}
}
