"use client";

import { Heading } from "@/components/ui/heading";
import { Button } from "@/components/ui/button";
import { Plus } from "lucide-react";
import { Separator } from "@/components/ui/separator";
import { useParams, useRouter } from "next/navigation";
import { DataTable } from "@/components/ui/data-table";
import { columns } from "./columns";
import { ApiList } from "@/components/ui/api-list";

export const BillBoardClient = () => {
    const router = useRouter();
    const params = useParams();
    const data = [{
        id: "1",
        label: "Billboard Fake Data",
        createdAt: "2021-10-10",

    },
    {
        id: "2",
        label: "Billboard Fake test",
        createdAt: "2021-10-10",

    }]
    return (
        <>
            <div className="flex items-center justify-between">
                <Heading
                    title={`Billboards (${data.length})`}
                    description="Manage billboards for your store"
                />
                <Button onClick={() => router.push('/billboards/new')}>
                    <Plus
                        className="mr-2 h-4 w-4"
                    />
                    Add New
                </Button>
            </div>
            <Separator />
            <DataTable
                searchKey="label"
                columns={columns}
                data={data}
            />
            <Heading title="API" description="API CALLS FOR Billboards" />
            <Separator />
            <ApiList
                entityName="billboards"
                entityIdName="billboardId"
            />
        </>
    )
}