import React from 'react'
import Fila from './FilaErrores'
import { Table } from 'semantic-ui-react'

function TablaErrores(props) {    
    return (
        <div>
            <div className="Tabla">
            <div className="ui segment container ">
                <Table celled inverted selectable>
                    <thead>
                        <tr>
                            {props.encabezados.map((dato) => (
                                <th>{dato}</th>
                            ))}
                        </tr>
                    </thead>
                    <tbody>
                        {props.data.map((dato, index) => (
                            <Fila
                                index={index}
                                Mensaje={dato.Mensaje}
                                Tipo={dato.Tipo}
                                Linea = {dato.Linea}
                                Columna = {dato.Columna}
                            />
                        ))}
                    </tbody>
                </Table>
            </div>
        </div>
        </div>
    )
}

export default TablaErrores
