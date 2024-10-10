// components/ClientWrapper.tsx
"use client";

import { ReactNode } from 'react';
import { AuthProvider } from '@/context/AuthContext';

const ClientWrapper = ({ children }: { children: ReactNode }) => {
    return <AuthProvider>{children}</AuthProvider>;
};

export default ClientWrapper;
