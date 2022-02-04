import {React, useState} from 'react'
import {Menu} from 'semantic-ui-react'
import '../Estilos/Nav.css'
import {saveAs} from 'file-saver'
import axios from 'axios'
//window.location.reload();

var realizado = false
const colores=[ 'green', 'yellow', 'orange', 'teal', 'violet']
const opciones =[ 'Guardar Archivo', 'Ejecutar', 'Reporte de Errores', 'Árbol AST', 'Reporte de Símbolos']



function NavBarIncio() {

    async function post(e){
        actual = localStorage.getItem('actual')
        lastChar = actual[actual.length -1];
        lastChar = parseInt(lastChar)
        datosPes = localStorage.getItem('datosPes')
        datosP = JSON.parse(datosPes);
        var entrada = datosP[lastChar];
        try{
            await axios.post("http://localhost:5000/ejecutar", {
                entrada
            })
            .then(response=>{
                var respuesta = "";
                for (let i = 0; i < response.data.message.length; i++) {
                    respuesta += response.data.message[i]+"\n"
                }
                console.log(respuesta)
                localStorage.setItem('consola', JSON.stringify([respuesta]))
            })
        }catch(error){
            console.log(error)
        }        
        window.location.reload()
    }
    
    const leerArchivo = (e) =>{
        const file = e.target.files[0];
        if(!file) return;
        const fileReader = new FileReader();

        fileReader.readAsText(file);

        fileReader.onload = () =>{
            var actual = localStorage.getItem('actual')
            var lastChar = actual[actual.length -1];
            lastChar = parseInt(lastChar)
            var datosPes = localStorage.getItem('datosPes')
            var datosP = JSON.parse(datosPes);
            datosP[lastChar] = fileReader.result;
            localStorage.setItem('datosPes', JSON.stringify(datosP))
            window.location.reload()
        }

        fileReader.onerror = () => {
            console.log( fileReader.error )
        }

    }

    const [activo, setactivo] = useState(colores[6])

    if (activo === "green"){
        if(!realizado){
            var actual = localStorage.getItem('actual')
            var lastChar = actual[actual.length -1];
            lastChar = parseInt(lastChar)
            var datosPes = localStorage.getItem('datosPes')
            var datosP = JSON.parse(datosPes);
            const blob = new Blob([datosP[lastChar]], {type: 'text/plain;charset=utf-8'})
            saveAs(blob, 'Pestaña'+lastChar+".txt")
            realizado = true
        }
        window.location.reload()
    }else if (activo === "yellow"){
        //localStorage.setItem('consola', JSON.stringify(["Bienvenido a mi proyecto de Compiladores 1"]))
        post();
    }else if (activo === "orange"){
        window.location.href = "http://localhost:3000/Errores"
    }else if (activo === "teal"){
        hacerAST()
    async function hacerAST(){
        var actual = localStorage.getItem('actual')
        var lastChar = actual[actual.length -1];
        lastChar = parseInt(lastChar)
        var datosPes = localStorage.getItem('datosPes')
        var datosP = JSON.parse(datosPes);
        var entrada = datosP[lastChar];
        await axios.post("http://localhost:5000/ast", {
            entrada
        })
        .then(response=>{
            console.log(response)
        })
        window.location.reload()
    }
    }else if (activo === "violet"){
        window.location.href = "http://localhost:3000/Simbolos"
    }
    return (
       <Menu inverted className="Nav" >
           
           <Menu.Item 
                key={"light blue"}
                name={"Cargar Archivo"}
                active={activo==='light blue'}
                color={'light blue'}
            >   
                <label for="selectedFile" id="etiqueta">Cargar Archivo</label>
                <input type="file" id="selectedFile" style={{display:"none"}} onChange={ leerArchivo } multiple="false"/>
            </Menu.Item>

           {colores.map((c,iterador)=>(
               <Menu.Item 
                    key={c}
                    name={opciones[iterador]}
                    active={activo===c}
                    color={c}
                    onClick={()=>setactivo(c)}
               />
           ))}
       </Menu>
    )
}

export default NavBarIncio
