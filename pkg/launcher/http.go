package launcher

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/influxdata/influxdb/v2"
	log "github.com/sirupsen/logrus"
	taskmodel "github.com/uvite/gvmdesk/pkg/model"
	"net/http"
	"strconv"
	"time"
)

// GetById Find by id structure
type GetById struct {
	ID uint64 `json:"id" form:"id"`
}

func (m *Launcher) runHTTP(opts *InfluxdOpts, router *gin.Engine) error {

	router.POST("/api/task", func(c *gin.Context) {
		time.Sleep(5 * time.Second)
		c.String(http.StatusOK, "Welcome Gin Server")
	})
	router.GET("/api/task", func(c *gin.Context) {
		//var info GetById
		//err := c.ShouldBindJSON(&info)
		//fmt.Println(info, err)

		idq := c.Query("id")

		id, err := strconv.ParseInt(idq, 10, 64)
		fmt.Println(id)
		a := taskmodel.TaskCreate{
			Flux: `option task = {name: "a task",every: 1h} from(bucket:"test") |> range(start:-1h)`,

			Status: string(taskmodel.TaskActive),
		}
		fmt.Println("[2]", a)

		Org := influxdb.Organization{Name: "genv-org", ID: 1}

		fmt.Printf("[org]%+v", Org)
		task, err := m.kvService.CreateTask(c, taskmodel.TaskCreate{
			Flux:           `option task = {name: "a task",every: 1h} from(bucket:"test") |> range(start:-1h)`,
			OrganizationID: 1,
			OwnerID:        1,
			Status:         string(taskmodel.TaskActive),
		})
		if err != nil {
			fmt.Println(err)
		}

		if err != nil {
			FailWithMessage("获取失败", c)
		}
		OkWithDetailed(task, "获取成功", c)
	})

	srv := &http.Server{
		Addr:    ":8888",
		Handler: router,
	}

	go func() {
		// 服务连接
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()
	return nil
}

//
//// runHTTP configures and launches a listener for incoming HTTP(S) requests.
//// The listener is run in a separate goroutine. If it fails to start up, it
//// will cancel the launcher.
//func (m *Launcher) runHTTP(opts *InfluxdOpts, handler nethttp.Handler, httpLogger *zap.Logger) error {
//	log := m.log.With(zap.String("service", "tcp-listener"))
//	fmt.Println("[33]", opts)
//	httpServer := &nethttp.Server{
//		Addr:              opts.HttpBindAddress,
//		Handler:           handler,
//		ReadHeaderTimeout: opts.HttpReadHeaderTimeout,
//		ReadTimeout:       opts.HttpReadTimeout,
//		WriteTimeout:      opts.HttpWriteTimeout,
//		IdleTimeout:       opts.HttpIdleTimeout,
//		ErrorLog:          zap.NewStdLog(httpLogger),
//	}
//	m.closers = append(m.closers, labeledCloser{
//		label:  "HTTP server",
//		closer: httpServer.Shutdown,
//	})
//
//	ln, err := net.Listen("tcp", opts.HttpBindAddress)
//	if err != nil {
//		log.Error("Failed to set up TCP listener", zap.String("addr", opts.HttpBindAddress), zap.Error(err))
//		return err
//	}
//	if addr, ok := ln.Addr().(*net.TCPAddr); ok {
//		m.httpPort = addr.Port
//	}
//	fmt.Println("[33]", m.httpPort)
//
//	m.wg.Add(1)
//
//	return nil
//}
