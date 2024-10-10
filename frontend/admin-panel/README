Here's the revised README that keeps your original content and adds a section on how to find the Firebase configuration values:

```markdown
# E-Shop Panel Admin Frontend

## Prerequisites

- **Node.js** (version 14 or later)
- **npm** (comes with Node.js)

## Backend Setup

1. Load the database using the SQL file:
   ```bash
   # Execute in your MySQL client
   source backend/datasources/mysql/schema/appli.sql;
   ```

## Frontend Setup

1. Install dependencies:
   ```bash
   cd frontend
   npm install
   npm install react-hot-toast axios
   ```

2. Create a `.env` file in the frontend root with your Firebase configuration:
   ```bash
   NEXT_PUBLIC_FIREBASE_API_KEY=
   NEXT_PUBLIC_FIREBASE_AUTH_DOMAIN=
   NEXT_PUBLIC_FIREBASE_PROJECT_ID=
   NEXT_PUBLIC_FIREBASE_STORAGE_BUCKET=
   NEXT_PUBLIC_FIREBASE_MESSAGING_SENDER_ID=
   NEXT_PUBLIC_FIREBASE_APP_ID=
   NEXT_PUBLIC_PV_PATH_API=http://localhost:8081/api/v1/private
   NEXT_PUBLIC_PB_PATH_API=http://localhost:8081/api/v1/public
   ```

3. Start the frontend server:
   ```bash
   npm run dev
   ```

## How to Obtain Firebase Configuration Values

To populate your `.env` file with the necessary Firebase configuration values, follow these steps:

1. **Access Firebase Console**:
   - Go to [Firebase Console](https://console.firebase.google.com/) and log in with your Google account.

2. **Select Your Project**:
   - Click on the Firebase project you created (in your case, it should be **e-shop-5caec**).

3. **Retrieve Firebase Configuration**:
   - In the left menu, click on the **Project Settings** gear icon.
   - Navigate to the **General** tab and scroll down to the **Your apps** section.
   - If you haven't already added a web app, click on **Add app** and select **Web**. Follow the prompts to register your web app.
   - After adding the app, you will see the Firebase SDK snippet containing the configuration details.

4. **Copy the Configuration Values**:
   - Copy the following values from the snippet provided in the Firebase Console:
     - `apiKey`
     - `authDomain`
     - `projectId`
     - `storageBucket`
     - `messagingSenderId`
     - `appId`

5. **Update Your `.env` File**:
   - Replace the placeholders in your `.env` file with the actual values you copied.
