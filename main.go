package main
import (
    "context"
    "fmt"
    "github.com/yevishev/restaurant-customer/db"
    "github.com/yevishev/restaurant-customer/handler"
    "log"
    "net"
    "net/http"
    "os"
    "os/signal"
    "syscall"
    "time"
)
func main() {
    addr := ":8080"
    //объявление адреса в локальной сети
    listener, err := net.Listen("tcp", addr)
    if err != nil {
        log.Fatalf("Error occurred: %s", err.Error())
    }

    dbUser, dbPassword, dbName :=
        os.Getenv("POSTGRES_USER"),
        os.Getenv("POSTGRES_PASSWORD"),
        os.Getenv("POSTGRES_DB")
    //создание подключения к бд psql
    database, err := db.Initialize(dbUser, dbPassword, dbName)
    if err != nil {
        log.Fatalf("Could not set up database: %v", err)
    }
    //при завершении сервера закрыть подключение к бд
    defer database.Conn.Close()
    //в структуру Server записывается набор хендлеров
    server := &http.Server{
        Handler: http.HandlerFunc(handler.Serve),
    }
    //для каждого хендлера создается горутина
    go func() {
        server.Serve(listener)
    }()
    
    defer Stop(server)
    log.Printf("Started server on %s", addr)
    //создаю канал с пропускной способностью 1, чтобы перехватывать системные сигналы
    ch := make(chan os.Signal, 1)
    //методом Notify передаю сигналы в созданный канал, меня интересуют только два,
    //поскольку они покрывают большинство сценариев завершения приложения
    //например Ctrl+C - SIGINT
    signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
    //вызовом <-ch я блокирую завершение приложения, поскольку в канал не записываются
    //сигналы SIGINT и SIGTERM
    //как только сигналы будут получены будут отработаны defer`ы
    log.Println(fmt.Sprint(<-ch))
    log.Println("Stopping API server.")
}

func Stop(server *http.Server) {
    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()
    if err := server.Shutdown(ctx); err != nil {
        log.Printf("Could not shut down server correctly: %v\n", err)
        os.Exit(1)
    }
}