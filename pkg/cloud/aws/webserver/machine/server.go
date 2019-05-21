package machine

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/serializer"
	"sigs.k8s.io/cluster-api-provider-aws/pkg/cloud/aws/actuators/machine"
	clusterv1 "sigs.k8s.io/cluster-api/pkg/apis/cluster/v1alpha1"
)

type Server struct {
	*http.ServeMux
	Actuator *machine.Actuator
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

	s := &Server{
		ServeMux: mux,
		Actuator: machine.NewActuator(machine.ActuatorParams{}),
		Decoder:  ud,
	}

	mux.HandleFunc("/delete", s.Delete)
	mux.HandleFunc("/exists", s.Exists)
	mux.HandleFunc("/update", s.Update)
	mux.HandleFunc("/create", s.Create)

	return s, nil
}

type Request struct {
	Cluster []byte
	Machine []byte
}

func (s *Server) Delete(w http.ResponseWriter, r *http.Request) {
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	request := &Request{}
	json.Unmarshal(b, request)

	cluster := &clusterv1.Cluster{}
	o, gvk, err := s.Decoder.Decode(request.Cluster, nil, cluster)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Println("cluster", o, gvk)
	machine := &clusterv1.Machine{}
	o, gvk, err = s.Decoder.Decode(request.Machine, nil, machine)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Println("machine", o, gvk)

	if err := s.Actuator.Delete(cluster, machine); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

type ExistsResponse struct {
	Exists bool
}

func (s *Server) Exists(w http.ResponseWriter, r *http.Request) {
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	request := &Request{}
	json.Unmarshal(b, request)

	cluster := &clusterv1.Cluster{}
	o, gvk, err := s.Decoder.Decode(request.Cluster, nil, cluster)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Println("cluster", o, gvk)
	machine := &clusterv1.Machine{}
	o, gvk, err = s.Decoder.Decode(request.Machine, nil, machine)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Println("machine", o, gvk)

	exists, err := s.Actuator.Exists(cluster, machine)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	response := &ExistsResponse{
		Exists: exists,
	}
	out, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write(out)
}

func (s *Server) Update(w http.ResponseWriter, r *http.Request) {
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	request := &Request{}
	json.Unmarshal(b, request)

	cluster := &clusterv1.Cluster{}
	o, gvk, err := s.Decoder.Decode(request.Cluster, nil, cluster)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Println("cluster", o, gvk)
	machine := &clusterv1.Machine{}
	o, gvk, err = s.Decoder.Decode(request.Machine, nil, machine)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Println("machine", o, gvk)

	if err := s.Actuator.Update(cluster, machine); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
func (s *Server) Create(w http.ResponseWriter, r *http.Request) {
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	request := &Request{}
	json.Unmarshal(b, request)

	cluster := &clusterv1.Cluster{}
	o, gvk, err := s.Decoder.Decode(request.Cluster, nil, cluster)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Println("cluster", o, gvk)
	machine := &clusterv1.Machine{}
	o, gvk, err = s.Decoder.Decode(request.Machine, nil, machine)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Println("machine", o, gvk)

	if err := s.Actuator.Create(cluster, machine); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
