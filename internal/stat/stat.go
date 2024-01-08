package stat

import (
	"encoding/json"
	"fmt"
	"net"
	"net/http"

	"github.com/ngaut/log"

	"github.com/mayooot/detect-gpu/pkg/detect"
)

type Option func(s *Stat)

// Stat is used to process http requests.
type Stat struct {
	l net.Listener

	Addr string
	Path string

	Client *detect.Client
}

func NewStat(opts ...Option) *Stat {
	s := &Stat{}
	for _, apply := range opts {
		apply(s)
	}
	return s
}

func WithAddr(Addr string) Option {
	return func(s *Stat) {
		s.Addr = Addr
	}
}

func WithPath(path string) Option {
	return func(s *Stat) {
		s.Path = path
	}
}

func WithClient(client *detect.Client) Option {
	return func(s *Stat) {
		s.Client = client
	}
}

func (st *Stat) Run() {
	var err error

	st.l, err = net.Listen("tcp", st.Addr)
	if err != nil {
		log.Errorf("detect server start failed, err: %v", err)
		return
	}
	log.Infof("detect server start success, listen on %s", st.Addr)
	log.Infof("detect gpu timeout: %d ms", st.Client.Timeout.Milliseconds())
	log.Info("ROUTES:")
	log.Infof("GET		-->		/api/v1/detect/gpu")

	srv := http.Server{}
	mux := http.NewServeMux()
	mux.HandleFunc(st.Path, func(w http.ResponseWriter, req *http.Request) {
		if req.Method != http.MethodGet {
			log.Errorf("method not allowed, method: %s", req.Method)
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
			return
		}

		gpuInfos, err := st.Client.DetectGpu()
		if err != nil {
			err = fmt.Errorf("detect.DetectGpu failed: %w", err)
			log.Errorf(err.Error())
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		resp, _ := json.Marshal(gpuInfos)
		w.WriteHeader(http.StatusOK)
		w.Write(resp)
		log.Infof("detect gpu success, gpu num: %d", len(gpuInfos))
	})

	srv.Handler = mux
	srv.Serve(st.l)
}
