Here’s a professional and clear README for your GitHub repository, written in English:

---

# Backend Setup Instructions

### Prerequisites:
- **Golang**: Version 1.23.2 or higher.

### Database Setup:
1. Load the database schema:
   - Import the SQL schema located at `backend/datasources/mysql/schema/appli.sql` into your MySQL database.

2. Configure environment variables:
   - Create a `.env` file at the root of the `backend` directory.
   - Add the following database access credentials to the file:

   ```plaintext
   MYSQL_APPLI_DB=shop-mysql
   MYSQL_USERNAME=root
   MYSQL_PASSWORD=
   MYSQL_HOST=localhost:3306
   MYSQL_APPLI_PORT=8080
   ```

### Firebase Setup:
1. Add your Firebase service account private key:
   - Place the private key file `firebase-service-account-key.json` in the path `backend/appli/datasources/firebase/firebase-service-account-key.json`.

2. How to obtain your Firebase service account key:
   - The `firebase-service-account-key.json` file contains the private key for your Firebase service account. Follow these steps to generate and download the file:

   **Steps to get your Firebase private key:**
   - Go to the [Firebase Console](https://console.firebase.google.com) and log in with your Google account.
   - Select the project you want to use.
   - Click on the gear icon next to your project name and select **Project Settings**.
   - Navigate to the **Service Accounts** tab.
   - Click on **Generate New Private Key** and a file named `firebase-service-account-key.json` will be automatically downloaded.

   **Important:** Store this file securely, as it contains sensitive information necessary for authenticating your application with Firebase.

### Running the Backend:

Once you've set up the environment and database, you can run the backend using the following command:

```bash
cd backend
go run main.go
```

 