package cmd

import (
	"github.com/sirupsen/logrus"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/clever-telemetry/miniloops/pkg"
	"github.com/clever-telemetry/miniloops/pkg/client"
	"github.com/clever-telemetry/miniloops/pkg/controllers"
	"github.com/clever-telemetry/miniloops/pkg/logr"
)

func init() {
	RootCmd.PersistentFlags().Int32P("log-level", "l", 4, "set logging level between 0 and 5")

	if err := viper.BindPFlags(RootCmd.PersistentFlags()); nil != err {
		log.WithError(err).Fatal("could not bind flags")
	}

	ctrl.SetLogger(logr.Logrus())
}

var RootCmd = &cobra.Command{
	Use: "miniloops",
	Run: rootCmdFn,
}

func rootCmdFn(cmd *cobra.Command, args []string) {
	if viper.GetInt32("log-level") > 4 {
		logrus.SetLevel(logrus.DebugLevel)
	}

	restConfig, err := ctrl.GetConfig()
	if err != nil {
		log.WithError(err).Fatal("cannot find kube configuration")
	}

	client.Set(restConfig)

	restConfig.QPS = float32(viper.GetFloat64("kubernetes.client.qps"))
	restConfig.Burst = viper.GetInt("kubernetes.client.burst")
	logrus.Infof("kubernetes client config (QPS: %f, Burst: %d)", restConfig.QPS, restConfig.Burst)

	mgr, err := ctrl.NewManager(restConfig, ctrl.Options{
		Scheme:                  pkg.Scheme,
		Logger:                  logr.Logrus(),
		MetricsBindAddress:      viper.GetString("manager.metrics.address"),
		LeaderElection:          false,
		LeaderElectionNamespace: "kube-system",
		LeaderElectionID:        "miniloops-leader",
		Host:                    viper.GetString("manager.address"),
		Port:                    viper.GetInt("manager.port"),
	})
	if err != nil {
		log.WithError(err).Fatal("cannot start manager")
	}

	client.EventRecorderFor = mgr.GetEventRecorderFor

	// Register controllers
	if err := controllers.Loops(mgr); err != nil {
		log.WithError(err).Fatal("cannot setup loops controller")
	}

	log.Infof("start manager")
	if err = mgr.Start(ctrl.SetupSignalHandler()); err != nil {
		log.WithError(err).Fatal("manager error")
	}
}
