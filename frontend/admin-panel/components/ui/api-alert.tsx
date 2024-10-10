"use client";


import { Alert, AlertTitle, AlertDescription } from "@/components/ui/alert";
import { Copy, Server } from "lucide-react";
import { Badge, BadgeProps } from "@/components/ui/badge";
import { Button } from "@/components/ui/button";
import toast from "react-hot-toast";

interface ApiAlertPrpos {
    title: string;
    description: string;
    variant: "public" | "admin" | "super_admin";
}

const textMap: Record<ApiAlertPrpos["variant"], string> = {
    public: "Public",
    admin: "Admin",
    super_admin: "Super Admin"
}

const variantMap: Record<ApiAlertPrpos["variant"], BadgeProps["variant"]> = {
    public: "outline",
    admin: "default",
    super_admin: "destructive"
}



export const ApiAlert: React.FC<ApiAlertPrpos> = ({
    title,
    description,
    variant = "public"
}) => {
    const onCopy = (description: string) => {
        navigator.clipboard.writeText(description);
        toast.success("API Route copied to the clipboard");
    }
    return (
        <Alert>
            <Server className="h4 w-4" />
            <AlertTitle className="flex items-center gap-x-2">
                {title}
                <Badge variant={variantMap[variant]}>
                    {textMap[variant]}
                </Badge>
            </AlertTitle>
            <AlertDescription className="mt-4 flex items-center justify-between">
                <code className="relative rounded bg-muted px-[0.3rem] font-mono text-sm font-semibold">
                    {description}
                </code>
                <Button variant="outline" size="icon" onClick={() => { onCopy(description) }} >
                    <Copy className="w-4 h-4" />
                </Button>
            </AlertDescription>
        </Alert>
    )
}