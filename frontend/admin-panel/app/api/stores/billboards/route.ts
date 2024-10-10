import { NextResponse } from 'next/server';

import { useAuthState } from "react-firebase-hooks/auth";
import { auth } from "@/app/firebase/config";

const [user] = useAuthState(auth);



export async function POST(req: Request) {
    try {
        const body = await req.json();

        const { label, imageUrl } = body;

        if (!user) {
            return new NextResponse("Unauthorized", { status: 401 });
        }

        if (!label) {
            return new NextResponse("Label is required", { status: 400 });
        }

        if (!imageUrl) {
            return new NextResponse("Image URL is required", { status: 400 });
        }

    } catch (error) {
        console.log(['STORES_POST', error])
        return new NextResponse("Internal Server Error", { status: 500 })
    }
}
