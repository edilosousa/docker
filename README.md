# COMANDOS DE CRIAR CONTAINERS
## docker run nome_do_container

# Lista os containers que estão rodando.
## docker ps

# Lista os container e quando ele parou
## docker ps -a

# Rodar o Ubuntu pra ficar editavel o Bash
## docker run -it ubuntu bash
## apt update -> atualiza a imagem do ubuntu

# Instalação do NGINX rodando na porta 8000
## docker run -p 8000:80 nginx

# Editar os arquivos do container
## docker exec -it nome_container bash
## dentro do container -> apt update --atualizar o linux dentro do container
## apt install vim -y -> instalação para editar os arquivos

# Listar images de containers
## docker images

# Deletar containers
## docker ps -a
## docker rm idcontainer
## docker images -> para pegar o ID
## docker rmi idcontainer

# Construir uma imagem apartir de um dockerfile
## docker build -t edilosousa/nginx:latest .
## docker run -p 8000:80 edilosousa/nginx:latest

# Compartilhamento no docker hub -- compartilhará na sua conta docker
## docker push nomedaimagem
## exemplo docker push edilosousa/nginx:latest