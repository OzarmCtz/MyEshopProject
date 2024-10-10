"use client"

import { useOrigin } from "@/hooks/use-origin";
import { useParams } from "next/navigation";
import { ApiAlert } from "./api-alert";


interface ApiListProps {
    entityName: string;
    entityIdName: string;
}


export const ApiList: React.FC<ApiListProps> = ({
    entityName,
    entityIdName
}) => {
    const params = useParams();
    const origin = useOrigin();

    const baseUrl = `${origin}/api`;

    return (
        <div>
            <ApiAlert
                title="GET"
                variant="public"
                description={`${process.env.NEXT_PUBLIC_PB_PATH_API}/${entityName}`}
            />
            <ApiAlert
                title="LIST"
                variant="public"
                description={`${process.env.NEXT_PUBLIC_PB_PATH_API}/${entityName}/{${entityIdName}}`}
            />
            <ApiAlert
                title="POST"
                variant="admin"
                description={`${process.env.NEXT_PUBLIC_PB_PATH_API}/${entityName}`}
            />
            <ApiAlert
                title="PUT"
                variant="admin"
                description={`${process.env.NEXT_PUBLIC_PB_PATH_API}/${entityName}/${entityIdName}`}
            />
            <ApiAlert
                title="DELETE"
                variant="admin"
                description={`${process.env.NEXT_PUBLIC_PB_PATH_API}/${entityName}/${entityIdName}`}
            />
        </div>
    )
}