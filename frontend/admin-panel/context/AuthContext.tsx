"use client";

import React, { createContext, useContext, useEffect, useState, ReactNode, Dispatch, SetStateAction } from 'react';
import { getAuth, onAuthStateChanged, User } from 'firebase/auth';
import { auth } from '@/app/firebase/config';
import Cookies from 'js-cookie';

interface AuthContextProps {
    user: User | null;
    setUser: Dispatch<SetStateAction<User | null>>;
    loading: boolean;
    token: string | null;
    email: string | null; // Ajoutez la propriété email ici
}

const AuthContext = createContext<AuthContextProps | undefined>(undefined);

export const AuthProvider = ({ children }: { children: ReactNode }) => {
    const [user, setUser] = useState<User | null>(null);
    const [loading, setLoading] = useState(true);
    const [token, setToken] = useState<string | null>(null);
    const [email, setEmail] = useState<string | null>(null);

    useEffect(() => {
        const unsubscribe = onAuthStateChanged(auth, async (user) => {
            console.log("AuthProvider - User state changed:", user);
            setUser(user);
            setLoading(false);
            // TODO , trés mauvaise pratique , changer le fonctionnement urgent !
            if (user) {
                const token = await user.getIdToken();
                const userEmail = user.email;
                setToken(token);
                setEmail(userEmail);
                Cookies.set('token', token);
                Cookies.set('email', userEmail!);
            } else {
                setToken(null);
                setEmail(null);  
                Cookies.remove('token');
                Cookies.remove('email');
            }
        });

        return () => unsubscribe();
    }, []);

    return (
        <AuthContext.Provider value={{ user, setUser, loading, token, email }}>
            {children}
        </AuthContext.Provider>
    );
};

export const useAuth = () => {
    const context = useContext(AuthContext);
    if (context === undefined) {
        throw new Error('useAuth must be used within an AuthProvider');
    }
    return context;
};
