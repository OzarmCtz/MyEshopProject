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
    email: string | null;
    isValidating: boolean;
    setIsValidating: Dispatch<SetStateAction<boolean>>;
}

const AuthContext = createContext<AuthContextProps | undefined>(undefined);

export const AuthProvider = ({ children }: { children: ReactNode }) => {
    const [user, setUser] = useState<User | null>(null);
    const [loading, setLoading] = useState(true);
    const [token, setToken] = useState<string | null>(null);
    const [email, setEmail] = useState<string | null>(null);
    const [isValidating, setIsValidating] = useState(false);

    useEffect(() => {
        const unsubscribe = onAuthStateChanged(auth, async (user) => {
            console.log("AuthProvider - User state changed:", user);
            
            // Si pas d'utilisateur, nettoyer tout
            if (!user) {
                setUser(null);
                setToken(null);
                setEmail(null);
                Cookies.remove('token');
                Cookies.remove('email');
                setLoading(false);
                return;
            }

            // Si utilisateur mais pas encore validé par le backend, ne pas le définir
            // Le hook useLogin se chargera de définir l'utilisateur après validation
            setLoading(false);
        });

        return () => unsubscribe();
    }, []);

    // Fonction pour définir l'utilisateur avec ses données (appelée par useLogin après validation)
    const setUserWithData = async (user: User) => {
        setUser(user);
        const token = await user.getIdToken();
        const userEmail = user.email;
        setToken(token);
        setEmail(userEmail);
        Cookies.set('token', token);
        Cookies.set('email', userEmail!);
    };

    // Fonction pour nettoyer les données utilisateur
    const clearUserData = () => {
        setUser(null);
        setToken(null);
        setEmail(null);
        Cookies.remove('token');
        Cookies.remove('email');
    };

    // Redéfinir setUser pour inclure la gestion des cookies et token
    const customSetUser: Dispatch<SetStateAction<User | null>> = (value) => {
        if (typeof value === 'function') {
            // Si c'est une fonction, l'appliquer à l'état actuel
            const newUser = value(user);
            if (newUser) {
                setUserWithData(newUser);
            } else {
                clearUserData();
            }
        } else {
            // Si c'est une valeur directe
            if (value) {
                setUserWithData(value);
            } else {
                clearUserData();
            }
        }
    };

    return (
        <AuthContext.Provider value={{ 
            user, 
            setUser: customSetUser, 
            loading, 
            token, 
            email, 
            isValidating, 
            setIsValidating 
        }}>
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