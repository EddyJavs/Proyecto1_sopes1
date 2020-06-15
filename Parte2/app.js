document.querySelector('#proc').addEventListener('click', function(){
            obtenerDatos();
        })

        function obtenerDatos(){

        	let url = 'http://127.0.0.1:3000/procesos';
        	
        	const api = new XMLHttpRequest();
        	api.open('GET', url); 
        	api.send();
        	
        	api.onreadystatechange = function() {
		    if (this.readyState == 4 && this.status == 200) {
		      //document.getElementById("demo").innerHTML = this.responseText;
		      let datos = JSON.parse(this.responseText);

		      //console.log(datos);

		      let resultado = document.querySelector('#res'); //
		      resultado.innerHTML = '';

		      for(let item of  datos){
		      	//${item.PADRE}|${item.PID}|${item.NOMBRE}|${item.STADO} 
		      	resultado.innerHTML +=` <tr>
							            <td>${item.PADRE}</td>
							            <td>${item.PID}</td>
							            <td>${item.NOMBRE}</td>
							            <td>${item.STADO} </td>
							          	</tr>`;
		      }
		    }
		  };
            console.log("clic reconocido")
        }