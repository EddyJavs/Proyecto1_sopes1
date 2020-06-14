package main
import (
	"fmt"
	"net/http"
	"io/ioutil"
	"strings"
	"strconv"
	"encoding/json"
	"os/exec"
)

type StructMemoria struct{
	MemoriaTotal int
	MemoriaUtilizada int
	MemoriaLibre int
	PorcentajeUtilizado int
}

type StructCpu struct{
	PorcentajeUsado float64
}

func main() {
	fmt.Println("hello world")
	http.HandleFunc("/memoria",ramInfo); //al meterme a la ruta /memoria ejecuta la funcion ramInfo
	//http.ListenAndServe(":3000",nil); 	
    http.HandleFunc("/cpu",cpuInfo); //al meterme a la ruta /memoria ejecuta la funcion ramInfo
	http.ListenAndServe(":3000",nil);
}

func ramInfo(w http.ResponseWriter, r *http.Request){
	fmt.Println("Ram info")
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

	func cpuInfo(w http.ResponseWriter, r *http.Request){
	//fmt.Println("cpu Info")
	//app := "top"
	//args := []string{ "-bn2", "| fgrep \"Cpu(s)\"", "| tail -1"}

	cmd := exec.Command("/bin/sh", "-c","top -bn2 | fgrep \"Cpu(s)\" | tail -1")
	stdout, err := cmd.Output()

if err != nil {
println(err.Error())
return
}

//print(string(stdout))

	listaInfo := strings.Split(string(stdout),",");//separo el archivo por saltos de linea
	PorcentajeLibre := strings.Replace((listaInfo[3])[1:5]," ","",-1)
	//fmt.Println(PorcentajeLibre)
	PorLibre, err1 := strconv.ParseFloat( PorcentajeLibre, 32);
	if err1 == nil{
		//fmt.Println(PorLibre)
		PorOcupado := (100-PorLibre)
		//fmt.Println(PorOcupado)
		memResponse := StructCpu{PorOcupado}
		jsonResponse, errorjson := json.Marshal(memResponse)
		if errorjson != nil{
			http.Error(w, errorjson.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.WriteHeader(http.StatusOK)
		w.Write(jsonResponse)
	}
	


}