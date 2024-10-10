import { NextResponse } from 'next/server';
import { cookies } from 'next/headers';


export async function POST(req: Request) {
    try {

        const cookieStore = cookies();
        const token = cookieStore.get('token')?.value;


        const body = await req.json();
        const { isc_name: name, isc_description: description, isc_picture_url: pictureUrl, icl_items_category_id: categoryId } = body;

        if (!name || !description || !pictureUrl || !categoryId) {
            return new NextResponse(JSON.stringify({ error: "All field value are required" }), { status: 400 });
        }


        if (!token) {
            return new NextResponse(JSON.stringify({ error: "You are not authenticated" }), { status: 401 });
        }

        try {
            const response = await fetch(`${process.env.NEXT_PUBLIC_PV_PATH_API}/items/sub/category?token=${token}`, {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json',
                },
                body: JSON.stringify(body)
            });

            const data = await response.json();

            if (response.status === 201) {
                return NextResponse.json(data, { status: 201 });
            } else {
                return new NextResponse(JSON.stringify(data), { status: response.status });
            }
        } catch (error) {
            console.error('Fetch error:', error);
            return new NextResponse(JSON.stringify({ error: "Unexpected error occurred during fetch" }), { status: 500 });
        }
    } catch (error) {
        console.error('Internal server error:', error);
        return new NextResponse(JSON.stringify({ error: "Internal Server Error" }), { status: 500 });
    }
}

export async function DELETE(req: Request) {
    try {

        const cookieStore = cookies();
        const token = cookieStore.get('token')?.value;


        const url = new URL(req.url);
        const subCategoryId = url.searchParams.get('id');

        if (!subCategoryId) {
            return new NextResponse(JSON.stringify({ error: "Category ID is required" }), { status: 400 });
        }

        if (!token) {
            return new NextResponse(JSON.stringify({ error: "You are not authenticated" }), { status: 401 });
        }

        const response = await fetch(`${process.env.NEXT_PUBLIC_PV_PATH_API}/items/sub/category/${subCategoryId}?token=${token}`, {
            method: 'DELETE',
        });

        const data = await response.json();

        if (response.status === 200) {
            return NextResponse.json(data, { status: 200 });
        } else {
            return new NextResponse(JSON.stringify(data), { status: response.status });
        }
    } catch (error) {
        return new NextResponse(JSON.stringify({ error: "Internal Server Error" }), { status: 500 });
    }
}

export async function PUT(req: Request) {
    try {

        const cookieStore = cookies();
        const token = cookieStore.get('token')?.value;


        const body = await req.json();
        const { isc_name: name, isc_description: description, isc_picture_url: pictureUrl } = body;

        const url = new URL(req.url);
        const subCategoryId = url.searchParams.get('id');

        if (!name || !description) {
            return new NextResponse(JSON.stringify({ error: "Name and description are required" }), { status: 400 });
        }


        if (!token) {
            return new NextResponse(JSON.stringify({ error: "You are not authenticated" }), { status: 401 });
        }

        const response = await fetch(`${process.env.NEXT_PUBLIC_PV_PATH_API}/items/sub/category/${subCategoryId}?token=${token}`, {
            method: 'PUT',
            headers: {
                'Content-Type': 'application/json',
            },
            body: JSON.stringify(body)
        });

        const data = await response.json();

        if (response.status === 200 || response.status === 201) {
            return NextResponse.json(data, { status: response.status });
        } else {
            return new NextResponse(JSON.stringify(data), { status: response.status });
        }
    } catch (error) {
        return new NextResponse(JSON.stringify({ error: "Internal Server Error" }), { status: 500 });
    }
}
