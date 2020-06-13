package main
import (
	"fmt"
	"net/http"
	"io/ioutil"
	"strings"
	"strconv"
	"encoding/json"
)

type StructMemoria struct{
	MemoriaTotal int
	MemoriaUtilizada int
	MemoriaLibre int
	PorcentajeUtilizado int
}

func main() {
	http.HandleFunc("/memoria",ramInfo); //al meterme a la ruta /memoria ejecuta la funcion ramInfo
	http.ListenAndServe(":3000",nil); 	
    fmt.Println("hello world")
}

func ramInfo(w http.ResponseWriter, r *http.Request){
	b, err := ioutil.ReadFile("/proc/meminfo");
	if err != nil {
		return;
	}
	str := string(b)
	listaInfo := strings.Split(string(str),"\n");//separo el archivo por saltos de linea
	memoriaTotal := strings.Replace((listaInfo[0])[10:24]," ","",-1)
	memoriaLibre := strings.Replace((listaInfo[1])[10:24]," ","",-1)  
	ramTotalKB, err1 := strconv.Atoi(memoriaTotal);
	ramLibreKB, err2 := strconv.Atoi(memoriaLibre);
	if err1 == nil && err2 ==nil{
		ramTotalMB := ramTotalKB/1024
		ramLibreMB := ramLibreKB/1024
		ramUtilizadaMB := ramTotalMB - ramLibreMB
		ramPorcentajeUtilizado:= (ramUtilizadaMB * 100) /ramTotalMB  
		memResponse := StructMemoria{ramTotalMB,ramUtilizadaMB,ramLibreMB,ramPorcentajeUtilizado}
		jsonResponse, errorjson := json.Marshal(memResponse)
		if errorjson != nil{
			http.Error(w, errorjson.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.WriteHeader(http.StatusOK)
		w.Write(jsonResponse)
	}else{
		return
	}

}