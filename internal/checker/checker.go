package checker

import (
	"context"
	"net/http"
	"sync"
	"time"
)

type Checker struct {
	Client      *http.Client
	Concurrency int
	Timeout     time.Duration
}

type Result struct {
	URL    string
	Status int
	OK     bool
	Error  string
	Source string
}

func New(concurrency int, timeout time.Duration) *Checker {
	return &Checker{
		Client: &http.Client{Timeout: timeout},
		Concurrency: concurrency,
		Timeout: timeout,
	}
}

func (c *Checker) Check(ctx context.Context, urls []string) []Result {
	results := make([]Result, len(urls))
	sem := make(chan struct{}, c.Concurrency)
	var wg sync.WaitGroup

	for i, url := range urls {
		wg.Add(1)
		go func(idx int, u string) {
			defer wg.Done()
			sem <- struct{}{}
			defer func() { <-sem }()

			req, _ := http.NewRequestWithContext(ctx, "HEAD", u, nil)
			resp, err := c.Client.Do(req)
			if err != nil {
				results[idx] = Result{URL: u, Error: err.Error()}
				return
			}
			defer resp.Body.Close()
			results[idx] = Result{URL: u, Status: resp.StatusCode, OK: resp.StatusCode < 400}
		}(i, url)
	}
	wg.Wait()
	return results
}
