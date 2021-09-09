# Api Identificar Mutante

La siguiente API esta desarrollada para ser cargada en un Lambda desde un archivo compilado .zip que se incluye en el repositorio.

Instrucciones de ejecución

- Descargar el archivo function.zip
- Desde aws crear un lamba en go y cargar function.zip
- Ir a apiGateway 
- Seleccionar API REST, Api Nueva y Crear 
- Crear el recurso /mutant
- Crear los metodos GET / y POST /mutant 
- Seleccionar funcion lambda
- Check en "Usar la integración de proxy Lambda"
- Seleccionar la funcion lambda cargada previamente


# URL 

- POST https://i2v7f7g5md.execute-api.us-east-2.amazonaws.com/meli/mutant
- GET https://i2v7f7g5md.execute-api.us-east-2.amazonaws.com/meli/

