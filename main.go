package main

import (
    "fmt"
    "net"
    "os"
    "google.golang.org/protobuf/proto"
    "github.com/FelipeFAlves/testecliente2/contact"
)

func main() {
    // Endereço e porta do servidor Python
    serverAddr := "200.235.90.15:8888"

    // Conectar ao servidor
    conn, err := net.Dial("tcp", serverAddr)
    if err != nil {
        fmt.Println("Erro ao conectar ao servidor:", err)
        os.Exit(1)
    }
    defer conn.Close()

    // Criar uma mensagem Pessoa
    pessoa := &contact.Pessoa{
        Nome:   "gabirel",
        Idade:  652,
        Cidade: "pinda",
    }

    // Serializar a mensagem Pessoa
    data, err := proto.Marshal(pessoa)
    if err != nil {
        fmt.Println("Erro ao serializar a mensagem:", err)
        return
    }

    // Enviar a mensagem serializada para o servidor
    _, err = conn.Write(data)
    if err != nil {
        fmt.Println("Erro ao enviar a mensagem para o servidor:", err)
        return
    }

    fmt.Println("Mensagem enviada para o servidor com sucesso!")

    // Receber a mensagem modificada do servidor
    modifiedData := make([]byte, 1024)
    n, err := conn.Read(modifiedData)
    if err != nil {
        fmt.Println("Erro ao receber a mensagem modificada:", err)
        return
    }

    // Desserializar a mensagem Pessoa modificada
    modifiedPessoa := &contact.Pessoa{}
    err = proto.Unmarshal(modifiedData[:n], modifiedPessoa)
    if err != nil {
        fmt.Println("Erro ao desserializar a mensagem modificada:", err)
        return
    }

    // Imprimir os detalhes da mensagem Pessoa modificada
    fmt.Println("Mensagem modificada recebida:")
    fmt.Println("Nome:", modifiedPessoa.GetNome())
    fmt.Println("Idade:", modifiedPessoa.GetIdade())
    fmt.Println("Cidade:", modifiedPessoa.GetCidade())
}