"use client";

import { Heading } from "@/components/ui/heading";
import { Button } from "@/components/ui/button";
import { Plus } from "lucide-react";
import { Separator } from "@/components/ui/separator";
import { useParams, useRouter } from "next/navigation";
import { DataTable } from "@/components/ui/data-table";
import { ApiList } from "@/components/ui/api-list";

import { UserColumn, columns } from "./columns";

interface UserClientProps {
    data: UserColumn[]
}


export const UserClient: React.FC<UserClientProps> = ({
    data
}) => {
    const router = useRouter();
    const params = useParams();

    return (
        <>
            <div className="flex items-center justify-between">
                <Heading
                    title={`Users (${data.length})`}
                    description="Manage Users"
                />
            </div>
            <Separator />
            <DataTable
                searchKey="email"
                columns={columns}
                data={data}
            />

        </>
    )
}