Api RestFull    
    API : c'est une regle de communication entre 2 application pr ex navig et le serveur http
    REST : c'est une architecture qui est basé sur les exposes du ressources, echange des donnees , accessibles via  URLS , manipulee vc des methodes http
    
    c'est quoi la difference entre API REST et un serveur web classique ?
        un serveur :
            recoit une requete
            executee le code
            envoyer une réponse
            
        donc la le point est comment le serveur répond au client
            si le serveur envoyer des donnees et laisse au client manipuler ces ressources la on parle sur l'utilisation des api restful
            si le serveur envoyer une vue prete 'page html' la on parle un serveur web classique
            si le serveur décidée a la place du client pr exemple la redirection des pages ect la on parle aussi un serveur web
            
        c/c ils se differencient au niveau d contexte et utilisation mais  au niveau technique se sont les memes ()
        
        RQ: la différence entre les deux est APIRestFul est statelesse chaque requete est indépendante cad a chaque fois on doit envoyer un token, cepandant srv classique garde une session souvient toi 
            
    
    Régles API RESTFUL 
        -stateless Le serveur ne garde aucune mémoire d’une requête à l’autre cad chaque requete doit contenir ses infos nécessaire
        -Utilisation des méthodes http
        -les routes représentent des ressources au contraire ds un serveur classique il represente des actions ou des pages ect 
        -réponse => des données sont de type json


    JSON
        javascript object Notation (c'est une juste format des données qui est basé sur js )
            Objets + tableau
            {
                key:value
            }
            []

    Principes d' API RESTFUL :
        REST : est une architecture , non techno
            repose sur 6 prinicpes:
                RESSOURCES:
                    ### ENDPOINTS se sont des routes (urls) sur le serveur à laquelle le client peut envoyer une requête pour obtenir une ressource ou effectuer une action.Chaque endpoint correspond à une action spécifique sur une ressource => URL + METHODE = endpoint distinct 
                    chaque ressource est identifiée par une URL >> /users ; /produits
                    RQ : le verbe ne va pas ds l'url => /createUsers

                HTTP METHODS: 
                    REST utilise ces méthodes pour spécifiée l'action 
                    GET    >> lire
                    POST   >> créer
                    PUT    >> modifier(tout)
                    PATCH  >> modifier(partiel)
                    DELETE >> Supprimer

                STATELESS:
                    - STATE est une information mémorisée  entre requetes 
                        un serveur web classique est STATEFULL il créer une session en mémoire 
                    -chaque requete doit contenir toutes les infos nécessaire 
                    -Pas de session coté serveur
                    -Token envoyé a chaque requete

                JSON:
                    -il renvoie une représentation json

                CODES HTTP:
                    200 >> OK
                    201 >> CREE
                    400 >> MAUVAISE REQUETE
                    401 >> NON AUTHENTIFIE
                    403 >> INTERDIT
                    404 >> NON TROUVE
                    500 >> ERREUR SERVEUR


## FLUX DE TRAVAIL
### Creation d'un serveur HTTP en GO
    http.ListenAndServe(":8080", nil)
        -sert a démarrer le  serveur http au port 8080
        -nil servent a utiliser le DefaultServeMux >> est un routeur qui associée les routes a leurs handlers 
        -Creation d'un routeur personnalisée on fait http.NewServeMux()

#### La question qu'on doit poser  comment le serveur écoute a ce port ???
    -le serveur GO cree un socket TCP  sur le port 8080
        TCP >> Transmission Control Protocol (protocle de communication réseau qui fonctionne au dessus IP permet transfer d message entre noeuds)
        SOCKET >> est une interface qui relie ton programme au réseau il définit adresse IP + PORT + Protocole 
         >>os réserve une structure interne 
         >> serveur associe ce socket IP + PORT
         >> serveur dit au sysyteme attends des connexions

### Structure du projet
    on va creer une interface web et un serveur http qui va répond au reqeutes de navigateur -client- 
    alors on va faire  un combinaison entre serveur http classique et APIREST 
        cad le serveur classique va communiquer entre un autre  serveur qui a génerer des api REST full  et on doit le récupérer a travers ce serveur classique qui va au dernier creation d'une interface web  par GO LANG

    * Récupération des ressources d'apres ces api Rest full
```GO
    /*
        http.Get => elle retourne un objet de type *http.Response
        elle contient 
        Status, StatusCode, Header, Body, ContentLength, Proto
        res.Body >> on doit faire ioutil.ReadAll(res.Body) ou json.NewDecoder(res.Body) pour récuperer le contenu 
        car res.Body n'est pas un texte , est un flux (stream)

        C/C le http est basé sur le protocole TCP qui permet ce protocole faire la transmission des paquets cad il segment le message
        donc Go  => res.Body est la porte de ces paquest cad les donnee ne vient pas en une seul fois 
        on donne l'exemple comme un robinet + l'eau + tuyau 
            l'eau => données
            tuyau => res.Body 
            robinet => le serveur 

            [Serveur Go (HTTP)]      [Client Go]
                    🔹 robinet 🔹
                    │
                    │    <- resp.Body = tuyau
                    │
            {"id":1,"nom":"Oumaima","email":"oumaima@mail.com"} = eau

            donc le serveur envoyer un flux d'octet 
            res.Body est une interface  io.ReadCloser qui combine 2 interfaces 

                :::: Je suis Un flux que tu peux me lire et fermer ::::
                type ReadCloser interface {
                    Read(p []byte) (n int, err error) => Lire les octets depuis le flux 
                    Close() error => fermer le flux
                }
    
            Buffer : est un espace mémoire temporaire où on peut stocker des données pendant qu’on les lit ou qu’on les écrit.
            
            
    */
    res, err := http.GET("https://groupietrackers.herokuapp.com/api/artists")
    if err != nil {
        ...
    }
```




