# Atualização da Versão do Go

Este guia fornece instruções sobre como atualizar a versão do Go em seu sistema Linux.

## Passo 1: Remoção da Versão Existente do Go

1. Remova o pacote existente do Go:
   ```
   sudo apt-get remove golang-go
   ```

2. Remova as dependências do pacote do Go:
   ```
   sudo apt-get remove --auto-remove golang-go
   ```

3. Desinstale a versão atual do Go:
   ```
   sudo rm -rvf /usr/local/go
   ```

## Passo 2: Instalação da Nova Versão do Go

1. Baixe o arquivo binário da versão específica do Go para o seu sistema:
   ```
   wget https://dl.google.com/go/go1.21.0.linux-amd64.tar.gz
   ```

   > Nota: Substitua a versão do Go acima pela versão correspondente ao seu sistema. A lista de versões pode ser encontrada [aqui](https://golang.org/dl/).

2. Extraia o arquivo do arquivo compactado:
   ```
   sudo tar -xvf go1.21.0.linux-amd64.tar.gz
   ```

3. Mova o diretório extraído para a localização desejada no sistema:
   ```
   sudo mv go /usr/local
   ```

   > Dica: A localização acima é recomendada para sistemas Linux.

## Passo 3: Configuração do Ambiente Go (Linux)

1. Configure a variável de ambiente do Go:
   Abra o arquivo `.profile` localizado no diretório home (~/.profile) e adicione as seguintes linhas:

   ```bash
   export GOROOT=/usr/local/go
   export GOPATH=$HOME/go
   export PATH=$GOPATH/bin:$GOROOT/bin:$PATH
   ```

   > `GOROOT` é o local onde o pacote do Go está instalado no seu sistema. `GOPATH` é o diretório de trabalho do seu projeto Go.

2. Atualize as variáveis de ambiente para que as alterações tenham efeito. Isso pode ser feito executando o seguinte comando no terminal ou fazendo login novamente no shell atual:
   ```
   source ~/.profile
   ```

## Passo 4: Verificação da Versão do Go e do Ambiente

1. Verifique a versão do Go instalada:
   ```
   go version
   ```

2. Verifique as variáveis de ambiente configuradas:
   ```
   go env
   ```

Agora você atualizou com sucesso a versão do Go em seu sistema e configurou o ambiente para desenvolvimento com Go.