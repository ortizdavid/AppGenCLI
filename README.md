## BASIC INFO:
- appgen - tool for generate applications
- Source code     https://www.github.com/ortizdavid/appgenerator
- Version         1.0.0

## DESCRIPTION:
    This tool helps create the structure of an application, including database
    For now it generates for Python and PHP

## USAGE:
``
**appgen** **-name** <application name> **-lang** <language> **-type** <type of project> **-db** <database>
``

##  COMANDS:
- **appgen**:     First command
- **list-langs**: List all supported languages and applications
- **help**:       Shows helps for appgen
- **-name**:      Project Name
- **-lang**:      Programming Language
- **-type**:      Type of application (mvc, api)
- **-db**:        database 

## EXAMPLES:

Create a MVC App with Python and MySQL
    ``
    appgen -name PythonWebMVC -lang python -type mvc -db mysql 
    ``

Creates an API with Python and Postgres
    ``
    appgen -name MyPythAPI -lang python -type api -db postgres    
    `` 

Create a MVC App with PHP and MySQL
    ``
    appgen -name PHPSimpleApp -lang php -type mvc -db mysql 
    ``

Show help comands
    ``
    appgen help 
    ``   

List all suportded languages
    ``
    appgen list-langs 
    ``                                                      

## AUTHOR:
- Name:         Ortiz de Arcanjo António David
- Phone:        +244 936 166 699
- Email:        ortizaad1994@gmail.com
- Github:       https://www.github.com/ortizdavid
- LinkedIn:     https://www.linkedin.com/in/ortiz-david