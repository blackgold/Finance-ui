package handle

import (
	"log"
	"net/http"
)

func get404() string {
	str404 := `
	 _  _    ___  _  _
	| || |  / _ \| || |
	| || |_| | | | || |_
	|__   _| | | |__   _|
           | | | |_| |  | |
	   |_|  \___/   |_|`
	return str404
}


func ApiSma200(w http.ResponseWriter, r *http.Request) {
	log.Println(r.RemoteAddr, r.Method, r.URL)
	data := `[
	   {"Stock":"AAPL","Status":"http://localhost:8080/images/up.png","Url":"http://finance.yahoo.com/echarts?s=AAPL+Interactive#{\"showArea\":false,\"bollingerUpperColor\":\"#9900ff\",\"bollingerLowerColor\":\"#ff00ff\",\"showSma\":true,\"smaColors\":\"#cc0000\",\"smaPeriods\":\"100\",\"smaWidths\":\"1\",\"smaGhosting\":\"0\",\"macdSignalColor\":\"#0000ff\",\"lineType\":\"line\",\"allowChartStacking\":true"},
	   {"Stock":"HPE","Status":"http://localhost:8080/images/down.png","Url":"http://finance.yahoo.com/echarts?s=HPE+Interactive#{\"showArea\":false,\"bollingerUpperColor\":\"#9900ff\",\"bollingerLowerColor\":\"#ff00ff\",\"showSma\":true,\"smaColors\":\"#cc0000\",\"smaPeriods\":\"100\",\"smaWidths\":\"1\",\"smaGhosting\":\"0\",\"macdSignalColor\":\"#0000ff\",\"lineType\":\"line\",\"allowChartStacking\":true"},
	   {"Stock":"HDP","Status":"http://localhost:8080/images/down.png","Url":"http://finance.yahoo.com/echarts?s=HDP+Interactive#{\"showArea\":false,\"bollingerUpperColor\":\"#9900ff\",\"bollingerLowerColor\":\"#ff00ff\",\"showSma\":true,\"smaColors\":\"#cc0000\",\"smaPeriods\":\"100\",\"smaWidths\":\"1\",\"smaGhosting\":\"0\",\"macdSignalColor\":\"#0000ff\",\"lineType\":\"line\",\"allowChartStacking\":true"},
	   {"Stock":"HPQ","Status":"http://localhost:8080/images/down.png","Url":"http://finance.yahoo.com/echarts?s=HPQ+Interactive#{\"showArea\":false,\"bollingerUpperColor\":\"#9900ff\",\"bollingerLowerColor\":\"#ff00ff\",\"showSma\":true,\"smaColors\":\"#cc0000\",\"smaPeriods\":\"100\",\"smaWidths\":\"1\",\"smaGhosting\":\"0\",\"macdSignalColor\":\"#0000ff\",\"lineType\":\"line\",\"allowChartStacking\":true"}]`
	w.Write([]byte(data))
}

func ApiEma50(w http.ResponseWriter, r *http.Request) {
	log.Println(r.RemoteAddr, r.Method, r.URL)
	data := `[
	   {"Stock":"AAPL","Status":"http://localhost:8080/images/up.png","Url":"http://finance.yahoo.com/echarts?s=AAPL+Interactive#{\"showArea\":false,\"bollingerUpperColor\":\"#9900ff\",\"bollingerLowerColor\":\"#ff00ff\",\"showSma\":true,\"smaColors\":\"#cc0000\",\"smaPeriods\":\"100\",\"smaWidths\":\"1\",\"smaGhosting\":\"0\",\"macdSignalColor\":\"#0000ff\",\"lineType\":\"line\",\"allowChartStacking\":true"},
	   {"Stock":"HPE","Status":"http://localhost:8080/images/down.png","Url":"http://finance.yahoo.com/echarts?s=HPE+Interactive#{\"showArea\":false,\"bollingerUpperColor\":\"#9900ff\",\"bollingerLowerColor\":\"#ff00ff\",\"showSma\":true,\"smaColors\":\"#cc0000\",\"smaPeriods\":\"100\",\"smaWidths\":\"1\",\"smaGhosting\":\"0\",\"macdSignalColor\":\"#0000ff\",\"lineType\":\"line\",\"allowChartStacking\":true"},
	   {"Stock":"HDP","Status":"http://localhost:8080/images/down.png","Url":"http://finance.yahoo.com/echarts?s=HDP+Interactive#{\"showArea\":false,\"bollingerUpperColor\":\"#9900ff\",\"bollingerLowerColor\":\"#ff00ff\",\"showSma\":true,\"smaColors\":\"#cc0000\",\"smaPeriods\":\"100\",\"smaWidths\":\"1\",\"smaGhosting\":\"0\",\"macdSignalColor\":\"#0000ff\",\"lineType\":\"line\",\"allowChartStacking\":true"},
	   {"Stock":"HPQ","Status":"http://localhost:8080/images/down.png","Url":"http://finance.yahoo.com/echarts?s=HPQ+Interactive#{\"showArea\":false,\"bollingerUpperColor\":\"#9900ff\",\"bollingerLowerColor\":\"#ff00ff\",\"showSma\":true,\"smaColors\":\"#cc0000\",\"smaPeriods\":\"100\",\"smaWidths\":\"1\",\"smaGhosting\":\"0\",\"macdSignalColor\":\"#0000ff\",\"lineType\":\"line\",\"allowChartStacking\":true"}]`
	w.Write([]byte(data))
}

func Page404(w http.ResponseWriter, r *http.Request) {
	log.Println( r.RemoteAddr, r.Method, r.URL)
	w.Write([]byte(get404()))
}