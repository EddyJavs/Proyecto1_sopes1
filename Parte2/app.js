document.querySelector('#proc').addEventListener('click', function(){
            obtenerDatosPadre();
        })

document.querySelector('#proc2').addEventListener('click', function(){
            obtenerDatosHijos();
        })

        function obtenerDatosPadre(){

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
		      resultado.innerHTML = `<thead>
                                      <tr>
                                      <th>PID</th>
                                      <th>NOMBRE</th>
                                      <th>ESTADO</th>
                                      </tr>
                                      </thead><tr>
                                      <tbody>`;

		      for(let item of  datos){
		      	//${item.PADRE}|${item.PID}|${item.NOMBRE}|${item.STADO} 
		      	if (item.PADRE == item.PID){
		      		resultado.innerHTML +=`<tr>
							            <td>${item.PID}</td>
							            <td>${item.NOMBRE}</td>
							            <td>${item.STADO} </td>
							          	<td><p>
									      <label>
									         <input type="checkbox"  onclick="eliminarProc(this,'${item.PID}')"/>
									        <span>Kill</span>
									      </label>
									    </p>
									    </td>
							          	</tr>`;
		      	}
		      	
		      }
		      resultado.innerHTML += `</tbody>`
		    }
		  };
		   //$(".padre").toggle();
		   //$('td:nth-child(1)').toggle();

            console.log("clic reconocido")
        }

        function obtenerDatosHijos(){

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
		      resultado.innerHTML = `<thead>
                                      <tr>
                                      <th>PADRE</th>
                                      <th>PID</th>
                                      <th>NOMBRE</th>
                                      <th>ESTADO</th>
                                      </tr>
                                      </thead><tr>
                                      <tbody>`;

		      for(let item of  datos){
		      	//${item.PADRE}|${item.PID}|${item.NOMBRE}|${item.STADO} 
		      	if (item.PADRE != item.PID){
		      		resultado.innerHTML +=`<tr>
		      							<td>${item.PADRE}</td>
							            <td>${item.PID}</td>
							            <td>${item.NOMBRE}</td>
							            <td>${item.STADO} </td>
							            <td><p>
									      <label>
									        <input type="checkbox"  onclick="eliminarProc(this,'${item.PID}')"/>
									        <span>Kill</span>
									      </label>
									    </p>
									    </td>
							          	</tr>`;
		      	}
		      	
		      }
		      resultado.innerHTML += `</tbody>`
		    }
		  };
		   //$(".padre").toggle();
		   //$('td:nth-child(1)').toggle();

            console.log("clic reconocido")
        }

function eliminarProc(tr,value) {

  alert("PDI -> " + value);

  let url = `http://127.0.0.1:3000/matarTask/${value}`;
        	
        	const api = new XMLHttpRequest();
        	api.open('GET', url); 
        	api.send();
        	
        	api.onreadystatechange = function() {
		    if (this.readyState == 4 && this.status == 200) {
		    	tr.disabled = true;
		    }}
}