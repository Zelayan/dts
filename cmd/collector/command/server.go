package command

import (
	"github.com/Zelayan/dts/api/server/router"
	"github.com/Zelayan/dts/api/server/router/collector"
	pb "github.com/Zelayan/dts/proto-gen/v1/dts"
	"log"
	"net"
	"net/http"
	"os/signal"
	"syscall"

	"fmt"
	"github.com/Zelayan/dts/cmd/collector/options"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
	"k8s.io/klog/v2"
	"os"
)

func NewServerCommand() *cobra.Command {
	opts, err := options.NewOptions()
	if err != nil {
		klog.Fatalf("unable to initialize command options: %v", err)
	}

	cmd := &cobra.Command{
		Use:  "dts-collector",
		Long: "",
		Run: func(cmd *cobra.Command, args []string) {
			if err = opts.Complete(); err != nil {
				fmt.Fprintf(os.Stderr, "%v\n", err)
				os.Exit(1)
			}
			if err = opts.Validate(); err != nil {
				fmt.Fprintf(os.Stderr, "%v\n", err)
				os.Exit(1)
			}
			if err = Run(opts); err != nil {
				fmt.Fprintf(os.Stderr, "%v\n", err)
				os.Exit(1)
			}
		},
	}
	return cmd

}

func Run(opts *options.Options) error {
	listen, err := net.Listen("tcp", opts.ComponentConfig.Default.Collector.Listen)
	if err != nil {
		return err
	}
	// TODO: 抽离出来注册路由
	server := grpc.NewServer()
	pb.RegisterCollectorServiceServer(server, collector.NewCollectorServer(opts))

	router.InstallRouter(opts)

	srv := &http.Server{
		Addr:    opts.ComponentConfig.Default.Query.Listen,
		Handler: opts.HttpEngine,
	}
	go func() {
		defer func() {
			err := recover()
			if err != nil {
				fmt.Printf("xx")
			}
		}()
		err = server.Serve(listen)
		if err != nil {
			log.Fatalf("failed to listen collector server: %v", err)
		} else {
			klog.Infof("collect listen: %d", opts.ComponentConfig.Default.Collector.Listen)
		}
	}()

	go func() {
		err := srv.ListenAndServe()
		if err != nil {
			klog.Fatalf("faild to listen query server: %v", err)
		}
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	klog.Infof("shutting collector server down...")
	// TODO 优雅的启停
	return nil
}
