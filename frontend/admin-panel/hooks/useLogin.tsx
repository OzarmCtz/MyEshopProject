// hooks/useLogin.ts
import { getAuth, signInWithEmailAndPassword } from 'firebase/auth';
import { useAuth } from '@/context/AuthContext';

export const useLogin = () => {
    const { setUser } = useAuth();
    const auth = getAuth();

    const login = async (email: string, password: string) => {
        try {
            const userCredential = await signInWithEmailAndPassword(auth, email, password);
            const token = await userCredential.user.getIdToken();
            const response = await fetch(`${process.env.NEXT_PUBLIC_PV_PATH_API}/dashboard?token=${token}`);

            if (response.status !== 200) {
                throw new Error('Access denied');
            }

            setUser(userCredential.user);
            return { user: userCredential.user };
        } catch (error) {
            console.error('Error signing in:', error);
            throw error; // Propager l'erreur à l'appelant
        }
    };

    return { login };
};
