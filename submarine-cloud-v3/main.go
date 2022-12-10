/*
 * Licensed to the Apache Software Foundation (ASF) under one or more
 * contributor license agreements.  See the NOTICE file distributed with
 * this work for additional information regarding copyright ownership.
 * The ASF licenses this file to You under the Apache License, Version 2.0
 * (the "License"); you may not use this file except in compliance with
 * the License.  You may obtain a copy of the License at
 *
 *    http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package main

import (
	"flag"
	"fmt"
	"go.uber.org/zap/zapcore"
	"k8s.io/client-go/tools/clientcmd"
	"os"

	// Import all Kubernetes client auth plugins (e.g. Azure, GCP, OIDC, etc.)
	// to ensure that exec-entrypoint and run can make use of them.
	_ "k8s.io/client-go/plugin/pkg/client/auth"

	"k8s.io/apimachinery/pkg/runtime"
	utilruntime "k8s.io/apimachinery/pkg/util/runtime"
	clientgoscheme "k8s.io/client-go/kubernetes/scheme"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/healthz"
	"sigs.k8s.io/controller-runtime/pkg/log/zap"

	istioscheme "istio.io/client-go/pkg/clientset/versioned/scheme"

	submarineapacheorgv1alpha1 "github.com/apache/submarine/submarine-cloud-v3/api/v1alpha1"
	"github.com/apache/submarine/submarine-cloud-v3/controllers"
	//+kubebuilder:scaffold:imports
)

var (
	// Flags generated by operator-sdk
	metricsAddr          string
	enableLeaderElection bool
	probeAddr            string

	// sigs.k8s.io/controller-runtime provides the --kubeconfig flag
	// Only required if out-of-cluster.
	// If set, will use the kubeconfig file at that location.
	// Otherwise will assume running in cluster and use the cluster provided kubeconfig.
	// kubeconfig string

	// Flags of submarine
	istioEnable             bool
	submarineGateway        string
	seldonIstioEnable       bool
	seldonGateway           string
	clusterType             string
	createPodSecurityPolicy bool
)

// Used for "source" field of events. Appears in the "FROM" column of `kubectl describe`
const controllerAgentName = "submarine-controller"

var (
	scheme   = runtime.NewScheme()
	setupLog = ctrl.Log.WithName("setup")
)

func init() {
	utilruntime.Must(clientgoscheme.AddToScheme(scheme))
	utilruntime.Must(istioscheme.AddToScheme(scheme))

	utilruntime.Must(submarineapacheorgv1alpha1.AddToScheme(scheme))
	//+kubebuilder:scaffold:scheme
}

func main() {
	// Setup flags
	// Flags generated by operator-sdk
	flag.StringVar(&metricsAddr, "metrics-bind-address", ":8080", "The address the metric endpoint binds to.")
	flag.StringVar(&probeAddr, "health-probe-bind-address", ":8081", "The address the probe endpoint binds to.")
	flag.BoolVar(&enableLeaderElection, "leader-elect", false,
		"Enable leader election for controller manager. "+
			"Enabling this will ensure there is only one active controller manager.")
	// Flags of submarine
	flag.BoolVar(&istioEnable, "istioenable", true, "Istio enable")
	flag.StringVar(&submarineGateway, "submarineateway", "", "Submarine gateway, used for server, minio, tensorboard, mlflow and notebook")
	flag.BoolVar(&seldonIstioEnable, "seldonistioenable", true, "Seldon istio enable")
	flag.StringVar(&seldonGateway, "seldongateway", "", "Seldon gateway, used for model serve")
	flag.StringVar(&clusterType, "clustertype", "kubernetes", "K8s cluster type, can be kubernetes or openshift")
	flag.BoolVar(&createPodSecurityPolicy, "createpsp", true, "Specifies whether a PodSecurityPolicy should be created. This configuration enables the database/minio/server to set securityContext.runAsUser")

	opts := zap.Options{
		Development: true,
		// format timestamp with 2006-01-02T15:04:05.000Z0700
		TimeEncoder: zapcore.ISO8601TimeEncoder,
	}
	opts.BindFlags(flag.CommandLine)
	flag.Parse()

	ctrl.SetLogger(zap.New(zap.UseFlagOptions(&opts)))

	// get current namespace
	loadingRules := clientcmd.NewDefaultClientConfigLoadingRules()
	kubeConfig := clientcmd.NewNonInteractiveDeferredLoadingClientConfig(loadingRules, &clientcmd.ConfigOverrides{})
	namespace, _, err := kubeConfig.Namespace()

	// if `seldonGateway` is empty, used ${namespace}/seldon-gateway
	// By default, the operator and seldon-gateway will be under the same namespace when deployed with helm
	if seldonGateway == "" {
		seldonGateway = fmt.Sprintf("%s/seldon-gateway", namespace)
	}
	// if `submarineGateway` is empty, used ${namespace}/seldon-gateway
	// By default, the operator and submarine-gateway will be under the same namespace when deployed with helm
	if submarineGateway == "" {
		submarineGateway = fmt.Sprintf("%s/submarine-gateway", namespace)
	}

	setupLog.Info("Starting submarine operator with ",
		"metrics-bind-address", &metricsAddr,
		"health-probe-bind-address", &probeAddr,
		"leader-elect", &enableLeaderElection,
		"namespace", namespace,
		"istioenable", &istioEnable,
		"submarineateway", &submarineGateway,
		"seldonistioenable", &seldonIstioEnable,
		"seldongateway", &seldonGateway,
		"clustertype", &clusterType,
		"createpsp", &createPodSecurityPolicy,
	)
	mgr, err := ctrl.NewManager(ctrl.GetConfigOrDie(), ctrl.Options{
		Scheme:                  scheme,
		MetricsBindAddress:      metricsAddr,
		Port:                    9443,
		HealthProbeBindAddress:  probeAddr,
		LeaderElectionNamespace: namespace,
		LeaderElection:          enableLeaderElection,
		LeaderElectionID:        "5d52732c.submarine.apache.org",
	})
	if err != nil {
		setupLog.Error(err, "unable to start manager")
		os.Exit(1)
	}

	if err = (&controllers.SubmarineReconciler{
		Client:                  mgr.GetClient(),
		Scheme:                  mgr.GetScheme(),
		Recorder:                mgr.GetEventRecorderFor(controllerAgentName),
		Log:                     ctrl.Log.WithName(controllerAgentName),
		Namespace:               namespace,
		IstioEnable:             istioEnable,
		SubmarineGateway:        submarineGateway,
		SeldonIstioEnable:       seldonIstioEnable,
		SeldonGateway:           seldonGateway,
		ClusterType:             clusterType,
		CreatePodSecurityPolicy: createPodSecurityPolicy,
	}).SetupWithManager(mgr); err != nil {
		setupLog.Error(err, "unable to create controller", "controller", "Submarine")
		os.Exit(1)
	}
	//+kubebuilder:scaffold:builder

	if err := mgr.AddHealthzCheck("healthz", healthz.Ping); err != nil {
		setupLog.Error(err, "unable to set up health check")
		os.Exit(1)
	}
	if err := mgr.AddReadyzCheck("readyz", healthz.Ping); err != nil {
		setupLog.Error(err, "unable to set up ready check")
		os.Exit(1)
	}

	setupLog.Info("starting manager")
	if err := mgr.Start(ctrl.SetupSignalHandler()); err != nil {
		setupLog.Error(err, "problem running manager")
		os.Exit(1)
	}
}
