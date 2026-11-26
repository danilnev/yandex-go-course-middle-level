
package main

import (
    "io"
    "log"
    "net/http"
    "time"
)

func StartServer(maxTimeout time.Duration) {
    // 1. Создаём маршрутизатор
    mux := http.NewServeMux()

    // 2. Создаём обработчик
    handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

        // 3. Создаём исходящий запрос с контекстом
        req, err := http.NewRequestWithContext(
            r.Context(),
            http.MethodGet,
            "http://localhost:8081/provideData",
            nil,
        )
        if err != nil {
            http.Error(w, "Bad request", http.StatusBadGateway)
            return
        }

        // 4. Отправляем запрос
        resp, err := http.DefaultClient.Do(req)
        if err != nil {
            if r.Context().Err() != nil {
                return
            }
            http.Error(w, "Upstream unavailable", http.StatusBadGateway)
            return
        }
        defer resp.Body.Close()

        // 5. Читаем данные
        data, err := io.ReadAll(resp.Body)
        if err != nil {
            if r.Context().Err() != nil {
                return
            }
            http.Error(w, "Read error", http.StatusBadGateway)
            return
        }

        // 6. Возвращаем ответ пользователю
        w.WriteHeader(resp.StatusCode)
        w.Write(data)
    })

    // 7. Ограничиваем время выполнения
    mux.Handle(
        "/readSource",
        http.TimeoutHandler(handler, maxTimeout, "Request timeout"),
    )

    // 8. Запускаем сервер
    if err := http.ListenAndServe("localhost:8080", mux); err != nil {
        log.Println(err)
    }
}
