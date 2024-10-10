// app/dashboard/layout.tsx
import { ReactNode } from 'react';
import ProtectedRoute from '@/components/ProtectedRoute';

const DashboardLayout = ({ children }: { children: ReactNode }) => {
    return (
        <ProtectedRoute>
            {children}
        </ProtectedRoute>
    );
};

export default DashboardLayout;
