package main

import (
	"log"
	"net/http"

	cu "github.com/achelovekov/collectorutils"
)

type PostReqHandler struct {
	*cu.PostReqHandler
}

func worker(src map[string]interface{}, ESClient cu.ESClient, ESIndex string, path cu.Path, mode int, filter cu.Filter, enrich cu.Enrich) {
	var pathIndex int
	header := make(map[string]interface{})
	buf := make([]map[string]interface{}, 0)
	pathPassed := make([]string, 0)
	var keysLeftFromPrevLayer bool

	/* 	cu.PrettyPrint(src) */

	/* 	fmt.Println(path.PathData, "len:", len(path.PathData))
	 */cu.FlattenMap(src, path, pathIndex, pathPassed, mode, header, &buf, filter, enrich, keysLeftFromPrevLayer)
	cu.ESPush(ESClient, ESIndex, buf)
}

func (prh *PostReqHandler) SysBgp(w http.ResponseWriter, httpRequest *http.Request) {
	src := cu.GetHttpBody(httpRequest)

	for i := range prh.ChunksProcessingPaths["sys/bgp"] {
		src := cu.CopyMap(src)
		go worker(src, prh.ESClient, prh.Config.ESIndex, prh.ChunksProcessingPaths["sys/bgp"][i], cu.Cadence, prh.Filter, prh.Enrich)
	}
}

func (prh *PostReqHandler) SysOspf(w http.ResponseWriter, httpRequest *http.Request) {
	src := cu.GetHttpBody(httpRequest)

	for i := range prh.ChunksProcessingPaths["sys/ospf"] {
		src := cu.CopyMap(src)
		go worker(src, prh.ESClient, prh.Config.ESIndex, prh.ChunksProcessingPaths["sys/ospf"][i], cu.Cadence, prh.Filter, prh.Enrich)
	}
}

func (prh *PostReqHandler) RIBHandler(w http.ResponseWriter, httpRequest *http.Request) {
	src := cu.GetHttpBody(httpRequest)

	for i := range prh.ChunksProcessingPaths["rib"] {
		src := cu.CopyMap(src)
		go worker(src, prh.ESClient, prh.Config.ESIndex, prh.ChunksProcessingPaths["rib"][i], cu.Native, prh.Filter, prh.Enrich)
	}
}

func (prh *PostReqHandler) MacAllHandler(w http.ResponseWriter, httpRequest *http.Request) {
	src := cu.GetHttpBody(httpRequest)

	for i := range prh.ChunksProcessingPaths["mac-all"] {
		src := cu.CopyMap(src)
		go worker(src, prh.ESClient, prh.Config.ESIndex, prh.ChunksProcessingPaths["mac-all"][i], cu.Native, prh.Filter, prh.Enrich)
	}
}

func (prh *PostReqHandler) AdjacencyHandler(w http.ResponseWriter, httpRequest *http.Request) {
	src := cu.GetHttpBody(httpRequest)

	for i := range prh.ChunksProcessingPaths["adjacency"] {
		src := cu.CopyMap(src)
		go worker(src, prh.ESClient, prh.Config.ESIndex, prh.ChunksProcessingPaths["adjacency"][i], cu.Native, prh.Filter, prh.Enrich)
	}
}

func (prh *PostReqHandler) EventHandler(w http.ResponseWriter, httpRequest *http.Request) {
	src := cu.GetHttpBody(httpRequest)

	for i := range prh.ChunksProcessingPaths["event"] {
		src := cu.CopyMap(src)
		go worker(src, prh.ESClient, prh.Config.ESIndex, prh.ChunksProcessingPaths["event"][i], cu.Event, prh.Filter, prh.Enrich)
	}
}

func (prh *PostReqHandler) VxlanSysEps(w http.ResponseWriter, httpRequest *http.Request) {
	src := cu.GetHttpBody(httpRequest)

	for i := range prh.ChunksProcessingPaths["vxlan:sys/eps"] {
		src := cu.CopyMap(src)
		go worker(src, prh.ESClient, prh.Config.ESIndex, prh.ChunksProcessingPaths["vxlan:sys/eps"][i], cu.Cadence, prh.Filter, prh.Enrich)
	}
}

func (prh *PostReqHandler) VxlanSysBD(w http.ResponseWriter, httpRequest *http.Request) {
	src := cu.GetHttpBody(httpRequest)

	for i := range prh.ChunksProcessingPaths["vxlan:sys/bd"] {
		src := cu.CopyMap(src)
		go worker(src, prh.ESClient, prh.Config.ESIndex, prh.ChunksProcessingPaths["vxlan:sys/bd"][i], cu.Cadence, prh.Filter, prh.Enrich)
	}
}

func (prh *PostReqHandler) SysIntfHandler(w http.ResponseWriter, httpRequest *http.Request) {
	src := cu.GetHttpBody(httpRequest)

	for i := range prh.ChunksProcessingPaths["interface:sys/intf"] {
		src := cu.CopyMap(src)
		go worker(src, prh.ESClient, prh.Config.ESIndex, prh.ChunksProcessingPaths["interface:sys/intf"][i], cu.Cadence, prh.Filter, prh.Enrich)
	}
}

func (prh *PostReqHandler) SysChHandler(w http.ResponseWriter, httpRequest *http.Request) {
	src := cu.GetHttpBody(httpRequest)

	for i := range prh.ChunksProcessingPaths["environment:sys/ch"] {
		src := cu.CopyMap(src)
		go worker(src, prh.ESClient, prh.Config.ESIndex, prh.ChunksProcessingPaths["environment:sys/ch"][i], cu.Cadence, prh.Filter, prh.Enrich)
	}
}

func (prh *PostReqHandler) sysProcHandler(w http.ResponseWriter, httpRequest *http.Request) {
	src := cu.GetHttpBody(httpRequest)

	for i := range prh.ChunksProcessingPaths["resources:sys/proc"] {
		src := cu.CopyMap(src)
		go worker(src, prh.ESClient, prh.Config.ESIndex, prh.ChunksProcessingPaths["resources:sys/proc"][i], cu.Cadence, prh.Filter, prh.Enrich)
	}
}

func (prh *PostReqHandler) sysProcSysHandler(w http.ResponseWriter, httpRequest *http.Request) {
	src := cu.GetHttpBody(httpRequest)

	for i := range prh.ChunksProcessingPaths["resources:sys/procsys"] {
		src := cu.CopyMap(src)
		go worker(src, prh.ESClient, prh.Config.ESIndex, prh.ChunksProcessingPaths["resources:sys/procsys"][i], cu.Cadence, prh.Filter, prh.Enrich)
	}
}

func main() {

	Config, Filter, Enrich := cu.Initialize("config.json")
	chunksProcessingPaths := cu.LoadChunksProcessingPaths(Config.ChunksDefinitionFile)

	ESClient, error := cu.ESConnect(Config.ESHost, Config.ESPort)
	if error != nil {
		log.Fatalf("error: %s", error)
	}

	postReqHandler := &PostReqHandler{&cu.PostReqHandler{ESClient: ESClient, Filter: Filter, Enrich: Enrich, Config: Config, ChunksProcessingPaths: chunksProcessingPaths, Mode: 2}}

	http.HandleFunc("/network/sys/bgp", postReqHandler.SysBgp)
	http.HandleFunc("/network/sys/ospf", postReqHandler.SysOspf)
	http.HandleFunc("/network/rib", postReqHandler.RIBHandler)
	http.HandleFunc("/network/mac-all", postReqHandler.MacAllHandler)
	http.HandleFunc("/network/adjacency", postReqHandler.AdjacencyHandler)
	http.HandleFunc("/network/EVENT-LIST", postReqHandler.EventHandler)
	http.HandleFunc("/network/vxlan:sys/eps", postReqHandler.VxlanSysEps)
	http.HandleFunc("/network/vxlan:sys/bd", postReqHandler.VxlanSysBD)
	http.HandleFunc("/network/interface:sys/intf", postReqHandler.SysIntfHandler)
	http.HandleFunc("/network/environment:sys/ch", postReqHandler.SysChHandler)
	http.HandleFunc("/network/resources:sys/proc", postReqHandler.sysProcHandler)
	http.HandleFunc("/network/resources:sys/procsys", postReqHandler.sysProcSysHandler)

	http.ListenAndServe(":11000", nil)
}
