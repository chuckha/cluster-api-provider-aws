package cluster

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/serializer"
	"sigs.k8s.io/cluster-api-provider-aws/pkg/cloud/aws/actuators/cluster"
	clusterv1 "sigs.k8s.io/cluster-api/pkg/apis/cluster/v1alpha1"
	"sigs.k8s.io/cluster-api/pkg/client/clientset_generated/clientset"
	"sigs.k8s.io/controller-runtime/pkg/client/config"
)

type Server struct {
	*http.ServeMux
	Actuator *cluster.Actuator
	Decoder  runtime.Decoder
}

func NewServer() (*Server, error) {
	scheme := runtime.NewScheme()
	if err := clusterv1.AddToScheme(scheme); err != nil {
		return nil, err
	}

	cf := serializer.NewCodecFactory(scheme)
	ud := cf.UniversalDecoder()
	mux := http.NewServeMux()

	cfg := config.GetConfigOrDie()
	cs, err := clientset.NewForConfig(cfg)
	if err != nil {
		return nil, err
	}

	s := &Server{
		ServeMux: mux,
		Actuator: cluster.NewActuator(cluster.ActuatorParams{
			Client:         cs.ClusterV1alpha1(),
			LoggingContext: "[cluster-actuator]",
		}),
		Decoder: ud,
	}

	mux.HandleFunc("/reconcile", s.Reconcile)
	return s, nil
}

func (s *Server) Reconcile(w http.ResponseWriter, r *http.Request) {
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	c := &clusterv1.Cluster{}
	o, gvk, err := s.Decoder.Decode(b, nil, c)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Println(o, gvk)
	if err := s.Actuator.Reconcile(c); err != nil {
		fmt.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
