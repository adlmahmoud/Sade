## 🎶 Projet Sade - L'Élégance Numérique

Ceci est un projet personnel et académique dédié à l'artiste **Sade**, visant à capturer l'ambiance calme, sophistiquée et profondément émotionnelle de sa musique dans un format numérique. Le site est construit comme un hommage à son art intemporel.

---

## 💻 Technologies Utilisées

* **Backend:** Go (Golang)
* **Frontend:** HTML5, CSS3, JavaScript
* **Données:** API Web de **Spotify**

---

## 🚀 Démarrage du Projet

Pour lancer l'application en local, vous devez suivre les étapes de configuration ci-dessous.

### 1. Prérequis

Assurez-vous d'avoir installé les éléments suivants sur votre machine :

* **Go (Golang)** : Vous pouvez le télécharger et l'installer depuis le [site officiel de Go](https://go.dev/dl/).

### 2. Configuration de l'API Spotify

Ce projet utilise l'**API Web de Spotify** pour récupérer les informations sur l'artiste, les albums et les titres. Vous devez obtenir vos propres identifiants (**Client ID** et **Client Secret**) en enregistrant votre application sur le [Dashboard des développeurs Spotify](https://developer.spotify.com/dashboard/).

Une fois vos identifiants obtenus, vous devez les définir comme **variables d'environnement** :

* **Pour Windows (PowerShell) :**
    ```powershell
    $Env:SPOTIFY_CLIENT_ID = "votre_client_id"
    $Env:SPOTIFY_CLIENT_SECRET = "votre_client_secret"
    ```
* **Pour macOS/Linux (Bash/Zsh) :**
    ```bash
    export SPOTIFY_CLIENT_ID="votre_client_id"
    export SPOTIFY_CLIENT_SECRET="votre_client_secret"
    ```

> ⚠️ **Note Importante :** Remplacez `"votre_client_id"` et `"votre_client_secret"` par les valeurs réelles fournies par Spotify.

### 3. Lancement de l'Application

1.  **Naviguez** dans le répertoire source de l'application :
    ```bash
    cd src
    ```
2.  **Lancez** le serveur Go en exécutant la commande suivante :
    ```bash
    go run main.go
    ```
3.  Le serveur devrait maintenant être en cours d'exécution. Ouvrez votre navigateur web et accédez à l'adresse indiquée (généralement `http://localhost:8080`).

---

## 👤 Auteur

* **[ADEL Mahmoud]** -https://github.com/adlmahmoud-
