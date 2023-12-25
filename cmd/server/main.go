package main

import (
	"flag"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	api "github.com/PandaGoL/api-project/internal/api/http"
	"github.com/PandaGoL/api-project/internal/database/postgres"
	"github.com/PandaGoL/api-project/internal/services/user"
	"github.com/PandaGoL/api-project/pkg/options"
	"github.com/PandaGoL/api-project/pkg/syslog"
	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
)

var (
	configName    string
	exitSignal    chan bool
	signalChannel chan os.Signal
	apiServer     *api.Server
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Warn(".env file not found")
	}
	flag.StringVar(&configName, "config", "api-project", "configuration file name")
	exitSignal = make(chan bool)
}

func Run() error {
	log.Info("Running App")
	flag.Parse()

	opt, err := options.LoadConfig(configName)
	if err != nil {
		log.Errorf("Unable to load configuration file %s: %s", configName, err)
		return err
	}

	if err := syslog.InitLog(); err != nil {
		log.Errorf("Unable to initialize log system: %s", err)
		return err
	}

	log.Debugf("-------> DNS: %s: \n", opt.DB.DSN())

	db, err := postgres.New(opt.DB)
	if err != nil {
		log.Fatalf("DB error: %s", err)
		return nil
	}
	err = db.Migrations()
	if err != nil {
		log.Fatalf("Migration error: %s", err)
		return nil
	}

	userService := user.NewUserService(db)

	apiServer = api.Init(userService)
	go func(srv *api.Server) {
		if err := srv.Serve(); err != nil && err != http.ErrServerClosed {
			log.WithError(err).Fatal("Unable to server HTTP API")
		} else if err == http.ErrServerClosed {
			log.Infof("HTTP server closed")
		}
	}(apiServer)

	time.Sleep(time.Second * 1)
	log.Infof("HTTP API server started on \"%s\"", options.Get().APIAddr)

	go initSignals(db)

	<-exitSignal

	return nil

}
func main() {
	err := Run()
	if err != nil {
		log.Errorf("Error running app: %s", err)
	}

}

func initSignals(db *postgres.PgStorage) {
	log.Info("Try to initialize signals...")
	signalChannel = make(chan os.Signal)
	signal.Notify(signalChannel, syscall.SIGTERM)
	signal.Notify(signalChannel, syscall.SIGINT)
	signal.Notify(signalChannel, syscall.SIGKILL)

	for {
		select {
		case s := <-signalChannel:
			switch s {
			case syscall.SIGTERM, syscall.SIGINT, syscall.SIGKILL:
				close(signalChannel)
				log.Warnf("We got %s, shutdown application...", s)
				_ = apiServer.Stop()
				db.Close()
				exitSignal <- true
				return
			}
		}
	}
}
