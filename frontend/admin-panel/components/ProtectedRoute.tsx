// components/ProtectedRoute.tsx
"use client";

import { ReactNode, useEffect } from 'react';
import { useRouter } from 'next/navigation';
import { useAuth } from '@/context/AuthContext';

const ProtectedRoute = ({ children }: { children: ReactNode }) => {
    const { user, loading } = useAuth();
    const router = useRouter();

    useEffect(() => {
        console.log("ProtectedRoute - Checking user:", user); // Debugging
        if (!loading && user === null) {
            router.push('/sign-in');
        }
    }, [user, loading, router]);

    if (loading || user === null) {
        return <div>Loading...</div>; // Spinner de chargement pendant le chargement de l'état utilisateur
    }

    return <>{children}</>;
};

export default ProtectedRoute;
