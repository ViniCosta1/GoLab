package middleware

import (
	"context"
	"log"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/influxdata/influxdb-client-go/v2/api/write"
	"github.com/vinicosta1/golab/db"
)

func InitInfluxDB() gin.HandlerFunc {
	return func(c *gin.Context) {
		db.InfluxDBConn()
		writeAPI := db.Client.WriteAPIBlocking(db.InfluxOrg, db.InfluxBucket)

		// 1. Marca o tempo de início ANTES de chamar a lógica da rota
        startTime := time.Now()

        // 2. Chama o próximo handler na cadeia (a lógica da sua rota)
        c.Next()

        // 3. O código abaixo só executa DEPOIS que a rota terminou
        duration := time.Since(startTime)

        // 4. Cria e envia o ponto de métrica para o InfluxDB
        p := write.NewPoint(
            "api_performance",
            map[string]string{
                "endpoint": c.Request.URL.Path, // Tag dinâmica com o endpoint real!
                "method":   c.Request.Method,   // Tag com o método HTTP (GET, POST, etc)
            },
            map[string]interface{}{
                "response_time_ms": float64(duration.Milliseconds()),
            },
            time.Now(),
        )

        // Envia a métrica (em uma goroutine para não bloquear a resposta)
        err := writeAPI.WritePoint(context.Background(), p)
		go func() {
            if err != nil {
                log.Printf("Erro ao enviar métrica para o InfluxDB: %v", err)
            }
        }()
	}
}