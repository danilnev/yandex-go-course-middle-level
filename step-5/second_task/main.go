package main

import (
	"context"
	"io"
	"net/http"
	"time"
)

type APIResponse struct {
	Data string
	StatusCode int 
}

func fetchAPI(ctx context.Context, url string, timeout time.Duration) (*APIResponse, error) {

	timeoutCtx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	req, err := http.NewRequestWithContext(timeoutCtx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err 
	}

	client := &http.Client{}

	resp, err := client.Do(req)

	

	if err != nil {
		if timeoutCtx.Err() == context.DeadlineExceeded {
			return nil, context.DeadlineExceeded
		}

		return nil, err 
	}

	defer resp.Body.Close()
	
	body, err := io.ReadAll(resp.Body)

	if err != nil {
		if timeoutCtx.Err() == context.DeadlineExceeded {
			return nil, context.DeadlineExceeded
		}

		return nil, err
	}

	result := &APIResponse{
		Data: string(body),
		StatusCode: resp.StatusCode,
	}

	return result, nil
}
