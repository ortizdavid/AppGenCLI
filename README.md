## BASIC INFO:
- appgen - tool for generate applications
- Source code     https://www.github.com/ortizdavid/appgenerator
- Version         1.0.0

## DESCRIPTION:
    This tool helps create the structure of an application, including database
    For now it generates for Python and PHP

## USAGE:
``
appgen -name <appname> -lang <language> -type <type of project> -db <database>
``

##  COMANDS:
appgen           First command
- list-langs      List all supported languages and applications
- help            Shows helps for appgen
- -name           Project Name
- -lang           Programming Language
- -type           Type of application (mvc, api)
- -db             database 

## EXAMPLES:
- appgen -name PythonWebMVC -lang python -type mvc -db mysql              Creates a MVC App with Python and MySQL
- appgen -name MyPythAPI -lang python -type api -db postgres              Creates an API with Python and Postgres
- appgen -name PHPSimpleApp -lang php -type mvc -db mysql                 Creates a MVC App with PHP and MySQL
- appgen help                                                             Shows help comands
- appgen list-langs                                                       Lists all suportded languages

## AUTHOR:
- Name:         Ortiz de Arcanjo António David
- Phone:        +244 936 166 699
- Email:        ortizaad1994@gmail.com
- Github:       https://www.github.com/ortizdavid
- LinkedIn:     https://www.linkedin.com/in/ortiz-david