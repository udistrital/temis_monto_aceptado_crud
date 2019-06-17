
# :ledger: temis_monto_aceptado_crud

CRUD para el módulo de gestión de información de las cuotas partes de pensionados - TEMIS.

El proyecto está escrito en lenguaje Go, generado mediante el <a href="https://beego.me/">framework BeeGo</a>

# Modelo de datos

![image](https://drive.google.com/uc?export=view&id=1r3Bs-4AodOPHIuEoUXkmIWt0v_mFk17e)

# Requerimientos

Go version >= 1.8.

<summary><h2> 🛠️ Preparación </h2></summary>
<details>

  # Opción 1

  ```shell 
      go get github.com/udistrital/temis_monto_aceptado_crud
  ```
  
  # Opción 2

  - Clonar el proyecto del repositorio de git en la carpeta local
  
  ```shell 
    cd ~go/src/github.com/udistrital 

    git clone https://github.com/udistrital/temis_monto_aceptado_crud.git
  ```
  
  Entramos a la carpeta del proyecto instalado

  ```shell 
      cd ~/temis_monto_aceptado_crud/
  ```

  Instalar dependencias del proyecto

  ```shell 
     go get
  ```
</details>

<summary><h2> Ejecución </h2></summary>
<details>


  Ubicado en la raíz del proyecto

  ```shell 
      cd ~/go/src/github.com/udistrital/temis_monto_aceptado_crud
  ```
   
  Ejecutar
  
  ```shell 
      bee run
  ```

  Si se quiere ejecutar el swagger

  ```shell 
      bee run -downdoc=true -gendoc=true
  ```
</details>