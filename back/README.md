# Application horraire Tisseo

Back-end Go d'une application POC d'affichage d'horraire pour exemplifier la mise en place d'un Infrastructure As Code Terraform/Ansible.

## Déploiement

Le livrable est `tisseo-back`.

Pour le re-compiler, Utilisez la commande: 
```bash
GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -ldflags="-s -w" -o tisseo-back ./cmd/tisseo
```

- GOOS: OS cible
- GOARCH: Architecture cible
- -o: Nom du livrable compilé

L'application est déployée sur le port 8080.

## Environnement

L'environnement cible doit avoir la variable d'environnement `TISSEO_BACK_SERV_NAME` mise en place pour pouvoir servir le nom du serveur sur /server
