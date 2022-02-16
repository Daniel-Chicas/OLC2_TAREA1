import React from 'react'

function Fila(props) {
    return (
        <>
            <tr key={props.index}>
                <td>{props.Mensaje}</td>
                <td>{props.Tipo}</td>
                <td>{props.Linea}</td>
                <td>{props.Columna}</td>
            </tr>
        </>
    )
}

export default Fila
