// components/AuthWrapper.tsx
"use client";

import { useEffect, useState } from 'react';
import { useRouter } from 'next/navigation';
import { useAuth } from '@/context/AuthContext';
import Navbar from './navbar';

const AuthWrapper = ({ children }: { children: React.ReactNode }) => {
    const { user, loading } = useAuth();
    const router = useRouter();

    useEffect(() => {
        if (!loading && !user) {
        }
    }, [user, loading, router]);

    if (loading) {
        return <div>Loading...</div>;
    }

    return (
        <>
            {user && <Navbar />}
            {children}
        </>
    );
};

export default AuthWrapper;
