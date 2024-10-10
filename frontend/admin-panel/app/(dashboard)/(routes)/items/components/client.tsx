"use client";

import { Heading } from "@/components/ui/heading";
import { Button } from "@/components/ui/button";
import { Plus } from "lucide-react";
import { Separator } from "@/components/ui/separator";
import { useParams, useRouter } from "next/navigation";
import { DataTable } from "@/components/ui/data-table";
import { ApiList } from "@/components/ui/api-list";

import { ItemColumn, columns } from "./columns";

interface ItemClientProps {
    data: ItemColumn[]
}


export const ItemClient: React.FC<ItemClientProps> = ({
    data
}) => {
    const router = useRouter();
    const params = useParams();

    return (
        <>
            <div className="flex items-center justify-between">
                <Heading
                    title={`Items (${data.length})`}
                    description="Manage Items"
                />
                <Button onClick={() => router.push('/items/new')}>
                    <Plus
                        className="mr-2 h-4 w-4"
                    />
                    Add New Item
                </Button>
            </div>
            <Separator />
            <DataTable
                searchKey="title"
                columns={columns}
                data={data}
            />

        </>
    )
}