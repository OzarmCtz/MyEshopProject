"use client";

import { Heading } from "@/components/ui/heading";
import { Button } from "@/components/ui/button";
import { Plus } from "lucide-react";
import { Separator } from "@/components/ui/separator";
import { useParams, useRouter } from "next/navigation";
import { DataTable } from "@/components/ui/data-table";
import { ApiList } from "@/components/ui/api-list";

import { SubCategoryColumn, columns } from "./columns";

interface SubCategoryClientProps {
    data: SubCategoryColumn[]
}


export const SubCategoryClient: React.FC<SubCategoryClientProps> = ({
    data
}) => {
    const router = useRouter();
    const params = useParams();

    return (
        <>
            <div className="flex items-center justify-between">
                <Heading
                    title={`Sub Categories (${data.length})`}
                    description="Manage Sub Categories"
                />
                <Button onClick={() => router.push('/sub-categories/new')}>
                    <Plus
                        className="mr-2 h-4 w-4"
                    />
                    Add New Sub Category
                </Button>
            </div>
            <Separator />
            <DataTable
                searchKey="name"
                columns={columns}
                data={data}
            />
            <ApiList
                entityName="category"
                entityIdName="categoryid"
            />

        </>
    )
}