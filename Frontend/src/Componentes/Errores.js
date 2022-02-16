import {React,  useState} from 'react'
import axios from 'axios';
import Tabla from './TablaErrores';
import '../Estilos/App.css'

function Errores() {
    const [Data, setData] = useState([])
    const encabezado=['MENSAJE', 'TIPO', 'LINEA', 'COLUMNA']
    //setData(response.data.message)

    tablaErrores()

    async function tablaErrores(){
        var actual = localStorage.getItem('actual')
        var lastChar = actual[actual.length -1];
        lastChar = parseInt(lastChar)
        var datosPes = localStorage.getItem('datosPes')
        var datosP = JSON.parse(datosPes);
        var text = datosP[lastChar];
        try{
            await axios.post("http://localhost:3001/Errores", {
                text
            })
            .then(response=>{
                //console.log(response.data.message)
                setData(response.data.message)
            })
        }catch(error){
            console.log(error)
        }  
    }

    return (
        <div className="GeneralLogin">
            <Tabla
                encabezados={encabezado}
                data = {Data}
            />
        </div>
    )
}

export default Errores
