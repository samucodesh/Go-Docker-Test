package main

import (
    "html/template"
    "net/http"
)

func main() {
    // Creamos un template HTML simple
    tmpl := template.Must(template.New("").Parse(`
        <!DOCTYPE html>
        <html lang="en">
        <head>
            <title>Test html template </title>
        </head>
        <body>
            <h1>Test works!</h1>
        </body>
        </html>
    `))

    // Creamos un manejador de solicitudes que devuelve el template
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        tmpl.Execute(w, nil)
    })

    // Escuchamos en el puerto 8080
    http.ListenAndServe(":8080", nil)
}
