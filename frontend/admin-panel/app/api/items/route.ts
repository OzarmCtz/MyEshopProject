import { NextResponse } from 'next/server';
import { cookies } from 'next/headers';


export async function POST(req: Request) {
    try {

        const cookieStore = cookies();
        const token = cookieStore.get('token')?.value;


        const body = await req.json();
        const { i_title: name, i_description: description, i_price: price, i_quantity: quantity, i_picture_url: pictureUrl, i_file_path: filePath, i_is_disabled: isDisabled, iscl_sub_category_id: subCategory } = body;

        // Create an object with the received values
        const receivedValues = { name, description, price, quantity, pictureUrl, filePath, isDisabled, subCategory };


        // TODO : DO THAT FOR ALL API ROUTES
        for (const [key, value] of Object.entries(receivedValues)) {
            if (value === null || value === undefined || value === '') {
                return new NextResponse(JSON.stringify({ error: `Field '${key}' is required and is missing.` }), { status: 400 });
            }
        }


        if (!token) {
            return new NextResponse(JSON.stringify({ error: "You are not authenticated" }), { status: 401 });
        }

        try {
            const response = await fetch(`${process.env.PV_PATH_API}/items?token=${token}`, {
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
        const itemId = url.searchParams.get('id');

        if (!itemId) {
            return new NextResponse(JSON.stringify({ error: "Item ID is required" }), { status: 400 });
        }

        if (!token) {
            return new NextResponse(JSON.stringify({ error: "You are not authenticated" }), { status: 401 });
        }

        const response = await fetch(`${process.env.PV_PATH_API}/items/${itemId}?token=${token}`, {
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
        const { i_title: name, i_description: description, i_price: price, i_quantity: quantity, i_picture_url: pictureUrl, i_file_path: filePath, i_is_disabled: isDisabled, iscl_sub_category_id: subCategory } = body;

        // Create an object with the received values
        const receivedValues = { name, description, price, quantity, pictureUrl, filePath, isDisabled, subCategory };

        for (const [key, value] of Object.entries(receivedValues)) {
            if (value === null || value === undefined || value === '') {
                return new NextResponse(JSON.stringify({ error: `Field '${key}' is required and is missing.` }), { status: 400 });
            }
        }


        const url = new URL(req.url);
        const itemId = url.searchParams.get('id');

        if (!name || !description) {
            return new NextResponse(JSON.stringify({ error: "Name and description are required" }), { status: 400 });
        }


        if (!token) {
            return new NextResponse(JSON.stringify({ error: "You are not authenticated" }), { status: 401 });
        }

        const response = await fetch(`${process.env.PV_PATH_API}/items/${itemId}?token=${token}`, {
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
