package stat

import (
	"encoding/json"
	"fmt"
	"net"
	"net/http"
	"time"

	"github.com/ngaut/log"

	"github.com/mayooot/detect-gpu/pkg/detect"
)

type Stat struct {
	l net.Listener
}

func (st *Stat) Run(port, pattern string, td time.Duration) {
	if len(port) == 0 || len(pattern) == 0 || td == 0 {
		return
	}

	var err error
	st.l, err = net.Listen("tcp", port)
	if err != nil {
		log.Errorf("detect server start failed, listen port: %s, err: %s", port, err)
		return
	}
	log.Infof("detect server start success, listen on %s", port)

	srv := http.Server{}
	mux := http.NewServeMux()
	mux.HandleFunc(pattern, func(w http.ResponseWriter, req *http.Request) {
		if req.Method != http.MethodGet {
			log.Errorf("method not allowed, method: %s", req.Method)
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
			return
		}

		gpuInfos, err := detect.DetectGpu(td)
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
