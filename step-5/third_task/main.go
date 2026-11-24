package main

import (
	"context"
	"io"
	"net/http"
	"sync"
	"time"
)

type APIResponse struct {
	URL        string // запрошенный URL
	Data       string // тело ответа
	StatusCode int    // код ответа
	Err        error  // ошибка, если возникла
}

func FetchAPI(ctx context.Context, urls []string, timeout time.Duration) []*APIResponse {

	resultResponses := make([]*APIResponse, len(urls))

	client := &http.Client{}

	var wg sync.WaitGroup

	for i, url := range urls {
		wg.Add(1)

		go func(i int, url string) {
			defer wg.Done()

			timeoutCtx, cancel := context.WithTimeout(ctx, timeout)
			defer cancel()

			req, err := http.NewRequestWithContext(timeoutCtx, http.MethodGet, url, nil)
			if err != nil {
				resultResponses[i] = &APIResponse{
					URL:        url,
					Data:       "",
					StatusCode: 0,
					Err:        err,
				}
				return
			}

			resp, err := client.Do(req)
			if err != nil {
				if timeoutCtx.Err() == context.DeadlineExceeded {
					resultResponses[i] = &APIResponse{
						URL:        url,
						Data:       "",
						StatusCode: 0,
						Err:        context.DeadlineExceeded,
					}
				} else {
					resultResponses[i] = &APIResponse{
						URL:        url,
						Data:       "",
						StatusCode: 0,
						Err:        err,
					}
				}
				return
			}

			defer resp.Body.Close()

			data, err := io.ReadAll(resp.Body)
			if err != nil {
				if timeoutCtx.Err() == context.DeadlineExceeded {
					resultResponses[i] = &APIResponse{
						URL:        url,
						Data:       "",
						StatusCode: 0,
						Err:        context.DeadlineExceeded,
					}
				} else {
					resultResponses[i] = &APIResponse{
						URL:        url,
						Data:       "",
						StatusCode: 0,
						Err:        err,
					}
				}
				return
			}

			resultResponses[i] = &APIResponse{
				URL:        url,
				Data:       string(data),
				StatusCode: resp.StatusCode,
				Err:        nil,
			}
		}(i, url)
	}

	wg.Wait()

	return resultResponses
}
