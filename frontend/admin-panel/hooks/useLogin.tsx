// hooks/useLogin.ts
import { getAuth, signInWithEmailAndPassword, signOut } from 'firebase/auth';
import { useAuth } from '@/context/AuthContext';
import { useRouter } from 'next/navigation';
import { useState } from 'react';

export const useLogin = () => {
    const { setUser } = useAuth();
    const auth = getAuth();
    const router = useRouter();
    const [isValidating, setIsValidating] = useState(false);

    const login = async (email: string, password: string) => {
        setIsValidating(true);
        
        try {
            // 1. Authentification Firebase
            const userCredential = await signInWithEmailAndPassword(auth, email, password);
            const token = await userCredential.user.getIdToken();
            
            // 2. Vérification des permissions via l'API backend AVANT de définir l'utilisateur
            const response = await fetch(`${process.env.NEXT_PUBLIC_PV_PATH_API}/dashboard?token=${token}`);
            
            if (response.status !== 200) {
                // Déconnecter de Firebase si pas autorisé
                await signOut(auth);
                throw new Error('Access denied - You are not authorized to access this dashboard');
            }
            
            // 3. SEULEMENT maintenant, définir l'utilisateur (après validation backend)
            setUser(userCredential.user);
            return { user: userCredential.user };
            
        } catch (error) {
            console.error('Error signing in:', error);
            
            // S'assurer que l'utilisateur est déconnecté en cas d'erreur
            try {
                await signOut(auth);
                setUser(null); // Forcer la réinitialisation du contexte
            } catch (signOutError) {
                console.error('Error signing out after failed login:', signOutError);
            }
            
            throw error; // Propager l'erreur à l'appelant
        } finally {
            setIsValidating(false);
        }
    };

    const logout = async () => {
        try {
            await signOut(auth);
            setUser(null);
            router.push('/login'); // Rediriger vers la page de login
        } catch (error) {
            console.error('Error signing out:', error);
            throw error;
        }
    };

    return { login, logout, isValidating };
};