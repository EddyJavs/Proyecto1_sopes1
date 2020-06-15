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
	"github.com/shirou/gopsutil/process"
	"os/exec"
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

type StructCpu struct{
	PorcentajeUsado float64
}


func main() {
	fmt.Println("hello world")
	router:= mux.NewRouter()

	router.HandleFunc("/memoria",ramInfo).Methods("GET");
	router.HandleFunc("/procesos",ProcessInfo).Methods("GET");
	router.HandleFunc("/matarTask/{id}",MatarProceso).Methods("GET");
	router.HandleFunc("/cpu",cpuInfo).Methods("GET"); //al meterme a la ruta /memoria ejecuta la funcion ramInfo
	//http.HandleFunc("/memoria",ramInfo); //al meterme a la ruta /memoria ejecuta la funcion ramInfo
	//http.HandleFunc("/procesos",ProcessInfo); //al meterme a la ruta /memoria ejecuta la funcion ramInfo
	//http.ListenAndServe(":3000",nil); 	
	

	log.Fatal(http.ListenAndServe(":3000",router))

    fmt.Println("hello world")
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

func MatarProceso(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)

	id := params["id"]
	//Pid,_ := strconv.Atoi(id)

	fmt.Println("Se eliminara proceso: ", id)

	list_processs,err := process.Processes()
	
	if err != nil {
        log.Println("Processes() Failed, are you using windows?")
        return
    }

    for _ , target := range list_processs{

        idProc := strconv.Itoa(int(target.Pid))
        fmt.Println("Pid_proc: ", idProc)
        if  idProc == id{ 
            fmt.Println("Proceso encontrado: ", id)
            target.Kill()
            break
        }else{
            fmt.Println("No se encontro proceso: ", id)
        }
        
    }

	//process2, err := list_processs.FindProcess(Pid)
	//if err != nil {
      //  http.Error(w, errorjson.Error(), http.StatusInternalServerError)
		//	return
    //}
    //fmt.Println("%s\n",process2)

	mapD := map[string]int{"status": 0}

	sonResponse, errorjson := json.Marshal(mapD)
		if errorjson != nil{
			http.Error(w, errorjson.Error(), http.StatusInternalServerError)
			return
	}
	fmt.Println("------")
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusOK)
	w.Write(sonResponse)

}