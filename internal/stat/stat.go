package stat

import (
	"encoding/json"
	"net"
	"net/http"

	"github.com/ngaut/log"

	"github.com/mayooot/detect-gpu/pkg/detect"
)

type Stat struct {
	l net.Listener
}

func (st *Stat) Run(port, pattern string) {
	if len(port) == 0 || len(pattern) == 0 {
		return
	}

	log.Infof("detect server listen on %s", port)

	var err error
	st.l, err = net.Listen("tcp", port)
	if err != nil {
		log.Errorf("listen detect port: %s err: %s", port, err)
		return
	}

	srv := http.Server{}
	mux := http.NewServeMux()
	mux.HandleFunc(pattern, func(w http.ResponseWriter, req *http.Request) {
		if req.Method != http.MethodGet {
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
			return
		}

		gpuInfos, err := detect.DetectGpu()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		resp, _ := json.Marshal(gpuInfos)
		w.WriteHeader(http.StatusOK)
		w.Write(resp)
	})

	srv.Handler = mux
	srv.Serve(st.l)
}
