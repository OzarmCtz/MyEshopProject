// app/layout.tsx
import type { Metadata } from "next";
import { Inter } from "next/font/google";
import "./globals.css";
import { ModalProvider } from "@/providers/modal-provider";
import { ToastProvider } from "@/providers/toast-provider";
import ClientWrapper from '@/components/ClientWrapper';
import { AuthProvider } from '@/context/AuthContext';
import AuthWrapper from '@/components/AuthWrapper'; // Importez le nouveau composant AuthWrapper

const inter = Inter({ subsets: ['latin'] });

export const metadata: Metadata = {
  title: 'Admin Dashboard',
  description: 'Admin Dashboard',
};

export default function RootLayout({
  children,
}: {
  children: React.ReactNode;
}) {
  return (
    <html lang="en">
      <body className={inter.className}>
        <ToastProvider />
        <ModalProvider />
        <AuthProvider>
          <ClientWrapper>
            <AuthWrapper>
              {children}
            </AuthWrapper>
          </ClientWrapper>
        </AuthProvider>
      </body>
    </html>
  );
}
