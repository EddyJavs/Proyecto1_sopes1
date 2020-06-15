package main
import (
	"fmt"
	"net/http"
	"io/ioutil"
	"strings"
	"strconv"
	"log"
	"encoding/json"
	"github.com/gorilla/mux"
)

type StructMemoria struct{
	MemoriaTotal int
	MemoriaUtilizada int
	MemoriaLibre int
	PorcentajeUtilizado int
}


    type Process struct{
        PADRE int32 `json:"PADRE"`
        PID int32 `json:"PID"`
        NOMBRE string `json:"NOMBRE"`
        STADO int32 `json:"STADO"`
        //HIJOS [] Process `json:"hijos"`
    }

func main() {
	fmt.Println("hello world")
	router:= mux.NewRouter()

	router.HandleFunc("/memoria",ramInfo).Methods("GET")
	router.HandleFunc("/procesos",ProcessInfo).Methods("GET")
	//http.HandleFunc("/memoria",ramInfo); //al meterme a la ruta /memoria ejecuta la funcion ramInfo
	//http.HandleFunc("/procesos",ProcessInfo); //al meterme a la ruta /memoria ejecuta la funcion ramInfo
	//http.ListenAndServe(":3000",nil); 	
	

	log.Fatal(http.ListenAndServe(":3000",router))

    fmt.Println("hello world")
}

func ramInfo(w http.ResponseWriter, r *http.Request){
	fmt.Println("MEMORIA")
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

func ProcessInfo(w http.ResponseWriter, req *http.Request){
	log.Println("PROCESOS")
 salida := ""
    var procesos []Process

  file, err := ioutil.ReadFile("/proc/cpu_201503699");
    if err != nil {
        return;
    }
    
    str := string(file)
    listaInfo := strings.Split(string(str),"\n");
   

    for i,v := range listaInfo {
        if i >= 12 && i < len(listaInfo)-2 {
                salida = salida + v + "\n"
            }
        }
  //------------------------------------------------------------

    listaInfo2 := strings.Split(string(salida),"\n");

    fmt.Println(len(listaInfo2))

    for i,value := range listaInfo2{
        raw :=Process{}
        if i < len(listaInfo2)-1 {
            in := []byte(value)
        if err2 := json.Unmarshal(in, &raw); err != nil {
            panic(err2)
        }
        procesos = append(procesos,raw)
        }
        
        fmt.Println(raw)   
    }

    //out, _ := json.Marshal(procesos)
    //println(string(out))
    jsonResponse, errorjson := json.Marshal(procesos)
		if errorjson != nil{
			http.Error(w, errorjson.Error(), http.StatusInternalServerError)
			return
		}

    w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResponse)
    //json.NewEncoder(w).Encode(procesos)

}