"use client";

import { cn } from "@/lib/utils";
import Link from "next/link"
import { useParams, usePathname } from "next/navigation"
export function MainNav({
    className,
    ...props
}: React.HTMLAttributes<HTMLElement>) {

    const pathname = usePathname();
    const params = useParams();

    const routes = [
        {
            href: `/`,
            label: 'Overview',
            active: pathname === `/`,
        },
        {
            href: `/billboards`,
            label: 'Billboards',
            active: pathname === `/billboards`,
        },
        {
            href: `/categories`,
            label: 'Categories',
            active: pathname === `/categories`,
        },
        {
            href: `/items`,
            label: 'Items',
            active: pathname === `/items`,
        },
        {
            href: `/users`,
            label: 'Users',
            active: pathname === `/users`,
        },
        {
            href: `/settings`,
            label: 'Settings',
            active: pathname === `/settings`,
        }
    ];
    return (
        <nav className={cn("flex items-center space-x-4 lg:space-x-6", className)}>
            {
                routes.map((route) => (
                    <Link
                        key={route.href}
                        href={route.href}
                        className={cn(
                            "text-sm font-medium transition-colors hover:text-primary",
                            route.active ? "text-black dark:text-white" : "text-muted-foreground"
                        )}
                    >
                        {route.label}
                    </Link>
                ))
            }
        </nav>

    )
};