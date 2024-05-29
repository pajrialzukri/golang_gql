# Backend

## Installation:

1. Mulai edit ```config.yaml``` (silakan copy dari ```config.template.yaml```), kemudian edit parameter sesuaikan dengan environment masing-masing.
    ```
    cp config.template.yaml config.yaml
    ```
2. Update/fix dependency:
    ```
    go mod tidy
    ```
3. Run
    ```
    go run server.go 8080
    ```

## Create new module:
1. Create graphqls file for the module: edit ```graph/graphqls/{module_name}.graphqls```
  
3. Generate model & resolvers:
    ```
    go run github.com/99designs/gqlgen generate
    ```
4. Resolvers implementation: edit ```graph/resolvers/{module_name}.resolvers.go```
  
6. Run:
    ```
    go run server.go 8080
    ```


## Penambahan db Sequence connector  
config.yaml :
dbsequser: root
dbseqpass: 
dbseqhost: localhost
dbseqport: 3306
dbseqname: his_seq
```
pemakaian
```
db.Handle.Query(q) 
di ganti menjadi
db.HandleSeq.Query(q)

db.HandleSeq.Exec(q)
diganti menjadi 
db.HandleSeq.Exed(q)

```